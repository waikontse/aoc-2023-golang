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

	allFits := utils.MapFunc1(puzzles, func(puzzle Puzzle) int {
		return checkNumberOfFits(puzzle)
	})

	return utils.SumInts(allFits)

	//for _, puzzle := range puzzles {
	//	fmt.Print("Puzzle string: ", puzzle.Gear)
	//	fmt.Printf(" has %d ? in it.\n", strings.Count(puzzle.Gear, "?"))
	//	// Map
	//
	//	fmt.Println("Number of fits: ", checkNumberOfFits(puzzle))
	//}
	//
	//return 0
}

func checkNumberOfFits(puzzle Puzzle) int {
	questionMarks := strings.Count(puzzle.Gear, "?")
	generatedMutations := generateMutationsStart([]string{"#", "."}, int64(questionMarks))
	fitCounts := int(0)
	for _, mutation := range generatedMutations {
		//	fmt.Println("Generated mutation: ", mutation)

		replacement := fillInReplacements(puzzle.Gear, mutation)
		isFit := checkFitness(replacement, puzzle.Positions)
		if isFit {
			fitCounts += 1
		}
	}

	return fitCounts
}

func fillInReplacements(target string, replacements string) string {
	var replacedString = target
	for i := 0; i < len(replacements); i++ {
		replacedString = strings.Replace(replacedString, "?", string([]rune(replacements)[i]), 1)
	}

	return replacedString
}

func checkFitness(target string, targets []int) bool {
	// Split string into '.' and '#'
	splittedTarget := strings.Split(target, ".")
	splittedTarget = utils.Filter(splittedTarget, func(subGroup string) bool {
		return subGroup != ""
	})

	//fmt.Println("Checking fitness: ", target, targets)
	//fmt.Println("Splitted targets: ", splittedTarget, len(splittedTarget))

	if len(targets) != len(splittedTarget) {
		return false
	}

	var targetFits = true
	for i := 0; i < len(targets); i++ {
		if len(splittedTarget[i]) != targets[i] {
			targetFits = false
			break
		}
	}

	return targetFits
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

	return utils.CompareSlices(expected, config)
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

func generateMutationsStart(choiceSet []string, maxDepth int64) []string {
	var sb = strings.Builder{}
	sb.Grow(int(maxDepth))
	var currentCombination = sb.String()

	return generateMutations(0, maxDepth, choiceSet, currentCombination, []string{})
}
func generateMutations(
	currentDepth int64,
	maxDepth int64,
	choiceSet []string,
	currentCombination string,
	accumulator []string) []string {

	if currentDepth == maxDepth {
		return append(accumulator, currentCombination)
	}

	var updatedAccumulator = accumulator
	for _, s := range choiceSet {
		var updatedCombination = currentCombination + s
		updatedAccumulator = generateMutations(currentDepth+1, maxDepth, choiceSet, updatedCombination, updatedAccumulator)
	}

	return updatedAccumulator
}
