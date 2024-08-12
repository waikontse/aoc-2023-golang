package day20

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"os"
	"slices"
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
	flipflops    map[string]*Flipflop
	conjunctions map[string]*Conjunction
	highSignals  int
	lowSignals   int
}

type UpdateEntry struct {
	value     int
	from      string
	component string
}

func solvePart1(filename string) int64 {
	circuit := parseIntoCircuit(filename)
	updateConjunctionInputs(circuit)

	result := simulateNRounds(circuit, 1)
	//printCircuit(circuit)

	return result
}

func printCircuit(circuit Circuit) {
	fmt.Println("***** Circuit *****")
	fmt.Println(circuit)

	for _, conjunction := range circuit.conjunctions {
		fmt.Printf("Conjunction: %+v\n", conjunction)
	}

	for _, flipflop := range circuit.flipflops {
		fmt.Printf("Flipflop: %+v\n", flipflop)
	}

	fmt.Println("***** Circuit *****")
}

func solvePart2(filename string) int {
	return 0
}

func parseIntoCircuit(filename string) Circuit {
	rawLines := utils.ParseFile(filename)

	var broadcast []string = nil
	var flipflops = make(map[string]*Flipflop)
	var conjunctions = make(map[string]*Conjunction)

	for _, line := range rawLines {
		if strings.HasPrefix(line, "broadcast") {
			broadcast = parseBroadcast(line)
		} else if strings.HasPrefix(line, "%") {
			flipflop := parseFlipflop(line)
			flipflops[flipflop.name] = &flipflop
		} else if strings.HasPrefix(line, "&") {
			conjunction := parseConjunction(line)
			conjunctions[conjunction.name] = &conjunction
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

func updateConjunctionInputs(circuit Circuit) {
	for name, conjunction := range circuit.conjunctions {

		// Search for all conjunctions linking to another conjunction
		for connectedToName, connectedToConjunction := range circuit.conjunctions {
			foundIndex := slices.IndexFunc(connectedToConjunction.connected, func(connectedTo string) bool {
				return connectedTo == name
			})

			if foundIndex != -1 {
				conjunction.inputs[connectedToName] = 0
			}
		}

		// Search for all the flipflops linking to another conjunction
		for flipflopName, flipflop := range circuit.flipflops {
			foundIndex := slices.IndexFunc(flipflop.connected, func(connectedTo string) bool {
				return connectedTo == name
			})

			if foundIndex != -1 {
				conjunction.inputs[flipflopName] = 0
			}
		}

		// Search for all the broadcast linking to a conjunction
		for _, broadcast := range circuit.broadcast {
			if broadcast == name {
				conjunction.inputs["broadcast"] = 0
			}
		}
	}
}

func parseBroadcast(broadcast string) []string {
	targets := strings.Replace(broadcast, "broadcaster -> ", "", 1)

	return strings.Split(targets, ", ")
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

func simulateNRounds(circuit Circuit, rounds int) int64 {
	lowCount := int64(0)
	highCount := int64(0)
	for round := 1; round <= rounds; round++ {
		newLowCount, newHighCount := simulateRound(circuit)
		lowCount += int64(newLowCount)
		highCount += int64(newHighCount)

		fmt.Printf("Round %d lowCount: %d highCount: %d\n", round, lowCount, highCount)
	}

	return lowCount * highCount
}

func simulateRound(circuit Circuit) (int, int) {
	// send a low signal to all the broadcasted elements
	frontier := make([]UpdateEntry, 0)
	for _, broadcastItem := range circuit.broadcast {
		frontier = append(frontier, UpdateEntry{value: 0, component: broadcastItem, from: "broadcast"})
	}

	lowCount := 1 + len(circuit.broadcast)
	highCount := 0
	for len(frontier) > 0 {
		// Update frontier
		currentUpdate := frontier[0]
		frontier = frontier[1:]

		// update the output
		connectedComponents, newOutput, hasChanged := updateOutputForComponent(circuit, currentUpdate)

		// We need to handle all the pulses in the same "generation" before going to handle the next generation

		//// update the item that is connected to the item
		//for _, component := range connectedComponents {
		//	frontier = append(frontier, UpdateEntry{value: newOutput, component: component, from: currentUpdate.component})
		//}

		if hasChanged {
			fmt.Printf("Output has changed for component: %s:%d\n", currentUpdate.component, newOutput)

			for _, connectedComponentName := range connectedComponents {
				frontier = append(frontier, UpdateEntry{
					value:     newOutput,
					component: connectedComponentName,
					from:      currentUpdate.component,
				})
			}

			// update the counts
			if newOutput == 0 {
				lowCount++
			} else {
				highCount++
			}
		}
	}

	return lowCount, highCount
}

func updateOutputForComponent(circuit Circuit, entry UpdateEntry) ([]string, int, bool) {
	fmt.Printf("Updating for entry: %+v\n", entry)

	// Update if item is a conjunction
	conjunction, hasConjunction := circuit.conjunctions[entry.component]
	var outputValue int
	var connectedComponents []string
	hasChanged := false
	if hasConjunction {
		conjunction.inputs[entry.from] = entry.value
		hasChangedSinceEval, newValue := evalOutputForConjunction(conjunction)
		conjunction.output = newValue
		outputValue = conjunction.output
		connectedComponents = conjunction.connected
		hasChanged = hasChangedSinceEval
	}

	// Update the flipflop
	flipflop, hasFlipflop := circuit.flipflops[entry.component]
	if hasFlipflop {
		flipflop.input = entry.value
		hasChangedSinceEval, newValue := evalOutputForFlipflop(flipflop)
		flipflop.output = newValue
		outputValue = newValue
		connectedComponents = flipflop.connected
		hasChanged = hasChangedSinceEval
	}

	// Return the list of connected components and its output value
	return connectedComponents, outputValue, hasChanged
}

func evalOutputForConjunction(conjunction *Conjunction) (bool, int) {
	fmt.Println("Evaluating conjunction: ", conjunction.name)

	sum := 0
	for _, v := range conjunction.inputs {
		sum += v
	}

	if sum == 0 {
		return true, 1
	} else if sum == len(conjunction.inputs) {
		return true, 0
	}

	return false, conjunction.output
}

func evalOutputForFlipflop(flipflop *Flipflop) (bool, int) {
	fmt.Println("Evaluation flipflop: ", flipflop.name)

	if flipflop.input == 0 {
		if flipflop.output == 0 {
			return true, 1
		}

		return true, 0
	}

	return false, flipflop.output
}
