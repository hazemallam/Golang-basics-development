package main
import (
	"fmt"
)
type customErr struct{
	lat string
	long string
	err error
}
func (c customErr) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", c.lat, c.long, c.err)
} 
func foo(e error){
	fmt.Println(e)
}
func main(){
	c1 := customErr{
		lat : "526.25N",
		long : "52.548W",
		err : fmt.Errorf("an error in custom "),
	}
	foo(c1)
}
