package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ParseFile(filenameWithPath string) []string {
	file, err := os.Open(filenameWithPath)
	failOnError(err)

	defer file.Close()

	var scanner = bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ParseFileForDay(day int) {

}

func SumInts(values []int) int {
	return ReduceFunc(values, func(i, j int) int {
		return i + j
	}, 0)
}

func SumStrings(values []string) int {
	mappedValues := MapFunc1(values, func(str string) int {
		value, _ := strconv.Atoi(str)
		return value
	})

	return SumInts(mappedValues)
}

func MapFunc1[T comparable, R comparable](values []T, transform func(T) R) []R {
	if len(values) == 0 {
		return []R{}
	}

	var result []R
	for _, value := range values {
		result = append(result, transform(value))
	}

	return result
}

func ReduceFunc[T comparable](values []T, reducer func(T, T) T, startValue T) T {
	if len(values) == 0 {
		return startValue
	}

	firstValue := reducer(values[0], startValue)
	var currentReducedValue = firstValue
	for _, value := range values[1:] {
		currentReducedValue = reducer(value, currentReducedValue)
	}

	return currentReducedValue
}
