package desktop

import (
	"strings"
)

// UnescapeString returns s in its original form.
func UnescapeString(s string) string {
	for i := len(subTable) - 1; i > -1; i-- {
		subRule := subTable[i]
		s = strings.ReplaceAll(s, subRule[1], subRule[0])
	}
	return s
}

// Unmarshal parses the encoded desktop entry and stores the result in the value
// pointed to by v.
func Unmarshal(data []byte, v any) error {
	return nil
}
