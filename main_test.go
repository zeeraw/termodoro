package pomodoro

import (
	"testing"
	"time"
)

const (
	layout string = "Jan 01 2006 at 15:04:01"
)

func TestNewPomodoro(t *testing.T) {
	n := NewPomodoro()
	if n.PomodoroActive != true {
		t.Fail()
		t.Log("\nExpected true but got", n.PomodoroActive)
	}
	t.Log("\n", n.PomodoroActive)
}

func TestGetCurrentTime(t *testing.T) {
	ti := GetCurrentTime().Format(layout)
	t.Log("\n", ti)
}

func TestSetStartTime(t *testing.T) {
	po := NewPomodoro()
	SetStartTime(po)
}

func TestGetStartTime(t *testing.T) {
	po := NewPomodoro()
	SetStartTime(po)
	st := GetStartTime(po)
	t.Log("\n", st)
}

func TestGetPomodoroDuration(t *testing.T) {
	po := NewPomodoro()
	SetStartTime(po)
	time.Sleep(1 * time.Microsecond)
	pd := GetPomodoroDuration(po)
	t.Log("\n", pd)
}

func TestAddPomodoro(t *testing.T) {
	ti := GetCurrentTime()
	pomodoro := AddPomodoro(ti)
	t.Log("\n", pomodoro)
}

func TestPomodoroTimer(t *testing.T) {
	var duration time.Duration = 1
	ti := Timer(duration, time.Microsecond)
	if ti != false {
		t.Fail()
	}
}

func TestFormatDate(t *testing.T) {
	date, time := FormatDate(2015, time.January, 1, 0, 0)
	t.Log("\nDate string:", date, "\ntime.Time format:", time)
	if date != "Jan 01 2015 at 00:00:01" {
		t.Fail()
	}
}
