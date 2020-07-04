package main

import "fmt"
const(
	_ = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main(){
	v := 2
	fmt.Printf("%d\t\t%b\n", v, v)
	x := v << 1
	fmt.Printf("%d\t\t%b\n", x, x)

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}