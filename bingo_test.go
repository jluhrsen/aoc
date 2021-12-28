package main

import (
	"testing"
)

var bingoCards []bingoCard
var bingoNumbers []int

func TestReadBingoInput(t *testing.T) {
	bingoCards, bingoNumbers := ReadBingoInput("./bingo_test_input.txt")
	if len(bingoCards) != 6 {
		t.Errorf("expected 6 bingo cards, but got %d", len(bingoCards))
	}
	if len(bingoNumbers) != 27 {
		t.Errorf("expected 27 bingo numbers, but got %d", len(bingoNumbers))
	}

	if bingoCards[2].numbers[2][3].number != 26 {
		t.Errorf("expected card 3, row 3, column 4 to be 26, but got %d", bingoCards[2].numbers[2][3].number)
	}
}

func TestWinningRow(t *testing.T) {
	bingoCards, bingoNumbers := ReadBingoInput("./bingo_test_input.txt")

	// 4th card in test data has the first row matching the first 5 numbers
	// so use that one for this test
	oneWinner := []bingoCard{bingoCards[3]}
	winningCard, _ := ProcessBingoCards(oneWinner, bingoNumbers, true)

	if winningCard.matchedLine != true {
		t.Errorf("bingo card should be a winner")
	}
}

func TestWinningColumn(t *testing.T) {
	bingoCards, bingoNumbers := ReadBingoInput("./bingo_test_input.txt")

	// 5th and 6th cards in test data has a column matching the last 5 numbers
	// so use that one for this test
	winnerOne := []bingoCard{bingoCards[4]}
	winnerTwo := []bingoCard{bingoCards[4]}

	winningCard, _ := ProcessBingoCards(winnerOne, bingoNumbers, true)
	if winningCard.matchedLine != true {
		t.Errorf("bingo card should be a winner")
	}
	winningCard, _ = ProcessBingoCards(winnerTwo, bingoNumbers, true)
	if winningCard.matchedLine != true {
		t.Errorf("bingo card should be a winner")
	}
}

func TestWinningCard(t *testing.T) {
	bingoCards, bingoNumbers := ReadBingoInput("./bingo_test_input.txt")

	ProcessBingoCards(bingoCards, bingoNumbers, true)

	if bingoCards[0].matchedLine == true {
		t.Errorf("bingo card 1 should not be a winner")
	}
	if bingoCards[1].matchedLine == true {
		t.Errorf("bingo card 2 should not be a winner")
	}
	if bingoCards[2].matchedLine == true {
		t.Errorf("bingo card 3 should not be a winner")
	}
	if bingoCards[3].matchedLine != true {
		t.Errorf("bingo card 4 should be a winner")
	}
	if bingoCards[4].matchedLine == true {
		t.Errorf("bingo card 5 should not be a winner")
	}
	if bingoCards[5].matchedLine == true {
		t.Errorf("bingo card 6 should not be a winner")
	}

}

func TestUnmarkedNumbersSsum(t *testing.T) {
	bingoCards, bingoNumbers := ReadBingoInput("./bingo_test_input.txt")

	bingoCards = []bingoCard{bingoCards[0], bingoCards[1], bingoCards[2]}
	winningCard, _ := ProcessBingoCards(bingoCards, bingoNumbers, true)

    unmarkedNumberSum := SumUnmarkedNumbers(winningCard)

    if unmarkedNumberSum != 188 {
    	t.Errorf("expected 188, but got %d", unmarkedNumberSum)
	}

}

func TestWinningScore(t *testing.T) {
	bingoCards, bingoNumbers := ReadBingoInput("./bingo_test_input.txt")

	bingoCards = []bingoCard{bingoCards[0], bingoCards[1], bingoCards[2]}
	winningCard, finalNumberSelected := ProcessBingoCards(bingoCards, bingoNumbers, true)

	unmarkedNumberSum := SumUnmarkedNumbers(winningCard)
    winningScore := unmarkedNumberSum * finalNumberSelected

	if winningScore != 4512 {
		t.Errorf("expected 188, but got %d", winningScore)
	}
}