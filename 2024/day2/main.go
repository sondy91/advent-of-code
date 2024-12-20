package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputData, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("Couldn't read input file")
	}

	safeReportCount := 0

	fileScanner := bufio.NewScanner(inputData)
	fileScanner.Split(bufio.ScanLines)
	var reports []string

	for fileScanner.Scan() {
		reports = append(reports, fileScanner.Text())
	}
	inputData.Close()

	for _, report := range reports {
		strLevels := strings.Split(report, " ")
		levels := make([]int, len(strLevels))
		//Convert strings to ints
		for i := range len(strLevels) {
			level, _ := strconv.Atoi(strLevels[i])
			levels[i] = level
		}
		reportSafe := true
		levelIncreasing := true
		levelChange := 0
		for i:=0; i < len(levels)-1; i++ {
			levelChange = levels[i+1] - levels[i]
			if i == 0 {
				levelIncreasing = levelChange > 0
			}
			if levelIncreasing != (levelChange > 0) {
				//fmt.Print("Report: ", reportIndex+1)
				//fmt.Print(" Report unsafe, Level Direction: ", levelIncreasing)
				//fmt.Println(" Current level direction: ", (levelChange > 0))
				reportSafe = false
				break
			}
			if levelChange == 0 {
				//fmt.Print("Report: ", reportIndex+1)
				//fmt.Println(" Report unsafe, Level change: ", levelChange)
				reportSafe = false
				break
			}
			if levelChange < 0 {
				levelChange = -levelChange
			}
			if levelChange < 1 || levelChange > 3 {
				//fmt.Print("Report: ", reportIndex+1)
				//fmt.Println(" Report unsafe, Level change: ", levelChange)
				reportSafe = false
				break
			}
		}
		if reportSafe {
			safeReportCount++
		}
	}
	fmt.Println("The number of safe reports is: ", safeReportCount)
}