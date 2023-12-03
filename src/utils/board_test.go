package utils

import "testing"

func TestPrint(t *testing.T) {
	board := CreateBoard[int](3, 3)

	Set(board, 0, 0, 1)
	Set(board, 0, 1, 2)
	Set(board, 0, 2, 3)

	Set(board, 1, 0, 4)
	Set(board, 1, 1, 5)
	Set(board, 1, 2, 6)

	Set(board, 2, 0, 7)
	Set(board, 2, 1, 8)
	Set(board, 2, 2, 9)

	Print(board)
}
