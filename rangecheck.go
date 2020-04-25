package nagios

import "fmt"

// RangeCheck is a kind of Nagios check with built-in warn/crit ranges
type RangeCheck struct {
	Check
	Warn *Range
	Crit *Range
}

// NewRangeCheck creates a new check with warning and critical ranges
func NewRangeCheck(warn *Range, crit *Range) *RangeCheck {
	return &RangeCheck{
		Warn: warn,
		Crit: crit,
	}
}

// NewRangeCheckParse creates a range check by parsing warning and critical as Nagios ranges
func NewRangeCheckParse(warn string, crit string) (*RangeCheck, error) {
	w, err := ParseRange(warn)
	if err != nil {
		return nil, fmt.Errorf("invalid warning range: %s", err)
	}

	c, err := ParseRange(crit)
	if err != nil {
		return nil, fmt.Errorf("invalid critical range: %s", err)
	}

	return NewRangeCheck(w, c), nil
}

// CheckValue updates the status based on available ranges
func (rc *RangeCheck) CheckValue(value float64) {
	if rc.Crit != nil && rc.Crit.InRange(value) {
		rc.Status = StatusCrit
	} else if rc.Warn != nil && rc.Warn.InRange(value) {
		rc.Status = StatusWarn
	} else {
		rc.Status = StatusOK
	}
}
