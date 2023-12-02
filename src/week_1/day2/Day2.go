package day2

import (
	"aoc-2023-golang/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Play struct {
	red   int
	green int
	blue  int
}

type Game struct {
	number int
	plays  []Play
}

func SolvePart1() {
	var rawLines = utils.ParseFile("day_2.txt")
	var parsedGames = parseIntoGames(rawLines)

	// Check which games can be played
	var allowablePlay = Play{red: 12, green: 13, blue: 14}
	var playableGames []int
	for _, game := range parsedGames {
		if isGamePlayable(game, allowablePlay) {
			playableGames = append(playableGames, game.number)
		}
	}

	sumOfIds := utils.SumInts(playableGames)
	fmt.Println("Sum of ids: ", sumOfIds)
}

func SolvePart2() {
	var rawLines = utils.ParseFile("day_2.txt")
	var parsedGames = parseIntoGames(rawLines)

	// Find the largest plays
	var largestPlays []Play
	for _, game := range parsedGames {
		largestPlays = append(largestPlays, findLargestPlays(game))
	}

	var powersOfPlays []int
	for _, play := range largestPlays {
		powersOfPlays = append(powersOfPlays, calcPowerOfPlay(play))
	}

	sumOfIds := utils.SumInts(powersOfPlays)
	fmt.Println("Sum of ids: ", sumOfIds)
}

func calcPowerOfPlay(play Play) int {
	return play.red * play.green * play.blue
}

func parseIntoGames(rawLines []string) []Game {
	var games []Game
	for _, rawLine := range rawLines {
		game := parseIntoGame(rawLine)
		games = append(games, game)
	}

	return games
}

func parseIntoGame(line string) Game {
	splitted := strings.Split(line, ":")

	first := splitted[0]
	gameNumber, _ := strconv.Atoi(first[5:])

	rest := splitted[1]
	splittedRest := strings.Split(rest, ";")
	parsedPlays := parseIntoPlays(splittedRest)

	return Game{number: gameNumber, plays: parsedPlays}
}

func parseIntoPlays(rawPlays []string) []Play {
	var parsedPlays []Play
	for _, rawPlay := range rawPlays {
		parsedPlay := parseIntoPlay(rawPlay)
		parsedPlays = append(parsedPlays, parsedPlay)
	}

	return parsedPlays
}

func parseIntoPlay(rawPlay string) Play {
	fmt.Println("Parsing play: ", rawPlay)
	splitCubes := strings.Split(rawPlay, ",")

	var parsedPlay Play
	for _, cube := range splitCubes {
		trimmedCube := strings.Trim(cube, " ")
		if strings.Contains(trimmedCube, "red") {
			redCubesCount, _ := strconv.Atoi(strings.Replace(trimmedCube, " red", "", 1))
			parsedPlay.red = redCubesCount
		} else if strings.Contains(trimmedCube, "green") {
			greenCubesCount, _ := strconv.Atoi(strings.Replace(trimmedCube, " green", "", 1))
			parsedPlay.green = greenCubesCount

		} else if strings.Contains(trimmedCube, "blue") {
			blueCubesCount, _ := strconv.Atoi(strings.Replace(trimmedCube, " blue", "", 1))
			parsedPlay.blue = blueCubesCount
		}
	}

	return parsedPlay
}

func isGamePlayable(game Game, allowablePlay Play) bool {
	var gameIsPlayable = true
	for _, play := range game.plays {
		gameIsPlayable = gameIsPlayable && isPlayAllowed(play, allowablePlay)
	}

	return gameIsPlayable
}

func isPlayAllowed(gamePlay Play, allowablePlay Play) bool {
	isRedPlayable := gamePlay.red <= allowablePlay.red
	isGreenPlayable := gamePlay.green <= allowablePlay.green
	isBluePlayable := gamePlay.blue <= allowablePlay.blue

	return isRedPlayable && isGreenPlayable && isBluePlayable
}

func findLargestPlays(game Game) Play {
	fmt.Println("Largest game: ", game)
	var largestPlay = Play{0, 0, 0}
	for _, play := range game.plays {
		if play.blue > largestPlay.blue {
			largestPlay.blue = play.blue
		}
		if play.red > largestPlay.red {
			largestPlay.red = play.red
		}
		if play.green > largestPlay.green {
			largestPlay.green = play.green
		}
	}

	fmt.Println("return largest play: ", largestPlay)
	return largestPlay
}
