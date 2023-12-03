package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

type numLocation struct {
	start pos
	end   pos
}

type numberMatch struct {
	value    int
	position numLocation
}

type matches struct {
	symbols []pos
	numbers []numberMatch
}

type extendedRange struct {
	min pos
	max pos
}

func main() {
	body, _ := os.ReadFile("day-03/input.txt")

	lines := strings.Split(string(body), "\n")
	symbolRegex := regexp.MustCompile("[^0-9.]")
	numRegex := regexp.MustCompile(`\d+`)
	m := matches{}
	var sum int

	// parse every line to register symbols and numbers in matches
	for y, line := range lines {
		symbolsIndices := symbolRegex.FindAllStringIndex(line, -1)
		numberIndices := numRegex.FindAllStringIndex(line, -1)
		numberValues := numRegex.FindAllString(line, -1)

		// register the position of each symbol in matches
		for _, indices := range symbolsIndices {
			m.symbols = append(m.symbols, pos{x: indices[0], y: y})
		}

		// register the value and start+end positions of each number in matches
		for i := 0; i < len(numberIndices); i++ {
			numIndices := numberIndices[i]
			numValue, _ := strconv.Atoi(numberValues[i])

			start := pos{x: numIndices[0], y: y}
			end := pos{x: numIndices[1] - 1, y: y}
			match := numberMatch{
				value:    numValue,
				position: numLocation{start: start, end: end},
			}

			m.numbers = append(m.numbers, match)

		}
	}

	// parse every matched number
Number:
	for _, numMatch := range m.numbers {
		y := numMatch.position.start.y
		startX := numMatch.position.start.x
		endX := numMatch.position.end.x
		numRange := createExtendedRange(y, startX, endX)

		// find if any symbol is within this range
		for _, symbol := range m.symbols {
			// add to sum and skip the remaining symbols
			if isInRange(symbol, numRange) {
				sum += numMatch.value
				continue Number
			}
		}
	}

	fmt.Println(sum)
}

func createExtendedRange(y int, startX int, endX int) extendedRange {
	return extendedRange{
		min: pos{x: startX - 1, y: y - 1},
		max: pos{x: endX + 1, y: y + 1},
	}
}

func isInRange(symbol pos, numRange extendedRange) bool {
	if symbol.x < numRange.min.x || symbol.x > numRange.max.x {
		return false
	}

	if symbol.y < numRange.min.y || symbol.y > numRange.max.y {
		return false
	}

	return true
}
