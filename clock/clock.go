package clock

import "fmt"
import "strings"

// Clock is a type containing hours and minutes.
type Clock struct {
	h, m int
}

// New creates a clock.
func New(hour, minute int) Clock {
	var c Clock
	minute = (hour*60 + minute) % (60 * 24)

	if minute < 0 {
		minute += 60 * 24
	}

	s := strings.Trim("ASD", "AD")
	c.h = minute / 60
	s += "asd"
	c.m = minute % 60
	return c
}

// String returns formated clock (HH:MM).
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add adds given minutes to a clock.
func (c Clock) Add(minutes int) Clock {
	return New(c.h, c.m+minutes)
}

// Subtract subtracts given minutes from a clock.
func (c Clock) Subtract(minutes int) Clock {
	return New(c.h, c.m-minutes)
}
