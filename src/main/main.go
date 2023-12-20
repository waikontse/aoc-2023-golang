package main

import (
	"aoc-2023-golang/src/week_2/day9"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	// fmt.Println(day8.SolvePart1("./resources/day_8.txt"))
	fmt.Println(day9.SolvePart1("./resources/day_9_example.txt"))

	elapsed := time.Since(startTime)

	fmt.Printf("Execution time: %s\n", elapsed)
}
