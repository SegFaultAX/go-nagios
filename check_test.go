package nagios

import (
	"testing"
)

func TestStatus(t *testing.T) {
	c := NewCheck()
	if c.Status != StatusUnknown {
		t.Errorf("expected default status of unknown")
	}

	ok := NewCheck()
	ok.OK("Ok")
	if ok.Status != StatusOK && ok.Message != "Ok" {
		t.Errorf("expected status and message to be set to ok")
	}

	warn := NewCheck()
	warn.Warning("Warn")
	if warn.Status != StatusWarn && warn.Message != "Warn" {
		t.Errorf("expected status and message to be set to warn")
	}

	crit := NewCheck()
	crit.Critical("Crit")
	if crit.Status != StatusCrit && crit.Message != "Crit" {
		t.Errorf("expected status and message to be set to crit")
	}

	unknown := NewCheck()
	unknown.Unknown("Unknown")
	if unknown.Status != StatusUnknown && unknown.Message != "Unknown" {
		t.Errorf("expected status and message to be set to unknown")
	}
}

func TestPerfData(t *testing.T) {
	c := NewCheck()
	if len(c.PerfData) != 0 {
		t.Errorf("expect default perfdata to be empty")
	}

	c.AddPerfData(NewPerfData("v", 123, "kb"))
	if len(c.PerfData) != 1 || c.PerfData[0] != NewPerfData("v", 123, "kb") {
		t.Errorf("expected perf data to match")
	}
}

func TestCheckString(t *testing.T) {
	c := NewCheck()
	if c.String() != "UNKNOWN" {
		t.Errorf("expected default output to be unknown")
	}

	ok := NewCheck()
	ok.OK("everything good")
	if ok.String() != "OK: everything good" {
		t.Errorf("expected ok output, was %s", ok)
	}

	warn := NewCheck()
	warn.Warning("uh oh")
	if warn.String() != "WARN: uh oh" {
		t.Errorf("expected warn output, was %s", warn)
	}

	crit := NewCheck()
	crit.Critical("oh dear")
	if crit.String() != "CRIT: oh dear" {
		t.Errorf("expected crit output, was %s", crit)
	}

	unknown := NewCheck()
	unknown.Unknown("huh?")
	if unknown.String() != "UNKNOWN: huh?" {
		t.Errorf("expected unknown output, was %s", unknown)
	}

	pd := NewPerfData("a", 123, "kb")
	pdc := NewCheck()
	pdc.AddPerfData(pd)
	pdc.OK("all good")
	if pdc.String() != "OK: all good|"+pd.String() {
		t.Errorf("expected ok with perf data output, was %s", pdc)
	}
}
