package main

import (
	"fmt"
)

func main() {
	slice := []string{"Lavoratorul", "de", "sisteme", "distribuite"}
	n := len(slice)
	
	for i := 0; i < n/2; i++ {
		slice[i], slice[n-i-1] = slice[n-i-1], slice[i]
	}
	fmt.Println(slice)
}