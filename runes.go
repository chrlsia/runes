package main

import "fmt"

func main(){
	// assign
	r1:='a' 
	r2:='\u03B2' // 2 bytes
	r10:='\U0001F334' // 4 bytes

	// display
	fmt.Println(r1) // 97, code point
	fmt.Println(r2) // 946, code point
	fmt.Printf("%U\n",r2) // U+03B2
	fmt.Printf("%c\n",r10) // 🌴
}