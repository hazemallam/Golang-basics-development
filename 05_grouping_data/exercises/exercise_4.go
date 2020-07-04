package main
import "fmt"
func main(){
	x := []int {42,43,44,45,46,47,48,49,50,51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	y := make([]string, 3, 3)
	y = []string {"aa", "bb", "cc"}
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	for i := 0; i<len(y); i++ {
		fmt.Println(i, y[i])
	}
}