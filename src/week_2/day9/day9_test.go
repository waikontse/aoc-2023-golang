package day9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1_example(t *testing.T) {
	var expected = 114
	result := SolvePart1("../../../resources/day_9_example.txt")

	assert.Equal(t, result, expected)
}

func TestSolvePart1(t *testing.T) {
	var expected = 1708206096
	result := SolvePart1("../../../resources/day_9.txt")

	assert.Equal(t, result, expected)
}

func TestSolvePart2_example(t *testing.T) {
	var expected = 2
	result := SolvePart2("../../../resources/day_9_example.txt")

	assert.Equal(t, result, expected)
}

func TestSolvePart2(t *testing.T) {
	var expected = 1050
	result := SolvePart2("../../../resources/day_9.txt")

	assert.Equal(t, result, expected)
}
