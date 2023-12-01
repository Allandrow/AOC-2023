package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const REGEX = `(\d)(?:.*(\d))?`

func logIfError(err error) {
	if err != nil {
		log.Fatalf("%s \n", err)
	}
}

func main() {
	r, _ := regexp.Compile(REGEX)
	body, err := os.ReadFile("input.txt")
	logIfError(err)

	lines := strings.Split(string(body), "\n")
	var sum int

	for _, line := range lines {
		result := r.FindString(line)

		num, err := strconv.Atoi(string(result[0]) + string(result[len(result)-1]))

		logIfError(err)
		sum += num
	}

	println(sum)
}
