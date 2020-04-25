package nagios

import (
	"math"
	"testing"
)

func TestPefDataString(t *testing.T) {
	examples := map[string]PerfData{
		"=0;;;;":       NewPerfData("", 0.0, ""),
		"a=1.23kb;;;;": NewPerfData("a", 1.23, "kb"),
		"a=1.23kb;~:7;~:15;0;100": NewPerfDataWithRanges(
			"a",
			1.23,
			"kb",
			NewRange(math.Inf(-1), 7.0, false),
			NewRange(math.Inf(-1), 15.0, false),
			0,
			100,
		),
	}

	for s, pd := range examples {
		if pd.String() != s {
			t.Errorf("expected %s, got %s", s, pd.String())
		}
	}
}
