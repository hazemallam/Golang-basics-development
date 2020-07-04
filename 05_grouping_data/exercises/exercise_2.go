package main
import "fmt"
func main(){
	slice := []int{41,42,43,44,45,46,47,48,49,50, 51}
	fmt.Println(slice[1:6])
	fmt.Println(slice[6:])
	fmt.Println(slice[3:8])
	fmt.Println(slice[2:7])
}