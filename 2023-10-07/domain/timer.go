package domain

import "time"

type Timer struct {
	//　外部から書き換えできないようにする
	count int64
}

func NewTimer(count int64) *Timer {
	return &Timer{
		count: count,
	}
}

func (t *Timer) Countdown() {
	t.count -= 1
}

func (t *Timer) CountStop() {}

func (t *Timer) GetDuration() time.Duration {
	return time.Duration(t.count) * time.Second
}
