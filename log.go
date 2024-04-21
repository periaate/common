package common

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"strings"
	"sync"
	"time"
)

func NewClog(w io.Writer, opts *Options) *slog.Logger {
	if opts == nil {
		opts = &Options{
			slog.LevelDebug,
			NewStyles(Styles{
				TimeStamp: false,
				Delim: [2]string{
					":" + Color(LightYellow, "<"),
					EndColor(">") + "; ",
				},
			})}
	}

	return slog.New(New(w, opts))
}

type Logger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Debug(msg string, args ...any)
}

type DummyLogger struct{}

func (DummyLogger) Error(_ string, _ ...any) {}
func (DummyLogger) Info(_ string, _ ...any)  {}
func (DummyLogger) Warn(_ string, _ ...any)  {}
func (DummyLogger) Debug(_ string, _ ...any) {}

type ClogHandler struct {
	Options
	Mut *sync.Mutex
	Out io.Writer
}

type Options struct {
	Level slog.Leveler
	Styles
}

func (l *ClogHandler) SetLogLoggerLevel(lvl slog.Level) {
	l.Level = lvl
}

type Styles struct {
	TimeStamp bool
	Delim     [2]string
}

func NewStyles(st Styles) Styles {
	if len(st.Delim[0]) == 0 && len(st.Delim[1]) == 0 {
		if st.Delim == [2]string{} {
			st.Delim[0] = "=\""
			st.Delim[1] = "\";"
		}
	}

	return st
}

func DefGetV(val slog.Value) string {
	switch val.Kind() {
	case slog.KindTime:
		return val.Time().Format(time.RFC3339Nano)
	case slog.KindInt64, slog.KindUint64, slog.KindFloat64:
		return HumanNumber(val.String())
	default:
		return val.String()
	}
}

func New(out io.Writer, opts *Options) *ClogHandler {
	h := &ClogHandler{Out: out, Mut: &sync.Mutex{}}
	if opts != nil {
		h.Options = *opts
	}
	if h.Level == nil {
		h.Level = slog.LevelInfo
	}
	return h
}

func (h *ClogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.Level.Level()
}

func (h *ClogHandler) WithGroup(name string) slog.Handler { return h }

func (h *ClogHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return h }

func (h *ClogHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)
	if !r.Time.IsZero() {
		buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time))
	}
	buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level))
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		buf = h.appendAttr(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)))
	}
	buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message))
	r.Attrs(func(a slog.Attr) bool {
		buf = h.appendAttr(buf, a)
		return true
	})
	buf = append(buf, '\n')
	h.Mut.Lock()
	defer h.Mut.Unlock()
	_, err := h.Out.Write(buf)
	return err
}

var m = map[string]string{
	"DEBUG": Colorize(White, "DEBUG"),
	"INFO":  Colorize(Cyan, "INFO"),
	"WARN":  Colorize(Yellow, "WARN"),
	"ERROR": Colorize(Red, "ERROR"),
}

func (h *ClogHandler) appendAttr(buf []byte, a slog.Attr) []byte {
	a.Value = a.Value.Resolve()
	if a.Equal(slog.Attr{}) {
		return buf
	}

	switch a.Key {
	case slog.LevelKey:
		buf = fmt.Appendf(buf, "%s", m[a.Value.String()])
		return buf
	case slog.TimeKey:
		if h.TimeStamp {
			buf = fmt.Appendf(buf, "%s", DefGetV(a.Value))
		}
		return buf
	case slog.SourceKey:
		ind := strings.LastIndex(a.Value.String(), "/")
		if ind != -1 {
			buf = fmt.Appendf(buf, " @ %s\t", a.Value.String()[ind+1:])
			return buf
		}
	}

	if res, ok := h.fmtKV(a); ok {
		buf = fmt.Appendf(buf, "%s", res)
	}
	return buf
}

func (st *Styles) fmtKV(val slog.Attr) (res string, ok bool) {
	return fmt.Sprintf("%s%s%s%s", val.Key, st.Delim[0], DefGetV(val.Value), st.Delim[1]), true
}
