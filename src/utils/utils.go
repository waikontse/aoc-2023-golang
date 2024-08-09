package utils

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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

func SumInts(values []int) int {
	return ReduceFunc(values, func(i, j int) int {
		return i + j
	}, 0)
}

func AllIntsAreZero(values []int) bool {
	for _, value := range values {
		if value != 0 {
			return false
		}
	}

	return true
}

func SumStrings(values []string) int {
	mappedValues := MapFunc1(values, func(str string) int {
		value, _ := strconv.Atoi(str)
		return value
	})

	return SumInts(mappedValues)
}

func MapFunc1[T any, R any](values []T, transform func(T) R) []R {
	if len(values) == 0 {
		return []R{}
	}

	var result []R
	for _, value := range values {
		result = append(result, transform(value))
	}

	return result
}

func MapFunc1WithIndex[T any, R any](values []T, transform func(int, T) R) []R {
	if len(values) == 0 {
		return []R{}
	}

	var result []R
	for i, value := range values {
		result = append(result, transform(i, value))
	}

	return result
}

func ReduceFunc[T any](values []T, reducer func(T, T) T, startValue T) T {
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

func ToOccurrenceMap(chars []string) map[string]int {
	occurrence := make(map[string]int, 0)

	for _, char := range chars {
		occurrence[char] += 1
	}

	return occurrence
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

func FindIndexes(s []string, target string) []int {
	var foundIndexes []int
	for i, str := range s {
		if strings.Contains(str, target) {
			foundIndexes = append(foundIndexes, i)
		}
	}
	return foundIndexes
}

func MapValues(m map[string]int) []int {
	occurrences := make([]int, 0)
	for _, v := range m {
		occurrences = append(occurrences, v)
	}

	return occurrences
}

func FindGCD(nums []int64) int64 {
	minValue, maxValue := nums[0], nums[0]

	for _, num := range nums {
		if num < minValue {
			minValue = num
		}
		if num > maxValue {
			maxValue = num
		}
	}

	return Gcd(minValue, maxValue)
}

func Gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / Gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Filter[T any](values []T, pred func(value T) bool) []T {
	filteredValues := make([]T, 0)

	for i := range values {
		if pred(values[i]) {
			filteredValues = append(filteredValues, values[i])
		}
	}

	return filteredValues
}

func Count[T any](values []T, predicate func(value T) bool) int {
	var count int = 0
	for i := range values {
		if predicate(values[i]) {
			count += 1
		}
	}

	return count
}

func ZipWith[T any, R any](lValues []T, rValues []R) []Pair[T, R] {
	pairs := make([]Pair[T, R], 0)

	for i := range lValues {
		newPair := Pair[T, R]{First: lValues[i], Second: rValues[i]}
		pairs = append(pairs, newPair)
	}

	return pairs
}

func CompareSlices[T comparable](left, right []T) bool {
	if len(left) != len(right) {
		return false
	}
	for i, v := range left {
		if v != right[i] {
			return false
		}
	}
	return true
}

func DiffInSlices[T comparable](left, right []T) int {
	diffs := 0
	for i, v := range left {
		if v != right[i] {
			diffs += 1
		}
	}

	return diffs
}

func Max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func MinValue(values []int) int {
	lowestValue := math.MaxInt32
	for _, value := range values {
		if value < lowestValue {
			lowestValue = value
		}
	}

	return lowestValue
}
