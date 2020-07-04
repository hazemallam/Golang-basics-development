package main
import (
	"io"
	"strings"
	"os"
	"fmt"
)
func main(){
	f, err := os.Create("try.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := strings.NewReader("James Bond")
	io.Copy(f, r)
}