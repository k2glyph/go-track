package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	timeUnix := float64(t.Unix())
	gigaSecond := math.Pow(10, 9)
	t = time.Unix(int64(timeUnix+gigaSecond), 0)
	return t
}
