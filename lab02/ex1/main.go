package main

import (
	"fmt"
)


func RutinaTrimis(ch chan string){
	ch <- "Salut!"
}

func main() {
	
	var ch = make(chan string)

	go RutinaTrimis(ch)


	fmt.Println(<-ch)

}