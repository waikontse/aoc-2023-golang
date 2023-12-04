package day4

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//In the above example, card 1 has five winning numbers (41, 48, 83, 86, and 17)
//and eight numbers you have (83, 86, 6, 31, 17, 9, 48, and 53).
//Of the numbers you have, four of them (48, 83, 17, and 86) are winning numbers!
// That means card 1 is worth 8 points (1 for the first match, then doubled three times for each of the three matches after the first).
//
//Card 2 has two winning numbers (32 and 61), so it is worth 2 points.
//Card 3 has two winning numbers (1 and 21), so it is worth 2 points.
//Card 4 has one winning number (84), so it is worth 1 point.
//Card 5 has no winning numbers, so it is worth no points.
//Card 6 has no winning numbers, so it is worth no points.
//
//So, in this example, the Elf's pile of scratchcards is worth 13 points.
//
//Take a seat in the large pile of colorful cards. How many points are they worth in total?
//

type Raffle struct {
	winningNumbers map[int]struct{}
	numbers        map[int]struct{}
}

type Game struct {
	Id      int
	Raffles Raffle
}

func SolvePart1() int {
	rawLines := utils.ParseFile(`./resources/day_4.txt`)

	games := ParseLinesIntoGames(rawLines)

	var sum float64 = 0
	for _, game := range games {
		winningInts := game.getWinningNumbers()
		if len(winningInts) == 0 {
			continue
		}

		fmt.Println("winning ints: ", winningInts)
		fmt.Println("ints length: ", winningInts)
		points := math.Pow(2, float64(len(game.getWinningNumbers())-1))

		sum += points
	}

	fmt.Println("answer: ", sum)

	return int(sum)
}

func SolvePart2() int {
	rawLines := utils.ParseFile("./resources/day_4.txt")
	games := ParseLinesIntoGames(rawLines)

	var collectedScorecards []int = make([]int, len(games)*2)
	var winningIntPerScoreCard []int = make([]int, len(games))
	for i, game := range games {
		winningIntPerScoreCard[i] = len(game.getWinningNumbers())
	}

	// Update collectedScoreCards
	utils.InitArray(collectedScorecards, 1)
	// 1. pick 1 score card, update the collectedScoreCards
	for i := 0; i < len(games); i++ {
		repeat := winningIntPerScoreCard[i]
		incrementBy := collectedScorecards[i]
		utils.IncrementArray(collectedScorecards, i+1, repeat, incrementBy)
	}

	sum := utils.SumInts(collectedScorecards[0:len(games)])

	return sum
}

func ParseLinesIntoGames(rawLines []string) []Game {
	var parsedGames []Game
	for _, line := range rawLines {
		parsedGames = append(parsedGames, ParseLineIntoGame(line))
	}

	return parsedGames
}

func ParseLineIntoGame(rawLine string) Game {
	// Break up line into 2 parts sep: ':'
	// 1. Parse the game id
	// 2. parse the line of raffles
	splitLine := strings.Split(rawLine, ":")

	gameId := ParseGameId(splitLine[0])
	raffle := ParseRaffleLine(splitLine[1])

	return Game{Id: gameId, Raffles: raffle}
}

func ParseGameId(rawGameId string) int {
	rawId := strings.Replace(rawGameId, "Card ", "", 1)
	value, _ := strconv.Atoi(rawId)

	return value
}

func ParseRaffleLine(rawRaffleLine string) Raffle {
	// split line
	// trim
	// separate the numbers for winning numbers
	// separate the numbers for raffle numbers
	splitLine := strings.Split(rawRaffleLine, "|")

	trimmedWinningNumbers := strings.TrimSpace(splitLine[0])
	splitWinningNumbers := strings.Split(trimmedWinningNumbers, " ")
	winningNumbers := utils.ToIntSlice(splitWinningNumbers)
	winningNumbersMap := utils.ConvertToIntMap(winningNumbers)

	trimmedRaffleNumbers := strings.TrimSpace(splitLine[1])
	splitRaffleNumbers := strings.Split(trimmedRaffleNumbers, " ")
	raffleNumbers := utils.ToIntSlice(splitRaffleNumbers)
	raffleNumbersMap := utils.ConvertToIntMap(raffleNumbers)

	return Raffle{winningNumbers: winningNumbersMap, numbers: raffleNumbersMap}
}

func (game Game) getWinningNumbers() []int {
	var matchedNumbers []int
	for k := range game.Raffles.numbers {
		_, found := game.Raffles.winningNumbers[k]

		if found {
			matchedNumbers = append(matchedNumbers, k)
		}
	}

	return matchedNumbers
}
