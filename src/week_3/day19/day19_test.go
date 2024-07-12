package day19

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1_example(t *testing.T) {
	expected := 19114
	result := SolvePart1("day19_example.txt")

	assert.Equal(t, expected, result)
}

func TestParseRule(t *testing.T) {
	test := "px{a<2006:qkq,m>2090:A,rfg}"
	spec := ParseRule(test)

	fmt.Println(spec)
}

func TestParseItem(t *testing.T) {
	test := "{x=787,m=2655,a=1222,s=2876}"
	item := ParseItem(test)

	assert.Equal(t, 787, item.x)
	assert.Equal(t, 2655, item.m)
	assert.Equal(t, 1222, item.a)
	assert.Equal(t, 2876, item.s)
}

func TestPart1(t *testing.T) {
	expected := 346230
	result := SolvePart1("day19.txt")

	assert.Equal(t, expected, result)
}

func TestPart2_example(t *testing.T) {
	expected := int64(167409079868000)
	result := SolvePart2("day19_example.txt")

	assert.Equal(t, expected, result)
}

func TestPart2_example_small(t *testing.T) {
	expected := int64(2952)
	result := SolvePart2("day19_small.txt")

	assert.Equal(t, expected, result)
}

func TestPart2(t *testing.T) {
	expected := int64(112074045986829)
	result := SolvePart2("day19.txt")

	assert.Equal(t, expected, result)
}
