package main
import (
	"log"
	"math"
	"fmt"
)
type sqrtError struct{
	lat string
	long string
	err error
}
func (s sqrtError) Error() string {
	return fmt.Sprintf("math error : %v %v %v", s.lat, s.long, s.err)
}
func main(){
	v, err := srooot(-10)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(v)
}

func srooot(f float64) (float64, error) {
	if f < 0{
		err := fmt.Errorf(" square root of negative error %v", f)
		return 0, sqrtError{"50,1235N", "99.254W", err}
	}
	return math.Sqrt(f), nil
}