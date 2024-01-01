package day13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1_example(t *testing.T) {
	expected := 405
	result := SolvePart1("../../../resources/day_13_example.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart1(t *testing.T) {
	expected := 29130
	result := SolvePart1("../../../resources/day_13.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart2_example(t *testing.T) {
	expected := 400
	result := SolvePart2("../../../resources/day_13_example.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart2(t *testing.T) {
	expected := 33438
	result := SolvePart2("../../../resources/day_13.txt")

	assert.Equal(t, expected, result)
}
