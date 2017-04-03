package main

import "fmt"

func half(n int) (int, bool) {
	var y bool
	y = (n%2 == 0)
	return n / 2, y
}
func main() {
	fmt.Print("1: ")
	fmt.Println(half(1))
	fmt.Print("2: ")
	fmt.Println(half(2))
}
