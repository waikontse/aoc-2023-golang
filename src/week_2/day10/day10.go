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
	board       utils.Board[string]
}

func (trackingInfo *TrackingInfo) MoveToNewPosition(newPosition utils.Position) {
	trackingInfo.PreviousPos, trackingInfo.CurrentPos = trackingInfo.CurrentPos, newPosition
	trackingInfo.stepsTaken += 1

	// Update the board
	trackingInfo.board.Set(newPosition.First, newPosition.Second, "*")
}

func SolvePart1(filename string) int {
	board := getBoard(filename)
	//startingPos := FindStartingPosition(&board)

	//utils.Print(&board)
	//fmt.Println(startingPos)

	trackingInfo := FindLargestLoop(&board, filename)

	return trackingInfo.stepsTaken / 2
}

func SolvePart2(filename string) int {
	board := getBoard(filename)
	trackingInfo := FindLargestLoop(&board, filename)
	cleanBoard(&trackingInfo.board, &board)
	//	utils.Print(&trackingInfo.board)
	foundDotLocations := utils.FindAllTargetInBoard(&trackingInfo.board, ".")
	foundDollarLocations := utils.FindAllTargetInBoard(&trackingInfo.board, "$")
	//utils.Print(&board)
	//fmt.Println("Found targets: ", len(foundDotLocations))

	// Find all the '.' in the board
	// Determine if '.' is inside the main loop
	dotPositionsWithinLoop := GetPositionsWithinLoop(foundDotLocations,
		&trackingInfo.board)
	dollarLocationsInsideLoop := GetPositionsWithinLoop(foundDollarLocations,
		&trackingInfo.board)

	fmt.Println("Found ground inside loop: ", len(dotPositionsWithinLoop))
	fmt.Println("Found loose pipes inside loop: ", len(dollarLocationsInsideLoop))

	//for _, p := range dollarLocationsInsideLoop {
	//	trackingInfo.board.Set(p.First, p.Second, ".")
	//}

	mergedPositionsWithinLoop := append(dotPositionsWithinLoop, dollarLocationsInsideLoop...)

	// Flood fill the map per position
	for i := range mergedPositionsWithinLoop {
		FloodFill(&trackingInfo.board, mergedPositionsWithinLoop[i], "^",
			func(target string) bool {
				return target == "." || target == "$"
			})
	}

	unFloodFillUnclosed(&trackingInfo.board)
	utils.Print(&trackingInfo.board)

	return len(utils.FindAllTargetInBoard(&trackingInfo.board, "^"))
}

func GetPositionsWithinLoop(locations []utils.Position, b *utils.Board[string]) []utils.Position {
	positionsInsideLoop := utils.MapFunc1(locations, func(p utils.Position) bool {
		return IsPositionInsideLoop(b, p)
	})
	positionsWithFound := utils.ZipWith(locations, positionsInsideLoop)
	positionsWithFoundTrue := utils.Filter(positionsWithFound, func(value utils.Pair[utils.Position, bool]) bool {
		return value.Second
	})

	return utils.MapFunc1(positionsWithFoundTrue, func(v utils.Pair[utils.Position, bool]) utils.Position {
		return v.First
	})
}

func cleanBoard(cleanMap *utils.Board[string], dirtyMap *utils.Board[string]) {
	// Clean all the remaining dangling pippes
	for y := 0; y < cleanMap.Height; y++ {
		for x := 0; x < cleanMap.Width; x++ {
			if cleanMap.Get(x, y) != "*" {
				if cleanMap.Get(x, y) != "." {
					// All the "." remain there, all the dangling pipes are
					// changed to $$
					cleanMap.Set(x, y, "$")
				}
			}
		}
	}

	// Combine both the cleanMap and the dirty map, so we can get a map with the main loop only
	for y := 0; y < cleanMap.Height; y++ {
		for x := 0; x < cleanMap.Width; x++ {
			if cleanMap.Get(x, y) == "*" {
				targetChar := dirtyMap.Get(x, y)
				cleanMap.Set(x, y, targetChar)
			}
		}
	}
}

func getBoard(filename string) utils.Board[string] {
	rawLines := utils.ParseFile(filename)
	return ParseRawLinesToBoard(rawLines)
}

func ParseRawLinesToBoard(rawLines []string) utils.Board[string] {
	width := len(rawLines[0])
	height := len(rawLines)

	board := utils.CreateBoard[string](width, height)
	fmt.Printf("Creating board with: %d:%d \n", width, height)

	// Fill in the board
	for y := range rawLines {
		for x, char := range strings.Split(rawLines[y], "") {
			board.Set(x, y, char)
		}
	}

	return board
}

func FindStartingPosition(board *utils.Board[string]) utils.Position {
	var x, y int
	for y = 0; y < board.Height; y++ {
		row := board.GetRow(y)
		for x = range row {
			if row[x] == "S" {
				return utils.Position{First: x, Second: y}
			}
		}
	}

	os.Exit(-1)
	return utils.Position{}
}

func FindLargestLoop(board *utils.Board[string], filename string) TrackingInfo {
	// Walk till we come back to the starting position
	startPos := FindStartingPosition(board)

	trackingInfo := TrackingInfo{
		StartingPos: startPos,
		PreviousPos: startPos,
		CurrentPos:  startPos,
		board:       getBoard(filename),
	}

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
	currentChar := board.Get(info.CurrentPos.First, info.CurrentPos.Second)
	topChar := board.GetTopChar(info.CurrentPos)
	bottomChar := board.GetBottomChar(info.CurrentPos)
	leftChar := board.GetLeftChar(info.CurrentPos)
	rightChar := board.GetRightChar(info.CurrentPos)

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

func IsPositionInsideLoop(b *utils.Board[string], p utils.Position) bool {
	column := b.GetColumn(p.First)
	row := b.GetRow(p.Second)

	// Check bottom has walls
	bottomSlice := column[p.Second+1:]
	bottomHits := CountTargetInList(bottomSlice, []string{"S", "F", "7", "L", "J", "-"})
	bottomIsOk := bottomHits%2 == 1

	// check top has walls
	topSlice := column[0:p.Second]
	topHits := CountTargetInList(topSlice, []string{"S", "F", "7", "L", "J", "-"})
	topIsOk := topHits%2 == 1

	// check left has walls
	leftSlice := row[0:p.First]
	leftHits := CountTargetInList(leftSlice, []string{"S", "F", "7", "L", "J", "|"})
	leftIsOk := leftHits%2 == 1

	// check right has walls
	rightSlice := row[p.First+1:]
	rightHits := CountTargetInList(rightSlice, []string{"S", "F", "7", "L", "J", "|"})
	rightIsOk := rightHits%2 == 1

	//return bottomIsOk && topIsOk && leftIsOk && rightIsOk
	isAlOk := (bottomHits > 0 && topHits > 0 && leftHits > 0 && rightHits > 0) &&
		(bottomIsOk || topIsOk || leftIsOk || rightIsOk)

	//if isAlOk {
	//	fmt.Println("Bottom hits: ", bottomHits)
	//	fmt.Println("top hits: ", topHits)
	//	fmt.Println("left hits: ", leftHits)
	//	fmt.Println("right hits: ", rightHits)
	//
	//}

	return isAlOk
}

func CountTargetInList(list []string, targets []string) int {
	var count int = 0
	for i := range list {
		for _, target := range targets {
			if list[i] == target {
				count += 1
			}
		}
	}

	return count
}

func FloodFill(
	b *utils.Board[string],
	p utils.Position,
	newChar string,
	pred func(target string) bool,
) {
	nextPositionsToVisit := make([]utils.Position, 0)
	nextPositionsToVisit = append(nextPositionsToVisit, p)

	//	fmt.Println("Floodfilling: ", p)

	for len(nextPositionsToVisit) != 0 {
		// pop item
		nextPosition := nextPositionsToVisit[0]
		nextPositionsToVisit = nextPositionsToVisit[1:]

		//		fmt.Println("filling next position: ", nextPosition)
		if b.Get(nextPosition.First, nextPosition.Second) == newChar {
			continue
		}

		// Color current position to 'newChar'
		b.Set(nextPosition.First, nextPosition.Second, newChar)

		// Add next items to color
		topChar := b.GetTopChar(nextPosition)
		// topChar == "." || topChar == "$" {
		if pred(topChar) {
			nextPositionsToVisit = append(nextPositionsToVisit, utils.Position{First: nextPosition.First, Second: nextPosition.Second - 1})
		}

		bottomChar := b.GetBottomChar(nextPosition)
		if pred(bottomChar) {
			nextPositionsToVisit = append(nextPositionsToVisit, utils.Position{First: nextPosition.First, Second: nextPosition.Second + 1})
		}

		leftChar := b.GetLeftChar(nextPosition)
		if pred(leftChar) {
			nextPositionsToVisit = append(nextPositionsToVisit, utils.Position{First: nextPosition.First - 1, Second: nextPosition.Second})
		}

		rightChar := b.GetRightChar(nextPosition)
		if pred(rightChar) {
			nextPositionsToVisit = append(nextPositionsToVisit, utils.Position{First: nextPosition.First + 1, Second: nextPosition.Second})
		}
	}
}

func unFloodFillUnclosed(b *utils.Board[string]) {
	// Keep un flood filling
	position, hasFound := findUnclosedSymbol(b, "^")
	for hasFound {
		fmt.Println("un-floodfilling: ", position)
		// Flood fill the position with a symbol
		FloodFill(b, position, "#",
			func(target string) bool {
				return target == "^"
			})

		// update the result
		position, hasFound = findUnclosedSymbol(b, "^")
	}
}

func findUnclosedSymbol(b *utils.Board[string], symbol string) (utils.Position, bool) {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			if b.Get(x, y) == symbol {
				potentialPosition := utils.Position{First: x, Second: y}
				if IsUnclosed(b, potentialPosition) {
					return potentialPosition, true
				}
			}
		}
	}

	return utils.Position{}, false
}

func IsUnclosed(b *utils.Board[string], p utils.Position) bool {
	column := b.GetColumn(p.First)
	row := b.GetRow(p.Second)

	// Check bottom has walls
	bottomSlice := column[p.Second+1:]
	bottomHits := CountTargetInList(bottomSlice, []string{"S", "F", "7", "L", "J", "-"})

	// check top has walls
	topSlice := column[0:p.Second]
	topHits := CountTargetInList(topSlice, []string{"S", "F", "7", "L", "J", "-"})

	// check left has walls
	leftSlice := row[0:p.First]
	leftHits := CountTargetInList(leftSlice, []string{"S", "F", "7", "L", "J", "|"})

	// check right has walls
	rightSlice := row[p.First+1:]
	rightHits := CountTargetInList(rightSlice, []string{"S", "F", "7", "L", "J", "|"})

	return bottomHits == 0 || topHits == 0 || leftHits == 0 || rightHits == 0
}
