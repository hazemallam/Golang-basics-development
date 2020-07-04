package main
import (
	"fmt"
	"time"
	"context"
)
func main(){
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// canceling goroutines with TimeOut() function
	// ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()

	select{
	case <-time.After(1 * time.Second):
		fmt.Println("over slept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}