package main

import (
	"fmt"
	"time"
)

const (
	// PomodoroLength defines the default length of a pomodoro
	// should be later sourced via .termodororc
	PomodoroLength time.Duration = 25

	// Layout represents the default time layout format to use
	Layout string = "Jan 1 2006 at 15:04:01"
)

func main() {}

// Pomodoro struct, not in use right now but will be, eventually
type Pomodoro struct {
	StartTime        time.Time
	PomodoroActive   bool
	PomodoroEnd      time.Time
	PomodoroDuration time.Time
}

// NewPomodoro creates a pomodoro object in memory
func NewPomodoro() *Pomodoro {
	return &Pomodoro{PomodoroActive: true}
}

// GetCurrentTime gets the current time
func GetCurrentTime() (t time.Time) {
	t = time.Now()
	return
}

// GetPomodoroDuration calculates how much time has passed since the pomodoro started
func GetPomodoroDuration(p *Pomodoro) (t time.Duration) {
	t = time.Since(p.StartTime)
	return
}

// SetStartTime sets the starting time of the pomodoro
// idealy I want to use this to set pomodoros in advance
func SetStartTime(p *Pomodoro) {
	p.StartTime = GetCurrentTime()
	fmt.Println("Pomodoro started at: ", p.StartTime.Format(Layout))
	return
}

// GetStartTime gets the time at which the pomodoro started
func GetStartTime(p *Pomodoro) (st time.Time) {
	st = p.StartTime
	return
}

// AddPomodoro calculates when the next Pomodoro should end
func AddPomodoro(t time.Time) (pomodoro time.Time) {
	pomodoro = t.Add(PomodoroLength * time.Minute)
	return
}

// PomodoroTimer counts down the time until active pomodoro ends
func PomodoroTimer(length, unit time.Duration) (active bool) {
	timer := time.NewTimer(length * unit)
	<-timer.C
	active = false
	fmt.Println("Pomodoro ended")
	return
}
