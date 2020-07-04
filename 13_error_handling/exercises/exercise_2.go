package main
import(
	"fmt"
	"encoding/json"
	"log"
)
type personn struct{
	First string
	Last string
	Sayings []string
}
func main(){
	p1 := personn{
		First : "James",
		Last : "Bond",
		Sayings : []string{
			"hello",
			"how are you",
			"call me",
		},
	}
	bs, err := toJSON(p1)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
	
}
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil{
		erro := fmt.Errorf("JSON DID NOT MARSHAL %v", err)
		return []byte{}, erro
	}
	return bs, nil

}