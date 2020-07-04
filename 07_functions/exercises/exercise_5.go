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
func (c Circle) areaa() float64 {
	return (3.14 * c.r * c.r)
}
type Shape interface{
	areaa() float64
}
func info(ss Shape){
	switch ss.(type){
	case Square:
		fmt.Println("the area of Square is : ", ss.(Square).areaa())
	case Circle:
		fmt.Println("the area of Circle is : ", ss.(Circle).areaa())

	}
}

func main(){
	sq := Square{
		width : 4, 
	}
	cir := Circle{
		r : 2,
	}
	info(sq)
	info(cir)
}