package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func logIfError(err error) {
	if err != nil {
		log.Fatalf("%s \n", err)
	}
}

func main() {
	body, err := os.ReadFile("sample.txt")

	logIfError(err)

	lines := strings.Split(string(body), "\n")

	for i, line := range lines {
		fmt.Printf("%v: %v\n", i, line)
	}
}
