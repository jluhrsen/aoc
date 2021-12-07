package main

import "testing"

/*
forward 5
down 5
forward 8
up 3
down 8
forward 2
*/

var movementInput = []movement{
	{"forward",5},
	{"down", 5},
	{"forward", 8},
	{"up", 3},
	{"down", 8},
	{"forward", 2},
}

var expectedHorizontalMovement = 15
var expectedVerticalMovement = 10
var expectedMultiplierOutput = 150

func TestPositionCalculatoin(t *testing.T) {

	horizontalMovement, verticalMovement := MoveSubmarine(movementInput)
	if horizontalMovement != expectedHorizontalMovement {
		t.Errorf("expected %d, but got %d",
			expectedHorizontalMovement, horizontalMovement)
	}
	if verticalMovement != expectedVerticalMovement {
		t.Errorf("expected %d, but got %d",
			expectedVerticalMovement, verticalMovement)
	}
	if (horizontalMovement * verticalMovement) !=  expectedMultiplierOutput {
		t.Errorf("expected %d, but got %d",
			expectedMultiplierOutput, (horizontalMovement * verticalMovement))

	}
}