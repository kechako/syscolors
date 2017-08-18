package main

import (
	"fmt"
	"io"

	colorable "github.com/mattn/go-colorable"
)

const gopher = "ʕ◔ϖ◔ʔ"

var fgs = []string{
	"m",
	"1m",
	"30m",
	"1;30m",
	"31m",
	"1;31m",
	"32m",
	"1;32m",
	"33m",
	"1;33m",
	"34m",
	"1;34m",
	"35m",
	"1;35m",
	"36m",
	"1;36m",
	"37m",
	"1;37m",
}
var bgs = []string{
	"40m",
	"41m",
	"42m",
	"43m",
	"44m",
	"45m",
	"46m",
	"47m",
}

func main() {
	w := colorable.NewColorableStdout()

	// print title
	io.WriteString(w, "System colors:\n")

	// column header
	io.WriteString(w, "             ")
	for _, bg := range bgs {
		fmt.Fprintf(w, "     %s", bg)
	}
	io.WriteString(w, "\n")

	for _, fg := range fgs {
		// row header
		fmt.Fprintf(w, " %5s  ", fg)

		fmt.Fprintf(w, "\x1b[%s %s  ", fg, gopher)

		for _, bg := range bgs {
			fmt.Fprintf(w, "\x1b[%s\x1b[%s %s \x1b[0m ", fg, bg, gopher)
		}
		io.WriteString(w, "\n")
	}
}
