package main

import "fmt"

//first way untyped constants
const a = 40
const b = 40.402
const c = "hello world"

//second way typed constants
const e int = 40
const d float64 = 40.25
const f string = "hello"

//another way to define constants
const(
	m = 40
	n = 52.5
	p = "hi"
)
const (
		sat = iota // used for incremt variable after it 
		sun
		mon
		tue
		wed
		thu
		fri
)
func main(){
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(e)
	fmt.Println(d)
	fmt.Println(f)

	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(p)

	fmt.Println(sat)
	fmt.Println(sun)
	fmt.Println(mon)
	fmt.Println(tue)
	fmt.Println(wed)
	fmt.Println(thu)
	fmt.Println(fri)
	
}