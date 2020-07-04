package main
import "fmt"
type Person struct{
	first string
	last string
	age int
}
func changeMe(p *Person){
	p.first = "firstName" // or (*p).first = "firstName"
	p.last = "lastName"	// or (*p).last = "lastName"
	p.age = 20			// or (*p).age = 20
}
func main(){
	p := Person{
		first : "James",
		last : "Bond",
		age : 30,
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}