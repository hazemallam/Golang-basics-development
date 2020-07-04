package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(string(bs))
	
	loginPassword := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err != nil{
		fmt.Println("YOU CAN NOT LOGGIN")
		return
	}
	fmt.Println("YOU LOGGED IN")

}
