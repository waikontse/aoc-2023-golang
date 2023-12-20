package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1Sample(t *testing.T) {
	result := SolvePart1("../../../resources/day_7_example.txt")
	expected := 6440

	assert.Equal(t, result, expected)
}

func TestSolvePart1(t *testing.T) {
	result := SolvePart1("../../../resources/day_7.txt")
	expected := 250232501

	assert.Equal(t, result, expected)
}

func TestSolvePart2Sample(t *testing.T) {
	result := SolvePart2("../../../resources/day_7_example.txt")
	expected := 5905

	assert.Equal(t, result, expected)
}

func TestSolvePart2(t *testing.T) {
	result := SolvePart2("../../../resources/day_7.txt")
	expected := 5905

	assert.Equal(t, result, expected)
}
