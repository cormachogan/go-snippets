package main

import (
	"fmt"
)

func main() {

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

	// Slice beyond len is ok, but not cap

	endless_summer := summer[:5]
	fmt.Printf("The new summer months are now %v\n\n", endless_summer)

	// See if two slices are identical

	if result := is_equal(summer, endless_summer); result != true {
		fmt.Printf("Summer is not equal to endless summer - %t\n\n", result)
	} else {
		fmt.Printf("Summer is equal to endless summer - %t\n\n", result)
	}

	// Slice beyond len is ok, but not cap

	new_summer := summer[:3]
	fmt.Printf("The new summer months are now %v\n\n", new_summer)

	// See if two slices are identical

	if result := is_equal(summer, new_summer); result != true {
		fmt.Printf("Summer is not equal to new summer - %t\n\n", result)
	} else {
		fmt.Printf("Summer is equal to new summer - %t\n\n", result)
	}

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
