package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := true
	d := 4.17

	fmt.Printf("%T : %d \t %T : %s \t %T : %t \t %T : %2.3f\n", a, a, b, b, c, c, d, d)
}
