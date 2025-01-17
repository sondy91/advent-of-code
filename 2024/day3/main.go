package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputData, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal("Couldn't read input file")
	}

	r, _ := regexp.Compile(`(?m)mul\([0-9]{1,3},[0-9]{1,3}\)`)
	total := 0
	commands := r.FindAllString(string(inputData), -1)

	for i := 0; i < len(commands); i++ {
		numbers := strings.Split(commands[i], ",")

		left := strings.TrimLeft(numbers[0], "mul(")
		right := strings.TrimRight(numbers[1], ")")
		leftInt, _ := strconv.Atoi(left)
		rightInt, _ := strconv.Atoi(right)
		total += (leftInt * rightInt)
	}
	fmt.Println("Total from multiplication commands: ", total)
	part2Total := part2(string(inputData))
	fmt.Println("Part 2: Total from multiplication commands: ", part2Total)
}

func part2(inputData string) int {
	segments := strings.Split(inputData, "don't()")
	total := findTotalInString(segments[0])
	for _, segment := range segments[1:] {
		// Split on do() and skip that first one since don't() disables them
		// and process any others that are enabled
		_, doSegments, _ := strings.Cut(segment, "do()")
		total += findTotalInString(doSegments)
	}
	return total
}

func findTotalInString(inputData string) int {
	r, _ := regexp.Compile(`(?m)mul\([0-9]{1,3},[0-9]{1,3}\)`)
	total := 0
	commands := r.FindAllString(inputData, -1)

	for i := 0; i < len(commands); i++ {
		numbers := strings.Split(commands[i], ",")

		left := strings.TrimLeft(numbers[0], "mul(")
		right := strings.TrimRight(numbers[1], ")")
		leftInt, _ := strconv.Atoi(left)
		rightInt, _ := strconv.Atoi(right)
		total += (leftInt * rightInt)
	}
	return total
}
