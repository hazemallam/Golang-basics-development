package main
import "fmt"

func main(){
	elements := []int {1,2,3,4,5,6,7,8,9}
	total := sum(elements...)
	fmt.Println(total)
	summation := barr(sum, elements...)
	fmt.Println(summation)
	n2 := summ(elements)
	fmt.Println(n2)

}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func barr(f func(x ...int) int, ele ...int) int{
	return f(ele...)
}

func summ(x []int) int{
	total := 0
	for _, v := range x {
		total += v
	}
	return total

}
