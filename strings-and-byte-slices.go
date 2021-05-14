// A look at the string and byte slice functions

package main

import (
	"fmt"
	"os"
	"strings"
	"bytes"

)

// Main
func main() {

	var has_ell, has_or bool

	s := "hello"
	b := []byte("world")

// String test - see if it contains a sub-string

	if has_ell = strings.Contains(s, "ell"); has_ell != true {
		fmt.Fprintf(os.Stderr, "String %s does not contain ell.\n", s)
	} else {
		fmt.Fprintf(os.Stdout, "String %s contains ell.\n", s)
	}

	if has_or = strings.Contains(s, "or"); has_or != true {
		fmt.Fprintf(os.Stderr, "String %s does not contain or.\n", s)
	} else {
		fmt.Fprintf(os.Stdout, "String %s contains or.\n", s)
	}

// Byte Slice test - see if it contains a sub-byte-slice

// Convert comparisson string to byte slice - cannot use sub-string to test in byte slice

	bs_ell := []byte("ell")
	bs_or  := []byte("or")

	if has_ell = bytes.Contains(b, bs_ell); has_ell != true {
		fmt.Fprintf(os.Stderr, "Byte Slice %b ( displayed as string - %[1]s ) does not contain ell.\n", b)
	} else {
		fmt.Fprintf(os.Stdout, "Byte Slice %b ( displayed as string - %[1]s ) contains ell.\n", b)
	}

	if has_or = bytes.Contains(b, bs_or); has_or != true {
		fmt.Fprintf(os.Stderr, "Byte Slice %b ( displayed as string - %[1]s ) does not contain or.\n", b)
	} else {
		fmt.Fprintf(os.Stdout, "Byte Slice %b ( displayed as string - %[1]s ) contains or.\n", b)
	}
}
