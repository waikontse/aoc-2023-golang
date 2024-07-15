package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type Board[T comparable] struct {
	Board  [][]T
	Width  int
	Height int
}

func ParseInputIntoBoard(filename string) Board[string] {
	rawLines := ParseFile(filename)

	return ParseRawLinesToBoard(rawLines)
}

func ParseRawLinesToBoard(rawLines []string) Board[string] {
	width := len(rawLines[0])
	height := len(rawLines)

	board := CreateBoard[string](width, height)
	fmt.Printf("Creating board with: %d:%d \n", width, height)

	// Fill in the board
	for y := range rawLines {
		for x, char := range strings.Split(rawLines[y], "") {
			board.Set(x, y, char)
		}
	}

	return board
}

func ParseInputIntoIntBoard(filename string) Board[int] {
	rawLines := ParseFile(filename)

	return ParseRawLinesToIntBoard(rawLines)
}

func ParseRawLinesToIntBoard(rawLines []string) Board[int] {
	width := len(rawLines[0])
	height := len(rawLines)

	board := CreateBoard[int](width, height)
	fmt.Printf("Creating board with: %d:%d \n", width, height)

	// Fill in the board
	for y := range rawLines {
		for x, char := range strings.Split(rawLines[y], "") {
			val, _ := strconv.Atoi(char)
			board.Set(x, y, val)
		}
	}

	return board
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
	fmt.Println()
}

func (board *Board[T]) Get(x, y int) T {
	return board.Board[x][y]
}

func (board *Board[T]) Set(x, y int, value T) {
	board.Board[x][y] = value
}

func (board *Board[T]) SetRowRange(xFrom, xTo, y int, value T) {
	for x := xFrom; x <= xTo; x++ {
		board.Set(x, y, value)
	}
}

func (board *Board[T]) GetRow(height int) []T {
	row := make([]T, board.Width)
	for i := range board.Board {
		row[i] = board.Board[i][height]
	}

	return row
}

func (board *Board[T]) SetRow(row []T, rowHeight int) {
	for i := range row {
		board.Set(i, rowHeight, row[i])
	}
}

func (board *Board[T]) SetColumnRange(yFrom, yTo, x int, value T) {
	for y := yFrom; y <= yTo; y++ {
		board.Set(x, y, value)
	}
}

func (board *Board[T]) GetColumn(width int) []T {
	return board.Board[width]
}

func (board *Board[T]) CanMoveTop(p Position) bool {
	return p.Second > 0
}

func (board *Board[T]) GetTopChar(p Position) T {
	if board.CanMoveTop(p) {
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

func (board *Board[T]) CanMoveLeft(p Position) bool {
	return p.First > 0
}

func (board *Board[T]) GetLeftChar(p Position) T {
	if board.CanMoveLeft(p) {
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

func FindAllTargetInBoard(board *Board[string], target string) []Position {
	var foundPositions []Position
	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			if board.Get(x, y) == target {
				foundPositions = append(foundPositions, Position{First: x, Second: y})
			}
		}
	}

	return foundPositions
}

func FillBoardWithValue(board *Board[string], value string) {
	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			board.Set(x, y, value)
		}
	}
}
