package algorithms

import (
	"aoc-2023-golang/src/utils"
	"math"
)

// CalculatePointsInArea Pick's theorem
func CalculatePointsInArea(positions []utils.Position) int64 {
	totalArea := math.Floor(CalculateArea(positions))

	subtract := int64((len(positions) / 2) - 1)
	totalPointsInArea := int64(totalArea) - subtract

	return totalPointsInArea
}

// CalculateArea Shoelace algorithm
func CalculateArea(positions []utils.Position) float64 {
	sum := 0
	for i := 0; i < len(positions)-1; i++ {
		pos1 := positions[i]
		pos2 := positions[i+1]

		sum += (pos1.First * pos2.Second) - (pos2.First * pos1.Second)
	}

	return math.Abs(float64(sum) / 2)
}
