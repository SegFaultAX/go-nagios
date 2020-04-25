package nagios

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExclusiveRange(t *testing.T) {
	r := NewRange(0, 100, false)

	if !r.InRange(101) {
		t.Errorf("expected 101 to be in %s", r)
	}

	if !r.InRange(-50) {
		t.Errorf("expected -50 to be in %s", r)
	}

	if r.InRange(r.Min) {
		t.Errorf("expected min to not be in %s", r)
	}

	if r.InRange(r.Max) {
		t.Errorf("expected max to not be in %s", r)
	}
}

func TestInclusiveRange(t *testing.T) {
	r := NewRange(0, 100, true)

	if !r.InRange(50) {
		t.Errorf("expected 50 to be in %s", r)
	}

	if !r.InRange(r.Min) {
		t.Errorf("expect min to be in %s", r)
	}

	if !r.InRange(r.Max) {
		t.Errorf("expected max to be in %s", r)
	}
}

func TestParseRange(t *testing.T) {
	assert := assert.New(t)

	examples := map[string]*Range{
		"10":  NewRange(0, 10, false),
		"@10": NewRange(0, 10, true),
		"~":   NewRange(0, math.Inf(1), false),
		"~:":  NewRange(math.Inf(-1), math.Inf(1), false),
	}

	for k, ex := range examples {
		r, err := ParseRange(k)
		if err != nil {
			t.Fatalf("unexpected error while parsing %s: %s", k, err)
		}
		assert.EqualValues(ex, r, "expected ranges to be equal")
	}
}

func TestInvalidParseRange(t *testing.T) {
	examples := map[string]string{
		"":        "not be empty",
		"1:2:3":   "more than 2 parts",
		":":       "invalid start",
		"123:abc": "invalid end",
		"456:123": "start must not be greater",
	}

	for k, m := range examples {
		_, err := ParseRange(k)
		if err == nil || !strings.Contains(err.Error(), m) {
			t.Errorf("expected parse to fail with error '%s', was '%s'", m, err)
		}
	}
}

func TestRangeString(t *testing.T) {
	examples := map[string]string{
		"1":           "1",
		"1:2":         "1:2",
		"~:~":         "~:~",
		"@-1.23:4.56": "@-1.23:4.56",
	}

	for k, ex := range examples {
		r, err := ParseRange(k)
		if err != nil {
			t.Fatalf("unexpected error while parsing %s: %s", k, err)
		}
		if r.String() != ex {
			t.Errorf("expected %s to encode as %s but was %s", k, ex, r.String())
		}
	}
}
