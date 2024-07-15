package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1_example(t *testing.T) {
	expected := 102
	result := SolvePart1("day17_example.txt")

	assert.Equal(t, expected, result)
}

func TestPart1(t *testing.T) {
	expected := 102
	result := SolvePart1("day17.txt")

	assert.Equal(t, expected, result)
}

func TestPart2_example(t *testing.T) {
	expected := 102
	result := SolvePart2("day17_example.txt")

	assert.Equal(t, expected, result)
}

func TestPart2(t *testing.T) {
	expected := 102
	result := SolvePart2("day17.txt")

	assert.Equal(t, expected, result)
}
