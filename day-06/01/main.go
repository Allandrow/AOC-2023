package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numRegex = regexp.MustCompile(`\d+`)

func main() {
	body, _ := os.ReadFile("day-06/input.txt")

	lines := strings.Split(string(body), "\n")

	times := numRegex.FindAllString(lines[0], -1)
	distances := numRegex.FindAllString(lines[1], -1)
	result := 1

	for i := 0; i < len(times); i++ {
		maxTime, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		var options int

		for j := 1; j < maxTime; j++ {
			if j*(maxTime-j) > distance {
				options++
			}
		}

		result *= options
	}

	fmt.Println(result)
}
