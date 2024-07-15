package day17

import (
	"aoc-2023-golang/src/algorithms"
	"aoc-2023-golang/src/utils"
)

func SolvePart1(filename string) int {
	board := utils.ParseInputIntoBoard(filename)

	utils.Print(&board)

	return 0
}

func SolvePart2(filename string) int {
	return 0
}

func parseFileToGraph(filename string) algorithms.Graph {
	board := utils.ParseInputIntoIntBoard(filename)

	graph := algorithms.Graph{
		Vertices: make([]string, 0),
		Edges:    make(map[string][]algorithms.Edge),
	}

	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			// Create a vertex

			// Add the edges for the vortex

			// TOP
			// LEFT
			// RIGHT
			// BOTTOM
		}
	}

	return graph
}
