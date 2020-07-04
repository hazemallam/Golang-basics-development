package main
import "fmt"
func main(){
	x := [][]string {{"James", "Bond", "Shaken, not stirred"},{"miss", "MoneyPenny", "Hello, James"}}
	for i, record := range x{
		fmt.Println("record : ", i)
		for index, value := range record {
			fmt.Printf("\t index position : %v\t value : %v\n", index , value)
		}
	}

	
}