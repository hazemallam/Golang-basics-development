package main

import "fmt"

func main() {
	switch { //this equivelent to switch true {}
	case false:
		fmt.Println("this should not print")
	case true:
		fmt.Println("this should print 1")
		fallthrough
	case (2 == 3):
		fmt.Println("this should not print")
	case (2 == 2):
		fmt.Println("this should print 2")
	default:
		fmt.Println("this is default")
	}

	n := "Bond"
	switch n {
	case "hello":
		fmt.Println("hello")
	case "Bond":
		fmt.Println("Bond")
	default:
		fmt.Println("default")
	}
	//multiple cases
	switch n {
	case "hello", "all":
		fmt.Println("hello, all")
	case "name", "Bond", "try":
		fmt.Println("name, Bond, try")
	default:
		fmt.Println("default")
	}
}
