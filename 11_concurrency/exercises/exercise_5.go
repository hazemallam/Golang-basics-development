package main

import (
	"runtime"
	"fmt"
	"sync/atomic"
	"sync"
)
func main(){
	var counter int64 = 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i<gs; i++{
		atomic.AddInt64(&counter, 1)
		fmt.Println("counter\t", atomic.LoadInt64(&counter))
		wg.Done()

	}
	wg.Wait()
	fmt.Println("end counter", counter)
}