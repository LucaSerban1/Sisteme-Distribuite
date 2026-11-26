package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

type Message struct {
	SenderID int
	Round    int
	Value    int
}

type Nod struct {
	id          int
	x           int
	s           int
	n           int
	isByzantine bool
	inbox       chan Message
}

func NewNod(id, valoareInitiala, s, n int) *Nod {
	return &Nod{
		id:    id,
		x:     valoareInitiala,
		s:     s,
		n:     n,
		inbox: make(chan Message, n),
	}
}

func (nod *Nod) sendMessage(dest *Nod, round, value int) {
	dest.inbox <- Message{SenderID: nod.id, Round: round, Value: value}
}

func (nod *Nod) Bcast(noduri []*Nod, round, value int) {
	for _, dest := range noduri {
		nod.sendMessage(dest, round, value)
	}
}

func (nod *Nod) MaliciousBcastState(noduri []*Nod) {
	for _, dest := range noduri {
		nod.sendMessage(dest, 1, rand.Intn(2))
	}
}

func (nod *Nod) MaliciousBcastTie(noduri []*Nod) {
	for _, dest := range noduri {
		nod.sendMessage(dest, 2, rand.Intn(2))
	}
}

func (nod *Nod) receiveMessages(cnt int) []int {
	vals := make([]int, 0, cnt)
	for i := 0; i < cnt; i++ {
		msg := <-nod.inbox
		vals = append(vals, msg.Value)
	}

	return vals
}

func calculeazaMajoritate(valori []int, n int) (int, int) {
	freqUnu := 0
	for _, valoare := range valori {
		if valoare == 1 {
			freqUnu++
		}
	}

	if freqUnu >= n/2+1 {
		return 1, freqUnu
	}

	return 0, len(valori) - freqUnu
}

func gasesteNoduriBizantine(noduri []*Nod, s int) []int {
	for _, nod := range noduri {
		nod.isByzantine = false
	}

	nrNoduriMalitioase := rand.Intn(s + 1)

	indiciAlesi := make(map[int]bool)
	cntIndiciAlesi := 0
	sliceIndici := make([]int, 0, nrNoduriMalitioase)
	for {
		if cntIndiciAlesi == nrNoduriMalitioase {
			break
		}

		index := rand.Intn(len(noduri))
		if !indiciAlesi[index] {
			noduri[index].isByzantine = true

			indiciAlesi[index] = true
			cntIndiciAlesi++

			sliceIndici = append(sliceIndici, index)
		}
	}

	return sliceIndici
}

func ruleazaInParalel(noduri []*Nod, metoda func(nod *Nod)) {
	var wg sync.WaitGroup
	for _, nod := range noduri {
		wg.Go(func() {
			metoda(nod)
		})
	}
	wg.Wait()
}

func runda1(noduri []*Nod) (majoritate []int, mults []int) {
	n := len(noduri)
	majoritate = make([]int, n)
	mults = make([]int, n)

	ruleazaInParalel(noduri, func(nod *Nod) {
		if nod.isByzantine {
			nod.MaliciousBcastState(noduri)
		} else {
			nod.Bcast(noduri, 1, nod.x)
		}
	})

	ruleazaInParalel(noduri, func(nod *Nod) {
		u := nod.receiveMessages(n)
		majoritate[nod.id], mults[nod.id] = calculeazaMajoritate(u, n)
	})

	return majoritate, mults
}

func runda2(noduri []*Nod, lider *Nod, majoritate []int, mults []int) {
	if lider.isByzantine {
		lider.MaliciousBcastTie(noduri)
	} else {
		lider.Bcast(noduri, 2, majoritate[lider.id])
	}

	ruleazaInParalel(noduri, func(nod *Nod) {
		tie := nod.receiveMessages(1)[0]

		if 2*mults[nod.id] > nod.n+2*nod.s {
			nod.x = majoritate[nod.id]
		} else {
			nod.x = tie
		}
	})
}

func valoareFinalaConsens(noduri []*Nod) int {
	valoare := -1
	for _, nod := range noduri {
		if nod.isByzantine {
			continue
		}

		if valoare == -1 {
			valoare = nod.x
		}
	}

	return valoare
}

func byzflood(noduri []*Nod, s int) int {
	n := len(noduri)
	for i := 0; i <= s+1; i++ {
		lider := noduri[i%n]

		gasesteNoduriBizantine(noduri, s)

		majoritate, mults := runda1(noduri)
		runda2(noduri, lider, majoritate, mults)
	}
	return valoareFinalaConsens(noduri)
}

func citireInput() (int, int, []int) {
	var n, s int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	n = citesteInt(scanner)
	s = citesteInt(scanner)

	valori := make([]int, n)
	for i := range valori {
		valori[i] = citesteInt(scanner)
	}

	return n, s, valori
}

func citesteInt(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		panic("input incomplet")
	}
	val, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return val
}

func creeazaNoduri(n, s int, valori []int) []*Nod {
	noduri := make([]*Nod, n)

	for i := range noduri {
		noduri[i] = NewNod(i, valori[i], s, n)
	}

	return noduri
}

func main() {
	n, s, valori := citireInput()
	noduri := creeazaNoduri(n, s, valori)

	fmt.Println("Final: ", byzflood(noduri, s))
}