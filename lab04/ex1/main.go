package main

import (
	"sync"
)

type Node struct {
	Value int
	maxValue int
}

func Flooding(n *Node, in <-chan int, out chan <- int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	max := n.Value

	send := func(v int) {
		out <- v
	}

	recv := func() int {
		v := <-in
		if v > max {
			max = v
		}
		return v
	}

	send(n.Value)
	for {
		primit := recv()
		if primit == n.Value {
			n.maxValue = primit
			break
		}
		send(primit)
	}
}

func main(){
	node1 := &Node{Value: 1, maxValue: -1}
	node2 := &Node{Value: 2, maxValue: -1}
	node3 := &Node{Value: 3, maxValue: -1}

	c12 := make(chan int, 1)
	c23 := make(chan int, 1)
	c31 := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(3)

	go Flooding(node1, c31,c12, &wg)
	go Flooding(node2, c12, c23, &wg)
	go Flooding(node3, c23, c31, &wg)

	wg.Wait()

	for _, node := range []*Node{node1, node2, node3} {
		println("Node value:", node.Value, "Max value:", node.maxValue)
	}
}	