package main
import "fmt"
type Person struct{
	first string
	last string
}
type SecretAgent struct{
	Person
	ltk bool
}
// this function is receving any abject of time SecretAgent which means it is a state in each sturct of type 
//SecretAgent
func (s SecretAgent) speak(){
	fmt.Println("I am ", s.first, s.last)
}
func main(){
	sa1 := SecretAgent{
		Person : Person{
			"James",
			"Bond",
		},
		ltk : true,
	}
	sa2 := SecretAgent{
		Person : Person{
			"Miss",
			"Moneypenny",
		},
		ltk : true,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()
}