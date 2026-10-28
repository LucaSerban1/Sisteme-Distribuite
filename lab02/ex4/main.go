package main

import "fmt"


func Subprogram(ch chan bool) {
	fmt.Println("Salut!")
	ch <- false
}

func main() {
	ch := make(chan bool)

	ruleaza := true

	go Subprogram(ch)

	for ruleaza{
		fmt.Println("Inca ruleaza!")
		ruleaza = <-ch
	}
	fmt.Println("Sfarsit!")

}