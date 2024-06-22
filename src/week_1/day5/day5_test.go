package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1Example(t *testing.T) {
	result := SolvePart1("../../../resources/day_5_example.txt")

	assert.Equal(t, result, int64(35))
}

func TestSolvePart1(t *testing.T) {
	result := SolvePart1("../../../resources/day_5.txt")

	assert.Equal(t, result, int64(346433842))
}

func TestSolvePart2Example(t *testing.T) {
	result := SolvePart2("../../../resources/day_5_example.txt")

	assert.Equal(t, result, int64(46))
}

func TestSolvePart2(t *testing.T) {
	result := SolvePart2("../../../resources/day_5.txt")

	assert.Equal(t, result, int64(60294664))
}
