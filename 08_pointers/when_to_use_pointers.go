package main
import "fmt"
func main(){
	x := 0
	foo(x)
	fmt.Println(x)

}
func foo(x int){
	fmt.Println(x)
	x = 42
	fmt.Println(x)
}

