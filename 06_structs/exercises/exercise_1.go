package main
import "fmt"
type Person struct{
	first_name string
	last_name string
	flavors []string
}
func main(){
	p1 := Person{
		first_name : "James",
		last_name : "Bond",
		flavors : []string{"mango", "orange"},
	}
	p2 := Person{
		first_name : "Miss ",
		last_name : "Moneypenny",
		flavors : []string{"lemon"},

	}
	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range p1.flavors {
		fmt.Println(i, v)

	}
}