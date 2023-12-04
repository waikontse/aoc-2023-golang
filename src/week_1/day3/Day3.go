package day3

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"strconv"
	"unicode"
)

func SolvePart1() {
	rawLines := utils.ParseFile("./resources/day_3.txt")

	//utils.PrintLines(rawLines)

	// Find all the numbers
	var allNumbers = FindAllNumbers(rawLines)
	//fmt.Println(allNumbers)

	// Find all the Symbols
	var allSymbols = FindAllSymbols(rawLines)
	fmt.Println(allSymbols)

	var allAdjacentNumbers []FoundNumber
	for _, symbol := range allSymbols {
		for _, number := range allNumbers {
			if number.IsAdjacentToSymbol(symbol) {
				allAdjacentNumbers = append(allAdjacentNumbers, number)
			}
			// TODO might be that number is already used for an earlier symbol (used multiple times)
		}
	}

	result := utils.SumInts(utils.MapFunc1(allAdjacentNumbers, func(t FoundNumber) int {
		return t.value
	}),
	)

	fmt.Println("All adjacent numbers")
	//fmt.Println(allAdjacentNumbers)

	fmt.Println("Result: ", result)
}

func SolvePart2() {
	rawLines := utils.ParseFile("./resources/day_3.txt")

	//utils.PrintLines(rawLines)

	// Find all the numbers
	var allNumbers = FindAllNumbers(rawLines)
	fmt.Println(allNumbers)

	// Find all the Symbols
	var allSymbols = FindAllSymbols(rawLines)
	fmt.Println(allSymbols)

	// Find all numbers touching a symbol
	// Filter all symbol '*'
	var sum int64 = 0
	var foundNumbers []FoundNumber
	for _, symbol := range allSymbols {
		// Skip all symbols not equal to "*"
		if symbol.symbol != "*" {
			continue
		}

		fmt.Println("Checking for symbol: ", symbol)

		for _, number := range allNumbers {
			if number.IsAdjacentToSymbol(symbol) {
				fmt.Println("Appending number: ", number)
				foundNumbers = append(foundNumbers, number)
			}
		}

		fmt.Println("len of numbers found: ", len(foundNumbers))
		// check if there are only 2 numbers
		// if yes, multiply and sum
		// if no, continue
		if len(foundNumbers) == 2 {
			fmt.Println("Multiplying: ", foundNumbers[0].value, foundNumbers[1].value)
			//fmt.Println("Old sum: ", sum)
			sum = sum + int64(foundNumbers[0].value*foundNumbers[1].value)
			//fmt.Println("New sum: ", sum)
			foundNumbers = make([]FoundNumber, 0)
		} else {
			foundNumbers = make([]FoundNumber, 0)
		}

		//fmt.Println("Len of foundNumbers: ", len(foundNumbers))
	}

	fmt.Println("Found number: ", sum)

}

func FindAllSymbols(rawLines []string) []Point {
	var foundSymbols []Point
	for i, line := range rawLines {
		foundSymbols = append(foundSymbols, FindSymbolsInLIne(line, i)...)
	}

	return foundSymbols
}

func FindAllNumbers(rawLines []string) []FoundNumber {
	var foundNumbers []FoundNumber
	for i, line := range rawLines {
		foundNumbers = append(foundNumbers, FindNumbersInLine(line, i)...)
	}

	return foundNumbers
}

func FindNumbersInLine(line string, row int) []FoundNumber {
	var hasFound = false
	var currentFoundNumber FoundNumber
	var collectedFoundNumbers []FoundNumber
	var currentNumberString string
	for i, char := range line {
		if unicode.IsNumber(char) {
			// Create a new FoundNumber if not found
			if !hasFound {
				currentFoundNumber = FoundNumber{from: Point{x: i, y: row}}
				currentNumberString = currentNumberString + string(char)
				hasFound = true
			} else {
				// Continue with the number if hasFound == true
				currentNumberString = currentNumberString + string(char)
			}
		} else {
			// Finish the found number if hasFound == true
			if hasFound {
				hasFound, currentNumberString, collectedFoundNumbers = finishParseNumber(hasFound, currentFoundNumber,
					row, i-1, currentNumberString, collectedFoundNumbers)
			}
		}
	}

	//  line ends in number, can be true iff hasFound == true
	if hasFound == true {
		hasFound, currentNumberString, collectedFoundNumbers = finishParseNumber(hasFound, currentFoundNumber, row,
			len(line)-1, currentNumberString, collectedFoundNumbers)
	}

	return collectedFoundNumbers
}

func FindSymbolsInLIne(line string, row int) []Point {
	var foundSymbols []Point

	for i, char := range line {
		if unicode.IsNumber(char) || char == '.' {
			continue
		}

		foundSymbols = append(foundSymbols, Point{symbol: string(char), x: i, y: row})
	}

	return foundSymbols
}

func finishParseNumber(
	hasFound bool,
	currentFoundNumber FoundNumber,
	row int,
	i int,
	currentNumberString string,
	collectedFoundNumbers []FoundNumber,
) (bool, string, []FoundNumber) {
	hasFound = false
	currentFoundNumber.to = Point{x: i, y: row}
	value, _ := strconv.Atoi(currentNumberString)
	currentFoundNumber.value = value
	currentNumberString = ""
	collectedFoundNumbers = append(collectedFoundNumbers, currentFoundNumber)
	return hasFound, currentNumberString, collectedFoundNumbers
}

func (foundNumber FoundNumber) IsAdjacentToSymbol(symbol Point) bool {
	var isXPointAdjacent bool = symbol.x >= foundNumber.from.x-1 && symbol.x <= foundNumber.to.x+1
	var isYPointAdjacent bool = symbol.y >= foundNumber.from.y-1 && symbol.y <= foundNumber.to.y+1

	return isXPointAdjacent && isYPointAdjacent
}

type Point struct {
	symbol string
	x      int
	y      int
}

type FoundNumber struct {
	value int
	from  Point
	to    Point
}
