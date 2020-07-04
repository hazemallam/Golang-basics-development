package main

import(
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func main(){

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
func foo(){
	fmt.Println("foo")
	wg.Done()
}
func bar(){
	wg.Done()
	fmt.Println("bar")
}