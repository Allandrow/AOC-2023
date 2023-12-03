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
	symbolRegex := regexp.MustCompile("[*]")
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

	// parse every symbol
	for _, symbol := range m.symbols {

		// check the number of nums it is within the extended range of
		adjacentNumbers := getAdjacentNumbers(symbol, m.numbers)

		if len(adjacentNumbers) != 2 {
			continue
		}

		mult := adjacentNumbers[0].value * adjacentNumbers[1].value
		sum += mult
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

func getAdjacentNumbers(symbol pos, numbers []numberMatch) []numberMatch {
	var resultRange []numberMatch
	for _, match := range numbers {
		y := match.position.start.y
		startX := match.position.start.x
		endX := match.position.end.x
		numRange := createExtendedRange(y, startX, endX)

		if isInRange(symbol, numRange) {
			resultRange = append(resultRange, match)
		}
	}

	return resultRange
}
