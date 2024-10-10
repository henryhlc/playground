package common

import (
	"fmt"
	"time"
)

func IntArgOrDefault(args []string, idx int, d int) (n int, useArg bool) {
	if idx >= len(args) {
		return d, false
	}
	var argN int
	_, err := fmt.Sscan(args[idx], &argN)
	if err != nil {
		return d, false
	}
	return argN, true
}

func StringArg(args []string, idx int) (s string, useArg bool) {
	if idx >= len(args) {
		return "", false
	}
	return args[idx], true
}

func TimeArg(args []string, idx int) (t time.Time, useArg bool) {
	if idx >= len(args) {
		return time.Now(), false
	}
	if args[idx] == "now" {
		return time.Now(), true
	}
	// Current date at given time
	t, err := time.Parse("15:04", args[idx])
	if err == nil {
		y, m, d := time.Now().Date()
		return time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, time.Local), true
	}
	// Current year at start of given date.
	t, err = time.Parse("1/2", args[idx])
	if err == nil {
		now := time.Now()
		return time.Date(now.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local), true
	}
	// Current year at start of given date and time.
	t, err = time.Parse("1/2 15:04", args[idx])
	if err == nil {
		now := time.Now()
		return time.Date(
			now.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), 0,
			0, time.Local), true
	}
	// Given year, month, day
	t, err = time.Parse("2006/01/02", args[idx])
	if err == nil {
		return t, true
	}
	// Given year, month, day, hour and minute.
	t, err = time.Parse("2006/01/02 15:04", args[idx])
	if err == nil {
		return t, true
	}
	return time.Now(), false
}

func DurationArg(args []string, idx int) (d time.Duration, useArg bool) {
	if idx >= len(args) {
		return time.Minute, false
	}
	d, err := time.ParseDuration(args[idx])
	if err != nil {
		return d, false
	}
	return d, true
}
