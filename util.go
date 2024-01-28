package interprog

import (
	"fmt"
	"strings"
	"time"
)

func formatDuration(d time.Duration) string {
	var parts []string

	hours := d / time.Hour
	d %= time.Hour
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}

	minutes := d / time.Minute
	d %= time.Minute
	if hours > 0 || minutes > 0 {
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}

	seconds := d / time.Second
	fraction := d % time.Second
	if fraction > 0 {
		parts = append(parts, fmt.Sprintf("%.3fs", float64(seconds)+float64(fraction)/float64(time.Second)))
	} else {
		parts = append(parts, fmt.Sprintf("%ds", seconds))
	}

	return strings.Join(parts, "")
}
