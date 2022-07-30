package common

import (
	"fmt"
	"time"
)

func UnixMicroTimeToHuman(timeepoch int64) string {
	t := time.Unix(timeepoch/1000000, 0)
	cur := time.Now()

	if cur.Sub(t) < 24*time.Hour {
		return "Today"
	}

	if cur.Sub(t) >= 24*time.Hour && cur.Sub(t) < 7*24*time.Hour {
		return "Last Week"
	}

	if cur.Sub(t) >= 7*24*time.Hour && cur.Sub(t) < 30*24*time.Hour {
		return "Last Month"
	}

	return fmt.Sprintf("%s %d %d", t.Month().String(), t.Day(), t.Year())
}
