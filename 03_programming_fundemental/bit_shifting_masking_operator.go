package main

import (
	"fmt"
	"strings"
	// "math/rand"
)

const(
	UPPER  = 1 // upper case
    LOWER  = 2 // lower case
    CAP    = 4 // capitalizes
    REV    = 8 // reverses
)

func main(){
	var x uint8 = 0xAC
	fmt.Printf("%b\n", x)
	x = x & 0xF0 // short assign x &= 0xF0
	fmt.Printf("%b\n", x)

	// & operator is useful to determine if a number is odd or even because the LSB for an odd number is 1 so 
	// if the the output is 1 the number is odd else it is an even number
	for i := 0; i<100; i++ {
		// num := rand.Int()
		if i & 1 == 1{
			fmt.Printf("%b\n", (i & 1))
			fmt.Printf("%d is odd\n", i)
		
		} else {
			fmt.Printf("%d is even\n", i)
		}
	}

	// | operator (bit masking or)
	var a uint8 = 0
	a |= 196
	fmt.Printf("%b\n", 196)
	fmt.Printf("%b\n", a)

	fmt.Println(procstr("HELLO PEOPLE!", LOWER|REV|CAP))

	// (^) XOR operator
	//determine if two integers have the differecne sign
	//Two integers a, b have the same signs when (a ^ b) ≥ 0 (or (a ^ b) < 0 for opposite sign)
	z, m := -12, 25
	fmt.Printf("z and m have same sign ? %t\n", (z ^ m) >= 0)

	// (^) Bitwise complement (NOT)
	var f byte = 0x0F
	fmt.Printf("%08b\n", f)
	fmt.Printf("%08b\n", ^f)

	//AND NOT operator
	//The next code snippet uses the AND NOT operator to clear 
	//the last four LSBs in variable a from 1010 1011 to 1010 0000.
	var andNot = 0xAB
	fmt.Printf("%08b\t%d\n", andNot, andNot)
	andNot &^= 0x0F
	fmt.Printf("%08b\t%d\n", andNot, andNot)
}
//Function will lower the cases for the string, reverse its order, and capitalize each word
func procstr(str string, conf byte) string {
    // reverse string
    rev := func(s string) string {
        runes := []rune(s)
        n := len(runes)
        for i := 0; i < n/2; i++ {
            runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
        }
        return string(runes)
    }
 
    // query config bits
    if (conf & UPPER) != 0 {
        str = strings.ToUpper(str)
    }
    if (conf & LOWER) != 0 {
        str = strings.ToLower(str)
    }
    if (conf & CAP) != 0 {
        str = strings.Title(str)
    }
    if (conf & REV) != 0 {
        str = rev(str)
    }
    return str
}