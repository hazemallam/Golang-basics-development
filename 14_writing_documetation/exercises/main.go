package main

import (
	"fmt"
	"computer_science/Golang/work_place/writing_documentation/exercises/dog"
)
type germanshipered struct{
	name string
	age int
}
func main(){
	siko := germanshipered{
		name : "Siko", 
		age : dog.Years(5),
	}
	
	
	fmt.Println(siko)
}