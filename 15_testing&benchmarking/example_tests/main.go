package main

import "fmt"

func main(){
	fmt.Println(MySum(5, 3))
	fmt.Println( MySum(5, 4))
	
}

//sum unlimited number of values
func MySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}