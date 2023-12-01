package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const REGEX = `(one|two|three|four|five|six|seven|eight|nine|\d)(?:.*(one|two|three|four|five|six|seven|eight|nine|\d))?`

var convertedNums = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func logIfError(err error) {
	if err != nil {
		log.Fatalf("%s \n", err)
	}
}

func convertedOrString(str string) string {
	if convertedNums[str] == "" {
		return str
	}

	return convertedNums[str]
}

func main() {
	r, _ := regexp.Compile(REGEX)
	body, err := os.ReadFile("day-01/input.txt")
	logIfError(err)

	lines := strings.Split(string(body), "\n")

	var sum int

	for _, line := range lines {
		result := r.FindStringSubmatch(line)[1:3]

		var str string = convertedOrString(result[0])

		if result[1] == "" {
			str += convertedOrString(result[0])
		} else {
			str += convertedOrString(result[1])
		}

		num, err := strconv.Atoi(string(str))
		logIfError(err)
		sum += num
	}

	println(sum)
}
