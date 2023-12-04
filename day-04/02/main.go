package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	body, _ := os.ReadFile("day-04/input.txt")
	regex := regexp.MustCompile(`\d+`)
	lines := strings.Split(string(body), "\n")
	var arr = make([]int, len(lines))
	var sum int

	for i, line := range lines {
		nums := map[string]bool{}
		var matches int
		var mult int

		for _, num := range regex.FindAllString(line, -1)[1:] {
			if !nums[num] {
				nums[num] = true
			} else {
				matches++
			}
		}

		mult = arr[i]

		for j := 0; j < matches; j++ {
			if index := i + j + 1; index < len(lines) {
				arr[index] = arr[index] + 1 + mult
			}
		}

		sum += 1 + mult
	}

	fmt.Println(sum)
}
