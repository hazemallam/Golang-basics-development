package main
import "fmt"
func main(){
	c := make(<-chan int)
	c = add()
	pull(c)
	fmt.Println("exited")
}

func add() <-chan int{
	c := make(chan int)
	go func (){
		for i := 0; i<100; i++{
			c <- i
		}
		close(c)
	}()
	return c
}

func pull(c <-chan int){
	for v := range c{
		fmt.Println(v)
	}
}