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
	
	total_distance := 0
	for i:=0; i < len(left); i++ {
		distance := math.Abs(float64(left[i])-float64(right[i]))
		total_distance += int(distance)
	}
	fmt.Println(total_distance)
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