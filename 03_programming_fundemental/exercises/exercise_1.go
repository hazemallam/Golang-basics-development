package main

import "fmt"

const(
	value = 40
	value2 string = "hello world"
	
)
const (

	a = iota + 2017
	b = iota + 2017
	c = iota + 2017
	d = iota + 2017
)

func main(){
	const x = 40
	value3  := `this a raw string literal "you see"`

	fmt.Printf("%d\n%b\n%#x\n%s\n", value, value, value, value2)
	fmt.Printf("%d\n%b\n%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\n%b\n%#x\n", y, y, y)
	fmt.Printf("%v\n", value3)

	fmt.Println(a, b, c, d)
}