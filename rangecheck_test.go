package nagios

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeCheckParse(t *testing.T) {
	assert := assert.New(t)

	c, err := NewRangeCheckParse("100", "1000")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	assert.Equal(NewRange(0, 100, false), c.Warn)
	assert.Equal(NewRange(0, 1000, false), c.Crit)

	_, err = NewRangeCheckParse("", "100")
	if err == nil {
		t.Errorf("expected range error")
	} else if !strings.Contains(err.Error(), "invalid warn") {
		t.Errorf("expected warn range error, got %s", err)
	}

	_, err = NewRangeCheckParse("100", "")
	if err == nil {
		t.Errorf("expected range error")
	} else if !strings.Contains(err.Error(), "invalid crit") {
		t.Errorf("expected crit range error, got %s", err)
	}
}

func TestRangeCheckTestValue(t *testing.T) {
	c := NewRangeCheck(NewRange(0, 100, false), NewRange(0, 1000, false))

	c.CheckValue(0)
	if c.Status != StatusOK {
		t.Errorf("expected status ok")
	}

	c.CheckValue(101)
	if c.Status != StatusWarn {
		t.Errorf("expected status warn")
	}

	c.CheckValue(1001)
	if c.Status != StatusCrit {
		t.Errorf("expected status crit")
	}
}
