package main

import "testing"

var diagnosticInput = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

var expectedGammaRate = 22
var expectedEpsilonRate = 9
var expectedPowerConsumption = 198

func TestPowerConsumptionRate(t *testing.T) {
	gammaRate, epsilonRate := CalculatePowerConsumption(diagnosticInput)
	if gammaRate != expectedGammaRate {
		t.Errorf("expected %d, but got %d",
			expectedGammaRate, gammaRate)
	}
	if epsilonRate != expectedEpsilonRate {
		t.Errorf("expected %d, but got %d",
			expectedEpsilonRate, epsilonRate)
	}
	if (gammaRate * epsilonRate) !=  expectedPowerConsumption {
		t.Errorf("expected %d, but got %d",
			expectedPowerConsumption, (gammaRate * epsilonRate))
	}
}