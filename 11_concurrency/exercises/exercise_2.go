package main

import "fmt"

type Person struct{
	Name string
}
func (p *Person) speak(){
	fmt.Println("my name is", p.Name)
}
type human interface{
	speak()
}
func saySomeThing(h human){
	// fmt.Println("saySomeThing", h.speak())
	fmt.Print("saySomeThing ")
	h.speak()
}
func main(){
	p1 := Person{
		Name : "James Bond",
	}
	saySomeThing(&p1)
}