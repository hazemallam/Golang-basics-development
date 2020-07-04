package main

import "fmt"

func main(){
	a := 97
	b := 3.2
	c := "hazem"
	var boolean bool = true
	fmt.Printf("default = %v\n", a)
	fmt.Printf("GO'S default = %#v\n", a)
	fmt.Printf("TYPE = %T\n", a)
	fmt.Printf(`just a "%%" = %%\n`)
	fmt.Printf("bool = %t\n", boolean)
	fmt.Printf("binary = %b\n", a)
	fmt.Printf("char  = %c\n", a)
	fmt.Printf("decimal = %d\n", a)
	fmt.Printf("octal = %o\n", a)
	// fmt.Printf("Octal = %O\n", a)
	fmt.Printf("single quoted = %q\n", a)
	fmt.Printf("hexadecimal = %x\n", a)
	fmt.Printf("Hexadecimal = %X\n", a)
	fmt.Printf("unicode %U\n", a)
	fmt.Printf("/////////////////////////\n")
	fmt.Printf("decimalless scientific %b\n", b)
	fmt.Printf("scientific notation %e\n", b)
	fmt.Printf("SCINETIFIC NOTATION %E\n", b)
	fmt.Printf("decimal point but no notation %f\n", b)
	fmt.Printf("synonym for %F\n", b)
	fmt.Printf("large exponets %g\n", b)
	fmt.Printf("LARGE EXPONETS %G\n", b)
	fmt.Printf("hexadecimal notation %x\n", b)
	fmt.Printf("HEXADECIMAL NOTATION %X\n", b)
	fmt.Printf("/////////////////////////\n")
	fmt.Printf("%s\n", c)
	fmt.Printf("%q\n", c)
	fmt.Printf("%x\n", c)
	fmt.Printf("%X\n", c)
	fmt.Printf("/////////////////////////\n")
	var trying float64 = 50.403
	fmt.Printf("%f\n", trying)
	fmt.Printf("%9f\n", trying)
	fmt.Printf("%.2f\n", trying)
	fmt.Printf("%0.2f\n", trying)
	fmt.Printf("%.0f\n", trying)


}