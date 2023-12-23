package day10

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"os"
	"strings"
)

// Description
//| is a vertical pipe connecting north and south.
//- is a horizontal pipe connecting east and west.
//L is a 90-degree bend connecting north and east.
//J is a 90-degree bend connecting north and west.
//7 is a 90-degree bend connecting south and west.
//F is a 90-degree bend connecting south and east.
//. is ground; there is no pipe in this tile.
//S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

type TrackingInfo struct {
	StartingPos utils.Position
	PreviousPos utils.Position
	CurrentPos  utils.Position
	stepsTaken  int
}

func (trackingInfo *TrackingInfo) MoveToNewPosition(newPosition utils.Position) {
	trackingInfo.PreviousPos, trackingInfo.CurrentPos = trackingInfo.CurrentPos, newPosition
	trackingInfo.stepsTaken += 1
}

func SolvePart1(filename string) int {
	rawLines := utils.ParseFile(filename)
	board := ParseRawLinesToBoard(rawLines)
	startingPos := FindStartingPosition(&board)

	utils.Print(&board)
	fmt.Println(startingPos)

	trackingInfo := FindLargestLoop(&board)

	return trackingInfo.stepsTaken / 2
}

func ParseRawLinesToBoard(rawLines []string) utils.Board[string] {
	width := len(rawLines[0])
	height := len(rawLines)

	board := utils.CreateBoard[string](width, height)

	// Fill in the board
	for y := range rawLines {
		for x, char := range strings.Split(rawLines[y], "") {
			utils.Set(&board, x, y, char)
		}
	}

	return board
}

func FindStartingPosition(board *utils.Board[string]) utils.Position {
	var x, y int
	for y = 0; y < board.Height; y++ {
		row := utils.GetRow(board, y)
		for x = range row {
			if row[x] == "S" {
				return utils.Position{First: x, Second: y}
			}
		}
	}

	os.Exit(-1)
	return utils.Position{}
}

func FindLargestLoop(board *utils.Board[string]) TrackingInfo {
	// Walk till we come back to the starting position
	startPos := FindStartingPosition(board)
	trackingInfo := TrackingInfo{StartingPos: startPos, PreviousPos: startPos, CurrentPos: startPos}

	// 1. Get the new position to move
	// 2. Move to the new position
	// 3. Is new position same as starting position?
	// 3a. Yes -> stop
	// 3b. No -> Go back to step 1

	nextPosition := FindNextMove(board, &trackingInfo)
	trackingInfo.MoveToNewPosition(nextPosition)
	nextPosition = FindNextMove(board, &trackingInfo)
	for startPos != nextPosition {
		trackingInfo.MoveToNewPosition(nextPosition)
		nextPosition = FindNextMove(board, &trackingInfo)
	}
	trackingInfo.MoveToNewPosition(nextPosition)

	return trackingInfo
}

func FindNextMove(board *utils.Board[string], info *TrackingInfo) utils.Position {
	// Get previous position
	// Get current position
	currentChar := utils.Get(board, info.CurrentPos.First, info.CurrentPos.Second)
	topChar := utils.GetTopChar(board, info.CurrentPos)
	bottomChar := utils.GetBottomChar(board, info.CurrentPos)
	leftChar := utils.GetLeftChar(board, info.CurrentPos)
	rightChar := utils.GetRightChar(board, info.CurrentPos)

	var nextMove utils.Position
	if currentChar == "S" {
		if topChar == "|" || topChar == "7" || topChar == "F" {
			nextMove = utils.Position{First: info.CurrentPos.First, Second: info.CurrentPos.Second - 1}
			return nextMove
		}

		if bottomChar == "|" || bottomChar == "L" || bottomChar == "J" {
			nextMove = utils.Position{First: info.CurrentPos.First, Second: info.CurrentPos.Second + 1}
			return nextMove
		}

		if rightChar == "-" || rightChar == "J" || rightChar == "7" {
			nextMove = utils.Position{First: info.CurrentPos.First + 1, Second: info.CurrentPos.Second}
			return nextMove
		}

		if leftChar == "-" || leftChar == "L" || leftChar == "F" {
			nextMove = utils.Position{First: info.CurrentPos.First - 1, Second: info.CurrentPos.Second}
			return nextMove
		}
	} else if currentChar == "|" {
		// Can only go up or down, depending on previous node
		topPos := info.CurrentPos.GetTop()
		bottomPos := info.CurrentPos.GetBottom()

		if info.PreviousPos == topPos {
			return bottomPos
		}
		return topPos
	} else if currentChar == "-" {
		// Can only go left or right, depending on previous node
		leftPos := info.CurrentPos.GetLeft()
		rightPos := info.CurrentPos.GetRight()
		if info.PreviousPos == leftPos {
			return rightPos
		}
		return leftPos
	} else if currentChar == "L" {
		// Can only go up or right, depending on previous node
		topPos := info.CurrentPos.GetTop()
		rightPos := info.CurrentPos.GetRight()
		if info.PreviousPos == topPos {
			return rightPos
		}
		return topPos
	} else if currentChar == "J" {
		// Can only go up or left, depending on previous node
		topPos := info.CurrentPos.GetTop()
		leftPos := info.CurrentPos.GetLeft()
		if info.PreviousPos == topPos {
			return leftPos
		}
		return topPos
	} else if currentChar == "7" {
		// Can only go down or left, depending on previous node
		downPos := info.CurrentPos.GetBottom()
		leftPos := info.CurrentPos.GetLeft()
		if info.PreviousPos == downPos {
			return leftPos
		}
		return downPos
	} else if currentChar == "F" {
		// Can only go down or right, depending on previous node
		downPos := info.CurrentPos.GetBottom()
		rightPos := info.CurrentPos.GetRight()
		if info.PreviousPos == downPos {
			return rightPos
		}
		return downPos
	} else {
		fmt.Println("unexpected char: ", currentChar)
		os.Exit(-1)
	}

	return nextMove
}
