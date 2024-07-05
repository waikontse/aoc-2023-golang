package day18

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Movement struct {
	direction string
	steps     int
	colour    string
}

func SolvePart1(filename string) int {
	rawLines := utils.ParseFile(filename)

	movements := parseRawlinesToMovements(rawLines)

	for _, movement := range movements {
		fmt.Println(movement)
	}

	return 0
}

func SolvePart2(filename string) int {
	return 0
}

func parseRawlinesToMovements(rawLines []string) []Movement {
	movements := make([]Movement, 0)
	for _, value := range rawLines {
		splitted := strings.Split(value, " ")
		steps, _ := strconv.Atoi(splitted[1])
		movement := Movement{
			direction: splitted[0],
			steps:     steps,
			colour:    splitted[2],
		}

		movements = append(movements, movement)
	}

	return movements
}
