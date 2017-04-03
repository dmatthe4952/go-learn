package main

import "fmt"

func main() {
	half := func(n int) (int, bool) {
		y := (n%2 == 0)
		return n / 2, y
	}
	fmt.Print("1: ")
	fmt.Println(half(1))
	fmt.Print("2: ")
	fmt.Println(half(2))
}
