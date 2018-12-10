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

func Run(input string) (string, string, error) {

	fmt.Println("This is day7")

	f, err := os.Open(input)

	//f, err := os.Open("day7/example.txt")
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	g := simple.NewDirectedGraph(0, 0)

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
		g.SetEdge(simple.Edge{simple.Node(s1[0]), simple.Node(s2[0]), 0})
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
		result1 = result1 + string(e.ID())
	}

	return result1, "", nil
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
