package main
import "fmt"

func main(){
	f := foo()
	b1, b2 := bar()
	fmt.Println(f)
	fmt.Println(b1, b2)

}
func foo () int{
	x:= 44
	return x
}
func bar() (int, string){
	x := 55
	s := "Hello"
	return x, s 
}