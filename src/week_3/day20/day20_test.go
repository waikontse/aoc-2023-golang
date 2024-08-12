package day20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solvePart1_example(t *testing.T) {
	expected := int64(32000000)

	result := solvePart1("day20_example.txt")

	assert.Equal(t, expected, result)
}

func Test_solvePart1_example2(t *testing.T) {
	expected := int64(11687500)

	result := solvePart1("day20_example_2.txt")

	assert.Equal(t, expected, result)
}

func Test_solvePart1(t *testing.T) {
	expected := 1000

	result := solvePart1("day20.txt")

	assert.Equal(t, expected, result)
}
