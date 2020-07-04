package main
import(
	"fmt"
	"sort"
	
)
func main(){
	xi := []int{9,8,7,6,5,4,3,2,1}
	si := []string{"d", "c", "e"}
	sort.Ints(xi)
	fmt.Println(xi)
	sort.Strings(si)
	fmt.Println(si)
}