package main
import "fmt"
func main(){
	for i := 10; i<=100; i++{
		fmt.Printf("%d\t%d\n", i, i%4)
	}
}