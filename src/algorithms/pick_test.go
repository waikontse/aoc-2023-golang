package algorithms

import (
	"aoc-2023-golang/src/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatePointsInArea(t *testing.T) {
	expected := int64(4)
	positions := make([]utils.Position, 0)

	positions = append(positions, utils.Position{First: 0, Second: 0})
	positions = append(positions, utils.Position{First: 1, Second: 0})
	positions = append(positions, utils.Position{First: 2, Second: 0})
	positions = append(positions, utils.Position{First: 3, Second: 0})

	positions = append(positions, utils.Position{First: 3, Second: 1})
	positions = append(positions, utils.Position{First: 3, Second: 2})
	positions = append(positions, utils.Position{First: 3, Second: 3})

	positions = append(positions, utils.Position{First: 2, Second: 3})
	positions = append(positions, utils.Position{First: 1, Second: 3})
	positions = append(positions, utils.Position{First: 0, Second: 3})

	positions = append(positions, utils.Position{First: 0, Second: 3})
	positions = append(positions, utils.Position{First: 0, Second: 2})
	positions = append(positions, utils.Position{First: 0, Second: 1})

	result := CalculatePointsInArea(positions)

	assert.Equal(t, expected, result)
}
