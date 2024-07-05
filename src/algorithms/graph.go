package algorithms

import "fmt"

type Edge struct {
	From int
	To   int
	Cost int
}

type Graph struct {
	Vertices []int
	Edges    [][]Edge
}

func (graph *Graph) getVertexCount() int {
	return len(graph.Vertices)
}

// Print functions
func (graph *Graph) printEdges() {
	for i := 0; i < len(graph.Edges); i++ {
		for y := 0; y < len(graph.Edges[i]); y++ {
			fmt.Print("Edge: ", graph.Edges[i][y])
		}
		fmt.Println("")
	}
}

func (graph *Graph) addVertex(number int) {
	graph.Vertices = append(graph.Vertices, number)
}
