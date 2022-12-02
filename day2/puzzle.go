package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	drawPoint = 3
	winPoint  = 6
)

var scores = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
}

var opponentOptions = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
}

var yourOptions = map[string]string{
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
}

var losesTo = map[string]string{
	"Rock":     "Scissors",
	"Scissors": "Paper",
	"Paper":    "Rock",
}

var wins = map[string]string{
	"Scissors": "Rock",
	"Paper":    "Scissors",
	"Rock":     "Paper",
}

var actions = map[string]func(opponent string) int{
	"X": func(op string) int { return scores[losesTo[op]] },         // should lose
	"Y": func(op string) int { return scores[op] + drawPoint },      // should draw
	"Z": func(op string) int { return scores[wins[op]] + winPoint }, // should win
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var scoreOne, scoreTwo int
	for scanner.Scan() {
		options := strings.Split(scanner.Text(), " ")

		your, opponent := yourOptions[options[1]], opponentOptions[options[0]]

		if isDraw(your, opponent) {
			scoreOne += drawPoint
		}

		if isWin(your, opponent) {
			scoreOne += winPoint
		}

		scoreOne += scores[your]

		scoreTwo += actions[options[1]](opponent)
	}

	fmt.Println(scoreOne, scoreTwo)
}

func isDraw(your, opponent string) bool {
	return your == opponent
}

func isWin(your, opponent string) bool {
	return losesTo[your] == opponent
}
