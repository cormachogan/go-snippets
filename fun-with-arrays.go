package main

import (
	"fmt"
)

func main() {

	var a [3]int

	// new array elements are initially set to 0
	for i, v := range a {
		fmt.Printf("%d, %d\n", i, v)
	}

	// add some data to the arrays
	a = [3]int{1, 2, 3}
	b := [3]string{"April", "May", "June"}

	for i, v := range a {
		fmt.Printf("%d, %d\n", i, v)
	}

	// display the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// display the indices only
	for i, _ := range a {
		fmt.Printf("%d\n", i)
	}

	for i, v := range b {
		fmt.Printf("%d, %s\n", i, v)
	}
}
