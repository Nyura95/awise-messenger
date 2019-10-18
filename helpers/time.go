package helpers

import "time"

// GetUtc for get date in UTC
func GetUtc() time.Time {
	return time.Time.UTC(time.Now())
}
