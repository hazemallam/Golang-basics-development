package main
import(
	"fmt"
	"log"
	"math"
)
type myError struct{
	lat string
	long string
	err error

}
func (m myError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", m.lat, m.long, m.err)

}
func main(){
	_, err := sroot(-10)
	if err != nil{
		log.Println(err)
	}
}
func sroot(f float64) (float64, error){
	if f < 0 {
		err := fmt.Errorf("norgate math error, square root of negative error %v", f)
		return 0, myError{"50,1235N", "99.254W", err}
	}
	return math.Sqrt(f), nil
}