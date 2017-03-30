package main

import "fmt"

const meterstoyards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter number of meters swam: ")
	fmt.Scan(&meters)
	var yards = meters * meterstoyards
	var feet = yards * 3
	var mi = feet / 5280
	fmt.Println("That is", yards, "yards or", mi, "mi. or", feet, "feet.")
}
