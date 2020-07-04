package word

import (
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T){
	for k, v := range UseCount("hello how are you you"){
		switch k {
		case "hello":
			if v != 1 {
				t.Error("got", v, "expected", 1)
			}
		case "how":
			if v != 1 {
				t.Error("got", v, "expected", 1)
			}
		case "are":
			if v != 1 {
				t.Error("got", v, "expected", 1)
			}
		case "you":
			if v != 2 {
				t.Error("got", v, "expected", 2)
			}
		}
	}
}

func TestCount(t *testing.T){
	n := Count("hello how are you")
	if n != 4 {
		t.Error("got", n, "expected", 4)
	}
}

func ExampleCount(){
	fmt.Println(Count("hello how are you"))
	//Output:
	//4
}

func BenchmarkCount(b *testing.B){
	for i := 0; i<b.N; i++{
		Count("hello how are you")
	}
}

func BenchmarkUseCount(b *testing.B){
	for i := 0; i<b.N; i++{
		UseCount("hello how are you")
	}
}
