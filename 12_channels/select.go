package main
import "fmt"
func main(){
	even := make(chan int)
	odd := make(chan int)
	quite := make(chan int)

	//send
	go send(even, odd, quite)

	//receive
	receive(even, odd, quite)
	fmt.Println("exited")
}

func send(e, o, q chan<- int){
	for i := 0; i<100; i++{
		if i % 2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// close(e)
	// close(o)
	close(q)
}

//receive 
func receive(e, o, q <-chan int){
	for {
		select {
		case v := <- e:
			fmt.Println("from the even channel", v)
		case v := <- o:
			fmt.Println("from the odd channel", v)
		case i, ok := <- q:
			if !ok{
				fmt.Println("from comma ok", i, ok)
				return
			} else {
			fmt.Println("from comma ok", i)
			}
		}
	}
}