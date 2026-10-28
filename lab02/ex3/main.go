package main

import (
	"fmt"
)


func RutinaTrimis(ch chan string){
	ch <- "Salut!"
	ch <- "Lume"
	close(ch)
}

func main() {
	
	var ch = make(chan string, 2)

	go RutinaTrimis(ch)


	for elem :=	 range ch {
		fmt.Println(elem)
	}

}