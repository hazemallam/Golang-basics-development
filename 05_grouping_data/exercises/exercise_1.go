package main

import "fmt"
func main(){
	var arr [5]int
	x := [5]int{5,6,7,8,9}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", arr)

	for i, v := range x{
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}