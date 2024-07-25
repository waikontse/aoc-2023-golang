package algorithms

import (
	"fmt"
)

type Edge struct {
	From string
	To   string
	Cost int
}

type Graph struct {
	// The number of vertices
	Vertices map[string]string
	// The edges from x to <n>
	Edges map[string][]Edge
}

func (graph *Graph) GetVertexCount() int {
	return len(graph.Vertices)
}

// Print functions
func (graph *Graph) PrintEdges() {
	for vertex, edges := range graph.Edges {
		fmt.Print("Vertex: ", vertex)
		for _, edge := range edges {
			fmt.Print(edge)
		}

		fmt.Println("")
	}
}

func (graph *Graph) AddVertex(name string) {
	graph.Vertices[name] = name
}

func (graph *Graph) AddEdge(vertex string, to string, cost int) {
	edge := Edge{From: vertex, To: to, Cost: cost}

	graph.Edges[vertex] = append(graph.Edges[vertex], edge)
}
