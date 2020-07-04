package main
import "fmt"
func main(){
	//defer function excuted when the containing function finshed excuting here main() function
	defer foo()
	bar()
	a()
	b()
	i := c()
	fmt.Println(i)

}
func foo(){
	fmt.Println("foo")
}
func bar(){
	fmt.Println("bar")
}
//A deferred function's arguments are evaluated when the defer statement is evaluated.
func a(){
	i := 0
	defer fmt.Println(i)
	i++
	return
	
}

//Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func b(){
	for i:=0; i<4; i++{
		defer fmt.Println(i)
	}
}
//Deferred functions may read and assign to the returning function's named return values.
func c() (i int){
	defer func(){
		i++
	}()
	return 1
}