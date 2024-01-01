package day12

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"regexp"
	"strings"
)

type Puzzle struct {
	Gear      string
	Positions []int
}

func SolvePart1(filename string) int {
	puzzles := ParseRawLines(utils.ParseFile(filename))

	fmt.Println("Puzzles: ", puzzles)
	for i := range puzzles {
		fmt.Println("splitted fillers: ", SplitIntoFillerSections(puzzles[i].Gear))
	}

	return 0
}

func SolvePart2(filename string) int {
	return 0
}

func ParseRawLines(rawLines []string) []Puzzle {
	puzzles := make([]Puzzle, len(rawLines))
	for i, line := range rawLines {
		puzzles[i] = parseRawLine(line)
	}

	return puzzles
}

func parseRawLine(rawLine string) Puzzle {
	splitted := strings.Split(rawLine, " ")
	numbers := strings.Split(splitted[1], ",")

	return Puzzle{
		Gear:      splitted[0],
		Positions: utils.ToIntSlice(numbers),
	}
}

func IsConfigurationCorrect(configuration string, expected []int) bool {
	// Convert # to numbers
	config := GetConfigurationForString(configuration)

	return utils.CompareIntSlices(expected, config)
}

func GetConfigurationForString(configuration string) []int {
	s := regexp.MustCompile("\\.+").Split(configuration, -1)

	fmt.Println(s)

	return utils.MapFunc1(s, func(v string) int {
		return len(v)
	})
}

func SplitIntoFillerSections(configuration string) []string {
	splitted := regexp.MustCompilePOSIX("\\.+").Split(configuration, -1)

	return utils.Filter(splitted, func(value string) bool {
		return len(value) != 0
	})
}
