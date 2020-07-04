package main

import "fmt"

func doSomething(x int) int{
	return 5 * x
}
func main(){
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()

	fmt.Println(<-ch)
}