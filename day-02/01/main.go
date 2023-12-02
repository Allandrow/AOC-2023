package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var MAX_COLOR_VALUES = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

const REGEX = `(?:(\d+) (red|green|blue))`

func main() {
	body, err := os.ReadFile("day-02/input.txt")
	logIfError(err)

	lines := strings.Split(string(body), "\n")
	regex := regexp.MustCompile(REGEX)

	var sum int

Line:
	for _, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			number, err := strconv.Atoi(match[1:][0])
			logIfError(err)
			colour := match[1:][1]

			if number > MAX_COLOR_VALUES[colour] {
				continue Line
			}
		}

		gameID := regexp.MustCompile(`\d+`).FindString(line)
		gameNum, err := strconv.Atoi(gameID)
		logIfError(err)

		sum += gameNum
	}

	fmt.Println(sum)
}

func logIfError(err error) {
	if err != nil {
		log.Fatalf("%s \n", err)
	}
}
