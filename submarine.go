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

func main() {

	day_1_inputValues := getIntsFromFile("./day_1_input.txt")
	day_2_inputValues := getMovementsFromFile("./day_2_input.txt")
	// day 1
	numDepthIncreases := CountDepthIncreases(day_1_inputValues)
	fmt.Printf("\nFound %d number of depth increases", numDepthIncreases)
	numDepthWindowIncreases := CountDepthWindowIncreases(day_1_inputValues, 3)
	fmt.Printf("\nFound %d number of depth increases", numDepthWindowIncreases)

	// day 2
	horizontalMovement, verticalMovement := MoveSubmarine(day_2_inputValues)
	fmt.Printf("\nMultiplier value for movement: %d",
		(horizontalMovement * verticalMovement))

}
