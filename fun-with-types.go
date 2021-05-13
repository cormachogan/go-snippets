// Examining some different data types and their behaviours

package main

import "fmt"
import "math"

func main() {
	fmt.Printf("INTEGERS\n")
	// unsigned int8 - range is 0 to 255
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)  // "255 0 1"

	// signed int8 - range is -128 to 127
	var i int8 = 127
	fmt.Println(i, i+1, i*i)  // "127 -128 1"
	fmt.Printf("\n")
  
  
	fmt.Printf("STRINGS\n")
	// printing an array in reverse order
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
	fmt.Printf("\n")
	for i := 0; i < len(medals); i++ {
		fmt.Println(medals[i])
	}
	fmt.Printf("\n")
  
// string parsing
	s := "hello, world"
	fmt.Println(len(s))	// 12
	fmt.Println(s[0], s[7])	// "104 119" ('h' and 'w')
	fmt.Println(s[0:5])	// New string hello
	fmt.Println(s[:5])	// New string hello
	fmt.Println(s[7:12])	// New string world - does not include byte 12
	fmt.Println(s[7:])	// New string world
	fmt.Println("goodbye" + s[5:])	// goodbye, world
	fmt.Printf("\n")

  
	fmt.Printf("OCTs And HEXs\n")
	// printing dec, oct and hex using the same operand 'o' and 'x'
	o := 0666
	fmt.Printf("%d, %[1]o, %#[1]o\n", o)
	x := int64(0xdeadbeef)
	fmt.Printf("%d, %[1]x, %#[1]x, %#[1]X\n", x)
	fmt.Printf("\n")
  
  
	fmt.Printf("FLOATs\n")
	for z :=0; z < 8; z++ {
		fmt.Printf("z = %d   e' = %8.3f\n", z, math.Exp(float64(z)))
	}
	fmt.Printf("\n")
  
  
	fmt.Printf("Unicode\n")
	fmt.Printf("\u4e16\u754c\n")
}
