// Package desktop implements encoding and decoding of Linux desktop entries.
//
// Refer to the following document for the complete format specification:
// https://specifications.freedesktop.org/desktop-entry-spec/desktop-entry-spec-latest.html
package desktop

// EscapeString returns a sanitized version of s.
func EscapeString(s string) string {
	return ""
}

// Marshal returns the Linux desktop entry encoding of v.
func Marshal(v any) ([]byte, error) {
	return nil, nil
}
