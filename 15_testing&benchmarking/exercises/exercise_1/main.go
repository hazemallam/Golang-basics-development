package main

import (
	
	"fmt"
	"computer_science/Golang/work_place/writing_documentation/exercises/word"
	
)
func main(){
	s := "hello how are you are name is how tell me your name"
	
	fmt.Println(word.Count(s))
	for k, v := range word.UseCount(s) {
		fmt.Println(v, k)
	}
	
}