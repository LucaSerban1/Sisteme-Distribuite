package main

import (
	"fmt"
)

func main(){
	array := [5]int{10,15,13,40,19}

	for i := 0; i < len(array); i++{
		for j := i + 1; j < len(array); j++{
			if array[i] > array[j]{
				aux := array[i]
				array[i] = array[j]
				array[j] = aux
			}
		}
	}
	fmt.Println(array)
}