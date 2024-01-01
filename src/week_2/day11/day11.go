package day11

import (
	"aoc-2023-golang/src/utils"
)

func SolvePart1(filename string) int {
	board := getBoard(filename)
	emptyRows := GetEmptyRows(&board)
	emptyColumns := GetEmptyColumns(&board)

	positions := utils.FindAllTargetInBoard(&board, "#")
	allPairs := CreatePairs(positions)

	sumDistance := 0
	for i := range allPairs {
		distance := CalculateDistanceBetweenPair(allPairs[i], emptyRows, emptyColumns, 2)
		sumDistance += distance
	}

	return sumDistance
}

func SolvePart2(filename string, expansion int) int {
	board := getBoard(filename)
	emptyRows := GetEmptyRows(&board)
	emptyColumns := GetEmptyColumns(&board)

	positions := utils.FindAllTargetInBoard(&board, "#")
	allPairs := CreatePairs(positions)

	sumDistance := 0
	for i := range allPairs {
		distance := CalculateDistanceBetweenPair(allPairs[i], emptyRows,
			emptyColumns, expansion)
		sumDistance += distance
	}

	return sumDistance
}

func getBoard(filename string) utils.Board[string] {
	rawLines := utils.ParseFile(filename)
	return utils.ParseRawLinesToBoard(rawLines)
}

func GetEmptyColumns(b *utils.Board[string]) []int {
	var emptyColumns []int
	for x := 0; x < b.Width; x++ {
		if IsEmpty(b.GetColumn(x)) {
			emptyColumns = append(emptyColumns, x)
		}
	}

	return emptyColumns
}

func GetEmptyRows(b *utils.Board[string]) []int {
	var emptyRows []int
	for y := 0; y < b.Height; y++ {
		if IsEmpty(b.GetRow(y)) {
			emptyRows = append(emptyRows, y)
		}
	}

	return emptyRows
}

func IsEmpty(values []string) bool {
	isEmpty := true
	for i := range values {
		if values[i] != "." {
			isEmpty = false
			break
		}
	}

	return isEmpty
}

func CreatePairs(values []utils.Position) []utils.Pair[utils.Position, utils.Position] {
	var pairs []utils.Pair[utils.Position, utils.Position]
	for i := 0; i < len(values); i++ {
		for y := i + 1; y < len(values); y++ {
			pairs = append(pairs, utils.Pair[utils.Position, utils.Position]{
				First:  values[i],
				Second: values[y],
			})
		}
	}

	return pairs
}

func CalculateDistanceBetweenPair(
	pair utils.Pair[utils.Position, utils.Position],
	emptyRows []int,
	emptyColumns []int,
	expansion int,
) int {
	var xFrom int
	if pair.First.First < pair.Second.First {
		xFrom = pair.First.First
	} else {
		xFrom = pair.Second.First
	}

	var xTo int
	if pair.First.First < pair.Second.First {
		xTo = pair.Second.First
	} else {
		xTo = pair.First.First
	}

	var yFrom int
	if pair.First.Second < pair.Second.Second {
		yFrom = pair.First.Second
	} else {
		yFrom = pair.Second.Second
	}

	var yTo int
	if pair.First.Second < pair.Second.Second {
		yTo = pair.Second.Second
	} else {
		yTo = pair.First.Second
	}

	// Crosses rows
	rowCrosses := CountMatches(xFrom, xTo, emptyColumns)

	// Crosses columns
	columnCrosses := CountMatches(yFrom, yTo, emptyRows)

	// calculate the distance
	xDiff := utils.AbsDiffInt(pair.First.First, pair.Second.First)
	yDiff := utils.AbsDiffInt(pair.First.Second, pair.Second.Second)

	// Adjust for the crosses
	adjustedXDiff := xDiff + (expansion-1)*rowCrosses
	adjustedYDiff := yDiff + (expansion-1)*columnCrosses

	return adjustedYDiff + adjustedXDiff
}

func CountMatches(from int, to int, values []int) int {
	count := 0
	for x := from; x <= to; x++ {
		if ExistsInSlice(x, values) {
			count += 1
		}
	}

	return count
}

func ExistsInSlice(target int, values []int) bool {
	isFound := false
	for _, value := range values {
		if target == value {
			isFound = true
			break
		}
	}

	return isFound
}
