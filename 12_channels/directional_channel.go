package main
import "fmt"
func main(){
	c := make(chan int) //receive and send channel(directional)
	fmt.Printf("%T\n", c)
	s := make(chan <- int, 2) // send only channel
	fmt.Printf("%T\n", s)

	r := make(<- chan int, 2) //receive only channel
	fmt.Printf("%T\n", r)

	//specific to general doesn't assign
	// c = s
	// c = r
	// fmt.Println((chan int)(s))
	// fmt.Println((chan int)(r))

	//general to specific assign
	// s = c
	// r = c
	// fmt.Println((<-chan int)(c))
	// fmt.Println((chan <- int)(c))

	//specific to specific doesn't assign
	// s = r
	// r = s

}