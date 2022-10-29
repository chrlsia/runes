package main

import "fmt"

func main(){
	s10 := "🌴 Siannas Χρήστος"

	// range to string gives us 2 things:
	// first the position of character
	// second the current character(=rune=code point)
	for pos, ele := range s10 {
		fmt.Printf("%d-%c and byte position %d\n", ele, ele, pos)
	}
}