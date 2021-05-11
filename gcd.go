// Snippet of code to return the greatest common divisor of 2 numbers
// The largest positive integer that divides into each of the integers provided

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

	fmt.Println("X (first arg) is ", nums[0])
	fmt.Println("Y (second arg) is ", nums[1])
	z := gcd(nums[0],nums[1])
	fmt.Println("Greatest Common Divisor is ", z)
}

func gcd(x, y int) int {
	for y != 0  {
		x,y = y, x%y
	}
	return x
}
