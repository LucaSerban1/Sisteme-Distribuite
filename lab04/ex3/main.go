package main

import (
	"sync"
)

type Message struct {
	Id int
	Vot int
}

type Node struct {
	Value int
	Vot int
	VotRezultat int
	intrari [] <- chan Message
	iesiri [] chan <- Message
}



func Flooding(n *Node, total int, wg *sync.WaitGroup) {
	defer wg.Done()
	

	msg := Message{Id: n.Value, Vot: n.Vot}

	inbox := make(chan Message, total * len(n.intrari))
	for _, in := range n.intrari {
		go func(in <-chan Message) {
			for v := range in{
				inbox <- v
			}
		}(in)
	}

	send := func(v Message) {
		for _, out := range n.iesiri {
			out <- v
		}
	}

	recv := func() Message {
		v := <-inbox
		return v
	}

	mapa := make(map[int]int)
	send(msg)
	
	vazut := 1
	mapa[msg.Id] = msg.Vot
	for vazut < total {
		v := recv()
		_ , ok := mapa[v.Id]
		if !ok{
			mapa[v.Id] = v.Vot
			send(v)
			vazut++
		}
	}

	suma := 0
	for _, v := range mapa {
		suma += v
	}
	if suma > total / 2 {
		n.VotRezultat = 1
	}else{
		n.VotRezultat = 0
	}
}

func main(){
	c12 := make(chan Message, 4)
	c23 := make(chan Message, 4)
	c31 := make(chan Message, 4)
	c14 := make(chan Message, 4)
	c42 := make(chan Message, 4)

	node1 := &Node{Value: 1, Vot: 0, intrari: []<-chan Message{c31} , iesiri: []chan<- Message{c12, c14}}
	node2 := &Node{Value: 2, Vot: 0, intrari: []<-chan Message{c12, c42} , iesiri: []chan<- Message{c23}}
	node3 := &Node{Value: 3, Vot: 0, intrari: []<-chan Message{c23} , iesiri: []chan<- Message{c31}}
	node4 := &Node{Value: 4, Vot: 1, intrari: []<-chan Message{c14} , iesiri: []chan<- Message{c42}}

	total := 4

	var wg sync.WaitGroup
	wg.Add(4)

	go Flooding(node1, total, &wg)
	go Flooding(node2, total, &wg)
	go Flooding(node3, total, &wg)
	go Flooding(node4, total, &wg)


	wg.Wait()

	for _, node := range []*Node{node1, node2, node3, node4} {
		println("Node value:", node.Value, "Vot nod:", node.Vot, "Vot Grup:", node.VotRezultat)
	}

}