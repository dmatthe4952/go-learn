package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {
	var x = 43
	fmt.Println(x)
	zero(&x)
	fmt.Println(x)
}
