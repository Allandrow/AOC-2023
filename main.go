package main

import (
	"fmt"
	"log"
	"os"
)

func logIfError(err error) {
	if err != nil {
		log.Fatalf("%s \n", err)
	}
}

func main() {
	body, err := os.ReadFile("sample.txt")

	logIfError(err)
	fmt.Println(string(body))
}
