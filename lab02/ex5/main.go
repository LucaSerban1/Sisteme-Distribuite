package main

import "fmt"

func receive(ch <-chan string) string {
	return <-ch
}

func send(ch chan<- string, msg string) {
	ch <- msg
	close(ch)
}

func main() {
	ch := make(chan string)

	go send(ch, "Salut!")
	msg := receive(ch)
	fmt.Println(msg)
}