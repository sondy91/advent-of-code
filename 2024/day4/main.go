package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	inputData, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("Couldn't read input file")
	}

	occurrences := 0
	fileScanner := bufio.NewScanner(inputData)
	fileScanner.Split(bufio.ScanLines)
	var reports []string

	for fileScanner.Scan() {
		reports = append(reports, fileScanner.Text())
	}
	inputData.Close()
	println("Occurrences of XMAS: ", occurrences)
}
