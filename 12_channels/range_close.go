package main
import "fmt"
func main(){
	c := make(chan int)

	//send
	go fooo(c)

	//receive
	barr(c)
	fmt.Println("exited")
}
//send
func fooo(c chan<- int){
	for i := 0; i<100; i++{
		c <- i
	}
	close(c)
}

//receive
func barr(c <-chan int){
	for v := range c{
		fmt.Println(v)
	}
}