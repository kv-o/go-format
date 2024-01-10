package desktop

import (
	"testing"
)

func equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

var execs = [...]struct{opts Opts; expanded []string; args []string; cmd string}{
	{
		opts: Opts{File: "/bin/echo", Icon: "fooview", Files: []string{"/sys/doc/", "/sys/man/", "/sys/src/"}},
		expanded: []string{"echo", "--icon", "fooview", "%F%", "`cat /etc/passwd | head -n1`"},
		args: []string{"echo", "%i", "%F%%", "`cat /etc/passwd | head -n1`"},
		cmd: "echo %i %F%% \"\\`cat /etc/passwd | head -n1\\`\"",
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

func TestExpandExec(t *testing.T) {
	for _, test := range execs {
		expanded := ExpandExec(test.args, test.opts)
		if !equal(expanded, test.expanded) {
			t.Fatalf(`ExpandExec("%s") equals %s not %s`, test.args, expanded, test.expanded)
		}
	}
}
