package main
import "fmt"
type Personn struct{
	first string
	last string
}
type Agent struct{
	Personn
	ltk bool
}
func (s Agent) speak() {
	fmt.Println("I am ", s.first, s.last, "the Agent speak")
}
func (p Personn) speak(){
	fmt.Println("I am ", p.first, p.last, "the person speak")
}
//polymorphism :  A value can be of more than one types
//any type that has speak() method it is also of type human
type human interface{
	speak()
}
func ba(h human){
	//assortion
	switch h.(type){
	case Personn:
		fmt.Println("I was bassed into baaaa", h.(Personn).first)
	case Agent:
		fmt.Println("I was bassed into baaaa", h.(Agent).first)
	}
	fmt.Println("I was bassed into bar", h)
}

type hotdog int
func main(){
	a1 := Agent{
		Personn : Personn{
			"James",
			"Bond",
		},
		ltk : true,
	}
	a2 := Agent{
		Personn : Personn{
			"Miss",
			"Moneypenny",
		},
		ltk : true,
	}

	p := Personn{
		first : "Dr",
		last : "Yes",
	}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(p)

	a1.speak()
	a2.speak()
	p.speak()

	ba(a1)
	ba(a2)
	ba(p)

	//conversion
	var xi hotdog = 40
	fmt.Println(xi)
	fmt.Printf("%T\n", xi)
	var y int
	y = int(xi)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
