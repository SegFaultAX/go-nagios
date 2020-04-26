package util

import "testing"

func TestPrettyFloat(t *testing.T) {
	v := 1.2345
	if s := PrettyFloat(v, 4); s != "1.2345" {
		t.Errorf("expected all available precision, got %s", s)
	}
	if s := PrettyFloat(v, 2); s != "1.23" {
		t.Errorf("expected less precision, got %s", s)
	}
	if s := PrettyFloat(v, 8); s != "1.2345" {
		t.Errorf("expected exactly 4 degrees of precision, got %s", s)
	}
}
