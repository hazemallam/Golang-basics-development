package main

import "fmt"

var boolean bool

func main(){
	a := 7
	b := 24
	fmt.Printf("%t\n", boolean)
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
}