package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type movement struct {
	direction string
	distance  int
}

type diagCounter struct {
	onesCount int
	zerosCount int
}

func getIntsFromFile(inputFile string) []int {

	fileContents := []int{}

	file, err := os.Open(inputFile)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line, _ :=	strconv.Atoi(scanner.Text())
			fileContents = append(fileContents, line)
		}
	}

	return fileContents
}

func getDiagnosticsFromFile(inputFile string) []string {

	fileContents := []string{}

	file, err := os.Open(inputFile)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line :=	scanner.Text()
			fileContents = append(fileContents, line)
		}
	}

	return fileContents
}

func getMovementsFromFile(inputFile string) []movement {

	fileContents := []movement{}

	file, err := os.Open(inputFile)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(file)
		var i = 0
		tmpMovement := movement{}
		for scanner.Scan() {
			input := strings.Fields(scanner.Text())
			tmpMovement.direction = input[0]
			tmpMovement.distance, _ = strconv.Atoi(input[1])
			fileContents = append(fileContents, tmpMovement)
			i += 1
		}
	}
	return fileContents
}

func CountDepthIncreases(measurements []int) int {
	var numDepthIncreases = 0

	for idx, m := range measurements {
		if idx == len(measurements) - 1 {
			break
		}
		if m < measurements[idx+1] {
			numDepthIncreases += 1
		}

	}

	return numDepthIncreases
}

func CountDepthWindowIncreases(measurements []int, windowSize int) int {
	// function is hardcoded to use a window size of 3. will need to
	// refactor to be able to use the argument as a variable size window
	var numDepthWindowIncreases = 0
	for idx, m := range measurements {
		if idx == len(measurements) - 3 {
			break
		}
		windowOne := m + measurements[idx+1] + measurements[idx+2]
		windowTwo := measurements[idx+1] + measurements[idx+2] + measurements[idx+3]
		if windowOne < windowTwo {
			numDepthWindowIncreases += 1
		}
	}

	return numDepthWindowIncreases
}

func MoveSubmarine(courseEntries []movement) (int, int) {

	var horizontalMovement = 0
	var totalVerticalMovement = 0
	var aim = 0

	for _, m := range courseEntries {
		if m.direction == "forward" {
			horizontalMovement += m.distance
			totalVerticalMovement += aim * m.distance
		} else {
			if m.direction == "down" {
				aim += m.distance
			} else {
				aim -= m.distance
			}
		}
	}
	return horizontalMovement, totalVerticalMovement

}


func FindMostCommonBit(diagnosticInput []string, bitPosition int) int {
	inputAnalysis := AnalyzeDiagnosticInputBitCounts(diagnosticInput)

	if inputAnalysis[bitPosition].onesCount >= inputAnalysis[bitPosition].zerosCount {
		return 1
	}
	return 0
}

func AnalyzeDiagnosticInputBitCounts(diagnosticInput []string) []diagCounter {
	var inputAnalysis = []diagCounter{}
	for i := 0; i < len(diagnosticInput[0]); i++ {
		inputAnalysis = append(inputAnalysis, diagCounter{0,0})
	}

	for _, line := range diagnosticInput {
		for idx, b := range line {
			if string(b) == "1" {
				inputAnalysis[idx].onesCount += 1
			} else {
				inputAnalysis[idx].zerosCount += 1
			}
		}
	}

	return inputAnalysis
}

func InputReducer(diagnosticInput []string, threshold string) (string, int) {

    var mcb int

	for i := 0; i < len(diagnosticInput[0]); i++ {
		mcb = FindMostCommonBit(diagnosticInput, i)
		tempBitStrings := []string{}
		for _, bitString := range diagnosticInput {
			tmpBit, _ := strconv.Atoi(string(bitString[i]))
			if tmpBit == mcb && threshold == "max" {
				tempBitStrings = append(tempBitStrings, bitString)
			} else if tmpBit != mcb && threshold == "min" {
				tempBitStrings = append(tempBitStrings, bitString)
			}
		}
		if len(tempBitStrings) == 1 {
            baseTenValue, _ := strconv.ParseInt(tempBitStrings[0], 2, 64)
			return tempBitStrings[0], int(baseTenValue)
		}
		diagnosticInput = tempBitStrings
	}

    return "error", 0
}

func CalculatePowerConsumption(diagnosticInput []string) (int, int) {

	inputAnalysis := AnalyzeDiagnosticInputBitCounts(diagnosticInput)

	var gammaBitString = ""
	var epsilonBitString = ""
	for _, bitCounts := range inputAnalysis {
		if bitCounts.onesCount > bitCounts.zerosCount {
			gammaBitString += "1"
			epsilonBitString += "0"
		} else {
			gammaBitString += "0"
			epsilonBitString += "1"
		}
	}
	gammaRate, _ := strconv.ParseInt(gammaBitString, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonBitString, 2, 64)

	return int(gammaRate), int(epsilonRate)

}