package main

import (
	"fmt"
	"sync"
)

func RutinaPrimit(ch chan string, wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println(<-ch)

}

func main(){
	ch := make(chan string)
	wg := sync.WaitGroup{}
	
	go RutinaPrimit(ch, &wg)
	wg.Add(1)

	ch <- "Salut!"

	wg.Wait()
}