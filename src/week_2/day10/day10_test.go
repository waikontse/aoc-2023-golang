package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1_example(t *testing.T) {
	var expected = 23
	result := SolvePart1("day_10_example.txt")

	assert.Equal(t, result, expected)
}

func TestSolvePart1(t *testing.T) {
	var expected = 7005
	result := SolvePart1("day_10.txt")

	assert.Equal(t, result, expected)
}

func TestSolvePart2_example(t *testing.T) {
	var expected = 4
	result := SolvePart2("day_10_example.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart2_example2(t *testing.T) {
	var expected = 8
	result := SolvePart2("day_10_example_2.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart2_example3(t *testing.T) {
	var expected = 10
	result := SolvePart2("day_10_example_3.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart2(t *testing.T) {
	var expected = 417
	result := SolvePart2("day_10.txt")

	assert.Equal(t, result, expected)
}
