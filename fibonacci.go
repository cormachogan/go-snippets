// Calculate the Fibonacci number (add all numbers up to requested number)

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var err error
	clargs := os.Args[1:]
	nums := make([]int, len(clargs))

// Without Error Handling

	for i :=0; i < len(clargs); i++ {
		nums[i], err =  strconv.Atoi(clargs[i])
	}

// With Error Handling

	for i :=0; i < len(clargs); i++ {
		if nums[i], err =  strconv.Atoi(clargs[i]); err != nil {
			panic(err)
		}
	}

	fmt.Println("Fibonacci number to calculate (first arg) is ", nums[0])
	z := fibonacci(nums[0])
	fmt.Println("Fibonacci number calculated is", z)
}

func fibonacci(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
