package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile(`\d+`)
var order = []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}

type mapping struct {
	destination int
	source      int
	srcRange    int
}

type seed struct {
	start int
	end   int
}

type mappers map[string][]mapping

func main() {
	body, _ := os.ReadFile("day-05/input.txt")
	blocks := strings.Split(string(body), "\n\n")

	seeds := setSeeds(blocks[0])
	mappers := setMappers(blocks[1:])

	var i int = 0

	// iterate over each location until a match is found in seeds: this will be the lowest location possible
Locations:
	for {
		humidity := applyMapping(i, mappers["location"])
		temperature := applyMapping(humidity, mappers["humidity"])
		light := applyMapping(temperature, mappers["temperature"])
		water := applyMapping(light, mappers["light"])
		fertilizer := applyMapping(water, mappers["water"])
		soil := applyMapping(fertilizer, mappers["fertilizer"])
		seed := applyMapping(soil, mappers["soil"])

		for _, seedRange := range seeds {
			if isInRange(seed, seedRange) {
				break Locations
			}
		}
		i++
	}

	fmt.Println(i)
}

func isInRange(num int, numRange seed) bool {
	return numRange.start <= num && num <= numRange.end
}

func applyMapping(num int, mapper []mapping) int {
	for _, instruction := range mapper {
		if num >= instruction.destination && num < instruction.destination+instruction.srcRange {
			return instruction.source - instruction.destination + num
		}
	}

	return num
}

func setSeeds(line string) []seed {
	var m = make([]seed, 0)
	var j int
	seeds := regex.FindAllString(line, -1)
	nums := getNumbersFromStrings(seeds)

	for j < len(nums) {
		start := nums[j]
		end := nums[j+1] + nums[j]
		s := seed{start: start, end: end}
		m = append(m, s)
		j += 2
	}

	return m
}

func setMappers(blocks []string) mappers {
	var m = make(mappers, 7)
	for i, orderType := range order {
		lines := strings.Split(blocks[i], "\n")[1:]
		var orderMappings = make([]mapping, len(lines))
		for i, line := range lines {
			nums := regex.FindAllString(line, -1)
			destination, _ := strconv.Atoi(nums[0])
			source, _ := strconv.Atoi(nums[1])
			srcRange, _ := strconv.Atoi(nums[2])
			orderTypeMapping := mapping{destination: destination, source: source, srcRange: srcRange}
			orderMappings[i] = orderTypeMapping
		}
		m[orderType] = orderMappings
	}

	return m
}

func getNumbersFromStrings(strs []string) []int {
	nums := make([]int, len(strs))

	for i, str := range strs {
		num, _ := strconv.Atoi(str)
		nums[i] = num
	}

	return nums
}
