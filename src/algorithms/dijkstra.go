package algorithms

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Entry struct {
	Name      string
	From      string
	Direction Direction
	Cost      int
	xLimit    int
	yLimit    int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int
}

type SeenEntry struct {
	Name      string
	Direction Direction
	Limit     int
}

//###########################################
// Priority queue implementation from go docs
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

func (graph *Graph) Dijkstra(startNode string, lowerLimits int, upperLimits int) (distances map[SeenEntry]int, err error) {
	startNode, ok := graph.Vertices[startNode]

	if !ok {
		return nil, fmt.Errorf("vertex %s not found in graph", startNode)
	}

	// Prepare the distance map
	distances = make(map[SeenEntry]int)
	for key := range graph.Vertices {
		for i := 0; i < 4; i++ {
			entryUp := SeenEntry{Name: key, Direction: UP, Limit: i}
			entryDown := SeenEntry{Name: key, Direction: DOWN, Limit: i}
			entryLeft := SeenEntry{Name: key, Direction: LEFT, Limit: i}
			entryRight := SeenEntry{Name: key, Direction: RIGHT, Limit: i}

			distances[entryUp] = math.MaxInt32
			distances[entryDown] = math.MaxInt32
			distances[entryLeft] = math.MaxInt32
			distances[entryRight] = math.MaxInt32
		}
	}

	startEntryVertical := SeenEntry{Name: startNode, Direction: LEFT, Limit: 3}
	//startEntryHorizontal := SeenEntry{Name: startNode, Direction: HORIZONTAL}

	distances[startEntryVertical] = 0
	//distances[startEntryHorizontal] = 0

	// prepare the seen nodes
	seen := make(map[SeenEntry]bool)
	seen[startEntryVertical] = true
	//seen[startEntryHorizontal] = true

	// Prepare the vertices
	h := make(PriorityQueue, 0)

	heap.Init(&h)
	heap.Push(&h, &Entry{Name: startNode, Direction: LEFT, xLimit: upperLimits, yLimit: upperLimits})

	for h.Len() != 0 {
		// Update the list of possible candidates
		currentShortestVertex := heap.Pop(&h).(*Entry)
		//fmt.Println("Starting a new loop with: ", currentShortestVertex.Name)
		for _, edge := range graph.Edges[currentShortestVertex.Name] {
			newPathCostDestination := distances[currentShortestVertex.toSeenEntry()] + edge.Cost
			// Can move vertically?
			isVertical := currentShortestVertex.isYMovementTo(edge)
			var destinationEntry SeenEntry
			if isVertical {
				//destinationEntry = edge.toVerticalEntry()
				if currentShortestVertex.isUpMovementTo(edge) {
					destinationEntry = edge.toUpEntry()
				}
				if currentShortestVertex.isDownMovementTo(edge) {
					destinationEntry = edge.toDownEntry()
				}
				destinationEntry.Limit = currentShortestVertex.yLimit - 1
			} else {
				//destinationEntry = edge.toHorizontalEntry()
				if currentShortestVertex.isLeftMovementTo(edge) {
					destinationEntry = edge.toLeftEntry()
				}
				if currentShortestVertex.isRightMovementTo(edge) {
					destinationEntry = edge.toRightEntry()
				}
				destinationEntry.Limit = currentShortestVertex.xLimit - 1
			}

			// Skip reversing directions
			if isOppositeDirection(currentShortestVertex, edge) {
				continue
			}

			//fmt.Printf("Trying for edge: %+v \n", edge)
			var distanceToDest = distances[destinationEntry]
			if !seen[currentShortestVertex.toSeenEntry()] || newPathCostDestination < distanceToDest {
				//fmt.Println("in 2")

				if canMoveToDestination(currentShortestVertex, edge, lowerLimits, upperLimits) {
					// update the heap
					newEntry := Entry{
						Name:      edge.To,
						From:      currentShortestVertex.Name,
						Cost:      newPathCostDestination,
						Direction: destinationEntry.Direction,
					}

					if isVertical {
						newEntry.xLimit = upperLimits
						newEntry.yLimit = currentShortestVertex.yLimit - 1
					} else {
						newEntry.xLimit = currentShortestVertex.xLimit - 1
						newEntry.yLimit = upperLimits
					}

					// Update the states to either v or h
					heap.Push(&h, &newEntry)
					//fmt.Printf("Adding entry: %+v \n", newEntry)
					distances[destinationEntry] = newPathCostDestination
					seen[destinationEntry] = true
				}
			}
		}
	}

	return distances, nil
}

func canMoveToDestination(from *Entry, destination Edge, lowerLimit int, upperLimit int) bool {
	canMoveToDestination := false
	if from.isXMovementTo(destination) {
		canMoveToDestination = from.xLimit > 0
	} else if from.isYMovementTo(destination) {
		canMoveToDestination = from.yLimit > 0
	}

	//fmt.Print("Moving from to: ", from, destination)
	//fmt.Println("  Can move to destination: ", canMoveToDestination)

	return canMoveToDestination
}

func (edge *Entry) isXMovementTo(destination Edge) bool {
	entryPos := strings.Split(edge.Name, ",")
	destinationPos := strings.Split(destination.To, ",")

	//fmt.Println("isXMovementTo: ", destinationPos[0], entryPos[0], entryPos[0] != destinationPos[0])

	return entryPos[0] != destinationPos[0]
}

func (edge *Entry) isYMovementTo(destination Edge) bool {
	entryPos := strings.Split(edge.Name, ",")
	destinationPos := strings.Split(destination.To, ",")

	return entryPos[1] != destinationPos[1]
}

func (edge *Entry) isLeftMovementTo(destination Edge) bool {
	entryPos := strings.Split(edge.Name, ",")
	destinationPos := strings.Split(destination.To, ",")

	leftX, _ := strconv.Atoi(entryPos[0])
	rightX, _ := strconv.Atoi(destinationPos[0])

	return (leftX - 1) == rightX
}

func (edge *Entry) isRightMovementTo(destination Edge) bool {
	entryPos := strings.Split(edge.Name, ",")
	destinationPos := strings.Split(destination.To, ",")

	leftX, _ := strconv.Atoi(entryPos[0])
	rightX, _ := strconv.Atoi(destinationPos[0])

	return (leftX + 1) == rightX
}

func (entry *Entry) isUpMovementTo(destination Edge) bool {
	entryPos := strings.Split(entry.Name, ",")
	destinationPos := strings.Split(destination.To, ",")

	leftY, _ := strconv.Atoi(entryPos[1])
	rightY, _ := strconv.Atoi(destinationPos[1])

	return (leftY - 1) == rightY
}

func (entry *Entry) isDownMovementTo(destination Edge) bool {
	entryPos := strings.Split(entry.Name, ",")
	destinationPos := strings.Split(destination.To, ",")

	leftY, _ := strconv.Atoi(entryPos[1])
	rightY, _ := strconv.Atoi(destinationPos[1])

	return (leftY + 1) == rightY
}

func (node *Entry) toSeenEntry() SeenEntry {
	entry := SeenEntry{Name: node.Name, Direction: node.Direction}

	if node.Direction == LEFT || node.Direction == RIGHT {
		entry.Limit = node.xLimit
	} else {
		entry.Limit = node.yLimit
	}

	return entry
}

func (edge *Edge) toLeftEntry() SeenEntry {
	return SeenEntry{Name: edge.To, Direction: LEFT}
}

func (edge *Edge) toRightEntry() SeenEntry {
	return SeenEntry{Name: edge.To, Direction: RIGHT}
}

func (edge *Edge) toUpEntry() SeenEntry {
	return SeenEntry{Name: edge.To, Direction: UP}
}

func (edge *Edge) toDownEntry() SeenEntry {
	return SeenEntry{Name: edge.To, Direction: DOWN}
}

func isOppositeDirection(entry *Entry, edge Edge) bool {
	if entry.Direction == UP && entry.isDownMovementTo(edge) {
		return true
	} else if entry.Direction == DOWN && entry.isUpMovementTo(edge) {
		return true
	} else if entry.Direction == LEFT && entry.isRightMovementTo(edge) {
		return true
	} else if entry.Direction == RIGHT && entry.isLeftMovementTo(edge) {
		return true
	} else {
		return false
	}
}
