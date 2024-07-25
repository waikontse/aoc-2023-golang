package day17

import (
	"aoc-2023-golang/src/algorithms"
	"aoc-2023-golang/src/utils"
	"fmt"
	"strconv"
)

func SolvePart1(filename string, limits int) int {
	board := utils.ParseInputIntoBoard(filename)

	utils.Print(&board)

	graph := parseFileToGraph(filename)
	//graph.PrintEdges()

	distances, _ := graph.Dijkstra("0,0", limits)

	for vertex, cost := range distances {
		fmt.Printf("Shortest path to %s: %d\n", vertex, cost)
	}

	return 0
}

func SolvePart2(filename string) int {
	return 0
}

func parseFileToGraph(filename string) algorithms.Graph {
	board := utils.ParseInputIntoIntBoard(filename)

	graph := algorithms.Graph{
		Vertices: make(map[string]string),
		Edges:    make(map[string][]algorithms.Edge),
	}

	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			// Create a vertex
			vertexName := strconv.Itoa(x) + "," + strconv.Itoa(y)
			graph.AddVertex(vertexName)

			// Add the edges for the vortex
			currentPosition := utils.Position{First: x, Second: y}

			// TOP
			if board.CanMoveTop(currentPosition) {
				topString := strconv.Itoa(x) + "," + strconv.Itoa(y-1)
				graph.AddEdge(vertexName, topString, board.Get(x, y-1))
			}

			// LEFT
			if board.CanMoveLeft(currentPosition) {
				leftString := strconv.Itoa(x-1) + "," + strconv.Itoa(y)
				graph.AddEdge(vertexName, leftString, board.Get(x-1, y))
			}

			// RIGHT
			if board.CanMoveRight(currentPosition) {
				rightString := strconv.Itoa(x+1) + "," + strconv.Itoa(y)
				graph.AddEdge(vertexName, rightString, board.Get(x+1, y))
			}
			// BOTTOM
			if board.CanMoveBottom(currentPosition) {
				bottomString := strconv.Itoa(x) + "," + strconv.Itoa(y+1)
				graph.AddEdge(vertexName, bottomString, board.Get(x, y+1))
			}
		}
	}

	return graph
}
