package main
import "fmt"
func main(){
	foo()
	bar("James")
	st := woo("Miss Moneypenny")
	fmt.Println(st)
	x, y := mouse("James", " Bond")
	fmt.Println(x)
	fmt.Println(y)
}

//function syntax
//func (r receiver) identifier(parameter) (return(s)) {code ...}

func foo(){
	fmt.Println("hello, from func foo")
}
//every thing in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("hello, ", s)
}

func woo(s string) string{
	return fmt.Sprint("Hello, from woo, ", s)
}

func mouse(fname, lname string) (string, bool){
	a := fmt.Sprint(fname, lname, ` says "hello"`)
	b := false
	return a, b
}
