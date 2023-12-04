package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	body, _ := os.ReadFile("day-04/input.txt")

	regex := regexp.MustCompile(`\d+`)
	var sum int

	for _, line := range strings.Split(string(body), "\n") {
		nums := map[string]bool{}
		var count int

		for _, num := range regex.FindAllString(line, -1)[1:] {
			if !nums[num] {
				nums[num] = true
			} else {
				count++
			}
		}

		if count > 0 {
			sum += int(math.Pow(2, float64(count-1)))
		}
	}

	fmt.Println(sum)
}
