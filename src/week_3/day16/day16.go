package day16

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"os"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type PositionWithDirection utils.Pair[utils.Position, Direction]

type TrackingInfo struct {
	Board            utils.Board[string]
	MarkedBoard      utils.Board[string]
	CurrentPositions []PositionWithDirection
	HasSeen          map[PositionWithDirection]bool
}

func SolvePart1(filename string) int {
	trackingInfo := initializeTrackingInfo(filename, utils.Position{}, RIGHT)

	moveAndMark(&trackingInfo)

	return len(utils.FindAllTargetInBoard(&trackingInfo.MarkedBoard, "#"))
}

func SolvePart2(filename string) int {
	max := 0
	board := utils.ParseInputIntoBoard(filename)

	// Find the max for the whole board
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			// check if x, y is an edge
			if y == 0 || y == board.Height-1 {
				if x == 0 || x == board.Width-1 {
					// try up
					newStartPosition := utils.Position{x, y}
					trackingInfo := initializeTrackingInfo(filename, newStartPosition, UP)
					moveAndMark(&trackingInfo)
					max = utils.Max(max, len(utils.FindAllTargetInBoard(&trackingInfo.MarkedBoard, "#")))

					// Try down
					trackingInfo = initializeTrackingInfo(filename, newStartPosition, DOWN)
					moveAndMark(&trackingInfo)
					max = utils.Max(max, len(utils.FindAllTargetInBoard(&trackingInfo.MarkedBoard, "#")))

					// try left
					trackingInfo = initializeTrackingInfo(filename, newStartPosition, LEFT)
					moveAndMark(&trackingInfo)
					max = utils.Max(max, len(utils.FindAllTargetInBoard(&trackingInfo.MarkedBoard, "#")))

					// try right
					trackingInfo = initializeTrackingInfo(filename, newStartPosition, RIGHT)
					moveAndMark(&trackingInfo)
					max = utils.Max(max, len(utils.FindAllTargetInBoard(&trackingInfo.MarkedBoard, "#")))

				} else {
					// we have to topmost or bottom most row
					if y == 0 {
						// we can only go down
						newStartPosition := utils.Position{x, y}
						trackingInfo := initializeTrackingInfo(filename, newStartPosition, DOWN)
						moveAndMark(&trackingInfo)
						max = utils.Max(max, len(utils.FindAllTargetInBoard(&trackingInfo.MarkedBoard, "#")))
					} else if y == board.Height-1 {
						// we can only go up
						newStartPosition := utils.Position{x, y}
						trackingInfo := initializeTrackingInfo(filename, newStartPosition, UP)
						moveAndMark(&trackingInfo)
						max = utils.Max(max, len(utils.FindAllTargetInBoard(&trackingInfo.MarkedBoard, "#")))
					}
				}
			} else {
				// Depending on X, we can only go left of right
				if x == 0 {
					// we can only go right
					newStartPosition := utils.Position{x, y}
					trackingInfo := initializeTrackingInfo(filename, newStartPosition, RIGHT)
					moveAndMark(&trackingInfo)
					max = utils.Max(max, len(utils.FindAllTargetInBoard(&trackingInfo.MarkedBoard, "#")))
				} else if x == board.Width-1 {
					// we can only go left
					newStartPosition := utils.Position{x, y}
					trackingInfo := initializeTrackingInfo(filename, newStartPosition, LEFT)
					moveAndMark(&trackingInfo)
					max = utils.Max(max, len(utils.FindAllTargetInBoard(&trackingInfo.MarkedBoard, "#")))
				}
			}

		}
	}

	return max
}

func initializeTrackingInfo(filename string, startPosition utils.Position, direction Direction) TrackingInfo {
	board := utils.ParseInputIntoBoard(filename)
	copyOfBord := utils.CreateBoard[string](board.Width, board.Height)
	utils.FillBoardWithValue(&copyOfBord, ".")

	trackingInfo := TrackingInfo{
		Board:            board,
		MarkedBoard:      copyOfBord,
		CurrentPositions: make([]PositionWithDirection, 0),
		HasSeen:          make(map[PositionWithDirection]bool),
	}

	trackingInfo.CurrentPositions = append(trackingInfo.CurrentPositions, PositionWithDirection{
		First:  startPosition,
		Second: direction,
	})

	return trackingInfo
}

func moveAndMark(info *TrackingInfo) {
	// For every position in the list, check if we can move 1 position
	for len(info.CurrentPositions) != 0 {
		// Get current positions
		var currentPosition = info.CurrentPositions[0]
		info.CurrentPositions = info.CurrentPositions[1:]

		// Skip if we have already seen the place and the same direction
		_, hasSeen := info.HasSeen[currentPosition]
		if hasSeen {
			continue
		}

		// Mark the current location
		position := currentPosition.First
		info.MarkedBoard.Set(position.First, position.Second, "#")
		info.HasSeen[currentPosition] = true

		move(info, currentPosition)
	}

	// 1. Mark the current location
	// 2. Make a move from the current location

	// 3. Check the return value of the move method
	// Do nothing
	// Add 1 location
	// Add 2 locations

}

func move(info *TrackingInfo, currentPosition PositionWithDirection) {
	switch currentPosition.Second {
	case UP:
		moveUp(info, currentPosition)
		break
	case DOWN:
		moveDown(info, currentPosition)
		break
	case LEFT:
		moveLeft(info, currentPosition)
		break
	case RIGHT:
		moveRight(info, currentPosition)
		break
	}
}

func moveUp(info *TrackingInfo, currentPosition PositionWithDirection) {
	position := currentPosition.First
	currentPositionChar := info.Board.Get(position.First, position.Second)

	if currentPositionChar == "." {
		tryMoveUp(info, currentPosition)
	} else if currentPositionChar == "-" {
		// Can go left?
		tryMoveLeft(info, currentPosition)
		// Can go right?
		tryMoveRight(info, currentPosition)
	} else if currentPositionChar == "|" {
		tryMoveUp(info, currentPosition)
	} else if currentPositionChar == "/" {
		// Can go right?
		tryMoveRight(info, currentPosition)
	} else if currentPositionChar == "\\" {
		// Can go left?
		tryMoveLeft(info, currentPosition)
	} else {
		fmt.Print("Unexpected char: " + currentPositionChar)
		os.Exit(-1)
	}
}

func moveDown(info *TrackingInfo, currentPosition PositionWithDirection) {
	position := currentPosition.First
	currentPositionChar := info.Board.Get(position.First, position.Second)

	if currentPositionChar == "." {
		tryMoveDown(info, currentPosition)
	} else if currentPositionChar == "-" {
		// Can move left?
		tryMoveLeft(info, currentPosition)

		// Can move right?
		tryMoveRight(info, currentPosition)
	} else if currentPositionChar == "|" {
		tryMoveDown(info, currentPosition)
	} else if currentPositionChar == "/" {
		// Can move left?
		tryMoveLeft(info, currentPosition)
	} else if currentPositionChar == "\\" {
		// Can move right?
		tryMoveRight(info, currentPosition)
	} else {
		fmt.Print("Unexpected char: " + currentPositionChar)
		os.Exit(-1)
	}
}

func moveLeft(info *TrackingInfo, currentPosition PositionWithDirection) {
	position := currentPosition.First
	currentPositionChar := info.Board.Get(position.First, position.Second)

	if currentPositionChar == "." {
		tryMoveLeft(info, currentPosition)
	} else if currentPositionChar == "-" {
		tryMoveLeft(info, currentPosition)
	} else if currentPositionChar == "|" {
		tryMoveUp(info, currentPosition)
		tryMoveDown(info, currentPosition)
	} else if currentPositionChar == "/" {
		tryMoveDown(info, currentPosition)
	} else if currentPositionChar == "\\" {
		tryMoveUp(info, currentPosition)
	} else {
		fmt.Print("Unexpected char: " + currentPositionChar)
		os.Exit(-1)
	}
}

func moveRight(info *TrackingInfo, currentPosition PositionWithDirection) {
	position := currentPosition.First
	currentPositionChar := info.Board.Get(position.First, position.Second)

	if currentPositionChar == "." {
		tryMoveRight(info, currentPosition)
	} else if currentPositionChar == "-" {
		tryMoveRight(info, currentPosition)
	} else if currentPositionChar == "|" {
		tryMoveUp(info, currentPosition)
		tryMoveDown(info, currentPosition)
	} else if currentPositionChar == "/" {
		tryMoveUp(info, currentPosition)
	} else if currentPositionChar == "\\" {
		tryMoveDown(info, currentPosition)
	} else {
		fmt.Print("Unexpected char: " + currentPositionChar)
		os.Exit(-1)
	}
}

func tryMoveLeft(info *TrackingInfo, currentPosition PositionWithDirection) {
	position := currentPosition.First
	// Can go left?
	if info.Board.CanMoveLeft(position) {
		newPosition := PositionWithDirection{
			First:  utils.Position{First: position.First - 1, Second: position.Second},
			Second: LEFT,
		}
		info.CurrentPositions = append(info.CurrentPositions, newPosition)
	}
}

func tryMoveRight(info *TrackingInfo, currentPosition PositionWithDirection) {
	position := currentPosition.First

	if info.Board.CanMoveRight(position) {
		newPosition := PositionWithDirection{
			First:  utils.Position{First: position.First + 1, Second: position.Second},
			Second: RIGHT,
		}
		info.CurrentPositions = append(info.CurrentPositions, newPosition)
	}
}

func tryMoveUp(info *TrackingInfo, currentPosition PositionWithDirection) {
	position := currentPosition.First

	if info.Board.CanMoveTop(position) {
		newPosition := PositionWithDirection{
			First:  utils.Position{First: position.First, Second: position.Second - 1},
			Second: UP,
		}
		info.CurrentPositions = append(info.CurrentPositions, newPosition)
	}
}

func tryMoveDown(info *TrackingInfo, currentPosition PositionWithDirection) {
	position := currentPosition.First

	if info.Board.CanMoveBottom(position) {
		newPosition := PositionWithDirection{
			First:  utils.Position{First: position.First, Second: position.Second + 1},
			Second: DOWN,
		}
		info.CurrentPositions = append(info.CurrentPositions, newPosition)
	}
}
