package logfilter

import "testing"

func TestParseLevel(t *testing.T) {
	for in, want := range map[string]Level{"debug": Debug, "INFO": Info, "Warn": Warn, "error": Error} {
		got, ok := ParseLevel(in)
		if !ok || got != want {
			t.Errorf("ParseLevel(%q) = %v, %v; want %v, true", in, got, ok, want)
		}
	}
	if _, ok := ParseLevel("loud"); ok {
		t.Error("ParseLevel(loud) ok = true; want false")
	}
}

func TestKeep(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		min      Level
		contains string
		invert   bool
		want     bool
	}{
		{"error passes warn floor", "ERROR disk full", Warn, "", false, true},
		{"info dropped at warn floor", "INFO started", Warn, "", false, false},
		{"no token counts as info", "plain line", Info, "", false, true},
		{"no token dropped at warn", "plain line", Warn, "", false, false},
		{"grep match", "WARN slow query", Debug, "slow", false, true},
		{"grep miss", "WARN slow query", Debug, "fast", false, false},
		{"invert drops a match", "WARN slow query", Debug, "slow", true, false},
		{"invert keeps a miss", "WARN slow query", Debug, "fast", true, true},
		{"invert is ignored without grep", "WARN slow query", Debug, "", true, true},
		{"invert still honours the level floor", "INFO started", Warn, "slow", true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Keep(tt.line, tt.min, tt.contains, tt.invert); got != tt.want {
				t.Errorf("Keep(%q, %v, %q, %v) = %v; want %v", tt.line, tt.min, tt.contains, tt.invert, got, tt.want)
			}
		})
	}
}
