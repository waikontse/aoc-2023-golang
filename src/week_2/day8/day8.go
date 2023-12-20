package day8

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"strings"
)

func SolvePart1(filename string) int {
	rawLines := utils.ParseFile(filename)
	game := ParseRawLinesIntoMap(rawLines)

	return walkTillWinning(game, "AAA", func(point string) bool {
		return point == "ZZZ"
	})
}

func SolvePart2(filename string) int64 {
	rawLines := utils.ParseFile(filename)
	game := ParseRawLinesIntoMap(rawLines)

	filteredNodes := FilterMap(game.atlas, func(node string) bool {
		return node[len(node)-1] == 'A'
	})

	var steps []int64
	for k := range filteredNodes {
		resultForNode := walkTillWinning(game, k, func(point string) bool {
			return point[len(point)-1] == 'Z'
		})

		fmt.Printf("Result for: %s %d\n", k, resultForNode)
		steps = append(steps, int64(resultForNode))
	}

	return utils.LCM(steps[0], steps[1], steps[2:]...)
}

func FilterMap(mapper map[string][]string, predicate func(node string) bool) map[string][]string {
	filteredValues := make(map[string][]string)
	for k, v := range mapper {
		if predicate(k) {
			filteredValues[k] = v
		}
	}

	return filteredValues
}

func walkTillWinning(game Game, startNode string, predicate func(point string) bool) int {
	currentNode := startNode
	var round = 0
	for !predicate(currentNode) {
		currentNode = walkMoveList(game, currentNode)
		round += 1
	}

	fmt.Println("Stopped at node: ", currentNode)

	return round * len(game.movesList)
}

func walkMoveList(game Game, start string) string {
	currentNode := start

	for _, command := range game.movesList {
		if command == 'R' {
			currentNode = game.atlas[currentNode][1]
		} else {
			currentNode = game.atlas[currentNode][0]
		}
	}

	return currentNode
}

type Game struct {
	movesList string
	atlas     map[string][]string
}

func ParseRawLinesIntoMap(rawLines []string) Game {
	movesList := rawLines[0:1]
	mapperLines := rawLines[2:]

	var mapper = make(map[string][]string)
	for _, line := range mapperLines {
		from, to := ParseRawLine(line)
		mapper[from] = to
	}

	return Game{movesList: movesList[0], atlas: mapper}
}

func ParseRawLine(rawLine string) (string, []string) {
	split := strings.Split(rawLine, " = ")

	return split[0], ParseLeftAndRight(split[1])
}

func ParseLeftAndRight(leftAndRight string) []string {
	left := strings.Replace(leftAndRight, "(", "", 1)
	right := strings.Replace(left, ")", "", 1)

	return strings.Split(right, ", ")
}
