package main
import "fmt"
func main(){
	x := sum(1,2,3,4,5,6,7,8,9)
	fmt.Println(x)
}

func sum(x ...int) int{
	total := 0
	for _, v := range x{
		total += v
	}
	return total
}