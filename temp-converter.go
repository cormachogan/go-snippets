// Converts a range of numeric arguments to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

)

type Celsius float64
type Fahrenheit float64


// Main
func main() {
	var err error
	clargs := os.Args[1:]
	floats := make([]float64, len(clargs))

	for j := 0; j < len(clargs); j++ {
		if floats[j], err = strconv.ParseFloat(clargs[j], 64); err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := Fahrenheit(floats[j])
		c := Celsius(floats[j])
		fmt.Printf("%s = %s, %s = %s\n",
			f, FToC(f), c, CToF(c))	// displaying as strings possible due to String methods below
	}
}

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }


// Method to enable string displays of Celsius
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }

// Method to enable string displays of Fahrenheit
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
