package nagios

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const epsilon = 1e-9

// Range represents a basic Nagios range object
// Note: Ranges are exclusive by default, eg x < min || x > max
type Range struct {
	Min       float64
	Max       float64
	Inclusive bool
}

// NewRange constructs a new Range
func NewRange(min, max float64, inclusive bool) *Range {
	return &Range{
		Min:       min,
		Max:       max,
		Inclusive: inclusive,
	}
}

// ParseRange creates a Range object from a Nagios-style Range string
// Ref: https://nagios-plugins.org/doc/guidelines.html#THRESHOLDFORMAT
func ParseRange(r string) (*Range, error) {
	val := strings.TrimSpace(r)
	if len(val) == 0 {
		return nil, fmt.Errorf("range must not be empty")
	}

	var (
		inclusive bool
		min       float64
		max       float64
		err       error
	)

	if strings.HasPrefix(val, "@") {
		inclusive = true
		val = val[1:]
	}

	parts := strings.Split(val, ":")
	if len(parts) > 2 {
		return nil, fmt.Errorf("range cannot have more than 2 parts")
	} else if len(parts) == 1 {
		parts = append([]string{"0"}, parts...)
	}

	if parts[0] == "~" {
		min = math.Inf(-1)
	} else {
		min, err = strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid start value %s", parts[0])
		}
	}

	if parts[1] == "~" || parts[1] == "" {
		max = math.Inf(1)
	} else {
		max, err = strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid end value %s", parts[1])
		}
	}

	if min > max {
		return nil, fmt.Errorf("start must not be greater than end")
	}

	return NewRange(min, max, inclusive), nil
}

func (r *Range) String() string {
	var res string
	if r.Inclusive {
		res += "@"
	}

	if r.Min == math.Inf(-1) {
		res += "~:"
	} else if r.Min != 0.0 {
		if mostlyIntegral(r.Min) {
			res += strconv.Itoa(int(r.Min)) + ":"
		} else {
			res += strconv.FormatFloat(r.Min, 'f', -1, 64) + ":"
		}
	}

	if r.Max == math.Inf(1) {
		res += "~"
	} else {
		if mostlyIntegral(r.Min) {
			res += strconv.Itoa(int(r.Max))
		} else {
			res += strconv.FormatFloat(r.Max, 'f', -1, 64)
		}
	}

	return res
}

// InRange checks whether a value is in range or not
func (r *Range) InRange(v float64) bool {
	if r.Inclusive {
		return v >= r.Min && v <= r.Max
	}
	return v < r.Min || v > r.Max
}

func mostlyIntegral(f float64) bool {
	_, frac := math.Modf(math.Abs(f))
	return frac < epsilon || frac > 1.0-epsilon
}
