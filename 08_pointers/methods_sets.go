/*
	Receivers     Values
	-----------------------
	(t T)		T and *T
	(r *T)		*T 

*/



package main
import (
	"fmt"
)
type Square struct{
	width float64
}
type Circle struct{
	r float64
}
func (s Square) areaa() float64 {
	return (s.width * s.width)
}
func (c *Circle) areaa() float64 {
	return (3.14 * c.r * c.r)
}
type Shape interface{
	areaa() float64
}
func info(ss Shape){
	/* this do mot work with methods sets
	// switch ss.(type){
	// case Square:
	// 	fmt.Println("the area of Square is : ", ss.(Square).areaa())
	// case Circle:
	// 	fmt.Println("the area of Circle is : ", ss.(Circle).areaa())

	// }
	*/
	fmt.Println(ss.areaa())
}

func main(){
	sq := Square{
		width : 4, 
	}
	cir := Circle{
		r : 2,
	}
	// fmt.Printf("%T\n", &cir)
	info(&sq)
	info(&cir)
}