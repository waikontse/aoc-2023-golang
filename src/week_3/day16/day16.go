package day16

import "aoc-2023-golang/src/utils"

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Position utils.Triple[int, int, Direction]

type TrackingInfo struct {
	Board            utils.Board[string]
	Copy             utils.Board[string]
	CurrentPositions []Position
}

func SolvePart1(filename string) int {
	board := utils.ParseInputIntoBoard(filename)
	copy := utils.CreateBoard[string](board.Width, board.Height)
	utils.FillBoardWithValue(&copy, ".")

	trackingInfo := TrackingInfo{
		Board:            board,
		Copy:             copy,
		CurrentPositions: make([]Position, 0),
	}

	utils.Print(&trackingInfo.Board)
	utils.Print(&trackingInfo.Copy)

	return 0
}

func SolvePart2() int {
	return 0
}

func moveAndMark(info *TrackingInfo) {
	// For every position in the list, check if we can move 1 position
}

func move(info *TrackingInfo, currentPosition Position) (int, Position) {
	switch currentPosition.Third {
	case UP:
		break
	case DOWN:
		break
	case LEFT:
		break
	case RIGHT:
		break
	}
}

func moveUp(info *TrackingInfo, currentPosition Position) (int, Position) {
	return -1, Position{}
}

func moveDown(info *TrackingInfo, currentPosition Position) (int, Position) {
	return -1, Position{}
}

func moveLeft(info *TrackingInfo, currentPosition Position) (int, Position) {
	return -1, Position{}
}

func moveRight(info *TrackingInfo, currentPosition Position) (int, Position) {
	return -1, Position{}
}
