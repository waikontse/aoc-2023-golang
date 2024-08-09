package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1_example(t *testing.T) {
	expected := 102
	result := SolvePart1("day17_example.txt", 3)

	assert.Equal(t, expected, result)
}

func TestPart1(t *testing.T) {
	expected := 847
	result := SolvePart1("day17.txt", 3)

	assert.Equal(t, expected, result)
}

func TestPart2_example(t *testing.T) {
	expected := 94
	result := SolvePart2("day17_example.txt", 4, 10)

	assert.Equal(t, expected, result)
}

func TestPart2_example2(t *testing.T) {
	expected := 71
	result := SolvePart2("day17_example2.txt", 4, 10)

	assert.Equal(t, expected, result)
}

func TestPart2(t *testing.T) {
	expected := 997
	result := SolvePart2("day17.txt", 4, 10)

	assert.Equal(t, expected, result)
}
