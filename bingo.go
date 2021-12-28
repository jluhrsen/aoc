package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bingoCard struct {
	matchedLine bool
	numbers [][]bingoNumber
}

type bingoNumber struct {
	number int
	selected bool
}

func ReadBingoInput(inputFile string) ([]bingoCard, []int) {

	bingoCards := []bingoCard{}
    pickedNumbers := []int{}

	file, err := os.Open(inputFile)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		firstLine := strings.Fields(scanner.Text())
		tmpNums := strings.Split(firstLine[0], ",")

		for i := 0; i < len(tmpNums); i++ {
			n, _ := strconv.Atoi(tmpNums[i])
			pickedNumbers = append(pickedNumbers, n)
		}

		var numCards = 0
		for scanner.Scan() {
			row := strings.Fields(scanner.Text())
			// assumption is that a blank line signifies a new bingo card is next
			if len(row) == 0 {
				numCards += 1
				bingoCards = append(bingoCards, bingoCard{})
				bingoCards[numCards - 1].numbers = [][]bingoNumber{
												{{-1,false}, {-1,false}, {-1,false}, {-1,false}, {-1,false}},
                            					{{-1,false}, {-1,false}, {-1,false}, {-1,false}, {-1,false}},
					                            {{-1,false}, {-1,false}, {-1,false}, {-1,false}, {-1,false}},
                            					{{-1,false}, {-1,false}, {-1,false}, {-1,false}, {-1,false}},
					                            {{-1,false}, {-1,false}, {-1,false}, {-1,false}, {-1,false}},
												}
			}

			for i := 0; i < 5; i++ {
				scanner.Scan()
				row = strings.Fields(scanner.Text())
				rowNumbers := []int{}
				for i := 0; i < len(row); i++ {
					n, _ := strconv.Atoi(row[i])
					rowNumbers = append(rowNumbers, n)
				}
				for idx, n := range rowNumbers {
					bingoCards[numCards - 1].numbers[i][idx].number = n
				}
			}
		}
	}
	return bingoCards, pickedNumbers
}

func ProcessBingoCards(bingoCards []bingoCard, bingoNumbers []int, exitOnWinner bool) (bingoCard, int) {
	// caller needs to validate if the returned card is a winner or not since
	// it's default value is the first card in the slice, just to initialize it
	// also, if there are multiple winners, the last one will be what is returned
	winner := bingoCards[0]
	for _, n := range bingoNumbers {
		for idx, bc := range bingoCards {
			MarkNumber(&bc, n)
			if WinningCard(bc) == true {
				bingoCards[idx].matchedLine = true
				winner = bingoCards[idx]
				if exitOnWinner == true {
					return winner, n
				}
			}
		}
	}

	return winner, -1
}

func MarkNumber(card *bingoCard, num int) {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if card.numbers[row][col].number == num {
				card.numbers[row][col].selected = true
			}
		}
	}
}

func WinningCard(card bingoCard) bool {
	for row := 0; row < 5; row++ {
		if card.numbers[row][0].selected == true &&
			card.numbers[row][1].selected == true &&
			card.numbers[row][2].selected == true &&
			card.numbers[row][3].selected == true &&
			card.numbers[row][4].selected == true {
			return true
		}
	}
	for col := 0; col < 5; col++ {
		if card.numbers[0][col].selected == true &&
			card.numbers[1][col].selected == true &&
			card.numbers[2][col].selected == true &&
			card.numbers[3][col].selected == true &&
			card.numbers[4][col].selected == true {
			return true
		}
	}

	return false
}

func SumUnmarkedNumbers(card bingoCard) int {
	sum := 0

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if card.numbers[row][col].selected == false {
				sum = sum + card.numbers[row][col].number
			}
		}
	}
	return sum
}