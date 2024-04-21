package common

import (
	"log/slog"
	"os"
	"testing"
)

type show struct {
	a string
	b []string
	c int
	f float64
}

func TestLog(t *testing.T) {
	ih := New(os.Stdout, &Options{
		slog.LevelDebug,
		NewStyles(Styles{
			TimeStamp: false,
			Delim: [2]string{
				":" + Color(LightYellow, "<"),
				EndColor(">") + "; ",
			},
		})})

	glog := slog.New(ih)

	// glog.Debug("testing", "abc", "dfg")
	s := show{
		a: "TestValue",
		b: []string{"SArrOne", "SArrTwo", "SArr3"},
		c: 2313421,
		f: 2142.14159265359,
	}

	glog.Error("begin flag capture", "name", s.a, "len", len(s.b), "args", s.b, "val", s.c, "fl", s.f)
	glog.Info("begin flag capture", "name", s.a, "len", len(s.b), "args", s.b, "val", s.c, "fl", s.f)
	glog.Warn("begin flag capture", "name", s.a, "len", len(s.b), "args", s.b, "val", s.c, "fl", s.f)
	glog.Debug("begin flag capture", "name", s.a, "len", len(s.b), "args", s.b, "val", s.c, "fl", s.f)
}
