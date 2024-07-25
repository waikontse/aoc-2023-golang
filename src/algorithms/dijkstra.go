package algorithms

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Entry struct {
	Name   string
	From   string
	Cost   int
	xLimit int
	yLimit int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int
}

type SeenEntry struct {
	Name      string
	Direction Direction
}

//###########################################
// Priority queue implementation fron go docs
//###########################################
// An Item is something we manage in a priority queue.

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Entry

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Entry)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	//item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

//###########################################

//type EntryHeap []Entry
//
//func (h EntryHeap) Len() int           { return len(h) }
//func (h EntryHeap) Less(i, j int) bool { return h[i].Cost < h[j].Cost }
//func (h EntryHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

//func (h *EntryHeap) Push(x any) {
//	// Push and Pop use pointer receivers because they modify the slice's length,
//	// not just its contents.
//	*h = append(*h, Entry{Name: x.(Entry).Name, Cost: x.(Entry).Cost, xLimit: x.(Entry).xLimit, yLimit: x.(Entry).yLimit})
//}

//func (h *EntryHeap) Pop() any {
//	old := *h
//	//n := len(old)
//	x := old[0]
//	*h = old[1:]
//	return x
//}

func (graph *Graph) Dijkstra(startNode string, limits int) (distances map[string]int, err error) {
	startNode, ok := graph.Vertices[startNode]

	if !ok {
		return nil, fmt.Errorf("vertex %s not found in graph", startNode)
	}

	// Prepare the distance map
	distances = make(map[string]int)
	for key := range graph.Vertices {
		distances[key+"v"] = math.MaxInt32
		distances[key+"h"] = math.MaxInt32
	}
	distances[startNode+"h"] = 0
	distances[startNode+"v"] = 0

	// prepare the seen nodes
	seen := make(map[SeenEntry]bool)
	seen[startNode] = true

	// Prepare the vertices
	h := make(PriorityQueue, 0)

	heap.Init(&h)
	heap.Push(&h, &Entry{Name: startNode + "h", xLimit: limits, yLimit: limits})
	//for key := range graph.Vertices {
	//	h.Push(Entry{Name: key, Cost: distances[key]})
	//}

	for h.Len() != 0 {
		// Find the lowest cost vertice first
		//slices.SortStableFunc(vertices, func(from string, to string) int {
		//	return distances[from] - distances[to]
		//})

		// Update the list of possible candidates
		currentShortestVertex := heap.Pop(&h).(*Entry)
		//fmt.Println("Starting a new loop with: ", currentShortestVertex.Name)
		nameForEdges := currentShortestVertex.Name[0 : len(currentShortestVertex.Name)-1]
		for _, edge := range graph.Edges[nameForEdges] {
			newPathCostDestination := distances[currentShortestVertex.Name] + edge.Cost
			// Can move vertically?
			if newPathCostDestination < distances[edge.To+"v"] {
				suffix := ""
				if canMoveToDestination(currentShortestVertex, edge) {
					// update the heap
					if currentShortestVertex.isXMovementTo(edge) {
						suffix = "h"
						heap.Push(&h, &Entry{
							Name:   edge.To + suffix,
							From:   currentShortestVertex.Name,
							Cost:   newPathCostDestination,
							xLimit: currentShortestVertex.xLimit - 1,
							yLimit: limits,
						})
					} else {
						suffix = "v"
						heap.Push(&h, &Entry{
							Name:   edge.To + suffix,
							From:   currentShortestVertex.Name,
							Cost:   newPathCostDestination,
							xLimit: limits,
							yLimit: currentShortestVertex.yLimit - 1,
						})
					}
				}

				// Update the states to either v or h
				distances[edge.To+suffix] = newPathCostDestination
				seen[edge.To+suffix] = true
			}

			// Can move horizontally
			if newPathCostDestination < distances[edge.To+"h"] {
				// TODO fill me in
			}
		}
	}

	return distances, nil
}

func canMoveToDestination(from *Entry, destination Edge) bool {
	canMoveToDestination := false
	if from.isXMovementTo(destination) {
		canMoveToDestination = from.xLimit > 0
	} else if from.isYMovementTo(destination) {
		canMoveToDestination = from.yLimit > 0
	}

	fmt.Print("Moving from to: ", from, destination)
	fmt.Println("  Can move to destination: ", canMoveToDestination)

	return canMoveToDestination
}

func (entry *Entry) isXMovementTo(destination Edge) bool {
	entryPos := strings.Split(entry.Name, ",")
	destinationPos := strings.Split(destination.To, ",")

	//fmt.Println("isXMovementTo: ", destinationPos[0], entryPos[0], entryPos[0] != destinationPos[0])

	return entryPos[0] != destinationPos[0]
}

func (entry *Entry) isYMovementTo(destination Edge) bool {
	entryPos := strings.Split(entry.Name, ",")
	destinationPos := strings.Split(destination.To, ",")

	//fmt.Println("isYMovementTo: ", destinationPos[1], entryPos[1], entryPos[1] != destinationPos[1])

	return entryPos[1] != destinationPos[1]
}
