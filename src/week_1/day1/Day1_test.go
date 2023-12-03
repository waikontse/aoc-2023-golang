package day1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCorrectCode(t *testing.T) {
	dir, _ := os.Getwd()
	fmt.Println("working dir: ", dir)

	var result = correctCode("5sixthree22fourfoursix", "522")

	assert.Equal(t, result, "5226")
}

func TestPart1Code(t *testing.T) {
	var result = solvePart1("../../../resources/day_1.txt")

	assert.Equal(t, result, 55816)
}

func TestPart2Code(t *testing.T) {
	var result = solvePart2("../../../resources/day_1.txt")

	assert.Equal(t, result, 54980)
}
