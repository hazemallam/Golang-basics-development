package mymath

import (
	"fmt"
	"testing"
)

func TestCenterdAvg(t *testing.T){
	type test struct{
		data []int
		answer float64
	}
	tests := []test{
		test{[]int{1,2,3,5}, 2.50},
		test{[]int{0,8,10,1000}, 9.00},
	}
	for _, v := range tests {
		x := CenterdAvg(v.data)
		if x != v.answer {
			t.Error("expected", v.answer, "got", x)
		}
	}
}

func ExampleCenterdAvg(){
	fmt.Printf("%.2f\n", CenterdAvg([]int{1,2,3,5}))
	//Output:
	//2.50
}

func BenchmarkCenterdAvg(b *testing.B){
	for i:=0; i<b.N; i++{
		CenterdAvg([]int{1,2,3,5})
	}
}
