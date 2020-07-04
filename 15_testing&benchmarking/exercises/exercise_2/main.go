package main

import (
	"fmt"
	"computer_science/Golang/work_place/writing_documentation/exercises/mymath"
)
func main(){
	xxi := gen()
	for _, v := range xxi{
		fmt.Printf("%.2f\n", mymath.CenterdAvg(v))
	}
}

func gen() [][]int {
	a := []int{1,2,3,4}
	b := []int{0,8,10,1000}
	c := []int{9000,4,5,6,10}
	d := []int{123,744,140,200}
	e := [][]int{a,b,c,d}
	return e
}