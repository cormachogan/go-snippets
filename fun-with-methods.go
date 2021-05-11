// Fun with types and methods

package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC	Celsius = -273.15
	FreezingC 	  Celsius = 0
	Boiling 	    Celsius = 100
)

func main() {
	var c Celsius
	var f Fahrenheit

// Do a few assignment tests  
	fmt.Println (c == 0)		// true
	fmt.Println (f >= 0)		// true
//	fmt.Println (c == f)		// compile error, mismatch
	fmt.Println (c == Celsius(f))	// true

  // If c was not previously declared, we would use ':=' instead of '=' here  
	c = FtoC(212.0)
	fmt.Println(c.String())
	fmt.Printf("print2 %v\n", c)
	fmt.Printf("print3 %s\n", c)
	fmt.Println(c)
	fmt.Printf("print5 %g\n", c)
	fmt.Println(float64(c))
}

// Convert Fahrenheit to Celsuis, receive type F, return type C

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9 )
}

// Celsuis parameter before the function name associates a method name String with the Celsius type

func (c Celsius) String() string {
	return fmt.Sprintf("method print %gC", c)
}
