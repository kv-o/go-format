package desktop

import (
	"os/exec"
)

// Opts represents the possible arguments that an Exec format specifier may
// expand to.
type Opts struct {
	// File will expand the `%f` format specifier. It is a single file path
	// (even if multiple file paths are provided to the application launcher. If
	// File is a URL, the file will be downloaded and `%f` will expand to the
	// path to the downloaded temporary file.
	File string
	// Like File, but will expand the `%F` format specifier, and may include
	// multiple files which will expand to several command-line arguments.
	Files []string
	// URL will expand the `%u` format specifier. It is a single URL. Locally
	// available files may either be passed as `file://` URLs or as a Unix file
	// path.
	URL string
	// Like URL, but will expand the `%U` format specifier and may include
	// multiple URLs which will expand to several command-line arguments.
	URLs []string
	// Icon will expand the `%i` format specifier. It represents the value of
	// the desktop entry Icon key.
	Icon string
	// Name will expand the `%c` format specifier. It represents the appropriate
	// localized value of the desktop entry Name key.
	Name string
	// Path will expand the `%k` format specifier. It represents the path to the
	// desktop entry as either a URI (if accessed through Linux-specific
	// filesystems not attached to the system file view), a Unix file path, or
	// an empty string if no location is known.
	Path string
}

// Command returns a ready-to-execute exec.Command from the unescaped, expanded
// arguments specified by args.
func Command(args []string) *exec.Cmd {
	return nil
}

// EscapeExec applies quoting and escaping rules to the command line arguments
// specified by args, and returns a single, sanitized command line.
func EscapeExec(args []string) string {
	return ""
}

// ExpandExec expands any format specifiers in the unescaped command-line
// arguments specified by args, using the provided information in opts.
func ExpandExec(args []string, opts Opts) string {
	return ""
}

// UnescapeExec reverses any quoting and escaping rules applied to the command
// line arguments in cmd, and returns a slice of command-line arguments.
func UnescapeExec(cmd string) []string {
	return nil
}
