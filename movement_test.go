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

var expectedHorizontalMovementWithAim = 15
var expectedVerticalMovementWithAim = 60
var expectedMultiplierOutputWithAim = 900

func TestPositionCalculationUsingAim(t *testing.T) {
	horizontalMovementWithAim, verticalMovementWithAim := MoveSubmarine(movementInput)
	if horizontalMovementWithAim != expectedHorizontalMovementWithAim {
		t.Errorf("expected %d, but got %d",
			expectedHorizontalMovementWithAim, horizontalMovementWithAim)
	}
	if verticalMovementWithAim != expectedVerticalMovementWithAim {
		t.Errorf("expected %d, but got %d",
			expectedVerticalMovementWithAim, verticalMovementWithAim)
	}
	if (horizontalMovementWithAim * verticalMovementWithAim) !=  expectedMultiplierOutputWithAim {
		t.Errorf("expected %d, but got %d",
			expectedMultiplierOutputWithAim, (horizontalMovementWithAim * verticalMovementWithAim))
	}
}