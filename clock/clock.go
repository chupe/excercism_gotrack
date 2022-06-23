package clock

import "fmt"

// Clock type here.
type Clock int

func normalize(m int) Clock {
	if m >= 24*60 {
		m %= 24 * 60
	}
	if m < 0 {
		m %= 24 * 60
		m += 24 * 60
	}
	return Clock(m)
}

func New(h, m int) Clock {
	return normalize(h*60 + m)
}

func (c Clock) Add(m int) Clock {
	return normalize(int(c) + m)
}

func (c Clock) Subtract(m int) Clock {
	return normalize(int(c) - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
