package day9

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"log"
	"slices"
	"strings"
)

func SolvePart1(filename string) int {
	rawLines := utils.ParseFile(filename)
	rows := ParseRawLinesToRows(rawLines)

	var sum = 0
	for i := range rows {
		allDiffRowsForI := GenerateAllDiffRows(rows[i])
		ExpandRows(allDiffRowsForI, 1)

		adding := allDiffRowsForI[0][len(allDiffRowsForI[0])-1]
		fmt.Println("Adding: ", adding)

		sum += adding

		ControlDiffRows(allDiffRowsForI)
	}

	return sum
}

func SolvePart2(filename string) int {
	rawLines := utils.ParseFile(filename)
	rows := ParseRawLinesToRows(rawLines)

	var sum = 0
	for i := range rows {
		slices.Reverse(rows[i])
		allDiffRowsForI := GenerateAllDiffRows(rows[i])
		ExpandRows(allDiffRowsForI, 1)

		adding := allDiffRowsForI[0][len(allDiffRowsForI[0])-1]
		fmt.Println("Adding: ", adding)

		sum += adding

		ControlDiffRows(allDiffRowsForI)
	}

	return sum
}

func ExpandRows(allDiffRows [][]int, expandCount int) {
	for i := 0; i < expandCount; i++ {

		// Start from last row, and expand it with the number from previous row
		currentExpander := 0
		for i := len(allDiffRows) - 1; i >= 0; i-- {
			lastNum := allDiffRows[i][len(allDiffRows[i])-1]
			currentExpander += lastNum

			// Add the number to the last position of current row
			allDiffRows[i] = append(allDiffRows[i], currentExpander)
		}
	}
}

func ParseRawLinesToRows(rawLines []string) [][]int {
	rows := make([][]int, 0)

	for i := range rawLines {
		parsedLine := ParseRawLineToRow(rawLines[i])
		rows = append(rows, parsedLine)
	}

	return rows
}

func ParseRawLineToRow(rawLine string) []int {
	trimmed := strings.TrimSpace(rawLine)
	split := strings.Split(trimmed, " ")

	return utils.ToIntSlice(split)
}

func GenerateAllDiffRows(row []int) [][]int {
	currentDiffs := row
	allDiffRows := make([][]int, 0)
	allDiffRows = append(allDiffRows, row)

	for !utils.AllIntsAreZero(currentDiffs) {
		currentDiffs = GenerateDiffRow(currentDiffs)
		allDiffRows = append(allDiffRows, currentDiffs)
	}

	return allDiffRows
}

func GenerateDiffRow(row []int) []int {
	var diffs []int
	for i := 0; i < len(row)-1; i++ {
		diff := row[i+1] - row[i]
		diffs = append(diffs, diff)
	}

	return diffs
}

func ControlDiffRows(diffRows [][]int) {
	slices.Reverse(diffRows)

	var sum = 0
	for i := range diffRows {
		sum += diffRows[i][len(diffRows[i])-3]
	}

	lastRow := len(diffRows) - 1
	control := diffRows[lastRow][len(diffRows[lastRow])-2]

	if control != sum {
		log.Fatal("Expected but incorrect. ", control, sum)
	}
}
