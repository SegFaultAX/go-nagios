package util

import (
	"strconv"
	"strings"
)

// PrettyFloat is like FormatFloat(v, 'f', -1, 64), except that it will
// resort to using a smaller precision when the available precision
// exceeds prec. Useful when formatting check outputs and performance
// data that don't need excessive degrees of precision.
func PrettyFloat(v float64, prec int) string {
	s := strconv.FormatFloat(v, 'f', -1, 64)
	parts := strings.Split(s, ".")
	if len(parts) > 1 && len(parts[1]) > prec {
		return strconv.FormatFloat(v, 'f', prec, 64)
	}
	return s
}
