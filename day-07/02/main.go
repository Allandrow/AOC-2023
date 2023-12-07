package main

import (
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type scoreAndBets struct {
	score int
	bet   int
}

func main() {
	body, _ := os.ReadFile("day-07/input.txt")
	lines := strings.Split(string(body), "\n")
	handScores := make([]scoreAndBets, len(lines))
	var sum int

	for i, line := range lines {
		values := strings.Split(line, " ")
		bet, _ := strconv.Atoi(values[1])
		hand := values[0]

		handScores[i].bet = bet
		handScores[i].score = getScore(hand)
	}

	slices.SortFunc(handScores, sortScore)
	for i, handScore := range handScores {
		sum += handScore.bet * (i + 1)
	}

	fmt.Println(sum)
}

var cardValues = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func getScore(hand string) int {
	cards := strings.Split(hand, "")
	occurences := make(map[string]int)
	max := 1
	var score int
	var wildcards int

	for i, card := range cards {
		cardValue := cardValues[card]

		if cardValue != 1 {
			occurences[card] = occurences[card] + 1
			if occurences[card] > max {
				max = occurences[card]
			}
		} else {
			wildcards++
		}

		exp := (4 - i) * 2
		score += cardValue * int(math.Pow10(exp))
	}

	max += wildcards

	switch len(occurences) {
	case 0:
		score += 6 * int(math.Pow10(10))
	case 1:
		score += 6 * int(math.Pow10(10))
	case 2:
		if max == 4 {
			score += 5 * int(math.Pow10(10))
		} else {
			score += 4 * int(math.Pow10(10))
		}
	case 3:
		if max == 3 {
			score += 3 * int(math.Pow10(10))
		} else {
			score += 2 * int(math.Pow10(10))
		}
	case 4:
		score += 1 * int(math.Pow10(10))
	}

	return score
}

func sortScore(a, b scoreAndBets) int {
	return cmp.Compare(a.score, b.score)
}
