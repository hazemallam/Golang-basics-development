package main

import "testing"

func TestMySum(t *testing.T){

	type test struct{
		data []int
		answer int
	}

	tests := []test{
		test{[]int{5, 3}, 8},
		test{[]int{5, 4}, 9},
		test{[]int{5, 5}, 10},
		test{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		x := MySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}