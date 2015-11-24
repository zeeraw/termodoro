package main

import (
	"time"
)

const (
	// PomodoroLength defines the default length of a pomodoro
	// should be later sourced via .termodororc
	PomodoroLength time.Duration = 25
)

func main() {

}

// Pomodoro interface implements getters and setters for active pomodoro
type Pomodoro struct {
	currentTime time.Time
	pomodoroEnd time.Time
}

// GetCurrentTime gets the current time
// Implements the Pomodoro struct
func (p *Pomodoro) GetCurrentTime() (t time.Time) {
	t = time.Now()
	p.currentTime = t
	return
}

// AddPomodoro calculates when the next Pomodoro should end
func (p *Pomodoro) AddPomodoro(t time.Time) (pomodoro time.Time) {
	pomodoro = t.Add(PomodoroLength * time.Minute)
	p.pomodoroEnd = pomodoro
	return
}

// TickTilPomodoroEnd counts down the time until active pomodoro ends
func (p *Pomodoro) TickTilPomodoroEnd(pomodoro time.Time) {
	// time.Ticker ??
}
