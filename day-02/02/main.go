package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const REGEX = `(?:(\d+) (red|green|blue))`

func main() {
	body, err := os.ReadFile("day-02/input.txt")
	logIfError(err)

	lines := strings.Split(string(body), "\n")
	regex := regexp.MustCompile(REGEX)

	var sum int

	for _, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)

		maxCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, match := range matches {
			number, err := strconv.Atoi(match[1:][0])
			logIfError(err)
			colour := match[1:][1]

			if maxCubes[colour] < number {
				maxCubes[colour] = number
			}

		}

		result := maxCubes["red"] * maxCubes["green"] * maxCubes["blue"]
		sum += result
	}

	fmt.Println(sum)
}

func logIfError(err error) {
	if err != nil {
		log.Fatalf("%s \n", err)
	}
}
