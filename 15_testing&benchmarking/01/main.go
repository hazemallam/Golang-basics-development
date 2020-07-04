package main

import "fmt"

func main(){
	fmt.Println("5 + 3 =", MySum(5, 3))
	fmt.Println("5 + 4 =", MySum(5, 4))
	fmt.Println("5 + 5 =", MySum(5, 5))

}

//sum unlimited number of values
func MySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}