package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	body, _ := os.ReadFile("day-09/input.txt")
	var result int

	for _, line := range strings.Split(string(body), "\n") {
		strs := regexp.MustCompile(`[-\d]+`).FindAllString(line, -1)
		nums := ArrStringToInt(strs)
		firstNum := nums[0] - getFirstNumSum(nums)
		result += firstNum
	}

	fmt.Println(result)

}

func ArrStringToInt(strs []string) []int {
	result := make([]int, len(strs))
	for i, str := range strs {
		num, _ := strconv.Atoi(str)
		result[i] = num
	}

	return result
}

// Recursive function to get the sum of last array values
func getFirstNumSum(nums []int) int {
	numMap := make(map[int]bool)
	slice := make([]int, len(nums)-1)

	for i := 0; i < len(slice); i++ {
		diff := nums[i+1] - nums[i]
		slice[i] = diff
		numMap[diff] = true
	}

	if len(numMap) == 1 {
		return slice[0]
	}

	sum := slice[0] - getFirstNumSum(slice)
	return sum
}
