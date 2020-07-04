package main
import "fmt"
func main(){
	// x := type{values} //composite literal
	x := []int{4,5,6,7,8}
	fmt.Println(x) 
	for index, value := range x {
		fmt.Println(index, value)
	}
	//slicing
	fmt.Println(x[1:3])
	fmt.Println(x[1:])
	fmt.Println(x[:])
	y := []int{1,2,3,4,5,6}
	t := y[1:3:6]
	m := []int{10, 11}
	fmt.Println(t, len(t))
	//append elements
	z := append(y, 9, 8)
	fmt.Println(z)
	//append slice 
	n := append(y, m...)
	fmt.Println(n)

	//deleting elements (there is no delete function delete in GOlang)
	y = append(y[:2], y[3:]...)
	fmt.Println(y)

	//using keyword make to make an underling array or slice 
	// variable := make([]T, length, capacity)

	hh := make([]int, 10, 12)
	fmt.Println(hh)
	fmt.Println(len(hh))
	fmt.Println(cap(hh))
	
	hh = append(hh, 1)
	fmt.Println(hh)
	fmt.Println(len(hh))
	fmt.Println(cap(hh))

	hh = append(hh, 2)
	fmt.Println(hh)
	fmt.Println(len(hh))
	fmt.Println(cap(hh))

	hh = append(hh, 3)
	fmt.Println(hh)
	fmt.Println(len(hh))
	fmt.Println(cap(hh))

	//multi dimentional slice

	_1 := []string {"hello", "world"}
	_2 := []string {"how", "are"}

	xb := [][]string {_1, _2}
	fmt.Println(xb)

}

//a slice allow us to group values of the same type together