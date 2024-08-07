package day5

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolvePart1(filename string) int64 {
	rawLines := utils.ParseFile(filename)

	seeds := parseSeeds(rawLines[0])
	parsedMappers := parseMappers(rawLines)

	var mappedLocations = utils.MapFunc1(seeds, func(seed int64) int64 {
		return SolveSeedToLocation(seed, parsedMappers)
	})

	return slices.Min(mappedLocations)
}

func SolvePart2(filename string) int64 {
	rawLines := utils.ParseFile(filename)

	seeds := parseSeeds(rawLines[0])
	parsedMappers := parseMappers(rawLines)
	seedPairs := []utils.Pair[int64, int64]{}
	for i := 0; i < len(seeds); i += 2 {
		var newPair = utils.Pair[int64, int64]{
			First:  seeds[i],
			Second: seeds[i+1],
		}

		seedPairs = append(seedPairs, newPair)
	}

	for _, seedPair := range seedPairs {
		fmt.Printf("First: %d Second: %d\n", seedPair.First, seedPair.Second)
	}

	var minLocation = int64(9223372036854775807)
	for _, seedPair := range seedPairs {
		for i := seedPair.First; i <= seedPair.First+seedPair.Second; i++ {
			var tempMin = SolveSeedToLocation(i, parsedMappers)
			if tempMin < minLocation {
				minLocation = tempMin
			}
		}
	}

	return minLocation
}

func SolveSeedToLocation(seed int64, mappers []Mapper) int64 {
	var currentResult int64 = seed
	for _, mapper := range mappers {
		currentResult = mapper.GetDestinationFromSource(currentResult)
	}

	return currentResult
}

type Entry struct {
	sourceStart      int64
	sourceStop       int64
	destinationStart int64
	destinationStop  int64
	length           int64
}

type Mapper struct {
	name    string
	entries []Entry
}

func (mapper Mapper) IsSourceMapped(source int64) bool {
	// check if source is in any ranges of the source entries
	var isMapped bool = false
	for _, entry := range mapper.entries {
		if source >= entry.sourceStart && source <= entry.sourceStop {
			isMapped = true
			break
		}
	}

	return isMapped
}

func (mapper Mapper) GetDestinationFromSource(source int64) int64 {
	if !mapper.IsSourceMapped(source) {
		return source
	}

	var destination int64
	for _, entry := range mapper.entries {
		if source >= entry.sourceStart && source <= entry.sourceStop {
			// Calculate the destination
			destination = entry.destinationStart + (source - entry.sourceStart)
			// Stop looking
			break
		}
	}

	return destination
}

func parseSeeds(rawSeeds string) []int64 {
	cleaned := strings.Replace(rawSeeds, "seeds: ", "", 1)
	splitted := strings.Split(cleaned, " ")

	return utils.MapFunc1(splitted, func(num string) int64 {
		converted, _ := strconv.ParseInt(num, 10, 64)
		return converted
	})
}

func parseMappers(rawLines []string) []Mapper {
	mapIndexes := utils.FindIndexes(rawLines, "map")
	var parsedMappers []Mapper
	for i := range mapIndexes {
		if i < (len(mapIndexes) - 1) {
			parsedMappers = append(parsedMappers, ParseMapper(rawLines[mapIndexes[i]:mapIndexes[i+1]-1]))
		} else {
			parsedMappers = append(parsedMappers, ParseMapper(rawLines[mapIndexes[i]:]))
		}
	}
	return parsedMappers
}

func ParseMapper(rawMapperLines []string) Mapper {
	mapperName := rawMapperLines[0]
	entries := ParseEntries(rawMapperLines[1:])

	return Mapper{
		name:    mapperName,
		entries: entries,
	}
}

func ParseEntries(rawEntries []string) []Entry {
	return utils.MapFunc1(rawEntries, func(entry string) Entry {
		return ParseEntry(entry)
	})
}

func ParseEntry(rawEntryLine string) Entry {
	splitted := strings.Split(strings.TrimSpace(rawEntryLine), " ")
	splitInts := utils.MapFunc1(splitted, func(val string) int64 {
		i, _ := strconv.ParseInt(val, 10, 64)

		return i
	})

	return Entry{
		sourceStart:      splitInts[1],
		sourceStop:       splitInts[1] + splitInts[2] - 1,
		destinationStart: splitInts[0],
		destinationStop:  splitInts[0] + splitInts[2] - 1,
		length:           splitInts[2],
	}
}
