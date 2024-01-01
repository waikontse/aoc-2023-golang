package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1_example(t *testing.T) {
	expected := 21
	result := SolvePart1("../../../resources/day_12_example.txt")

	assert.Equal(t, expected, result)
}

func TestSolvePart1(t *testing.T) {
	expected := 21
	result := SolvePart1("../../../resources/day_12.txt")

	assert.Equal(t, expected, result)
}

func TestGetConfigurationForString(t *testing.T) {
	str := GetConfigurationForString("#.#...###")
	expected := []int{1, 1, 3}

	assert.Equal(t, expected, str)
}

func TestIsConfigurationCorrect(t *testing.T) {
	b := IsConfigurationCorrect("##...#....######", []int{2, 1, 6})
	expected := true

	assert.Equal(t, expected, b)
}

func TestIsConfigurationCorrect_false(t *testing.T) {
	b := IsConfigurationCorrect("#...#....######", []int{2, 1, 6})
	expected := false

	assert.Equal(t, expected, b)
}
