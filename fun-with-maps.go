//Maps - also known as hash tables

package main

import (
	"fmt"
	"sort"
	"reflect"
)

func main() {

	ages := map[string]int{
		"alice":	31,
		"charlie":	34,
	}

	var names []string

	var newer_ages = map[string]int{}

	fmt.Printf("Current age of Alice is %d\n", ages["alice"])
	ages["alice"] = 32
	fmt.Printf("New age of Alice is %d\n\n", ages["alice"])

	// new entries initialize to 0
	ages["bob"]++
	fmt.Printf("Age of Bob is %d\n", ages["bob"])

	// range based loop
	fmt.Printf("\nAll names and ages:\n")
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// range based loop - ordered based on name using sort. Strings on an array

	fmt.Printf("\nAll names and ages, sorted by name:\n")

	// only require keys, age is not needed
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)

	// only require names, set index to blank
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	// Check if element exists
	fmt.Printf("\nChecking existence of elements:\n")
	age, ok := ages["bob"]
	if !ok {
		fmt.Printf("Bob element status not found: got %t\n", ok)
	} else {
		fmt.Printf("Bob element status found: got %t. Age is %d.\n", ok, age)
	}

	// Check if element exists - shorthand
	if dudeage, dudeok := ages["dude"]; !ok {
		fmt.Printf("Dude element status not found: got %t\n", dudeok)
	} else {
		fmt.Printf("Dude element status found: got %t. Age is %d.\n", dudeok, dudeage)
	}

	// Let's try to make a copy of a map to do a comparisson - first the wrong way

	// Failed method to create a unique copy of the original map
	new_ages := ages

	// range based loop
	fmt.Printf("\nAll names and ages:\n")
	for name, age := range new_ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	fmt.Printf("\nComparing two maps:\n")

	if result := reflect.DeepEqual(ages, new_ages); result != true {
		fmt.Printf("ages is not equal to new_ages - %t\n\n", result)
	} else {
		fmt.Printf("ages is equal to new_ages - %t\n\n", result)
	}

	// Try to force a mismatch - this doesn't do it since it also updates the original map

	new_ages["jake"] = 55

	// range based loop
	fmt.Printf("\nAll names and ages in new map:\n")
	for name, age := range new_ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	fmt.Printf("\nAll names and ages in original map:\n")
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	fmt.Printf("\n\nNote the contents of the two maps is the same - need to make unique map\n\n")

	// Let's try to make another copy of a map to do a comparisson - this time the right way

	// Make a unique copy of the original map
	for k, v := range ages {
		newer_ages[k] = v
	}

	// Force a mismatch

	newer_ages["emer"] = 20

	// range based loop
	fmt.Printf("\nAll names and ages in new map:\n")
	for name, age := range newer_ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	fmt.Printf("\nAll names and ages in original map:\n")
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	fmt.Printf("\nComparing two maps:\n")

	if result := reflect.DeepEqual(ages, newer_ages); result != true {
		fmt.Printf("ages is not equal to new ages - %t\n\n", result)
	} else {
		fmt.Printf("ages is equal to new_ages - %t\n\n", result)
	}

	fmt.Printf("\nComparing two maps - home grown method:\n")

	if result := is_equal(ages, newer_ages); result != true {
		fmt.Printf("ages is not equal to new ages - %t\n\n", result)
	} else {
		fmt.Printf("ages is equal to new_ages - %t\n\n", result)
	}
}

// Maps are not comparable - you need to perform the comparison yourself

func is_equal(x,y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		yv, ok := y[k]
		if yv != xv || !ok {
			return false
		}
	}
	return true
}
