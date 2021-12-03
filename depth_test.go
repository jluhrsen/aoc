package main

import "testing"

var input = []int{199,200,208,210,200,207,240,269,260,263}
var expectedDepthIncreases = 7
var expectedDepthWindowIncreases = 5
var windowSize = 3

func TestDepthIncreaseCount(t *testing.T) {

	numDepthIncreases := CountDepthIncreases(input)
	if numDepthIncreases != expectedDepthIncreases {
		t.Errorf("expected %d, but got %d",
			expectedDepthIncreases, numDepthIncreases)
	}

}

func TestDepthWindowIncreaseCount(t *testing.T) {

	numDepthWindowIncreases := CountDepthWindowIncreases(input, windowSize)
	if numDepthWindowIncreases != expectedDepthWindowIncreases {
		t.Errorf("expected %d, but got %d",
			expectedDepthWindowIncreases, numDepthWindowIncreases)
	}

}