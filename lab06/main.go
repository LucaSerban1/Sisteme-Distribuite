package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Info struct {
	ID    int
	Value int
}

type Memory map[Info]struct{}

type Node struct {
	ID        int
	Neighbors []int
	M         Memory
	Crashed   bool
}

const prob = 0.3


func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0
	for src := 0; src < n; src++ {
		// BFS din nodul src
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		queue := []int{src}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}
		for _, d := range dist {
			if d == -1 {
				return -1 // avem un nod la care nu putem ajunge din src
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

// tareConexFara verifică dacă subgraful indus de nodurile încă corecte,
// fără nodul `exclus`, rămâne tare conex.
func tareConexFara(nodes []*Node, exclus int) bool {
	idx := map[int]int{} // id original -> index în subgraf
	for _, nd := range nodes {
		if !nd.Crashed && nd.ID != exclus {
			idx[nd.ID] = len(idx)
		}
	}
	if len(idx) == 0 {
		return false
	}
	sub := make([][]int, len(idx))
	for _, nd := range nodes {
		i, ok := idx[nd.ID]
		if !ok {
			continue
		}
		for _, v := range nd.Neighbors {
			if j, ok := idx[v]; ok {
				sub[i] = append(sub[i], j)
			}
		}
	}
	return diametru(sub) != -1
}

func copieMemorie(m Memory) Memory {
	c := make(Memory, len(m))
	for k := range m {
		c[k] = struct{}{}
	}
	return c
}

func formatMemorie(m Memory) string {
	lista := make([]Info, 0, len(m))
	for k := range m {
		lista = append(lista, k)
	}
	sort.Slice(lista, func(a, b int) bool {
		if lista[a].ID != lista[b].ID {
			return lista[a].ID < lista[b].ID
		}
		return lista[a].Value < lista[b].Value
	})
	parti := make([]string, len(lista))
	for i, x := range lista {
		parti[i] = fmt.Sprintf("(%d, %d)", x.ID, x.Value)
	}
	return strings.Join(parti, " ")
}

func egale(a, b Memory) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}
	return true
}

func citesteInput() (n, s int, tokens []int, adj [][]int, err error) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	linii := []string{}
	for sc.Scan() {
		l := strings.TrimSpace(sc.Text())
		if l != "" {
			linii = append(linii, l)
		}
	}
	if len(linii) < 2 {
		return 0, 0, nil, nil, fmt.Errorf("input incomplet")
	}

	atoi := func(f []string) ([]int, error) {
		r := make([]int, len(f))
		for i, x := range f {
			v, e := strconv.Atoi(x)
			if e != nil {
				return nil, fmt.Errorf("valoare invalidă %q", x)
			}
			r[i] = v
		}
		return r, nil
	}

	prima, err := atoi(strings.Fields(linii[0]))
	if err != nil || len(prima) < 2 {
		return 0, 0, nil, nil, fmt.Errorf("linia 1 trebuie să conțină n și s")
	}
	n, s = prima[0], prima[1]

	tokens, err = atoi(strings.Fields(linii[1]))
	if err != nil || len(tokens) != n {
		return 0, 0, nil, nil, fmt.Errorf("linia 2 trebuie să conțină %d token-uri", n)
	}
	if len(linii) < 2+n {
		return 0, 0, nil, nil, fmt.Errorf("lipsesc liste de adiacență (aștept %d)", n)
	}

	adj = make([][]int, n)
	for i := 0; i < n; i++ {
		vec, e := atoi(strings.Fields(linii[2+i]))
		if e != nil {
			return 0, 0, nil, nil, e
		}
		for _, v := range vec {
			if v < 0 || v >= n {
				return 0, 0, nil, nil, fmt.Errorf("vecin invalid %d pentru nodul %d", v, i)
			}
		}
		adj[i] = vec
	}
	return n, s, tokens, adj, nil
}

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	n, s, tokens, adj, err := citesteInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Eroare:", err)
		os.Exit(1)
	}

	diam := diametru(adj)
	if diam == -1 {
		fmt.Fprintln(os.Stderr, "Eroare: graful nu este tare conex")
		os.Exit(1)
	}

	// Inițializare: Mi(0) = {(i, vi)}
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{ID: i, Neighbors: adj[i], M: Memory{{i, tokens[i]}: {}}}
	}

	// Câte o inbox pe nod; un nod primește cel mult n mesaje pe iterație.
	inbox := make([]chan Memory, n)
	for i := range inbox {
		inbox[i] = make(chan Memory, n)
	}

	limita := (s + 1) * diam
	crashuri := 0
	t := 0
	for {
		victima := -1
		if crashuri < s && rng.Float64() < prob {
			candidati := []int{}
			for _, nd := range nodes {
				if !nd.Crashed && tareConexFara(nodes, nd.ID) {
					candidati = append(candidati, nd.ID)
				}
			}
			if len(candidati) > 0 {
				victima = candidati[rng.Intn(len(candidati))]
				crashuri++
			}
		}

		var wg sync.WaitGroup
		for _, nd := range nodes {
			if nd.Crashed {
				continue
			}
			destinatari := nd.Neighbors
			if nd.ID == victima {
				destinatari = nil
				for _, v := range nd.Neighbors {
					if rng.Intn(2) == 0 {
						destinatari = append(destinatari, v)
					}
				}
			}
			wg.Add(1)
			go func(nd *Node, dest []int) {
				defer wg.Done()
				for _, v := range dest {
					inbox[v] <- copieMemorie(nd.M)
				}
			}(nd, destinatari)
		}
		wg.Wait() // barieră: toate mesajele iterației t au fost trimise

		if victima != -1 {
			nodes[victima].Crashed = true
			fmt.Fprintf(os.Stderr, "[t=%d] nodul %d a intrat în crash\n", t, victima)
		}

		// Pașii 2-3: fiecare nod primește mesajele și face Mi(t+1) = Mi(t) ∪ U.
		for _, nd := range nodes {
			wg.Add(1)
			go func(nd *Node) {
				defer wg.Done()
				for {
					select {
					case m := <-inbox[nd.ID]:
						if nd.Crashed {
							continue // un nod căzut nu mai procesează nimic
						}
						for k := range m {
							nd.M[k] = struct{}{}
						}
					default:
						return
					}
				}
			}(nd)
		}
		wg.Wait()

		// Pasul 4: condiția de oprire.
		if t > limita {
			break
		}
		// Pasul 5
		t++
	}

	// Afișare: memoria finală a unui nod corect.
	var corect *Node
	for _, nd := range nodes {
		if !nd.Crashed {
			corect = nd
			break
		}
	}

	consens := true
	for _, nd := range nodes {
		if !nd.Crashed && !egale(nd.M, corect.M) {
			consens = false
		}
	}

	fmt.Println("Diametru:", diam)
	fmt.Println("Iteratia de oprire:", t)
	fmt.Println(formatMemorie(corect.M))
	fmt.Fprintf(os.Stderr, "Crash-uri: %d/%d, consens intre nodurile corecte: %v\n", crashuri, s, consens)
}