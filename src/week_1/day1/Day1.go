package day1

import (
	"aoc-2023-golang/src/utils"
	"cmp"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func solvePart1(filenameWithPath string) int {
	var lines = utils.ParseFile(filenameWithPath)

	newLines := extractNumbersFromLines(lines)

	var newInts []string
	for _, newLine := range newLines {
		newInts = append(newInts, convertToIntStr(newLine))
	}

	return utils.SumStrings(newInts)
}

func solvePart2(filenameWithPath string) int {
	var lines = utils.ParseFile(filenameWithPath)
	newLines := extractNumbersFromLines(lines)

	var newCorrectedInts []string
	for i, newLine := range newLines {
		newCorrectedInts = append(newCorrectedInts, correctCode(lines[i], newLine))
	}

	var newCorrectedInts2 []string
	for _, newLine := range newCorrectedInts {
		newCorrectedInts2 = append(newCorrectedInts2, convertToIntStr(newLine))
	}

	return utils.SumStrings(newCorrectedInts2)
}

func extractNumbersFromLines(lines []string) []string {
	var newLines []string
	for _, line := range lines {
		newLines = append(newLines, extractNumbersFromLine(line))
	}

	return newLines
}

func extractNumbersFromLine(str string) string {
	ignoreAlpha := func(r rune) rune {
		if unicode.IsNumber(r) {
			return r
		} else {
			return rune(-1)
		}
	}

	return strings.Map(ignoreAlpha, str)
}

func convertToIntStr(str string) string {
	if len(str) == 0 {
		log.Fatal("Gotten a string of 0 len")
	}

	if len(str) == 1 {
		return str + str
	}

	firstChar := string(str[0])
	lastChar := string(str[len(str)-1])

	return firstChar + lastChar
}

type Pair struct {
	value int
	index int
}

func extractWordsFromLine(original string) []Pair {
	var numbers = []string{One, Two, Three, Four, Five, Six, Seven, Eight, Nine}

	// Parse words in the original string
	var parsedWords []Pair
	for i, numberStr := range numbers {
		if strings.Contains(original, numberStr) {
			var newPairFirst = Pair{value: i + 1, index: strings.Index(original, numberStr)}
			var newPairLast = Pair{value: i + 1, index: strings.LastIndex(original, numberStr)}
			parsedWords = append(parsedWords, newPairFirst)
			parsedWords = append(parsedWords, newPairLast)
		}
	}

	// Sort the parsed words based on index found
	sort.Slice(parsedWords, func(i, j int) bool {
		return parsedWords[i].index < parsedWords[j].index
	})

	return parsedWords
}

func correctCode(original string, parsedNumbers string) string {
	parsedWords := extractWordsFromLine(original)
	// Just return if no words found
	if len(parsedWords) == 0 {
		return parsedNumbers
	}

	// Determine the 1st and last parsedNumbers of the string
	// If string is empty
	// If string is not empty
	var correctedNumbers = parsedNumbers
	if len(parsedNumbers) == 0 {
		for _, e := range parsedWords {
			correctedNumbers += strconv.Itoa(e.value)
		}
	} else {
		// index of 1st parsedNumbers
		var indexOfFirstParsedNumber = strings.Index(original, string(parsedNumbers[0]))
		var earliestParsedWord = slices.MinFunc(parsedWords, func(a, b Pair) int {
			return cmp.Compare(a.index, b.index)
		})

		if earliestParsedWord.index < indexOfFirstParsedNumber {
			correctedNumbers = strconv.Itoa(earliestParsedWord.value) + correctedNumbers
		}

		// index of last parsedNumbers
		var indexOfSecondParsedNumber = strings.LastIndex(original, string(parsedNumbers[len(parsedNumbers)-1]))
		var latestParsedWord = slices.MaxFunc(parsedWords, func(a, b Pair) int {
			return cmp.Compare(a.index, b.index)
		})

		if latestParsedWord.index > indexOfSecondParsedNumber {
			correctedNumbers = correctedNumbers + strconv.Itoa(latestParsedWord.value)
		}
	}

	return correctedNumbers
}

const (
	One   string = "one"
	Two          = "two"
	Three        = "three"
	Four         = "four"
	Five         = "five"
	Six          = "six"
	Seven        = "seven"
	Eight        = "eight"
	Nine         = "nine"
)
