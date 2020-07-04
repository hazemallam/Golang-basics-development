package main

import (
	"fmt"
	"runtime"
)

func main() {

	var num1 int8  //-128 to 127
	var num2 int16 //-32768 to 32767
	var num3 int32 //-2147483648 to 2147483647
	var num4 int64 //-9223372036854775808 to 9223372036854775807

	var num5 uint8  //0 to 255
	var num6 uint16 //0 to 65535
	var num7 uint32 //0 to 4294967295
	var num8 uint64 //0 to 18446744073709551615

	var num9 float32  // 32-bit floating point numbers
	var num10 float64 // 64-bit floating point numbers

	var num11 complex64  //complex numbers with float32 and imaginary numbers
	var num12 complex128 //complex numbers with float64 and imaginary numbers

	var num13 byte //alias for uint8
	var num14 rune //alias for int32

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	var num15 uint //either 32 or 64 bits
	var num16 int  //either 32 or 64 bits

	fmt.Printf("%T\n", num1)
	fmt.Printf("%T\n", num2)
	fmt.Printf("%T\n", num3)
	fmt.Printf("%T\n", num4)
	fmt.Printf("%T\n", num5)
	fmt.Printf("%T\n", num6)
	fmt.Printf("%T\n", num7)
	fmt.Printf("%T\n", num8)
	fmt.Printf("%T\n", num9)
	fmt.Printf("%T\n", num10)
	fmt.Printf("%T\n", num11)
	fmt.Printf("%T\n", num12)
	fmt.Printf("%T\n", num13)
	fmt.Printf("%T\n", num14)
	fmt.Printf("%T\n", num15)
	fmt.Printf("%T\n", num16)

}
