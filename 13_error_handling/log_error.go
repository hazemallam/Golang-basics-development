package main
import (
	"log"
	"fmt"
	"os"
)
func main(){
	f, err := os.Create("log.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil{
		// fmt.Println("err happened: ", err)
		log.Println("err happened: ", err)
		// log.Fatalln(err)
		// log.Panic(err)
		// panic(err)
	}
	defer f2.Close()
	fmt.Println("check the log.txt file in the directory")
	
}