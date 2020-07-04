package main
import "fmt"
func main(){

	x := []int {1,2,3,4,5,6,7,8,9}
	output := sum(x...)
	fmt.Println(output)

	  greetings("hello, ", "James", "Miss moneypenny")
	
}
// x ...int //means 0 or unlimited arguments
func sum(x ...int) int {
	total := 0
	for _, v := range x{
		total += v
	}
	return total
}

func greetings(prefix string, who ...string) {
	for _, v := range who {
		fmt.Println(prefix, v)
	}	
}