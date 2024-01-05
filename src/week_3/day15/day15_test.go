package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1_example(t *testing.T) {
	expected := 1320
	result := SolvePart1("day_15_example.txt")

	assert.Equal(t, expected, result)
}

func TestCalculateHash(t *testing.T) {
	expected := 52
	result := CalculateHash("HASH")

	assert.Equal(t, expected, result)
}

func TestSolvePart1(t *testing.T) {
	expected := 517551
	result := SolvePart1("day_15.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart2_example(t *testing.T) {
	expected := 145
	result := SolvePart2("day_15_example.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart2(t *testing.T) {
	expected := 286097
	result := SolvePart2("day_15.txt")

	assert.Equal(t, expected, result)
}
