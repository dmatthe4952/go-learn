package main

import (
	"fmt"
)

func isPallindromic(n int) bool {
	string := fmt.Sprintf("%v", n)
	length := len(string)
	p := true
	if length%2 == 0 {
		for i := 0; i < length/2; i++ {
			first := string[length-(i+1)]
			last := string[i]
			if first != last {
				p = false
			}
		}
	} else {
		for i := 0; i <= length/2; i++ {
			first := string[length-(i+1)]
			last := string[i]
			if first != last {
				p = false
			}
		}
	}
	return p
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Println(isPallindromic(n))
}
