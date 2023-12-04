package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	// Arrange
	var expected = 25010

	// Act
	result := SolvePart1()

	// Assert
	assert.Equal(t, expected, result)
}

func TestSolvePart2(t *testing.T) {
	// Arrange
	var expected = 9924412

	// Act
	result := SolvePart1()

	// Assert
	assert.Equal(t, expected, result)
}
