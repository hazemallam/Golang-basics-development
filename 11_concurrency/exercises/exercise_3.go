package main
import (
	"fmt"
	"sync"
	"runtime"
)
func main(){
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i<gs; i++ {
		go func(){
			fmt.Println("first", counter)
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("second", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end", counter)
	
}