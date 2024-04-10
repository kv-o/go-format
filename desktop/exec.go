package desktop

import (
	"fmt"
	"strings"
)

var execSubTable = [...][2]string{
	{`\`, "\\\\"},
	{"`", "\\`"},
	{"$", "\\$"},
}

var reserved = [...]string{
	` `, "\t", "\n", `"`, `'`, `\`, `>`, `<`, `~`, `|`,
	`&`, `;`, `$`, `*`, `?`, `#`, `(`, `)`, "`",
}

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

// EscapeExec applies quoting and escaping rules to the command line arguments
// specified by args, and returns a single, sanitized command line. The
// resulting command line must be sanitized as a string with EscapeString before
// embedding into a Linux desktop entry.
func EscapeExec(args []string) string {
	cmd := ""
	for i, s := range args {
		changed := false
		for _, r := range reserved {
			if strings.Index(s, r) != -1 {
				changed = true
				break
			}
		}
		for _, subRule := range execSubTable {
			s = strings.ReplaceAll(s, subRule[0], subRule[1])
		}
		if i != 0 {
			cmd += " "
		}
		if changed {
			cmd += fmt.Sprintf(`"%s"`, s)
		} else {
			cmd += fmt.Sprintf("%s", s)
		}
	}
	return cmd
}

// ExpandExec expands any format specifiers in the unescaped command-line
// arguments specified by args, using the provided information in opts.
func ExpandExec(args []string, opts Opts) []string {
	var newArgs []string
	fflag := false
	uflag := false
	Fflag := false
	Uflag := false
	for _, arg := range args {
		switch arg {
		case "%F":
			if !Fflag {
				newArgs = append(newArgs, opts.Files...)
				Fflag = true
			}
		case "%U":
			if !Uflag {
				newArgs = append(newArgs, opts.URLs...)
				Uflag = true
			}
		case "%i":
			newArgs = append(newArgs, "--icon", opts.Icon)
		default:
			if !fflag {
				arg = strings.Replace(arg, "%f", opts.File, 1)
				fflag = true
			}
			if !uflag {
				arg = strings.Replace(arg, "%u", opts.URL, 1)
				uflag = true
			}
			arg = strings.ReplaceAll(arg, "%c", opts.Name)
			arg = strings.ReplaceAll(arg, "%k", opts.Path)
			arg = strings.ReplaceAll(arg, "%%", "%")
			newArgs = append(newArgs, arg)
		}
	}
	return newArgs
}

// UnescapeExec reverses any quoting and escaping rules applied to the command
// line arguments in cmd, and returns a slice of command-line arguments.
func UnescapeExec(cmd string) []string {
	args := []string{}
	// split into args
	/*for i, arg := range args {
		arg = 
		args[i] = arg
	}*/
	return args
}
