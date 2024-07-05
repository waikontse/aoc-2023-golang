package day18

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1_example(t *testing.T) {
	expected := 62
	result := SolvePart1("day18_example.txt")

	assert.Equal(t, expected, result)
}

func TestPart1(t *testing.T) {
	expected := 102
	result := SolvePart1("day18.txt")

	assert.Equal(t, expected, result)
}

func TestPart2_example(t *testing.T) {
	expected := 102
	result := SolvePart2("day18_example.txt")

	assert.Equal(t, expected, result)
}

func TestPart2(t *testing.T) {
	expected := 102
	result := SolvePart2("day18.txt")

	assert.Equal(t, expected, result)
}
