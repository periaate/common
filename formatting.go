package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func HumanNumber(num string) string {
	ind := strings.Index(num, ".")
	var temp string
	if ind != -1 {
		temp = num[ind:]
		num = num[:ind]
	}

	for i := len(num) - 3; i > 0; i -= 3 {
		num = num[:i] + "," + num[i:]
	}

	return num + temp
}

func RelativeTime(s string) string {
	inp, err := time.Parse(time.DateTime, s)
	if err != nil {
		return ""
	}

	dur := time.Since(inp)
	days := dur.Hours() / 24
	switch {
	case days <= 31:
		if days < 1 {
			return "Today"
		}
		return strconv.Itoa(int(days)) + " days ago"
	case days <= 365:
		months := int(days / 30)
		if months == 1 {
			return "1 month ago"
		}
		return strconv.Itoa(months) + " months ago"
	default:
		years := int(days / 365)
		if years == 1 {
			return "1 year ago"
		}
		fmt.Println(days, years)
		return strconv.Itoa(years) + " years ago"
	}
}

const (
	reset = "\033[0m"

	Black        = 30
	Red          = 31
	Green        = 32
	Yellow       = 33
	Blue         = 34
	Magenta      = 35
	Cyan         = 36
	LightGray    = 37
	DarkGray     = 90
	LightRed     = 91
	LightGreen   = 92
	LightYellow  = 93
	LightBlue    = 94
	LightMagenta = 95
	LightCyan    = 96
	White        = 97
)

func Color(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s", strconv.Itoa(colorCode), v)
}

func EndColor(v string) string { return fmt.Sprintf("%s%s", v, reset) }

func Colorize(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}
