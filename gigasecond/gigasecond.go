// Package gigasecond contains time exercise
package gigasecond

import "time"

// AddGigasecond adds gigasecond to given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
