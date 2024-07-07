package day19

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Item struct {
	x int
	m int
	a int
	s int
}

type Spec struct {
	operand1 string
	operand2 int
	opertor  string
	result   string
}

type EvalSpec struct {
	Name  string
	Specs []Spec
}

type Engine struct {
	rules    map[string]EvalSpec
	items    []Item
	accepted []Item
}

func SolvePart1(filename string) int {
	engine := parseRawInputToItemsAndEngine(filename)

	//fmt.Println(engine)

	for _, item := range engine.items {
		evalItem(&engine, item)
	}

	var sum = 0
	for _, item := range engine.accepted {
		sum += item.getTotalPoints()
	}

	return sum
}

func SolvePart2(filename string) int {
	engine := parseRawInputToItemsAndEngine(filename)

	// build all the rules up with their transitions to check
	// which rules end up in Accepted

	// filter out the lists of rules that will always fail

	frontier := make([]EvalSpec, 0)
	frontier = append(frontier, engine.rules["in"])

	for len(frontier) > 0 {
		currentNode := frontier[0]
		frontier = frontier[1:]

		// check where we can continue or is accepted
		for _, spec := range currentNode.Specs {
			if spec.result == "R" {
				continue
			} else if spec.result == "A" {
				fmt.Println("We have accepted in rule: ", currentNode.Name)
			} else {
				frontier = append(frontier, engine.rules[spec.result])
			}
		}
	}

	fmt.Println(engine.rules)

	return 0
}

func parseRawInputToItemsAndEngine(filename string) Engine {
	rawLines := utils.ParseFile(filename)

	splitPosition := slices.IndexFunc(rawLines, func(line string) bool {
		return line == ""
	})

	evalSpecs := make([]EvalSpec, 0)
	for _, rule := range rawLines[0:splitPosition] {
		evalSpec := ParseRule(rule)

		if !evalSpec.isDeadend() {
			evalSpecs = append(evalSpecs, evalSpec)
		}
	}

	items := make([]Item, 0)
	for _, item := range rawLines[splitPosition+1:] {
		parsedItem := ParseItem(item)
		items = append(items, parsedItem)
	}

	return Engine{
		rules:    evalSpecsToMap(evalSpecs),
		items:    items,
		accepted: make([]Item, 0),
	}
}

func evalSpecsToMap(specs []EvalSpec) map[string]EvalSpec {
	result := make(map[string]EvalSpec)
	for _, spec := range specs {
		result[spec.Name] = spec
	}

	return result
}

func ParseRule(line string) EvalSpec {
	splitted := strings.Split(line, "{")
	ruleName := splitted[0]

	cleanedFullspec := strings.Replace(splitted[1], "}", "", -1)

	rawSpecs := strings.Split(cleanedFullspec, ",")
	parsedSpecs := parseSpecs(rawSpecs)

	return EvalSpec{
		Name:  ruleName,
		Specs: parsedSpecs,
	}
}

func parseSpecs(lines []string) []Spec {
	specs := make([]Spec, 0)
	for _, line := range lines {
		specs = append(specs, parseSpec(line))
	}

	return specs
}

func parseSpec(line string) Spec {
	var splitted []string = nil
	var operator string = ""

	if strings.Contains(line, "<") {
		splitted = strings.Split(line, "<")
		operator = "<"
	} else if strings.Contains(line, ">") {
		splitted = strings.Split(line, ">")
		operator = ">"
	} else {
		operator = "jmp"
	}

	// return if only the jump operation
	if operator == "jmp" {
		return Spec{
			opertor: operator,
			result:  line,
		}
	}

	operandAndResult := strings.Split(splitted[1], ":")
	operand2, _ := strconv.Atoi(operandAndResult[0])

	return Spec{
		operand1: splitted[0],
		operand2: operand2,
		opertor:  operator,
		result:   operandAndResult[1],
	}
}

func ParseItem(line string) Item {
	cleanedItem := strings.Replace(line, "{", "", -1)
	cleanedItem = strings.Replace(line, "}", "", -1)
	splitted := strings.Split(cleanedItem, ",")

	xValue, _ := strconv.Atoi(strings.Split(splitted[0], "=")[1])
	mValue, _ := strconv.Atoi(strings.Split(splitted[1], "=")[1])
	aValue, _ := strconv.Atoi(strings.Split(splitted[2], "=")[1])
	sValue, _ := strconv.Atoi(strings.Split(splitted[3], "=")[1])

	return Item{
		x: xValue,
		m: mValue,
		a: aValue,
		s: sValue,
	}
}

func evalItem(engine *Engine, item Item) {
	ifRule(engine, item, engine.rules["in"])
}

func ifRule(engine *Engine, item Item, spec EvalSpec) {
	// Fetch the rule spec
	// Evaluate the rule spec

	// Operators:
	// <
	// >
	// jmp
	for _, spec := range spec.Specs {
		if spec.opertor == "<" {
			if item.getOperand(spec.operand1) < spec.operand2 {
				if spec.result == "A" {
					fmt.Println("Accepting item")
					engine.accepted = append(engine.accepted, item)
					break
				} else if spec.result == "R" {
					// do nothing
					fmt.Println("Rejecting item item")
					break
				}

				// Go to the next rule with name
				fmt.Println("going to rule: ", spec.result)
				ifRule(engine, item, engine.rules[spec.result])
				break
			} else {
				// Continue with the next rule
				continue
			}
		} else if spec.opertor == ">" {
			if item.getOperand(spec.operand1) > spec.operand2 {
				if spec.result == "A" {
					fmt.Println("Accepting item")
					engine.accepted = append(engine.accepted, item)
					break
				} else if spec.result == "R" {
					// do nothing
					fmt.Println("Rejecting item")
					break
				}

				// Go to the next rule with name
				fmt.Println("going to rule: ", spec.result)
				ifRule(engine, item, engine.rules[spec.result])
				break
			} else {
				// Continue with the next rule
				continue
			}
		} else if spec.opertor == "jmp" {
			// Jump to the rule if not equals A or R
			if spec.result == "A" {
				fmt.Println("Accepting item")
				engine.accepted = append(engine.accepted, item)
				break
			} else if spec.result == "R" {
				// Rejected
				fmt.Println("Rejected item: ", item)
				break
			} else {
				fmt.Println("Jumping to rule: ", spec.result)
				ifRule(engine, item, engine.rules[spec.result])
				break
			}

		} else {
			fmt.Println("unknown opertor:", spec.opertor)
			os.Exit(-1)
		}
	}
}

func (item *Item) getOperand(operand string) int {
	if operand == "x" {
		return item.x
	} else if operand == "m" {
		return item.m
	} else if operand == "a" {
		return item.a
	} else if operand == "s" {
		return item.s
	}

	os.Exit(-1)
	return 0
}

func (item *Item) getTotalPoints() int {
	return item.x + item.m + item.a + item.s
}

func (evalSpec *EvalSpec) isDeadend() bool {
	var isDeadEnd = true
	for _, spec := range evalSpec.Specs {
		if spec.result != "R" {
			isDeadEnd = false
			break
		}
	}

	return isDeadEnd
}
