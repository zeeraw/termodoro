package main

import (
	"testing"
	// "time"
)

func TestGetCurrentTime(t *testing.T) {
	layout := "Jan 1 2006 at 3:04pm"
	ti := GetCurrentTime().Format(layout)
	t.Log(ti)
}
