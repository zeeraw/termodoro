package main

import (
	"fmt"
	"time"
)

const (
	// PomodoroLength defines the default length of a pomodoro
	// should be later sourced via .termodororc
	PomodoroLength time.Duration = 25
)

func main() {

}

// Pomodoro struct, not in use right now but will be, eventually
type Pomodoro struct {
	StartTime        time.Time
	PomodoroActive   bool
	PomodoroEnd      time.Time
	PomodoroDuration time.Time
}

// GetCurrentTime gets the current time
func GetCurrentTime() (t time.Time) {
	t = time.Now()
	return
}

// GetPomodoroDuration calculates how much time has passed since the pomodoro started
func (p *Pomodoro) GetPomodoroDuration(starttime time.Time) (t time.Duration) {
	t = time.Since(p.StartTime)
	return
}

// SetStartTime sets the starting time of the pomodoro
// idealy I want to use this to set pomodoros in advance
func (p *Pomodoro) SetStartTime() (t time.Time) {
	t = time.Now()
	p.StartTime = t
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
	// PomodoroActive = true
	<-timer.C
	active = false
	fmt.Println("Timer expired! This pomodoro has ended!")
	// PomodoroActive = false
	return
}
