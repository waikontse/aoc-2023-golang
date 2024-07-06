package day18

import (
	"aoc-2023-golang/src/algorithms"
	"aoc-2023-golang/src/utils"
	"strconv"
	"strings"
)

type Movement struct {
	direction string
	steps     int
	colour    string
}

type TrackingInfo struct {
	StartingPos utils.Position
	CurrentPos  utils.Position
	stepsTaken  int64
	positions   []utils.Position
}

func SolvePart1(filename string) int64 {
	movements, trackingInfo := prepareInput(filename)

	for _, movement := range movements {
		walk(movement, &trackingInfo)
	}

	return algorithms.CalculatePointsInArea(trackingInfo.positions) + trackingInfo.stepsTaken
}

func SolvePart2(filename string) int64 {
	movements, trackingInfo := prepareInput(filename)

	for _, movement := range movements {
		transformedMovement := transformMovement(movement)
		walk(transformedMovement, &trackingInfo)
	}

	return algorithms.CalculatePointsInArea(trackingInfo.positions) + trackingInfo.stepsTaken
}

func prepareInput(filename string) ([]Movement, TrackingInfo) {
	rawLines := utils.ParseFile(filename)

	movements := parseRawlinesToMovements(rawLines)
	size := determineSize(movements)
	trackingInfo := TrackingInfo{
		StartingPos: utils.Position{First: size / 2, Second: size / 2},
		CurrentPos:  utils.Position{First: size / 2, Second: size / 2},
		stepsTaken:  0,
		positions:   make([]utils.Position, 0),
	}

	trackingInfo.positions = append(trackingInfo.positions, utils.Position{First: size / 2, Second: size / 2})
	return movements, trackingInfo
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

func determineSize(movements []Movement) int {
	dMovements := utils.Filter(movements, func(value Movement) bool {
		return value.direction == "D"
	})

	uMovements := utils.Filter(movements, func(value Movement) bool {
		return value.direction == "U"
	})

	lMovements := utils.Filter(movements, func(value Movement) bool {
		return value.direction == "L"
	})

	rMovements := utils.Filter(movements, func(value Movement) bool {
		return value.direction == "R"
	})

	sum := 0
	for _, movement := range dMovements {
		sum += movement.steps
	}
	for _, movement := range uMovements {
		sum += movement.steps
	}
	for _, movement := range lMovements {
		sum += movement.steps
	}
	for _, movement := range rMovements {
		sum += movement.steps
	}

	return sum
}

func walk(movement Movement, t *TrackingInfo) {
	// 1) Update the steps taken
	// 2) Update the current position
	// 3) Update the positions

	t.stepsTaken += int64(movement.steps)
	currPos := t.CurrentPos
	if movement.direction == "D" {
		t.CurrentPos.Second = currPos.Second + movement.steps

		for y := currPos.Second + 1; y <= currPos.Second+movement.steps; y++ {
			newPosition := utils.Position{First: currPos.First, Second: y}
			t.positions = append(t.positions, newPosition)
		}

	} else if movement.direction == "U" {
		t.CurrentPos.Second = currPos.Second - movement.steps

		for y := currPos.Second - 1; y >= currPos.Second-movement.steps; y-- {
			newPosition := utils.Position{First: currPos.First, Second: y}
			t.positions = append(t.positions, newPosition)
		}

	} else if movement.direction == "L" {
		t.CurrentPos.First = currPos.First - movement.steps

		for x := currPos.First - 1; x >= currPos.First-movement.steps; x-- {
			newPosition := utils.Position{First: x, Second: currPos.Second}
			t.positions = append(t.positions, newPosition)
		}

	} else if movement.direction == "R" {
		t.CurrentPos.First = currPos.First + movement.steps

		for x := currPos.First + 1; x <= currPos.First+movement.steps; x++ {
			newPosition := utils.Position{First: x, Second: currPos.Second}
			t.positions = append(t.positions, newPosition)
		}
	}
}

func transformMovement(movement Movement) Movement {
	// first 1 strings are hex
	// last char is direction

	steps, _ := strconv.ParseInt(movement.colour[2:7], 16, 64)
	direction := convertNumbertoDirection(movement.colour[7:8])

	return Movement{
		direction: direction, steps: int(steps), colour: "",
	}
}

func convertNumbertoDirection(value string) string {
	if value == "0" {
		return "R"
	} else if value == "1" {
		return "D"
	} else if value == "2" {
		return "L"
	} else if value == "3" {
		return "U"
	}

	return "ERROR"
}
