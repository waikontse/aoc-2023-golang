package day20

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"os"
	"strings"
)

type Flipflop struct {
	name      string
	input     int
	output    int
	connected []string
}

type Conjunction struct {
	name      string
	inputs    map[string]int
	output    int
	connected []string
}

type Circuit struct {
	broadcast    []string
	flipflops    map[string]Flipflop
	conjunctions map[string]Conjunction
	highSignals  int
	lowSignals   int
}

type UpdateEntry struct {
	value     int
	from      string
	component string
}

func solvePart1(filename string) int {
	circuit := parseIntoCircuit(filename)

	fmt.Println(circuit)

	return 0
}

func solvePart2(filename string) int {
	return 0
}

func parseIntoCircuit(filename string) Circuit {
	rawLines := utils.ParseFile(filename)

	var broadcast []string = nil
	var flipflops = make(map[string]Flipflop)
	var conjunctions = make(map[string]Conjunction)

	for _, line := range rawLines {
		if strings.HasPrefix(line, "broadcast") {
			broadcast = parseBroadcast(line)
		} else if strings.HasPrefix(line, "%") {
			flipflop := parseFlipflop(line)
			flipflops[flipflop.name] = flipflop
		} else if strings.HasPrefix(line, "&") {
			conjunction := parseConjunction(line)
			conjunctions[conjunction.name] = conjunction
		} else {
			fmt.Println("Failed to parse line:", line)
			os.Exit(-1)
		}
	}

	circuit := Circuit{
		broadcast:    broadcast,
		flipflops:    flipflops,
		conjunctions: conjunctions,
	}

	return circuit
}

func parseBroadcast(broadcast string) []string {
	targets := strings.Replace(broadcast, "broadcast -> ", "", 1)

	return strings.Split(targets, ",")
}

func parseFlipflop(flipflop string) Flipflop {
	parts := strings.Split(flipflop, " -> ")
	name := parts[0][1:]
	connected := strings.Split(parts[1], ",")

	return Flipflop{
		name:      name,
		connected: connected,
		input:     0,
		output:    0,
	}
}

func parseConjunction(conjunction string) Conjunction {
	parts := strings.Split(conjunction, " -> ")
	name := parts[0][1:]
	connected := strings.Split(parts[1], ",")

	return Conjunction{
		name:      name,
		connected: connected,
		inputs:    make(map[string]int),
		output:    0,
	}
}

func simulateRound(circuit Circuit) {
	// send a low signal to all the broadcasted elements
	frontier := make([]UpdateEntry, 0)
	for _, broadcastItem := range circuit.broadcast {
		frontier = append(frontier, UpdateEntry{value: 0, component: broadcastItem, from: "broadcast"})
	}

	for len(frontier) > 0 {
		// Update frontier
		currentUpdate := frontier[0]
		frontier = frontier[1:]

		// update the output
		connectedComponents, newOutput := updateOutputForComponent(circuit, currentUpdate)

		// update the item that is connected to the item
		for _, component := range connectedComponents {
			frontier = append(frontier, UpdateEntry{value: newOutput, component: component, from: currentUpdate.component})
		}

		// update the counts
		// TODO
	}
}

func updateOutputForComponent(circuit Circuit, entry UpdateEntry) ([]string, int) {
	// Update if item is a conjunction
	conjunction, hasConjunction := circuit.conjunctions[entry.component]
	if hasConjunction {
		conjunction.inputs[entry.from]
	}

	// Update the flipflop
	flipflop, hasFlipflop := circuit.flipflops[entry.component]
	if hasFlipflop {

	}
}
