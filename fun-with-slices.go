package main

import (
	"fmt"
)

func main() {

	var runes []rune

	months := [...]string{
		1: "January",
		2: "February",
		3: "MArch",
		4: "April",
		5: "May",
		6: "June",
		7: "July",
		8: "August",
		9: "September",
		10: "October",
		11: "November",
		12: "December"}

	// create some slices from the array
	
	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Printf("The months in Q2 are %v\n", Q2)
	fmt.Printf("The summer months are %v\n", summer)

	//Find the common month(s) between Q2 and summer

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both Q2 and summer\n\n", s)
			}
		}
	}

	// Slice beyond len is ok, but not cap - slice is grown

	endless_summer := summer[:5]
	fmt.Printf("The new summer months are now %v\n\n", endless_summer)

	// See if two slices are identical

	if result := is_equal(summer, endless_summer); result != true {
		fmt.Printf("Summer is not equal to endless summer - %t\n\n", result)
	} else {
		fmt.Printf("Summer is equal to endless summer - %t\n\n", result)
	}

	// Slices do not support comparisons

	new_summer := summer[:3]
	fmt.Printf("The new summer months are now %v\n\n", new_summer)

	// Special code to check if two slices are identical

	if result := is_equal(summer, new_summer); result != true {
		fmt.Printf("Summer is not equal to new summer - %t\n\n", result)
	} else {
		fmt.Printf("Summer is equal to new summer - %t\n\n", result)
	}

	// The append function - print out the 11 runes in the string literal
	// Append call could cause a reallocation, and assignment may not be reflected in new slice
	// Thus common to assign result to the same slice variable, i.e. runes

	for _, r := range "Hello,World" {
		runes = append(runes, r)
	}

	fmt.Printf("Content as runes (Unicode code point) is %c\n", runes)
	fmt.Printf("Content as integers is %d\n", runes)
	fmt.Printf("Content as octals is %o\n", runes)
	fmt.Printf("Content as hexadecimal is %X\n", runes)
	fmt.Printf("Content as binary is %b\n", runes)
	fmt.Printf("Content as quotes runes is %q\n", runes)

	fmt.Printf("\nAppend to a slice\n")
	runes = append(runes, 'X', 'X', 'X')
	fmt.Printf("Content as quotes runes is %q\n", runes)

	fmt.Printf("\nAppend the same slice to itself\n")
	runes = append(runes, runes...)
	fmt.Printf("Content as quotes runes is %q\n", runes)

	fmt.Printf("\nRemove element from a slice\n")
	mysl := []int{5, 6, 7, 8, 9}
	fmt.Println(mysl)
	fmt.Println(remove(mysl,2))
}

// Slices are not comparable - you need to perform the comparison yourself

func is_equal(x,y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// Remove element from middle of slice, preserving the order

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
