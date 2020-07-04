package main
import "fmt"
func main(){
	
	func(){
		fmt.Println("hello, from anonymous function")
	}()

	val := func(x int) int {
		return x * x
	}
	x := val(5)
	fmt.Println("this from val func", x)
}