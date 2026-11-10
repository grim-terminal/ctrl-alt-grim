// Package logfilter decides which log lines to keep.
package logfilter

import "strings"

type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
)

// ParseLevel maps a flag value to a Level; ok is false for an unknown name.
func ParseLevel(s string) (Level, bool) {
	switch strings.ToLower(s) {
	case "debug":
		return Debug, true
	case "info":
		return Info, true
	case "warn":
		return Warn, true
	case "error":
		return Error, true
	}
	return Info, false
}

// lineLevel returns the highest level token found in a line.
// A line with no level token counts as Info.
func lineLevel(line string) Level {
	upper := strings.ToUpper(line)
	switch {
	case strings.Contains(upper, "ERROR"):
		return Error
	case strings.Contains(upper, "WARN"):
		return Warn
	case strings.Contains(upper, "DEBUG"):
		return Debug
	}
	return Info
}

// Keep reports whether a line is at or above min and matches the text
// (an empty text matches every line). With invert, the text match is flipped:
// only lines that do not contain it are kept. invert has no effect when
// contains is empty.
func Keep(line string, min Level, contains string, invert bool) bool {
	if lineLevel(line) < min {
		return false
	}
	if contains == "" {
		return true
	}
	return strings.Contains(line, contains) != invert
}
