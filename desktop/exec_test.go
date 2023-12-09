package desktop

import (
	"testing"
)

var execs = [...]struct{args []string; cmd string}{
	{
		args: []string{"echo", "`cat /etc/passwd | head -n1`"},
		cmd: "echo \"\\`cat /etc/passwd | head -n1\\`\"",
	},
}

func TestEscapeExec(t *testing.T) {
	for _, test := range execs {
		escaped := EscapeExec(test.args)
		if escaped != test.cmd {
			t.Fatalf(`EscapeExec("%s") equals %s not %s`, test.args, escaped, test.cmd)
		}
	}
}
