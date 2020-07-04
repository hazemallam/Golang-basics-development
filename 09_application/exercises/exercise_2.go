package main
import (
	"fmt"
	"encoding/json"
)
type persong struct{
	First string
	Last string
	Age int
	Sayings []string
}
func main(){
	s := `[{"First":"James", "Last":"Bond", "Age":30, "Sayings":["how are you", "tell me your name", "please"]}, {"First":"Miss", "Last":"Moneypenny", "Age":20, "Sayings":["how are you", "tell me your name", "please"]}]`
	var people []persong
	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)

	}	
	fmt.Println(people)
	for i, per := range people{
		fmt.Println("PERSON NUMBER", i)
		fmt.Println("\t", per.First, per.Last, per.Age)
		for _, saying := range per.Sayings{
			fmt.Println("\t\t",saying)
		}
	}
}