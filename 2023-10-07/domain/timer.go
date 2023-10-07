package domain

type Timer struct {
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
