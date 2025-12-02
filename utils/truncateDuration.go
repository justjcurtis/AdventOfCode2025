package utils

import "time"

func TruncateToDynamicUnit(duration time.Duration) time.Duration {
	units := []time.Duration{
		time.Hour,
		time.Minute,
		time.Second,
		time.Millisecond,
		time.Microsecond,
		time.Nanosecond,
	}

	for _, unit := range units {
		if duration >= unit {
			return duration.Truncate(unit)
		}
	}

	return duration
}
