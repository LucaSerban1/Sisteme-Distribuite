package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	slice := make([]float32, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	var index int
	fmt.Scan(&index)

	Slicefinal := append(slice[:index], slice[index+1:]...)
	fmt.Println(Slicefinal)
}
