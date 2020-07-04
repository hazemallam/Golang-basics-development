package main
import (
	"fmt"
	"io"
	"os"
)
func main(){
	fmt.Println("hello, world")
	fmt.Fprintln(os.Stdout, "hello, world")
	io.WriteString(os.Stdout, "hello, world")
}