package utils

import "fmt"

type Board[T comparable] struct {
	Board  [][]T
	Width  int
	Height int
}

func CreateBoard[T comparable](width, height int) Board[T] {
	// Create the X-axis
	matrix := make([][]T, width)
	// Create the Y-axis
	for i := range matrix {
		matrix[i] = make([]T, height)
	}

	return Board[T]{Board: matrix, Width: width, Height: height}
}

func Print[T comparable](board *Board[T]) {
	for y := 0; y < board.Width; y++ {
		for x := 0; x < board.Height; x++ {
			fmt.Print(board.Board[x][y])
		}
		fmt.Println()
	}
}

func Get[T comparable](board *Board[T], x, y int) T {
	return board.Board[x][y]
}

func Set[T comparable](board *Board[T], x, y int, value T) {
	board.Board[x][y] = value
}

func GetRow[T comparable](board *Board[T], height int) []T {
	row := make([]T, board.Width)
	for i := range board.Board {
		row[i] = board.Board[i][height]
	}

	return row
}

func GetColumn[T comparable](board *Board[T], width int) []T {
	return board.Board[width]
}

func CanMoveTop(p Position) bool {
	return p.Second > 0
}

func GetTopChar(board *Board[string], p Position) string {
	if CanMoveTop(p) {
		return Get(board, p.First, p.Second-1)
	}

	return ""
}

func CanMoveBottom(board *Board[string], p Position) bool {
	return p.Second < (board.Height - 1)
}

func GetBottomChar(b *Board[string], p Position) string {
	if CanMoveBottom(b, p) {
		return Get(b, p.First, p.Second+1)
	}

	return ""
}

func CanMoveLeft(p Position) bool {
	return p.First > 0
}

func GetLeftChar(b *Board[string], p Position) string {
	if CanMoveLeft(p) {
		return Get(b, p.First-1, p.Second)
	}

	return ""
}

func CanMoveRight(board *Board[string], p Position) bool {
	return p.First < (board.Width - 1)
}

func GetRightChar(b *Board[string], p Position) string {
	if CanMoveRight(b, p) {
		return Get(b, p.First+1, p.Second)
	}

	return ""
}
