package main

import "fmt"

func main() {
	var dividend int
	var divisor int
	fmt.Print("Enter a small number: ")
	fmt.Scan(&divisor)
	fmt.Print("Enter a larger number: ")
	fmt.Scan(&dividend)
	fmt.Println("Dividing the second number by the first leaves a remainder of:", dividend%divisor)
}
