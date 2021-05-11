//Fun with pointers
package main

import (
	"fmt"
)

func main() {
	var y, z int

	x := 1

	p := &x		// p is assigned to address of x

	fmt.Println("*p = ", *p) // *p displays the contents of the address of x, i.e the value of variable x

	*p = 2		// *p updates the contents of the address of x, i.e the value of variable x

	fmt.Println("x = ", x)

	*p++ 		// *p updates the contents of the address of x, i.e the value of variable x

	fmt.Println("x = ", x)

	fmt.Println("*p = ", *p)

	// Some simple pointer equivalence - 2 pointers are equal if the point to same var, or both are nil

	fmt.Println(&y == &y, &y == &z, &y == nil) // Should return true, false, false
	fmt.Println(&y, &z)

	// passing a pointer to a function - function can update variable that was indirectly passed

	v := 1
	incr(&v)		// v is now 2
	fmt.Println("incr(&v) = ", incr(&v))	// v is now 3

	// new Function creates an unnamed variable

	r := new(int)
	fmt.Println("r = ", *r)		// r is now 0
	*r = 2
	fmt.Println("r = ", *r)		// r is now 2

}

func incr(p2 *int) int {
	*p2++		// *p2 updates the contents of the address of v, i.e. the value of variable v
	return *p2
}
