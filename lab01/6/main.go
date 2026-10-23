package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Andrei": 190,
		"Maria": 165,
		"Ion": 182,
		"Elena": 192,
		"George": 178,
		"Ana": 165,
	}

	type Person struct {
		nume  string
		inaltime int
	}

	var slice []Person

	for key, value := range m {
		slice = append(slice, Person{nume: key, inaltime: value})
	}

	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i].inaltime < slice[j].inaltime {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i].nume, slice[i].inaltime)
	}
}