package main

import (
	"fmt"
)

func main() {
	c := genn()
	receivee(c)

	fmt.Println("about to exit")
}

func genn() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
func receivee(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
