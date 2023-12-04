package main

import (
	"aoc-2023-golang/src/week_1/day4"
	"fmt"
	"time"
)

func main() {
	//day3.SolvePart1()
	//day3.SolvePart2()

	startTime := time.Now()

	fmt.Println(day4.SolvePart2())

	elapsed := time.Since(startTime)

	fmt.Printf("Execution time: %s\n", elapsed)
}
