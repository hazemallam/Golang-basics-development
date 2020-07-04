package main
import "fmt"
func main(){
	normalFunction()
	
	//Anonymous function : A function that has no name
	func(){
		fmt.Println("hello, from Anonymous function without arguments")
	}()

	func (x int){
		fmt.Println("Hello, from Anonymous function that has one parameter", x)
	}(42)
}
func normalFunction(){
	fmt.Println("hello, from normal function")
}