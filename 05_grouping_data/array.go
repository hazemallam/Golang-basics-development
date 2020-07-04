package main

import "fmt"
func main(){
	var arr [2][3]int
	fmt.Printf("%v\n", arr)
	for i := 0; i<2; i++ {
		for j := 0; j<3; j++ {
			fmt.Println(arr[i][j])
		}
	}
}

