package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

//creating my own variable
type myOwnType int

//creating underlying type

var l myOwnType
var m int

func main() {

	s := fmt.Sprintf("%v\n%v\n%v", x, y, z)
	fmt.Println(s)

	//printing the value of the variable x
	fmt.Printf("%v\n", l)
	//printing the type of the variable x
	fmt.Printf("%T\n", l)
	//assinging value to the variable x
	l = 42
	fmt.Printf("%d\n", l)
	//converting (casting) the variable l to int
	m = int(l)
	fmt.Printf("%d\n", m)
	fmt.Printf("%T\n", m)

}
