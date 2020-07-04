package main
import (
	"fmt"
	"sync"
)
func main(){
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)
	for i := 0; i<gs; i++ {
		go func(){
			mu.Lock()
			fmt.Println("first", counter)
			v := counter
			v++
			counter = v
			fmt.Println("second", counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end", counter)
	
}