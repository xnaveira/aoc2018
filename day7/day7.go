package day7

import (
	"bufio"
	"fmt"
	"github.com/gonum/graph"
	"github.com/gonum/graph/simple"
	"log"
	"os"
	"sort"
)

var example bool

func Run(input string) (string, string, error) {

	fmt.Println("This is day7")

	example = false

	var f *os.File
	var err error
	var workers int
	if example {
		f, err = os.Open("day7/example.txt")
		workers = 2
	} else {
		f, err = os.Open(input)
		workers = 5
	}
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	g := simple.NewDirectedGraph(0, 0)
	gg := simple.NewDirectedGraph(0, 0)

	for scanner.Scan() {
		var s1, s2 string
		_, err := fmt.Sscanf(scanner.Text(), "Step %s must be finished before step %s can begin", &s1, &s2)
		if err != nil {
			log.Fatal(err)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(s1, " ", s2)
		g.SetEdge(simple.Edge{
			simple.Node(formatInput(int(s1[0]), example)),
			simple.Node(formatInput(int(s2[0]), example)),
			0})
		gg.SetEdge(simple.Edge{
			simple.Node(formatInput(int(s1[0]), example)),
			simple.Node(formatInput(int(s2[0]), example)),
			0})
	}

	transitiveReduction(g)

	//First try, didnt work
	//sorted, err := topo.SortStabilized(g, func(nodes []graph.Node) {
	//	sort.SliceStable(nodes, func(i, j int) bool {
	//		return nodes[i].ID() < nodes[j].ID()
	//	})
	//})

	sorted, err := kahn(g)
	if err != nil {
		return "", "", err
	}

	result1 := ""
	for _, e := range sorted {
		result1 = result1 + string(formaOutput(e.ID(), example))
	}

	//transitiveReduction(gg)

	steps, err := kahnWorkers(gg, workers)
	if err != nil {
		return "", "", err
	}

	return result1, string(steps), nil
}

func transitiveReduction(g *simple.DirectedGraph) {
	for _, x := range g.Nodes() {
		for _, y := range g.Nodes() {
			for _, z := range g.Nodes() {
				if g.HasEdgeFromTo(x, y) && g.HasEdgeFromTo(y, z) {
					if g.HasEdgeFromTo(x, z) {
						g.RemoveEdge(g.Edge(x, z))
					}
				}
			}
		}
	}
}

func kahnWorkers(g *simple.DirectedGraph, w int) (int, error) {
	S := []graph.Node{}
	seconds := -1
	type worker struct {
		n graph.Node
		u int
	}
	working := []worker{}
	visited := map[graph.Node]bool{}
	for _, n := range g.Nodes() {
		if len(g.To(n)) == 0 {
			S = append(S, n)
		}
	}
	for {
		if len(S) == 0 && len(working) == 0 {
			break
		}
		seconds++
		tmpS := []graph.Node{}
		for _, m := range S {
			if len(working) < w {
				if !visited[m] {
					parents := g.To(m)
					parentsInWorking := false
					for _, w := range working {
						for _, p := range parents {
							if w.n == p {
								parentsInWorking = true
							}
						}
					}
					numOfParents := len(parents)
					visitedParents := 0
					for k, v := range visited {
						for _, p := range parents {
							if v == true && k == p {
								visitedParents++
							}
						}
					}
					parentsInVisited := numOfParents == visitedParents
					if !parentsInWorking && parentsInVisited {
						working = append(working, worker{m, m.ID()})
						sort.SliceStable(working, func(i, j int) bool {
							return working[i].n.ID() < working[j].n.ID()
						})
						visited[m] = true
					}
				}
			} else {
				tmpS = append(tmpS, m)
			}
		}
		S = tmpS
		tmpworking := []worker{}
		for i, ww := range working {
			fmt.Printf("Second: %d. Working on: ", seconds)
			for _, w := range working {
				fmt.Printf("%s %d ", string(formaOutput(w.n.ID(), example)), w.u)
			}
			fmt.Printf("\n")
			working[i].u--
			if working[i].u == 0 {
				S = append(S, g.From(ww.n)...)
				sort.SliceStable(S, func(i, j int) bool {
					return S[i].ID() < S[j].ID()
				})
			} else {
				tmpworking = append(tmpworking, working[i])
			}
		}
		working = tmpworking

	}
	if len(g.Edges()) != 0 {
		return 0, fmt.Errorf("the graph ahs at least one cycle: %v", g.Edges())
	} else {
		return seconds, nil
	}

}

func kahn(g *simple.DirectedGraph) ([]graph.Node, error) {
	S := []graph.Node{}
	for _, n := range g.Nodes() {
		if len(g.To(n)) == 0 {
			S = append(S, n)
		}
	}
	L := []graph.Node{}
	var n graph.Node
	for {
		if len(S) == 0 {
			break
		}
		n = S[0]
		S = S[1:]
		L = append(L, n)
		for _, m := range g.From(n) {
			g.RemoveEdge(simple.Edge{n, m, 0})
			if len(g.To(m)) == 0 {
				S = append(S, m)
			}
		}
		//Modified kahn slightly so it respects order
		sort.SliceStable(S, func(i, j int) bool {
			return S[i].ID() < S[j].ID()
		})
	}
	if len(g.Edges()) != 0 {
		return nil, fmt.Errorf("the graph ahs at least one cycle")
	} else {
		return L, nil
	}

}

func formatInput(i int, e bool) int {
	if e {
		return i - 64
	} else {
		return i - 64 + 60
	}
}

func formaOutput(i int, e bool) int {
	if e {
		return i + 64
	} else {
		return i + 64 - 60
	}
}
