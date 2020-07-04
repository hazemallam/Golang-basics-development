package main
import "fmt"
type Person struct{
	first string
	last string 
	age int
}
func (p Person) speak(){
	fmt.Println("I am", p.first, p.last, "and my age is", p.age)
}
func main(){
	p1 :=  Person{
		first : "James",
		last : "Bond",
		age : 24,
	}
	p1.speak()
	
}