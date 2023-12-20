package day8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	expected := 20221
	result := SolvePart1("../../../resources/day_8.txt")

	assert.Equal(t, result, expected)
}

func TestSolvePart2(t *testing.T) {
	var expected int64 = 14616363770447
	result := SolvePart2("../../../resources/day_8.txt")

	assert.Equal(t, result, expected)
}
