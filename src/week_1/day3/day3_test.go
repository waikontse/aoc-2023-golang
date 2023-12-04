package day3

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseNumbersInLine(t *testing.T) {
	// Arrange
	result := FindNumbersInLine("467..114.9", 0)

	// Act
	// Assert
	assert.Len(t, result, 3)
	fmt.Println(result)
}

func TestParseNumbersInLine2(t *testing.T) {
	// Arrange
	result := FindNumbersInLine("467..11489", 0)

	// Act
	// Assert
	assert.Len(t, result, 2)
	fmt.Println(result)
}

func TestFindSymbolsInLIne(t *testing.T) {
	// Arrange
	result := FindSymbolsInLIne(".12$.*45..8.", 0)
	expected1 := Point{x: 0, y: 3}
	expected2 := Point{x: 0, y: 5}

	// Act
	// Assert
	assert.Len(t, result, 2)
	assert.Equal(t, result[0], expected1)
	assert.Equal(t, result[1], expected2)
}
