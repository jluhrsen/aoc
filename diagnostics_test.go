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

var diagnosticInput2 = []string{
	"001",
	"010",
	"100",
	"101",
}

var expectedGammaRate = 22
var expectedEpsilonRate = 9
var expectedPowerConsumption = 198
var expectedO2GeneratorRating = 23
var expectedCO2ScrubberRating = 10
var expectedLifeSupportRating = 230

func TestMaxInputReducer(t *testing.T) {

	bitString, _ := InputReducer(diagnosticInput2, "max")

	if bitString != "101" {
		t.Errorf("expected 101, but got %s", bitString)
	}
}
func TestMinInputReducer(t *testing.T) {

	bitString, _ := InputReducer(diagnosticInput2, "min")

	if bitString != "001" {
		t.Errorf("expected 001, but got %s", bitString)
	}
}
func TestLifeSupportRating(t *testing.T) {

	_, O2GeneratorRating := InputReducer(diagnosticInput, "max")
	_, CO2ScrubberRating := InputReducer(diagnosticInput, "min")
    LifeSupportRating := O2GeneratorRating * CO2ScrubberRating

	if O2GeneratorRating != expectedO2GeneratorRating {
		t.Errorf("expected %d, but got %d", expectedO2GeneratorRating, O2GeneratorRating)
	}
	if CO2ScrubberRating != expectedCO2ScrubberRating {
		t.Errorf("expected %d, but got %d", expectedCO2ScrubberRating, CO2ScrubberRating)
	}
	if LifeSupportRating != expectedLifeSupportRating {
		t.Errorf("expected %d, but got %d", expectedLifeSupportRating, LifeSupportRating)
	}
}

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

func TestMostCommonBits(t *testing.T) {

	mcb := FindMostCommonBit(diagnosticInput, 0)
	if mcb != 1 {
		t.Errorf("expected 1, but got %d", mcb)
	}
	mcb = FindMostCommonBit(diagnosticInput, 4)
	if mcb != 0 {
		t.Errorf("expected 0, but got %d", mcb)
	}
	mcb = FindMostCommonBit(diagnosticInput2, 0)
	if mcb != 1 {
		t.Errorf("expected 1, but got %d", mcb)
	}
	mcb = FindMostCommonBit(diagnosticInput2, 1)
	if mcb != 0 {
		t.Errorf("expected 0, but got %d", mcb)
	}
	mcb = FindMostCommonBit(diagnosticInput2, 2)
	if mcb != 1 {
		t.Errorf("expected 1, but got %d", mcb)
	}
}