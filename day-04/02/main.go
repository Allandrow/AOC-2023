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

	for _, line := range lines {
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

		mult = arr[0]
		arr = arr[1:]

		for i := 0; i < matches; i++ {
			arr[i] = arr[i] + 1 + mult
		}

		sum += 1 + mult
	}

	fmt.Println(sum)
}
