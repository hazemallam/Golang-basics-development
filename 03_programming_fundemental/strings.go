package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){

	//first example 
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println(sample)
	// loop through the string byte by byte not char by char 
	for i := 0; i<len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Printf("%x\n", sample)
	fmt.Printf("% x\n", sample)
	//this flag causes the output to escape non-printable sequences
	fmt.Printf("%q\n", sample)
	//This flag causes the output to escape not only non-printable sequences, but also any non-ASCII bytes
	fmt.Printf("%+q\n", sample)

	bs := []byte(sample)
	fmt.Printf("%v\n", bs)

	for i := 0; i<len(sample); i++ {
		fmt.Printf("%q", sample[i])
	}
	fmt.Printf("\n")
	//SECOND EXAMPLE
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s\n", placeOfInterest)

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q\n", placeOfInterest)

	fmt.Printf("hex bytes : ")
	for i := 0; i<len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
	const strangeString = "日本語"
	for i, w := 0, 0; i<len(strangeString); i += w {
		runeValue, width := utf8.DecodeRuneInString(strangeString[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}

	
}



//Go source code is always UTF-8.
//A string holds arbitrary bytes.
//A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
//Those sequences represent Unicode code points, called runes.
//No guarantee is made in Go that characters in strings are normalized.
