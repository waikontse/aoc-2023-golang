package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	result := SolvePart1("../../../resources/day_5_example.txt")

	assert.Equal(t, result, 32)
}
