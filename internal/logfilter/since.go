package logfilter

import (
	"strings"
	"time"
)

// Since reports whether a line is at or after cutoff. The timestamp is the
// first space-separated field, in RFC3339. A zero cutoff keeps everything,
// and so does a line without a parsable timestamp.
func Since(line string, cutoff time.Time) bool {
	if cutoff.IsZero() {
		return true
	}
	field, _, _ := strings.Cut(line, " ")
	ts, err := time.Parse(time.RFC3339, field)
	if err != nil {
		return true
	}
	return ts.After(cutoff)
}
