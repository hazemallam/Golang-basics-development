package main

import (
	"io/ioutil"
	"os"
	"fmt"
)
func main(){
	f, err := os.Open("try.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}