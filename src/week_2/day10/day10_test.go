package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1_example(t *testing.T) {
	var expected = 8
	result := SolvePart1("../../../resources/day_10_example.txt")

	assert.Equal(t, result, expected)
}

func TestSolvePart1(t *testing.T) {
	var expected = 8
	result := SolvePart1("../../../resources/day_10.txt")

	assert.Equal(t, result, expected)
}
