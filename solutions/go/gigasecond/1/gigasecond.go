// Package gigasecond has the functionality to add one gigasec to the given time
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds 1e9 seconds to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add((1e9 * time.Second))
}
