package main
import "fmt"
func main(){
	f := func(){
		fmt.Println("My first func expression")
	}

	funcWithReturn := func(s string) string{
		x := fmt.Sprint("hello ", s)
		return x
	}

	f()
	z := funcWithReturn("James Bond")
	fmt.Println(z)
	
}