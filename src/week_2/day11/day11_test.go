package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1_example(t *testing.T) {
	expected := 374
	result := SolvePart1("../../../resources/day_11_example.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart1(t *testing.T) {
	expected := 9609130
	result := SolvePart1("../../../resources/day_11.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart2_example(t *testing.T) {
	expected := 8410
	result := SolvePart2("../../../resources/day_11_example.txt", 100)

	assert.Equal(t, expected, result)
}

func TestSolvePart2(t *testing.T) {
	expected := 702152204842
	result := SolvePart2("../../../resources/day_11.txt", 1000000)

	assert.Equal(t, expected, result)
}
