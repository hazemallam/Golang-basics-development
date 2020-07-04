package main
import "fmt"

type Personn struct{
	first_name string
	last_name string
	flavors []string
}
func main(){
	p1 := Personn{
		first_name : "James",
		last_name : "Bond",
		flavors : []string{"mango", "orange"},
	}
	p2 := Personn{
		first_name : "Miss ",
		last_name : "Moneypenny",
		flavors : []string{"lemon"},

	}
	m := map[string]Personn{
		p1.first_name: p1,
		p2.first_name: p2,
	}
	fmt.Println(m)
	for k, v := range m{
		fmt.Println(k, v)
	}
}