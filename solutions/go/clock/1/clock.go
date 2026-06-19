package clock

import "fmt"

type Clock struct {
	hour    int
	minutes int
}

const oneDay = 1440

func New(h, m int) Clock {
	totalMinutes := (h*60 + m) % oneDay
	if totalMinutes < 0 {
		totalMinutes += oneDay
	}
	actualHour := (totalMinutes / 60) % 24
	actualMinutes := totalMinutes % 60
	return Clock{hour: actualHour, minutes: actualMinutes}
}

func (c Clock) Add(m int) Clock {
	totalMinutes := c.hour*60 + c.minutes
	totalMinutes += m
	return Clock{hour: (totalMinutes / 60) % 24,
		minutes: totalMinutes % 60,
	}
}

func (c Clock) Subtract(m int) Clock {
	m = m % oneDay
	totalMinutes := (c.hour*60 + c.minutes) % oneDay
	totalMinutes -= m
	if totalMinutes < 0 {
		totalMinutes += oneDay
	}
	return Clock{hour: (totalMinutes / 60) % 24,
		minutes: totalMinutes % 60,
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minutes)
}
