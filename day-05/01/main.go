package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numRegex = regexp.MustCompile(`\d+`)

func main() {
	body, _ := os.ReadFile("day-05/input.txt")

	blocks := strings.Split(string(body), "\n\n")

	seedLine := numRegex.FindAllString(blocks[0], -1)
	values := make([]int, len(seedLine))
	mapper := setMapper(blocks[1:])

	// set initial values
	for i, seed := range seedLine {
		num, _ := strconv.Atoi(seed)
		values[i] = num
	}

	for _, step := range mapper {
	Nums:
		for j, num := range values {
			for _, instructions := range step {
				if isInRange(num, instructions) {
					if instructions[0] < instructions[1] {
						values[j] -= int(math.Abs(float64(instructions[0] - instructions[1])))
					} else {
						values[j] += int(math.Abs(float64(instructions[0] - instructions[1])))
					}
					continue Nums
				}
			}
		}
	}

	result := values[0]
	for _, num := range values[1:] {
		if result > num {
			result = num
		}
	}

	fmt.Println(result)
}

func setMapper(blocks []string) [][][]int {
	mapper := make([][][]int, len(blocks))

	for i, block := range blocks {
		lines := strings.Split(block, "\n")[1:]
		mapper[i] = make([][]int, len(lines))

		for j, line := range lines {
			mapper[i][j] = make([]int, 0)
			for _, ref := range numRegex.FindAllString(line, -1) {
				num, _ := strconv.Atoi(ref)
				mapper[i][j] = append(mapper[i][j], num)
			}
		}
	}

	return mapper
}

func isInRange(value int, nums []int) bool {
	return value >= nums[1] && value < nums[1]+nums[2]
}
