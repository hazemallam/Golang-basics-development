package main
import "fmt"
func main(){
	x := 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	fooo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)


}
func fooo(y *int){
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
	
}