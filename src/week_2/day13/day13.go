package day13

import (
	"aoc-2023-golang/src/utils"
	"fmt"
)

func SolvePart1(filename string) int {
	rawLines := utils.ParseFile(filename)

	boards := ParseRawLines(rawLines)

	//for i := range boards {
	//	utils.Print(&boards[i])
	//}

	//pair, width := FindLargestMirrorHorizontal(&boards[0])
	//fmt.Println("width: ", width, pair)
	//
	//pairVert, height := FindLargestMirrorVertical(&boards[1])
	//fmt.Println("height: ", height, pairVert)

	var sum int = 0
	for i := range boards {
		fmt.Println("Checking for for board: ", i)
		sum += DeterminePointsForBoard(&boards[i], 1, 100)
	}

	return sum
}

func SolvePart2(filename string) int {
	return 0
}

func DeterminePointsForBoard(b *utils.Board[string], multiplierForHorizontal int, multiplierForVertical int) int {
	pairHor, widthHor := FindLargestMirrorHorizontal(b)
	pairVer, heightVer := FindLargestMirrorVertical(b)

	fmt.Println("Horizontal: ", widthHor, pairHor)
	fmt.Println("Vertical: ", heightVer, pairVer)

	// What happens when
	//if widthHor > heightVer {
	//	fmt.Println("Horizontal won")
	//	return pairHor.Second * multiplierForHorizontal
	//} else {
	//	fmt.Println("Vertical won")
	//	return pairVer.Second * multiplierForVertical
	//}

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

	var previousIndexEmptyLine int = -1
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

func FindLargestMirrorHorizontal(b *utils.Board[string]) (utils.Pair[int, int], int) {
	var maxWidth int = 0
	var pairForMaxWidth utils.Pair[int, int]
	for i := 0; i < b.Width; i++ {
		leftMost := i
		rightMost := i + 1
		currentWidth := 0

		for leftMost >= 0 && rightMost < b.Width {
			leftColumn := b.GetColumn(leftMost)
			rightColumn := b.GetColumn(rightMost)

			//	fmt.Printf("Checking: %d %d\n", leftMost, rightMost)

			if AreLinesSame(leftColumn, rightColumn) {
				currentWidth += 1
			}

			// update the indexes
			leftMost -= 1
			rightMost += 1
		}

		//fmt.Printf("end Checking: %d %d\n", leftMost, rightMost)
		if (leftMost == -1 || rightMost == b.Width) && currentWidth > 0 {
			//	fmt.Println("settings column")
			maxWidth = currentWidth
			pairForMaxWidth.First = i
			pairForMaxWidth.Second = i + 1
		}
	}

	return pairForMaxWidth, maxWidth
}

func FindLargestMirrorVertical(b *utils.Board[string]) (utils.Pair[int, int], int) {
	var maxHeight int = 0
	var pairForMaxHeight utils.Pair[int, int]
	for i := 0; i < b.Height; i++ {
		topMost := i
		bottomMost := i + 1
		currentHeight := 0

		for topMost >= 0 && bottomMost < b.Height {
			topRow := b.GetRow(topMost)
			bottomRow := b.GetRow(bottomMost)

			//	fmt.Printf("Checking: %d %d\n", topMost, bottomMost)

			if AreLinesSame(topRow, bottomRow) {
				currentHeight += 1
			} else {
				// stop searching if not matching.
				break
			}

			// update the indexes
			topMost -= 1
			bottomMost += 1
		}

		// Check for largest width so far
		//fmt.Println("topmost: ", topMost, bottomMost)
		if (topMost == -1 || bottomMost == b.Height) && currentHeight > 0 {
			//	fmt.Println("setting vertical")
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
