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
	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			fmt.Print(board.Board[x][y])
		}
		fmt.Println()
	}
}

func (board *Board[T]) Get(x, y int) T {
	return board.Board[x][y]
}

func (board *Board[T]) Set(x, y int, value T) {
	board.Board[x][y] = value
}

func (board *Board[T]) GetRow(height int) []T {
	row := make([]T, board.Width)
	for i := range board.Board {
		row[i] = board.Board[i][height]
	}

	return row
}

func (board *Board[T]) GetColumn(width int) []T {
	return board.Board[width]
}

func CanMoveTop(p Position) bool {
	return p.Second > 0
}

func (board *Board[T]) GetTopChar(p Position) T {
	if CanMoveTop(p) {
		return board.Get(p.First, p.Second-1)
	}

	return *new(T)
}

func (board *Board[T]) CanMoveBottom(p Position) bool {
	return p.Second < (board.Height - 1)
}

func (board *Board[T]) GetBottomChar(p Position) T {
	if board.CanMoveBottom(p) {
		return board.Get(p.First, p.Second+1)
	}

	return *new(T)
}

func CanMoveLeft(p Position) bool {
	return p.First > 0
}

func (board *Board[T]) GetLeftChar(p Position) T {
	if CanMoveLeft(p) {
		return board.Get(p.First-1, p.Second)
	}

	return *new(T)
}

func (board *Board[T]) CanMoveRight(p Position) bool {
	return p.First < (board.Width - 1)
}

func (board *Board[T]) GetRightChar(p Position) T {
	if board.CanMoveRight(p) {
		return board.Get(p.First+1, p.Second)
	}

	return *new(T)
}
