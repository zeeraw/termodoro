package main

import (
	"testing"
	"time"
)

const (
	layout string = "Jan 1 2006 at 15:04:01"
)

func TestNewPomodoro(t *testing.T) {
	n := NewPomodoro()
	if n.PomodoroActive != true {
		t.Fail()
		t.Log("Expected true but got", n.PomodoroActive)
	}
	t.Log(n.PomodoroActive)
}

func TestGetCurrentTime(t *testing.T) {
	ti := GetCurrentTime().Format(layout)
	t.Log(ti)
}

func TestSetStartTime(t *testing.T) {
	po := NewPomodoro()
	SetStartTime(po)
}

func TestGetStartTime(t *testing.T) {
	po := NewPomodoro()
	SetStartTime(po)
	st := GetStartTime(po)
	t.Log(st)
}

func TestGetPomodoroDuration(t *testing.T) {
	po := NewPomodoro()
	SetStartTime(po)
	time.Sleep(5 * time.Second)
	pd := GetPomodoroDuration(po)
	t.Log(pd)
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
