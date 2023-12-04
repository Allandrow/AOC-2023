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

		for _, num := range regex.FindAllString(line, -1)[1:] {
			if !nums[num] {
				nums[num] = true
			} else {
				matches++
			}
		}

		for j := 0; j < matches; j++ {
			if index := i + j + 1; index < len(lines) {
				arr[index] = arr[index] + 1 + arr[i]
			}
		}

		sum += 1 + arr[i]
	}

	fmt.Println(sum)
}
