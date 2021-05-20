// Dedup prints only one instance of each line; duplicates are removed.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool) // a set of strings
	fmt.Printf("Please insert some lines of text - terminate with a single . character\n")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		// If we have not see this line before, then its value/element will be false
		if !seen[line] && line != "." {
			seen[line] = true
			fmt.Printf("%s is a never before seen line\n", line)
		} else {
			if line == "." {
				os.Exit(0)
			} else {
				fmt.Printf("%s has been seen before\n", line)
			}
		}
	}
}
