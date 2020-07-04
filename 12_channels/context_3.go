package main
import (
	"fmt"
	"context"
	"time"
	"runtime"
)
func main(){
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Error check 1 : ", ctx.Err())
	fmt.Println("number of Goroutiens 1 : ", runtime.NumGoroutine())

	go func (){
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("Error check 2 : ", ctx.Err())
	fmt.Println("number of Goroutiens 2 : ", runtime.NumGoroutine())
	
	fmt.Println("about to cancel")

	cancel()
	fmt.Println("canceled context")
	time.Sleep(time.Second * 2)
	fmt.Println("Error check 3 : ", ctx.Err())
	fmt.Println("number of Goroutiens 3 : ", runtime.NumGoroutine())
}
