package day15

import (
	"aoc-2023-golang/src/utils"
	"slices"
	"strconv"
	"strings"
)

type Lens struct {
	Label    string
	Strength int
}

type LensBoxes struct {
	Boxes [256][]Lens
}

func SolvePart1(filename string) int {
	input := ReadInput(filename)

	return CalculateAllHashes(input)
}

func SolvePart2(filename string) int {
	input := ReadInput(filename)
	boxes := LensBoxes{}

	UpdateBoxesWithInstructions(&boxes, input)

	return CalculateAllBoxesFocusPower(&boxes)
}

func ReadInput(filename string) []string {
	rawLine := utils.ParseFile(filename)

	return strings.Split(rawLine[0], ",")
}

func CalculateAllHashes(values []string) int {
	sum := 0
	for i := range values {
		sum += CalculateHash(values[i])
	}

	return sum
}

func CalculateHash(value string) int {
	sum := 0
	for _, i := range value {
		sum += int(i)
		sum = sum * 17
		sum = sum % 256
	}

	return sum
}

func UpdateBoxesWithInstructions(boxes *LensBoxes, instructions []string) {
	for i := range instructions {
		UpdateBoxesWithInstruction(boxes, instructions[i])
	}
}

func UpdateBoxesWithInstruction(boxes *LensBoxes, instruction string) {
	// Determine instruction: = or -
	// Split the instruction
	// Determine with box to use
	// Perform operation

	if strings.Contains(instruction, "=") {
		UpdateBoxesWithEquals(boxes, instruction)
	} else {
		UpdateBoxesWithMinus(boxes, instruction)
	}
}

func UpdateBoxesWithMinus(boxes *LensBoxes, instruction string) {
	split := strings.Split(instruction, "-")
	label := split[0]
	location := CalculateHash(label)
	index := slices.IndexFunc(boxes.Boxes[location], func(lens Lens) bool {
		return lens.Label == label
	})

	if index != -1 {
		boxes.Boxes[location] = slices.Delete(boxes.Boxes[location], index, index+1)
	}
}

func UpdateBoxesWithEquals(boxes *LensBoxes, instruction string) {
	split := strings.Split(instruction, "=")
	location := CalculateHash(split[0])
	label := split[0]

	index := slices.IndexFunc(boxes.Boxes[location], func(lens Lens) bool {
		return lens.Label == label
	})

	lensValue, _ := strconv.Atoi(split[1])
	if index != -1 {
		boxes.Boxes[location][index].Strength = lensValue
	} else {
		newLens := Lens{Label: label, Strength: lensValue}
		boxes.Boxes[location] = append(boxes.Boxes[location], newLens)
	}
}

func CalculateAllBoxesFocusPower(boxes *LensBoxes) int {
	sum := 0

	for i := range boxes.Boxes {
		sum += CalculateBoxFocusPower(boxes.Boxes[i], i+1)
	}

	return sum
}

func CalculateBoxFocusPower(lenses []Lens, boxNumber int) int {
	// calculate boxNumber * slotNumber * focal strength
	sum := 0
	for i := range lenses {
		sum += boxNumber * (i + 1) * lenses[i].Strength
	}

	return sum
}
