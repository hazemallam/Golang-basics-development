package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gennn(q)

	receiveee(c, q)

	fmt.Println("about to exit")
}
func receiveee(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}

	}

}

func gennn(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}
