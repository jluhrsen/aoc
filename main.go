package main

import (
	"fmt"
)

func main() {

	getIntsFromFile("./day_1_input.txt")
	day_1_inputValues := getIntsFromFile("./day_1_input.txt")
	day_2_inputValues := getMovementsFromFile("./day_2_input.txt")
	day_3_inputValues := getDiagnosticsFromFile("./day_3_input.txt")

	// day 1
	numDepthIncreases := CountDepthIncreases(day_1_inputValues)
	fmt.Printf("\nFound %d number of depth increases", numDepthIncreases)
	numDepthWindowIncreases := CountDepthWindowIncreases(day_1_inputValues, 3)
	fmt.Printf("\nFound %d number of depth increases", numDepthWindowIncreases)

	// day 2
	horizontalMovement, verticalMovement := MoveSubmarine(day_2_inputValues)
	fmt.Printf("\nMultiplier value for movement: %d",
		(horizontalMovement * verticalMovement))

	// day 3
	gammaRate, epsilonRate := CalculatePowerConsumption(day_3_inputValues)
	fmt.Printf("\nPower Consumption is: %d",
		(gammaRate * epsilonRate))
	_, O2GeneratorRating := InputReducer(day_3_inputValues, "max")
	_, CO2ScrubberRating := InputReducer(day_3_inputValues, "min")
	// print LifeSupportRating
	fmt.Printf("\nLife Support Rating is: %d", O2GeneratorRating * CO2ScrubberRating)

	// day 4
	bingoCards, bingoNumbers := ReadBingoInput("./day_4_input.txt")
	winningCard, finalNumberSelected := ProcessBingoCards(bingoCards, bingoNumbers, true)
	unmarkedNumberSum := SumUnmarkedNumbers(winningCard)
	winningScore := unmarkedNumberSum * finalNumberSelected
	fmt.Printf("\nWinning score is %d", winningScore)
}
