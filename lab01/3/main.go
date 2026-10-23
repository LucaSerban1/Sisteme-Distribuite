package main

import (
	"fmt"
)

func main() {
	vector := [5]int{10, 15, 13, 40, 19}
	var slice []int

	for i := 0; i < len(vector); i++ {
		if vector[i]%2 == 0 {
			slice = append(slice, vector[i])
		}
	}
	fmt.Println(slice)
}