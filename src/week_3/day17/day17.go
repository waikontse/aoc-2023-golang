package day17

import (
	"aoc-2023-golang/src/algorithms"
	"aoc-2023-golang/src/utils"
	"fmt"
	"strconv"
)

func SolvePart1(filename string, upperLimit int) int {
	board := utils.ParseInputIntoBoard(filename)

	graph := parseFileToGraph(filename)
	distances, _ := graph.Dijkstra("0,0", 0, upperLimit)

	return getShortestPathToPosition(fmt.Sprintf("%d,%d", board.Width-1, board.Height-1), distances)
}

func SolvePart2(filename string, lowerLimit int, upperLimit int) int {
	board := utils.ParseInputIntoBoard(filename)

	graph := parseFileToGraph(filename)
	distances, _ := graph.Dijkstra("0,0", lowerLimit, upperLimit)

	return getShortestPathToPosition(fmt.Sprintf("%d,%d", board.Width-1, board.Height-1), distances)
}

func getShortestPathToPosition(position string, distances map[algorithms.SeenEntry]int) int {
	// 1. Filter all the entries in distances map with the same name
	// 2. Get the lowest value of the distances

	distancesForMatchingEntries := make([]int, 0)
	for k, v := range distances {
		if k.Name == position {
			distancesForMatchingEntries = append(distancesForMatchingEntries, v)
		}
	}

	return utils.MinValue(distancesForMatchingEntries)
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
