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
	println("Input length: ", len(inputData))
	if err != nil {
		log.Fatal("Couldn't read input file")
	}

	r, _ := regexp.Compile(`(?m)mul\([0-9]{1,3},[0-9]{1,3}\)`)
	total := 0
	commands := r.FindAllString(string(inputData), -1)
	println(len(commands))
	for i := 0; i < len(commands); i++ {
		numbers := strings.Split(commands[i], ",")

		left := strings.TrimLeft(numbers[0], "mul(")
		right := strings.TrimRight(numbers[1], ")")
		leftInt, _ := strconv.Atoi(left)
		rightInt, _ := strconv.Atoi(right)
		total += (leftInt * rightInt)
	}
	fmt.Println("Total from multiplication commands: ", total)
}
