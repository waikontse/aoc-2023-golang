package day14

import (
	"aoc-2023-golang/src/utils"
	"fmt"
)

func SolvePart1(filename string) int {
	board := ParseInputIntoBoard(filename)

	MoveNorth(&board)

	return CalculateTotalWeightForBoard(&board)
}

func SolvePart2(filename string) int {
	board := ParseInputIntoBoard(filename)

	sampleSize := 200
	samples := make([]int, sampleSize)
	for i := 0; i < sampleSize; i++ {
		MoveWholeCycle(&board)
		samples[i] = CalculateTotalWeightForBoard(&board)
	}

	start, periodicity, bucketSamples := FindPeriodicity(samples)
	bucket := (1000000000 - start) % (periodicity)

	return bucketSamples[bucket]
}

func ParseInputIntoBoard(filename string) utils.Board[string] {
	rawLines := utils.ParseFile(filename)

	return utils.ParseRawLinesToBoard(rawLines)
}

func MoveWholeCycle(b *utils.Board[string]) {
	MoveNorth(b)
	MoveWest(b)
	MoveSouth(b)
	MoveEast(b)
}

func MoveNorth(b *utils.Board[string]) {
	for i := 0; i < b.Width; i++ {
		MoveToStartOfRow(b.GetColumn(i))
	}
}

func MoveWest(b *utils.Board[string]) {
	for i := 0; i < b.Height; i++ {
		row := b.GetRow(i)
		MoveToStartOfRow(row)
		b.SetRow(row, i)
	}
}

func MoveSouth(b *utils.Board[string]) {
	for i := 0; i < b.Width; i++ {
		column := b.GetColumn(i)
		MoveToEndOfRow(column)
	}
}

func MoveEast(b *utils.Board[string]) {
	for i := 0; i < b.Height; i++ {
		row := b.GetRow(i)
		MoveToEndOfRow(row)
		b.SetRow(row, i)
	}
}

func MoveToStartOfRow(row []string) {
	for i := range row {
		// We can move the rock
		if row[i] == "O" {
			canMoveTo, isSameLocation := findNextAvailableLocationTowardsStart(row, i)
			if !isSameLocation {
				// Change the row in param
				row[canMoveTo] = "O"
				row[i] = "."
			}
		}
	}
}

func MoveToEndOfRow(row []string) {
	startLocation := len(row) - 1
	for i := range row {
		// We can move the rock
		currentLocation := startLocation - i
		if row[currentLocation] == "O" {
			canMoveTo, isSameLocation := findNextAvailableLocationTowardsEnd(row, currentLocation, len(row))
			if !isSameLocation {
				row[canMoveTo] = "O"
				row[currentLocation] = "."
			}
		}
	}
}

func findNextAvailableLocationTowardsStart(row []string, startLocation int) (int, bool) {
	location := 0
	for i := startLocation - 1; i >= 0; i-- {
		// stop when we reached a rock/cube
		if row[i] == "O" || row[i] == "#" {
			location = i + 1
			break
		}
	}

	return location, location == startLocation
}

func findNextAvailableLocationTowardsEnd(row []string, startLocation int, end int) (int, bool) {
	location := end - 1
	for i := startLocation + 1; i < end; i++ {
		// stop when we reached a rock/cube
		if row[i] == "O" || row[i] == "#" {
			location = i - 1
			break
		}
	}

	return location, location == startLocation
}

func CalculateTotalWeightForBoard(b *utils.Board[string]) int {
	sum := 0
	for i := 0; i < b.Height; i++ {
		sum += CalculateTotalWeightForRow(b.Height-i, b.GetRow(i))
	}

	return sum
}

func CalculateTotalWeightForRow(rowLoad int, row []string) int {
	return rowLoad * utils.Count(row, func(value string) bool {
		return value == "O"
	})
}

func FindPeriodicity(values []int) (int, int, []int) {
	periodicity := 0
	startIndex := 0

	for x := 0; x < len(values) && periodicity == 0; x++ {
		first := values[x]
		second := values[x+1]
		third := values[x+2]

		for i := x + 3; i < len(values)-2; i++ {
			if first == values[i] && second == values[i+1] && third == values[i+2] {
				fmt.Println("Found a pattern starting at: ", x+1)
				// found a pattern
				periodicity = i - x
				startIndex = x + 1
				break
			}
		}
	}

	return startIndex, periodicity, values[startIndex-1 : startIndex+periodicity-1]
}
