package main
import "fmt"

func main(){
	f1()
	defer f2()
	f3()

}
func f1(){
	fmt.Println("function 1")
}

func f2(){
	fmt.Println("function 2")
}
func f3(){
	fmt.Println("function 3")
}