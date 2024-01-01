package day13

import (
	"aoc-2023-golang/src/utils"
	"fmt"
)

func SolvePart1(filename string) int {
	rawLines := utils.ParseFile(filename)
	boards := ParseRawLines(rawLines)
	predicate := func(val int) bool {
		return val == 0
	}

	var sum = 0
	for i := range boards {
		fmt.Println("Checking for for board: ", i)
		sum += DeterminePointsForBoard(&boards[i], 1, 100, predicate)
	}

	return sum
}

func SolvePart2(filename string) int {
	rawLines := utils.ParseFile(filename)
	boards := ParseRawLines(rawLines)
	predicate := func(val int) bool {
		return val == 1
	}

	var sum = 0
	for i := range boards {
		fmt.Println("Checking for for board: ", i)
		sum += DeterminePointsForBoard(&boards[i], 1, 100, predicate)
	}

	return sum
}

func DeterminePointsForBoard(
	b *utils.Board[string],
	multiplierForHorizontal int,
	multiplierForVertical int,
	pred func(val int) bool,
) int {
	pairHor, widthHor := FindLargestMirrorHorizontal(b, pred)
	pairVer, heightVer := FindLargestMirrorVertical(b, pred)

	sum := 0
	if widthHor > 0 {
		sum += pairHor.Second * multiplierForHorizontal
	}

	if heightVer > 0 {
		sum += pairVer.Second * multiplierForVertical
	}

	return sum
}

func ParseRawLines(rawLines []string) []utils.Board[string] {
	// Find all empty lines
	// Create the slices
	// Parse into board
	var previousIndexEmptyLine = -1
	var indexOfEmptyLine int
	parsedBoards := make([]utils.Board[string], 0)
	for indexOfEmptyLine = 0; indexOfEmptyLine < len(rawLines); indexOfEmptyLine++ {
		if rawLines[indexOfEmptyLine] == "" {
			boardSlice := rawLines[previousIndexEmptyLine+1 : indexOfEmptyLine]
			parsedBoards = append(parsedBoards, ParseRawLine(boardSlice))
			previousIndexEmptyLine = indexOfEmptyLine
		}
	}

	// parse the last one
	boardSlice := rawLines[previousIndexEmptyLine+1 : indexOfEmptyLine]
	parsedBoards = append(parsedBoards, ParseRawLine(boardSlice))

	return parsedBoards
}

func ParseRawLine(rawLine []string) utils.Board[string] {
	return utils.ParseRawLinesToBoard(rawLine)
}

func FindLargestMirrorHorizontal(b *utils.Board[string], pred func(val int) bool) (utils.Pair[int, int], int) {
	var maxWidth = 0
	var pairForMaxWidth utils.Pair[int, int]
	for i := 0; i < b.Width; i++ {
		leftMost := i
		rightMost := i + 1
		currentWidth := 0
		totalDiffs := 0

		for leftMost >= 0 && rightMost < b.Width {
			leftColumn := b.GetColumn(leftMost)
			rightColumn := b.GetColumn(rightMost)

			if !AreLinesSame(leftColumn, rightColumn) {
				totalDiffs += utils.DiffInSlices(leftColumn, rightColumn)
			}

			// update the indexes
			currentWidth += 1
			leftMost -= 1
			rightMost += 1
		}

		if pred(totalDiffs) && currentWidth > 0 {
			maxWidth = currentWidth
			pairForMaxWidth.First = i
			pairForMaxWidth.Second = i + 1
		}
	}

	return pairForMaxWidth, maxWidth
}

func FindLargestMirrorVertical(b *utils.Board[string], pred func(val int) bool) (utils.Pair[int, int], int) {
	var maxHeight = 0
	var pairForMaxHeight utils.Pair[int, int]
	for i := 0; i < b.Height; i++ {
		topMost := i
		bottomMost := i + 1
		currentHeight := 0
		totalDiffs := 0

		for topMost >= 0 && bottomMost < b.Height {
			topRow := b.GetRow(topMost)
			bottomRow := b.GetRow(bottomMost)

			if !AreLinesSame(topRow, bottomRow) {
				totalDiffs += utils.DiffInSlices(topRow, bottomRow)
			}

			// update the indexes
			currentHeight += 1
			topMost -= 1
			bottomMost += 1
		}

		// Check for largest width so far
		if pred(totalDiffs) && currentHeight > 0 {
			maxHeight = currentHeight
			pairForMaxHeight.First = i
			pairForMaxHeight.Second = i + 1
		}
	}

	return pairForMaxHeight, maxHeight
}

func AreLinesSame(left []string, right []string) bool {
	return utils.CompareSlices(left, right)
}
