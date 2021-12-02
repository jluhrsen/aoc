package main

import "testing"

var input = []int{199,200,208,210,200,207,240,269,260,263}
var expectedOutput = 7

func TestDepthIncreaseCount(t *testing.T) {

	numDepthIncreases := CountDepthIncreases(input)
	if numDepthIncreases != expectedOutput {
		t.Errorf("expected %d, but got %d", numDepthIncreases, expectedOutput)
	}

}