package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID		int
	Name		string
	Address		string
	DoB		time.Time
	Position	string
	Salary		int
	ManagerID	int
}

type Point struct {
	X, Y int
}

type Circle struct {
	//Center Point
	Point		// fields with type and no name are anonymous
	Radius int
}

type Wheel struct {
	//Circle Circle
	Circle		// fields with type and no name are anonymous
	Spokes int
}


func main() {

	var dilbert Employee

	// fields in structs are accessed using dot notation
	dilbert.Salary += 5000
	dilbert.Position = "Code Monkey"

	fmt.Printf("Dilbert's Salary is %v\n", dilbert.Salary)
	fmt.Printf("Dilbert's Position is %s\n", dilbert.Position)

	// field can also be accessed via its address / pointer
	position := &dilbert.Position
	*position = "Senior " + *position

	fmt.Printf("Dilbert's Position is now %s\n", *position)
	fmt.Printf("Dilbert's Position is now %s\n", dilbert.Position)

	// dot notation also works with a pointer to a struct
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"

	fmt.Printf("Dilbert's Position is now %s\n", employeeOfTheMonth.Position)
	fmt.Printf("Dilbert's Position is now %s\n", dilbert.Position)

	// Struct Comparisson is possible

	fmt.Println("Comparisson check:", dilbert == *employeeOfTheMonth)

	// larger struct types are usually passed to / returned from functions indirectly using a pointer
	dilbert.Salary += Bonus(&dilbert, 10)
	fmt.Printf("Dilbert's Salary is now %v\n", dilbert.Salary)

	// Pointer also needed if we are modifying struct value since function only recieves copy of arg, not actual arg
	Raise(&dilbert)
	fmt.Printf("Dilbert's Salary is now %v\n", dilbert.Salary)

	// Embedded Structs and Anonymous Fields
	var w Wheel

	//w.Circle.Center.X = 8 //assignment without anonymous fields
	w.X = 8			//assignment with anonymous fields
	//w.Circle.Center.Y = 8	//assignment without anonymous fields
	w.Y = 8			//assignment with anonymous fields
	//w.Circle.Radius = 5	//assignment without anonymous fields
	w.Radius = 5		//assignment with anonymous fields
	w.Spokes = 20

	// We can refer to the name at the leaves of the implicit tree without the intervening names
	// The fields do have names however - the names of the types
	fmt.Printf("w.Circle.Point.X = %d\n", w.Circle.Point.X)
	fmt.Printf("w.Circle.Point.X = %d\n", w.X)
	fmt.Printf("w.Circle.Point.Y = %d\n", w.Circle.Point.Y)
	fmt.Printf("w.Circle.Point.Y = %d\n", w.Y)
	fmt.Printf("w.Circle.Radius = %d\n", w.Circle.Radius)
	fmt.Printf("w.Circle.Radius = %d\n", w.Radius)

	// liternal syntax for creating populated struct

	w1 := Wheel{Circle{Point{8,8}, 5}, 20 }
	fmt.Println("Comparisson check (wheel1):", w == w1)
	// # adverb prints values in a format similar to GO syntax
	fmt.Printf("%#v\n", w1)

	w2 := Wheel{Circle{Point{9,9}, 5}, 20 }
	fmt.Println("Comparisson check (wheel2):", w == w2)
	// # adverb prints values in a format similar to GO syntax
	fmt.Printf("%#v\n", w2)
}


// Calaculate bonus based on salary and percent - return bonus
func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

// Award a raise
func Raise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}
