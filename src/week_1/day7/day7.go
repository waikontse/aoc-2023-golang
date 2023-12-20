package day7

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Value         []string
	OriginalValue string
	Bid           int
	Strength      HandStrength
	Rank          int
	Occurence     map[string]int
}

func SolvePart1(filename string) int {
	var replacementStrings = map[string]string{
		"A": "Z",
		"K": "Y",
		"Q": "X",
		"J": "W",
		"T": "V",
	}
	rawLines := utils.ParseFile(filename)

	parsedHands := utils.MapFunc1(rawLines, func(rawLine string) Hand {
		return ParseRawLineToHand(rawLine, replacementStrings)
	})

	for i := range parsedHands {
		parsedHands[i].FillOccurrenceMap()
		parsedHands[i].DetermineHandStrength()
	}

	RankAllHands(parsedHands)

	sum := utils.MapFunc1(parsedHands, func(h Hand) int {
		return h.Bid * h.Rank
	})

	result := utils.SumInts(sum)

	return result
}

func SolvePart2(filename string) int {
	var replacementStrings = map[string]string{
		"A": "Z",
		"K": "Y",
		"Q": "X",
		"J": "1",
		"T": "V",
	}
	rawLines := utils.ParseFile(filename)

	parsedHands := utils.MapFunc1(rawLines, func(rawLine string) Hand {
		return ParseRawLineToHand(rawLine, replacementStrings)
	})

	for i := range parsedHands {
		parsedHands[i].FillOccurrenceMap()
		parsedHands[i].DetermineHandStrength()
		fmt.Println(parsedHands[i])
		parsedHands[i].UpdateHandStrengthIfPossible()
		fmt.Println(parsedHands[i])
	}

	RankAllHands(parsedHands)

	sum := utils.MapFunc1(parsedHands, func(h Hand) int {
		return h.Bid * h.Rank
	})

	result := utils.SumInts(sum)

	return result
}

func ParseRawLineToHand(rawLine string, mapper map[string]string) Hand {
	split := strings.Split(rawLine, " ")
	points, _ := strconv.ParseInt(split[1], 10, 64)
	normalizedHandValue := NormalizeValue(split[0], mapper)

	return Hand{
		Value:         strings.Split(normalizedHandValue, ""),
		OriginalValue: normalizedHandValue,
		Bid:           int(points),
	}
}

func (hand *Hand) FillOccurrenceMap() {
	hand.Occurence = utils.ToOccurrenceMap(hand.Value)
}

func (hand *Hand) DetermineHandStrength() {
	// Determine the hand string
	// 1. Put the chars into a map
	// 2. Figure out the strength
	// Five of a kinds: map has length 1

	// Four of a kind: map has length 2 (4 + 1)
	// Full house: map has length 2 (3 + 2)

	// Three of a kind: map has length 3
	// two pair: map has length 3

	// one pair: map has length 4

	// high card: map has length 5

	occurences := len(hand.Occurence)
	maxOccurrence := slices.Max(utils.MapValues(hand.Occurence))
	if occurences == 1 {
		hand.Strength = FIVE_OF_A_KIND
	} else if occurences == 4 {
		hand.Strength = ONE_PAIR
	} else if occurences == 5 {
		hand.Strength = HIGH_CARD
	} else if occurences == 2 {
		// Determine full house or four of a kind
		if maxOccurrence == 4 {
			hand.Strength = FOUR_OF_A_KIND
		} else {
			hand.Strength = FULL_HOUSE
		}
	} else if occurences == 3 {
		// Determine three of a kind or two pair
		if maxOccurrence == 3 {
			hand.Strength = THREE_OF_A_KIND
		} else {
			hand.Strength = TWO_PAIR
		}
	}
}

func (hand *Hand) UpdateHandStrengthIfPossible() {
	// Get hand strength
	// Get the total number of jokers
	// Can we upgrade based on hand strength and jokers?

	totalJokers, found := hand.Occurence["1"]
	if !found {
		return
	}

	if totalJokers == 1 {
		switch hand.Strength {
		case HIGH_CARD:
			hand.Strength = ONE_PAIR
		case ONE_PAIR:
			hand.Strength = THREE_OF_A_KIND
		case TWO_PAIR:
			hand.Strength = FULL_HOUSE
		case THREE_OF_A_KIND:
			hand.Strength = FOUR_OF_A_KIND
		case FULL_HOUSE:
			os.Exit(-1)
		case FOUR_OF_A_KIND:
			hand.Strength = FIVE_OF_A_KIND
		case FIVE_OF_A_KIND:
			os.Exit(-1)
		}
	} else if totalJokers == 2 {
		switch hand.Strength {
		case HIGH_CARD:
			os.Exit(-1)
		case ONE_PAIR:
			hand.Strength = THREE_OF_A_KIND
		case TWO_PAIR:
			hand.Strength = FOUR_OF_A_KIND
		case THREE_OF_A_KIND:
			os.Exit(-1)
		case FULL_HOUSE:
			hand.Strength = FIVE_OF_A_KIND
		case FOUR_OF_A_KIND:
			os.Exit(-1)
		case FIVE_OF_A_KIND:
			os.Exit(-1)
		}
	} else if totalJokers == 3 {
		switch hand.Strength {
		case HIGH_CARD:
			os.Exit(-1)
		case ONE_PAIR:
			os.Exit(-1)
		case TWO_PAIR:
			os.Exit(-1)
		case THREE_OF_A_KIND:
			hand.Strength = FOUR_OF_A_KIND
		case FULL_HOUSE:
			hand.Strength = FIVE_OF_A_KIND
		case FOUR_OF_A_KIND:
			os.Exit(-1)
		case FIVE_OF_A_KIND:
			os.Exit(-1)
		}
	} else if totalJokers == 4 {
		switch hand.Strength {
		case HIGH_CARD:
			os.Exit(-1)
		case ONE_PAIR:
			os.Exit(-1)
		case TWO_PAIR:
			os.Exit(-1)
		case THREE_OF_A_KIND:
			os.Exit(-1)
		case FULL_HOUSE:
			os.Exit(-1)
		case FOUR_OF_A_KIND:
			hand.Strength = FIVE_OF_A_KIND
		case FIVE_OF_A_KIND:
			os.Exit(-1)
		}
	}
}

type HandStrength int

const (
	FIVE_OF_A_KIND HandStrength = iota
	FOUR_OF_A_KIND
	FULL_HOUSE
	THREE_OF_A_KIND
	TWO_PAIR
	ONE_PAIR
	HIGH_CARD
)

func NormalizeValue(rawHand string, mapper map[string]string) string {
	var normalizedHand string = rawHand
	for k, v := range mapper {
		normalizedHand = strings.Replace(normalizedHand, k, v, -1)
	}

	return normalizedHand
}

type ByHand []*Hand

func (a ByHand) Len() int      { return len(a) }
func (a ByHand) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByHand) Less(i, j int) bool {
	return strings.Compare(a[i].OriginalValue, a[j].OriginalValue) == -1
}

func RankAllHands(hands []Hand) {
	// 1. Split the hands into strength group
	// 2. Rank all of them
	fiveOfAKind := utils.Filter(hands, func(h Hand) bool {
		if h.Strength == FIVE_OF_A_KIND {
			return true
		} else {
			return false
		}
	})

	fourOfAKind := utils.Filter(hands, func(h Hand) bool {
		if h.Strength == FOUR_OF_A_KIND {
			return true
		} else {
			return false
		}
	})

	fullHouse := utils.Filter(hands, func(h Hand) bool {
		if h.Strength == FULL_HOUSE {
			return true
		} else {
			return false
		}
	})

	threeOfAKind := utils.Filter(hands, func(h Hand) bool {
		if h.Strength == THREE_OF_A_KIND {
			return true
		} else {
			return false
		}
	})

	twoPair := utils.Filter(hands, func(h Hand) bool {
		if h.Strength == TWO_PAIR {
			return true
		} else {
			return false
		}
	})

	pair := utils.Filter(hands, func(h Hand) bool {
		if h.Strength == ONE_PAIR {
			return true
		} else {
			return false
		}
	})

	highCard := utils.Filter(hands, func(h Hand) bool {
		if h.Strength == HIGH_CARD {
			return true
		} else {
			return false
		}
	})

	startIndex := RankHandsWithinGroup(highCard, 1)
	startIndex = RankHandsWithinGroup(pair, startIndex)
	startIndex = RankHandsWithinGroup(twoPair, startIndex)
	startIndex = RankHandsWithinGroup(threeOfAKind, startIndex)
	startIndex = RankHandsWithinGroup(fullHouse, startIndex)
	startIndex = RankHandsWithinGroup(fourOfAKind, startIndex)
	startIndex = RankHandsWithinGroup(fiveOfAKind, startIndex)
}

func RankHandsWithinGroup(handsWithSameStrength []*Hand, handRankStartIndex int) int {
	sort.Sort(ByHand(handsWithSameStrength))

	var i int
	for i = 0; i < len(handsWithSameStrength); i++ {
		handsWithSameStrength[i].Rank = handRankStartIndex + i
	}

	return i + handRankStartIndex
}
