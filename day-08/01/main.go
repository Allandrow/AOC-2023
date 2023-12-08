package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Values map[string][]string

func main() {
	var rl = map[string]int{"L": 0, "R": 1}
	body, _ := os.ReadFile("day-08/input.txt")
	blocks := strings.Split(string(body), "\n\n")

	instructions := blocks[0]
	lines := strings.Split(blocks[1], "\n")

	values := setValueTuples(lines)
	current := "AAA"
	var i int

	for current != "ZZZ" {
		index := i % len(instructions)
		instruction := instructions[index]
		current = values[current][rl[string(instruction)]]
		i++
	}

	fmt.Println(i)
}

func setValueTuples(lines []string) Values {
	m := Values{}
	regex := regexp.MustCompile(`\w+`)
	for _, line := range lines {
		strings := regex.FindAllString(line, -1)
		vals := make([]string, 2)
		vals[0] = strings[1]
		vals[1] = strings[2]
		m[strings[0]] = vals
	}

	return m
}
