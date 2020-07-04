package main
import "fmt"
func main(){
	//printing years
	bd := 1996
	count := 0
	for bd <= 2020 {
		fmt.Printf("%v\n", bd)
		bd++
		count++
	}
	fmt.Println(count)
}