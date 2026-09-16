// Package printer wraps stdout writes so command output stays testable.
package printer

import (
	"fmt"
	"io"
	"os"
)

// Stdout is the writer used by the printer helpers; overridable in tests.
var Stdout io.Writer = os.Stdout

// Println prints args to Stdout followed by a newline.
func Println(a ...any) {
	fmt.Fprintln(Stdout, a...)
}

// Printf prints a formatted string to Stdout.
func Printf(format string, a ...any) {
	fmt.Fprintf(Stdout, format, a...)
}

// Fprintf prints a formatted string to the given writer.
func Fprintf(w io.Writer, format string, a ...any) {
	fmt.Fprintf(w, format, a...)
}
