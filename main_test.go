package main

import "testing"

const (
	layout string = "Jan 1 2006 at 15:04:01"
)

func (p *Pomodoro) TestGetCurrentTime(t *testing.T) {
	ti := p.GetCurrentTime().Format(layout)
	t.Log(ti)
}

func (p *Pomodoro) TestAddPomodoro(t *testing.T) {
	ti := p.GetCurrentTime()
	pomodoro := p.AddPomodoro(ti)
	t.Log(pomodoro)
}

func TestTickTilPomodoroEnd(t *testing.T) {
	// time.Ticker() ??
}
