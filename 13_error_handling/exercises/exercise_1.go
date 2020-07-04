package main
import(
	"fmt"
	"encoding/json"
	"log"
)
type person struct{
	First string
	Last string
	Sayings []string
}
func main(){
	p1 := person{
		First : "James",
		Last : "Bond",
		Sayings : []string{
			"hello",
			"how are you",
			"call me",
		},
	}
	bs, err := json.Marshal(p1)
	if err != nil{
		log.Fatalln("JSON DID NOT MARSHAL ", err)
	}
	fmt.Println(string(bs))
}