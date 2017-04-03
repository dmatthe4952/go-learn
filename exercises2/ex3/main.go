package main

import "fmt"

func greatest(list []int) int {
	var hold int
	for x := 0; x < len(list); x++ {
		if list[x] > hold {
			hold = list[x]
		}
	}
	return hold
}

func main() {
	aSlice := []int{50, 22, 46, 39, 70, 66, 2, 5}
	fmt.Println("Of the list ", aSlice)
	fmt.Println("The greatest number is:", greatest(aSlice))
}
