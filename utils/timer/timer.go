package timer

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Timer struct {
	time float32
}

func NewTimer() Timer {
	return Timer{
		time: 0.0,
	}
}

func (t *Timer) Start(time float32) {
	t.time = time
}

func (t *Timer) Update() {
	if t.time > 0 {
		t.time -= rl.GetFrameTime()
	}
}

func (t *Timer) IsFinished() bool {
	return t.time <= 0
}
