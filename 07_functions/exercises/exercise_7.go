package main
import "fmt"
func main(){
	
	val := ret()
	val2 := val(5)
	fmt.Println(val2)

}

func ret() func(x int)int{
	return func(x int) int{
		return x * x
	}
}