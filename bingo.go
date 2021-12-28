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

var initializedCard = bingoCard{
	matchedLine: false,
	numbers: [][]bingoNumber{
		{{-1,false}, {-1,false}, {-1,false}, {-1,false}, {-1,false}},
		{{-1,false}, {-1,false}, {-1,false}, {-1,false}, {-1,false}},
		{{-1,false}, {-1,false}, {-1,false}, {-1,false}, {-1,false}},
		{{-1,false}, {-1,false}, {-1,false}, {-1,false}, {-1,false}},
		{{-1,false}, {-1,false}, {-1,false}, {-1,false}, {-1,false}},
	},
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
	//  all cards will be marked.
	//
	//  the card returned will be:
	//  - if no winner is found, then an initiazed card with unselecte numbers
	//    all set to -1.
	//  - exitOnWinner is true, the first winning card found is returned and.
	//  - exitOnWinner is false, the last winning card that is found will be
	//    returned.
	winningCard := initializedCard
	winningNumber := -1
	for _, n := range bingoNumbers {
		for idx, bc := range bingoCards {
			// since we process all cards with all numbers we don't want to
			// mark any numbers on a card that has already proven to be a
			// winner
			if bingoCards[idx].matchedLine != true {
				MarkNumber(&bc, n)
			}
			if WinningCard(bc) == true && bingoCards[idx].matchedLine != true {
				bingoCards[idx].matchedLine = true
				// if we want to find the last winning card, we need to keep
				// updating winners as we find them
				if exitOnWinner == true && winningNumber == -1 {
					winningCard = bingoCards[idx]
					winningNumber = n
				} else if exitOnWinner == false {
					winningCard = bingoCards[idx]
					winningNumber = n
				}
			}
		}
	}

	return winningCard, winningNumber
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