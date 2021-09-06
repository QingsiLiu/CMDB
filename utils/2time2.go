package utils

import (
	"strings"
	"time"
)

const (
	TimeLayout = "2006-01-02 15:04"
)

func String2Time(t string) *time.Time {
	timeNow, err := time.Parse(TimeLayout, strings.ReplaceAll(t, "T", " "))
	if err != nil {
		return nil
	}
	return &timeNow
}

func Time2String(t *time.Time) string {
	if t == nil {
		return "--"
	}
	return t.Format(TimeLayout)
}
