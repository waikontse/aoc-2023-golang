package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1_example(t *testing.T) {
	expected := 46
	result := SolvePart1("day16_example.txt")

	assert.Equal(t, expected, result)
}

func TestPart1(t *testing.T) {
	expected := 7472
	result := SolvePart1("day16.txt")

	assert.Equal(t, expected, result)
}

func TestPart2_example(t *testing.T) {
	expected := 51
	result := SolvePart2("day16_example.txt")

	assert.Equal(t, expected, result)
}

func TestPart2(t *testing.T) {
	expected := 7716
	result := SolvePart2("day16.txt")

	assert.Equal(t, expected, result)
}
