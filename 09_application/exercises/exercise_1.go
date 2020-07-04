package main
import (
	"fmt"
	"encoding/json"
)
type user struct{
	First string
	Age int
}
func main(){
	u1 := user{
		First : "James",
		Age : 30,
	}
	u2 := user{
		First : "Moneypenny",
		Age : 20,
	}
	users := []user{u1, u2}
	fmt.Println(users)
	bs, err := json.Marshal(users)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}