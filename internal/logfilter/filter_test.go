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
		want     bool
	}{
		{"error passes warn floor", "ERROR disk full", Warn, "", true},
		{"info dropped at warn floor", "INFO started", Warn, "", false},
		{"no token counts as info", "plain line", Info, "", true},
		{"no token dropped at warn", "plain line", Warn, "", false},
		{"grep match", "WARN slow query", Debug, "slow", true},
		{"grep miss", "WARN slow query", Debug, "fast", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Keep(tt.line, tt.min, tt.contains); got != tt.want {
				t.Errorf("Keep(%q, %v, %q) = %v; want %v", tt.line, tt.min, tt.contains, got, tt.want)
			}
		})
	}
}
