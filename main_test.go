package main

import (
	"testing"
	"time"
)

const (
	layout string = "Jan 1 2006 at 15:04:01"
)

type FakePomodoro struct {
	StartTime        time.Time
	CurrentTime      time.Time
	PomodoroActive   bool
	PomodoroEnd      time.Time
	PomodoroDuration time.Time
}

func TestGetCurrentTime(t *testing.T) {
	ti := GetCurrentTime().Format(layout)
	t.Log(ti)
}

func TestAddPomodoro(t *testing.T) {
	ti := GetCurrentTime()
	pomodoro := AddPomodoro(ti)
	t.Log(pomodoro)
}

func TestPomodoroTimer(t *testing.T) {
	var duration time.Duration = 1
	ti := PomodoroTimer(duration, time.Microsecond)
	if ti != false {
		t.Fail()
	}
}
