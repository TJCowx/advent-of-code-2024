package utils

import (
	"fmt"
	"time"
)

type Timer struct {
	start      time.Time
	end        time.Time
	hasStarted bool
}

func BuildTimer() Timer {
	return Timer{start: time.Now(), end: time.Now()}
}

func (t *Timer) Start() {
	if t.hasStarted {
		fmt.Println("Restarting Timer!")
	}
	t.start = time.Now()
	t.hasStarted = true
}

func (t *Timer) End() {
	if !t.hasStarted {
		fmt.Println("TIMER HAS NOT STARTED")
		return
	}

	t.end = time.Now()
	t.hasStarted = false
}

func (t *Timer) TimeLapsed() string {
	return fmt.Sprintf("Timer ran for: %s", t.end.Sub(t.start))
}
