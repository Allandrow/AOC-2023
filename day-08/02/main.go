package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Values map[string][]string

func main() {
	body, _ := os.ReadFile("day-08/input.txt")
	blocks := strings.Split(string(body), "\n\n")

	instructions := blocks[0]
	lines := strings.Split(blocks[1], "\n")

	values, starters := setValueTuplesAndStarters(lines)
	steps := make([]int, len(starters))

	// Get the # of steps required to reach Z for each starting position
	for i, starter := range starters {
		count := stepsToZ(starter, instructions, values)
		steps[i] = count
	}

	fmt.Println("result: ", PPCM(steps...))
}

func stepsToZ(pos string, instructions string, values Values) int {
	var rl = map[string]int{"L": 0, "R": 1}
	var steps int
	for !strings.HasSuffix(pos, "Z") {
		index := steps % len(instructions)
		instruction := instructions[index]
		pos = values[pos][rl[string(instruction)]]
		steps++
	}

	return steps
}

func setValueTuplesAndStarters(lines []string) (Values, []string) {
	m := Values{}
	var starts []string
	regex := regexp.MustCompile(`\w+`)
	for _, line := range lines {
		strs := regex.FindAllString(line, -1)
		vals := make([]string, 2)
		pos := strs[0]
		vals[0] = strs[1]
		vals[1] = strs[2]
		m[pos] = vals

		if strings.HasSuffix(pos, "A") {
			starts = append(starts, pos)
		}
	}

	return m, starts
}

func PGCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func PPCM(ints ...int) int {
	a := ints[0]
	b := ints[1]
	arr := ints[2:]
	result := a * b / PGCD(a, b)

	for i := 0; i < len(arr); i++ {
		result = PPCM(result, arr[i])
	}

	return result
}
