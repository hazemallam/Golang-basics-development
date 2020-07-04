//TAKING VALUES FROM MANY CHANNELS AND PUTTING THOSE VALUES ONTO ONE CHANNEL
package main 
import (
	"sync"
	"fmt"
)
func main(){
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go sendd(even, odd)

	go receivee(even, odd, fanin)

	for v := range fanin{
		fmt.Println(v)
	}

	fmt.Println("exited")
}
func sendd(e, o chan<- int){
	for i := 0; i<10; i++{
		if i % 2 == 0{
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}
func receivee(e, o <-chan int, fanin chan<- int){
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		for v := range e{
			fanin <- v
		}
		wg.Done()
	}()
	go func(){
		for v := range o{
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}