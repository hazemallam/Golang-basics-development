package main
import "testing"
func TestMySum(t *testing.T){
	x := MySum(5, 3)
	if x != 8 {
		t.Error("Expected 8 Got", x)
	}
	
}

