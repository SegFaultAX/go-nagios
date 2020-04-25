package nagios

import (
	"fmt"
	"os"
	"strings"
)

// Status describes a Nagios exit code
type Status struct {
	Label    string
	ExitCode int
}

// Nagios return status codes
var (
	StatusOK      = Status{"OK", 0}
	StatusWarn    = Status{"WARN", 1}
	StatusCrit    = Status{"CRIT", 2}
	StatusUnknown = Status{"UNKNOWN", 3}
)

// Check is the core Nagios check holder
type Check struct {
	Status   Status
	Message  string
	PerfData []PerfData
}

// NewCheck initializes a new Check in an unknown state
func NewCheck() *Check {
	return &Check{
		Status: StatusUnknown,
	}
}

// SetMessage sets check message (with sprintf formatting)
func (c *Check) SetMessage(msg string, vs ...interface{}) {
	c.Message = fmt.Sprintf(msg, vs...)
}

// AddPerfData adds Nagios performance data to the check result
func (c *Check) AddPerfData(pd PerfData) {
	c.PerfData = append(c.PerfData, pd)
}

// OK sets the message and return status to OK
func (c *Check) OK(msg string, vs ...interface{}) {
	c.SetMessage(msg, vs...)
	c.Status = StatusOK
}

// Warning sets the message and return status to Warning
func (c *Check) Warning(msg string, vs ...interface{}) {
	c.SetMessage(msg, vs...)
	c.Status = StatusWarn
}

// Critical sets the message and return status to Critical
func (c *Check) Critical(msg string, vs ...interface{}) {
	c.SetMessage(msg, vs...)
	c.Status = StatusCrit
}

// Unknown sets the message and return status to Unknown
func (c *Check) Unknown(msg string, vs ...interface{}) {
	c.SetMessage(msg, vs...)
	c.Status = StatusUnknown
}

func (c *Check) String() string {
	msg := c.Status.Label
	if c.Message != "" {
		msg += fmt.Sprintf(": %s", c.Message)
	}

	if len(c.PerfData) > 0 {
		var pd []string
		for _, d := range c.PerfData {
			pd = append(pd, d.String())
		}
		msg = fmt.Sprintf("%s|%s", msg, strings.Join(pd, "|"))
	}

	return msg
}

// Done prints the output of the check and exits with the appropriate code
func (c *Check) Done() {
	fmt.Println(c)
	os.Exit(c.Status.ExitCode)
}
