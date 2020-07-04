//callback function : passing a function as an argument to another function
package main
import "fmt"
func main(){
	xi := []int {1,2,3,4,5,6,7,8,9}
	returning := sum(xi...)
	fmt.Println("summing all numbers = ", returning)
	xv := even(sum, xi...)
	fmt.Println("summing even numbers = ", xv)
	zv := odd(sum, xi...)
	fmt.Println("summing odd numbers = ", zv)
}
//summing numbers
func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
//summing even numbers
func even(f func(x ...int)int, vi ...int) int{
	var yi []int
	for _, v := range vi{
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

//summing odd numbers
func odd(f func(x ...int)int, in ...int) int {
	var oddNumbers []int
	for _, v := range in {
		if v&1 == 1 { // odd number
			oddNumbers = append(oddNumbers, v)
		}
	}
	return f(oddNumbers...)
}