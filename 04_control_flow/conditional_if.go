package main

import "fmt"

func main(){
	if true{
		fmt.Println(1)

	}
	if false{
		fmt.Println(2)
	}
	if !true{
		fmt.Println(3)
	}
	if !false {
		fmt.Println(4)
	}
	//initialization statement
	if x := 42; x == 2 {
		fmt.Println("in if condition")
	}
	fmt.Println("here")
	fmt.Println("there")
	//	fmt.Println(x) this can't work because x in if scoop

	//if else statement
	x := 40
	if x == 30 {
		fmt.Println("our value is equal 30")
	} else if x == 40 {
		fmt.Println("our value is equal 40")
	
	} else {
		fmt.Println("our value is not equal 40")
	}
}