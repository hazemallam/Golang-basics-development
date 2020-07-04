package main
import (
	"fmt"
	"encoding/json"
)
type personn struct{
	First string `json:First` //TAGS CAN BE REMOVED BUT FIELDS NAMES MUST BE STARTED WITH CAPITAL LETTER
	Last string `json:"Last"`
	Age int `json:"Age"`
}
func main(){

	s := `[{"First":"James","Last":"Bond","Age":40},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []personn
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}	
	fmt.Println("all of the data", people)
	for i, v := range people{
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}