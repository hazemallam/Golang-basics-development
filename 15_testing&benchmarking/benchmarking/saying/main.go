package main

import "fmt"

func main(){

	fmt.Println(Greet("James"))
}

func Greet(s string) string {
	return fmt.Sprint("Welcome dear ", s)
}
