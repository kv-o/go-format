// Package desktop implements encoding and decoding of Linux desktop entries.
//
// Refer to the following document for the complete format specification:
// https://specifications.freedesktop.org/desktop-entry-spec/desktop-entry-spec-latest.html
package desktop

import (
	"strings"
)

var subTable = [...][2]string{
	{`\`,  `\\`},
	{` `,  `\s`},
	{"\n", `\n`},
	{"\t", `\t`},
	{"\r", `\r`},
	{`;`,  `\;`},
}

// EscapeString returns a sanitized version of s.
func EscapeString(s string) string {
	for _, subRule := range subTable {
		s = strings.ReplaceAll(s, subRule[0], subRule[1])
	}
	return s
}

// Marshal returns the Linux desktop entry encoding of v.
func Marshal(v any) ([]byte, error) {
	return nil, nil
}
