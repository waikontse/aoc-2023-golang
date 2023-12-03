package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Code(t *testing.T) {
	var result = SolvePart1("../../../resources/day_2.txt")

	assert.Equal(t, result, 2683)
}

func TestPart2Code(t *testing.T) {
	var result = SolvePart2("../../../resources/day_2.txt")

	assert.Equal(t, result, 49710)
}
