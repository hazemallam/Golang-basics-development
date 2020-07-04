package main

import "fmt"

func main(){
	//print unicode
	unicode := 65 
	for i := 1 ; i<=26; i++{
		fmt.Printf("%v\n", i)
		for j := 0 ; j<3; j++{
			fmt.Printf("\t%#U\n", unicode)
		}
		unicode++
	}
}