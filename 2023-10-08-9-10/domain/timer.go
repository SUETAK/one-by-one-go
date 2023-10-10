package domain

import "time"

type Timer struct {
	//　外部から書き換えできないようにする
	count      int64
	state      string
	keyMonitor *KeyMonitor
}

func NewTimer(count int64) *Timer {
	return &Timer{
		count:      count,
		keyMonitor: NewKeyMonitor(),
	}
}

func (t *Timer) Countdown() {
	if t.state == "stop" {
		return
	}
	if t.count == 0 {
		return
	}
	time.Sleep(1 * time.Second)
	t.count -= 1
}

func (t *Timer) CountStop() {
	t.state = "stop"
}

func (t *Timer) CountStart() {
	t.state = "start"
}

func (t *Timer) GetState() string {
	return t.state
}

func (t *Timer) GetDuration() time.Duration {
	return time.Duration(t.count) * time.Second
}

func (t *Timer) GetKeyMonitor() *KeyMonitor {
	return t.keyMonitor
}
