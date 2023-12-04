package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PrintLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
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

func MapFunc1WithIndex[T comparable, R comparable](values []T, transform func(int, T) R) []R {
	if len(values) == 0 {
		return []R{}
	}

	var result []R
	for i, value := range values {
		result = append(result, transform(i, value))
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

type Pair[T comparable, R comparable] struct {
	First  T
	Second R
}

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
}

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func ToIntSlice(intStrings []string) []int {
	var ints []int
	for _, intString := range intStrings {
		if len(intString) == 0 {
			continue
		}
		result, _ := strconv.Atoi(intString)
		ints = append(ints, result)
	}

	return ints
}

func ConvertToIntMap(numbers []int) map[int]struct{} {
	numberMap := make(map[int]struct{})
	for _, number := range numbers {
		numberMap[number] = struct{}{}
	}

	return numberMap
}

func InitArray(ints []int, initValue int) {
	for i := 0; i < len(ints); i++ {
		ints[i] = initValue
	}
}

func IncrementArray(array []int, startPosition int, repeat int, incrementBy int) {
	for i := startPosition; i < startPosition+repeat; i++ {
		array[i] = array[i] + incrementBy
	}
}
