//Closure : limiting the scoop of a variable to specified block of code
package main
import "fmt"

var y int
func main(){

	this()
	//narrowing the scoop of variable x
	{ //code block
		x := 0 
		x++
		fmt.Println(x)
		fmt.Println("end of scoop")
	}

	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	
}

func this(){
	fmt.Println("inside this")
	y++
	fmt.Println(y)
	fmt.Println("end of this")

}

//example
func incrementor() func() int {
	var x int
	return func () int {
		x++
		return x
	}
}