package main

import "fmt"

func main(){
	s1 := "S i Χ ρ 🌴"
	b1 := []byte(s1)
	fmt.Printf("Bytes of \"S i Χ ρ 🌴\" are:\n")
	fmt.Println(b1) //[83 32 105 32 206 167 32 207 129 32 240 159 140 180]
	// each rune(=code point) is from 1 to 4 bytes
}