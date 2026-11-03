package main

import (
	"math/rand"
	"sync"
	"sort"
)

type LamportClock struct{ counter int }

func (l *LamportClock) Tick() { l.counter++ }

func (l *LamportClock) Send() { l.counter++ }

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.counter++
}

type Message struct{ sender , ts int }

type Event struct {
	process int
	internal_counter int
	kind string // "internal", "send", "recv"
}

// TODO: retuneaza true daca a precede b
func less(a, b Event) bool { 
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

func process(id int, in <-chan Message, out chan<- Message, events chan<- Event, wg *sync.WaitGroup) {
	defer wg.Done()
	clock := &LamportClock{}

	send := func() {
		tick := rand.Intn(3) + 1
		for i := 0; i < tick; i++ {
			clock.Tick()
			events <- Event{id, clock.counter, "internal"}
		}
		clock.Send()
		events <- Event{id, clock.counter, "send"}
		out <- Message{id, clock.counter}
	}

	recv := func() {
		tick := rand.Intn(3) + 1
		for i := 0; i < tick; i++ {
			clock.Tick()
			events <- Event{id, clock.counter, "internal"}
		}
		msg := <-in
		clock.Receive(msg.ts)
		events <- Event{id, clock.counter, "recv"}
	}

	if id == 1 {
		send()
		recv()
	} else {
		recv()
		send()
	}
}


func main() {
	c12 := make(chan Message, 1)
	c23 := make(chan Message, 1)
	c31 := make(chan Message, 1)
	events := make(chan Event, 100)

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go process(1, c31, c12, events, wg)
	go process(2, c12, c23, events, wg)
	go process(3, c23, c31, events, wg)

	wg.Wait()
	close(events)

	var eventList []Event
	for e := range events {
		eventList = append(eventList, e)
	}

	sort.Slice(eventList, func(i, j int) bool {
		return less(eventList[i], eventList[j])
	})

	for _, e := range eventList {
		println("ts = ", e.internal_counter, "| P ", e.process, " | ", e.kind)
	}
}