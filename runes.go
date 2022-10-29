package main

import "fmt"

func main(){
	s1 := "S i Χ ρ 🌴"
	fmt.Printf("\nCode points of \"S i Χ ρ 🌴\" are:\n")
	r3 := []rune(s1)
	fmt.Printf("%U\n", r3) //[U+0053 U+0020 U+0069 U+0020 U+03A7 U+0020 U+03C1 U+0020 U+1F334]
	// each character is a rune(=code point)
}