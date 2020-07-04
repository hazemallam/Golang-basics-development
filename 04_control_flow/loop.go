package main

import "fmt"

func main(){
	//initial for loop
	for i := 0; i<10; i++{
		fmt.Printf("the outer loop : %d\n", i)
		for j := 0; j<3; j++{
			fmt.Printf("\t\tthe inner loop : %d\n", j)
		}
	}
	//for can work as a while
	x := 1
	for x <= 10{
		fmt.Println(x)
		x++
	}
	fmt.Println("Done.")
	
	//infinite loop and break it using break statement
	z := 1
	for{
		if z > 9{
			break
		}
		fmt.Println(z)
		z++
	}
	fmt.Println("done.")
	//using continue statement
	for i := 0; i<10; i++ {
		if i & 1 == 1{
			continue
		}
		fmt.Println(i)
	}
	//printing unicode 
	for i := 33; i<=122; i++ {
		fmt.Printf("%v\t%#U\n",i, i)
	}
}

