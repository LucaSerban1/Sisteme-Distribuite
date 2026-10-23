package main

import (
	"fmt"
)


func main() {
	var n int	
	fmt.Scan(&n)

	var s int
	fmt.Scan(&s)

	array := make([]int, n)
	m := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])

		if j, ok := m[s - array[i]]; ok{
			fmt.Println(j, i)
		}

		m[array[i]] = i
	}
	
	
}