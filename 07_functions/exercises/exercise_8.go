package main
import "fmt"
func main(){

	value := []int {1,2,3,4,5,6,7,8,9}
	output := odd(gam3, value...)
	fmt.Println("the output is : ", output)

}
func odd(f func(va ...int) int ,x ...int) int{
	var oddNumber []int
	for _, v := range x{
		if v & 1 == 1{
			oddNumber = append(oddNumber, v)
		}
	}
	return f(oddNumber...)
}
func gam3(x ...int) int{
	toal := 0
	for _, v := range x{
		toal += v
	}
	return toal
}