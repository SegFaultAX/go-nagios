package nagios

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/segfaultax/go-nagios/util"
)

// PerfData is Nagios performance data
type PerfData struct {
	Label string
	Value float64
	Units string
	Warn  *Range
	Crit  *Range
	Min   *float64
	Max   *float64
}

// NewPerfData creates a simple Nagios Performance Data object with no ranges
func NewPerfData(label string, value float64, units string) PerfData {
	return PerfData{
		Label: label,
		Value: value,
		Units: units,
	}
}

// NewPerfDataWithRanges creates a complete Nagios Performance Data object with all range information
func NewPerfDataWithRanges(label string, value float64, units string, warn, crit *Range, min, max float64) PerfData {
	return PerfData{
		Label: label,
		Value: value,
		Units: units,
		Warn:  warn,
		Crit:  crit,
		Min:   &min,
		Max:   &max,
	}
}

func (pd PerfData) String() string {
	s := fmt.Sprintf("%s=%s%s", pd.Label, util.PrettyFloat(pd.Value, 6), pd.Units)
	args := make([]string, 4)
	if pd.Warn != nil {
		args[0] = pd.Warn.String()
	}
	if pd.Crit != nil {
		args[1] = pd.Crit.String()
	}
	if pd.Min != nil {
		args[2] = strconv.FormatFloat(*pd.Min, 'f', -1, 64)
	}
	if pd.Max != nil {
		args[3] = strconv.FormatFloat(*pd.Max, 'f', -1, 64)
	}
	return fmt.Sprintf("%s;%s", s, strings.Join(args, ";"))
}
