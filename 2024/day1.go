package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
	"strconv"
)

func main() {
	left, right := read_input()
	
	slices.Sort(left)	
	slices.Sort(right)
	occurrenceMap := make(map[int]int)
	for _, num := range right {
		occurrenceMap[num] = occurrenceMap[num] + 1
	}
	
	similarityScore := 0
	totalDistance := 0
	for i:=0; i < len(left); i++ {
		similarityScore += left[i] * occurrenceMap[left[i]]
		distance := math.Abs(float64(left[i])-float64(right[i]))
		totalDistance += int(distance)
	}
	fmt.Println("Total distance: ", totalDistance)
	fmt.Println("Similarity score: ", similarityScore)
}

func read_input() ([]int, []int) {
	input_data, err := os.Open("day1_input.txt")
	check(err)
	fileScanner := bufio.NewScanner(input_data)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	input_data.Close()

	left := make([]int, len(fileLines))
	right := make([]int, len(fileLines))

	for i, line := range fileLines {
		line_numbers := strings.Fields(line)
		left[i], _ = strconv.Atoi(line_numbers[0])
		right[i], _ = strconv.Atoi(line_numbers[1])
	}

	return left, right
}

func check(err error) {
	if err != nil {
        log.Fatal(err)
    }
}