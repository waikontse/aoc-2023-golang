package utils

import "fmt"

type Board[T comparable] struct {
	Board  [][]T
	width  int
	height int
}

func CreateBoard[T comparable](width, height int) Board[T] {
	// Create the X-axis
	matrix := make([][]T, width)
	// Create the Y-axis
	for i := range matrix {
		matrix[i] = make([]T, height)
	}

	return Board[T]{Board: matrix, width: width, height: height}
}

func Print[T comparable](board Board[T]) {
	for y := 0; y < board.width; y++ {
		for x := 0; x < board.height; x++ {
			fmt.Print(board.Board[x][y])
			fmt.Print(",")
		}
		fmt.Println()
	}
}

func Get[T comparable](board Board[T], x, y int) T {
	return board.Board[x][y]
}

func Set[T comparable](board Board[T], x, y int, value T) {
	board.Board[x][y] = value
}

func GetRow[T comparable](board Board[T], height int) []T {
	row := make([]T, board.width)
	for i := range board.Board {
		row[i] = board.Board[i][height]
	}

	return row
}

func GetColumn[T comparable](board Board[T], width int) []T {
	return board.Board[width]
}
