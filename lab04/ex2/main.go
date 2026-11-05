package main

import (
	"sync"
)

type Node struct {
	Value int
	maxValue int
	intrari [] <- chan int
	iesiri [] chan <- int
}

func Flooding(n *Node, total int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	max := n.Value

	inbox := make(chan int, total * len(n.intrari))
	for _, in := range n.intrari {
		go func(in <-chan int){
			for v := range in{
				inbox <- v
			}
		}(in)
	}

	send := func(v int) {
		for _, out := range n.iesiri {
			out <- v
		}
	}

	recv := func() int {
		v := <-inbox
		if v > max {
			max = v
		}
		return v
	}

	mapa := make(map[int]bool)
	send(n.Value)
	
	vazut := 1
	mapa[n.Value] = true
	for vazut < total {
		v := recv()
		if !mapa[v]{
			mapa[v] = true
			send(v)
			vazut++
		}
	}
	n.maxValue = max
}

func main(){
	c12 := make(chan int, 4)
	c23 := make(chan int, 4)
	c31 := make(chan int, 4)
	c14 := make(chan int, 4)
	c42 := make(chan int, 4)

	node1 := &Node{Value: 1, maxValue: -1, intrari: []<-chan int{c31} , iesiri: []chan<- int{c12, c14}}
	node2 := &Node{Value: 2, maxValue: -1, intrari: []<-chan int{c12, c42} , iesiri: []chan<- int{c23}}
	node3 := &Node{Value: 3, maxValue: -1, intrari: []<-chan int{c23} , iesiri: []chan<- int{c31}}
	node4 := &Node{Value: 4, maxValue: -1, intrari: []<-chan int{c14} , iesiri: []chan<- int{c42}}

	total := 4

	var wg sync.WaitGroup
	wg.Add(4)

	go Flooding(node1, total, &wg)
	go Flooding(node2, total, &wg)
	go Flooding(node3, total, &wg)
	go Flooding(node4, total, &wg)


	wg.Wait()

	for _, node := range []*Node{node1, node2, node3, node4} {
		println("Node value:", node.Value, "Max value:", node.maxValue)
	}

}