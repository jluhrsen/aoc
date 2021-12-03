package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInput(inputFile string) []int {

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

func main() {

	inputValues := getInput("./day_1_input.txt")
	// fmt.Print(inputValues)
	numDepthIncreases := CountDepthIncreases(inputValues)
	fmt.Printf("Found %d number of depth increases", numDepthIncreases)
	numDepthWindowIncreases := CountDepthWindowIncreases(inputValues, 3)
	fmt.Printf("Found %d number of depth increases", numDepthWindowIncreases)

}
