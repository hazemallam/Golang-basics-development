package main
import "fmt"
func main(){
	output := funcReturningString()
	fmt.Println(output)

	x := funcReturningFunction()
	fmt.Printf("%T\n", x)
	returning := x()
	fmt.Println(returning)
	fmt.Printf("%T\n", returning)

	//we can also run the funcReturningFunction() like this
	fmt.Println(funcReturningFunction()())
}

func funcReturningString() string{
	x := "Hello, world"
	return x
}
func funcReturningFunction() func() int{
	return func() int{
		return 555
	}
}