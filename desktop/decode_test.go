package desktop

import (
	"testing"
)

func TestUnescapeString(t *testing.T) {
	for _, test := range escapes {
		unescaped := UnescapeString(test[1])
		if unescaped != test[0] {
			t.Fatalf(`UnescapeString("%s") equals %s not %s`, test[1], unescaped, test[0])
		}
	}
}
