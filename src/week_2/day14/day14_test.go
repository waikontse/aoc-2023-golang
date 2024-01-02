package day14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1_example(t *testing.T) {
	expected := 136
	result := SolvePart1("day_14_example.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart1(t *testing.T) {
	expected := 109466
	result := SolvePart1("day_14.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart2_example(t *testing.T) {
	expected := 64
	result := SolvePart2("day_14_example.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart2(t *testing.T) {
	expected := 94585
	result := SolvePart2("day_14.txt")

	assert.Equal(t, expected, result)
}
