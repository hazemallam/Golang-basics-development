//struct is a data structure to composite variables with different types
package main
import "fmt"
type Person struct {
	first string
	last string
	age int

}
type SecretAgent struct{
	Person
	first string
	ltk bool
}
func main(){
	p1 := Person{
		first : "James",
		last : "Bond",
		age : 20,
	}
	p2 := Person{
		first : "Miss",
		last  : "Moneypenny",
		age : 30,
	}
	sa := SecretAgent{
		Person : Person{
			first : "first",
			last : "last",
			age : 40,
		},
		first : "collesion",
		ltk : true,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.last)
	fmt.Println(p2.age)

	fmt.Println(sa)
	fmt.Println(sa.Person.first)
	

}