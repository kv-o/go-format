package desktop

import (
	"testing"
)

var escapes = [...][2]string{
	{"\\ \r\n\t;", `\\\s\r\n\t\;`},
	{"ms\r\ndos", `ms\r\ndos`},
	{"foo", `foo`},
}

func TestEscapeString(t *testing.T) {
	for _, test := range escapes {
		escaped := EscapeString(test[0])
		if escaped != test[1] {
			t.Fatalf(`EscapeString("%s") equals %s not %s`, test[0], escaped, test[1])
		}
	}
}
