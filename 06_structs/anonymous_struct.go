package main
import "fmt"
func main(){
	//Anonymous struct
	p1 := struct{
		first string
		last string
		age int

	}{
		first : "James",
		last :"Bond",
		age : 30,
	}
	fmt.Println(p1)
}