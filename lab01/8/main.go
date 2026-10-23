package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Ana Popescu": 1,
		"Mihai Ionescu": 2,
		"Bogdan Radu": 1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile": 2,
		"Ioana Enache": 1,
	}

	m_v2 := map[int][]string{}

	for k, v := range m {
		m_v2[v] = append(m_v2[v], k)
	}
	fmt.Println(m_v2)
}