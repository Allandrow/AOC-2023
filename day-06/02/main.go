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

	time, _ := strconv.Atoi(strings.Join(numRegex.FindAllString(lines[0], -1), ""))
	distance, _ := strconv.Atoi(strings.Join(numRegex.FindAllString(lines[1], -1), ""))
	var result int

	for i := 1; i < time; i++ {
		if i*(time-i) > distance {
			result++
		}
	}

	fmt.Println(result)
}
