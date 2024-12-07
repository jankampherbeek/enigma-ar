/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package analysis

import (
	"enigma-ar/domain"
	"math"
	"testing"
)

func TestCalcHarmonicsHappyFlow(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 4, Position: 2.0},
		{Id: 8, Position: 3.3},
	}
	harmNr := 2.0
	expected := []domain.SinglePosition{
		{Id: 1, Position: 200.0},
		{Id: 4, Position: 4.0},
		{Id: 8, Position: 6.6},
	}

	hCalc := NewHarmonicsCalculation()
	result, err := hCalc.CalcHarmonics(positions, harmNr)
	if err != nil {
		t.Fatalf("Harmonics calculation failed returned unexpected error %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("Expected %d results, got %d", len(expected), len(result))
	}

	for i := range result {
		difference := math.Abs(result[i].Position - expected[i].Position)
		if difference > 1e-8 {
			t.Errorf("Position mismatch at index %d: got %f, want %f",
				i, result[i].Position, expected[i].Position)
		}
		if result[i].Id != expected[i].Id {
			t.Errorf("ID mismatch at index %d: got %d, want %d",
				i, result[i].Id, expected[i].Id)
		}
	}
}

func TestCalcHarmonicsHarmNrTooLarge(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 4, Position: 2.0},
		{Id: 8, Position: 3.3},
	}
	harmNr := 220_000.0
	hCalc := NewHarmonicsCalculation()
	result, err := hCalc.CalcHarmonics(positions, harmNr)
	if err == nil {
		t.Errorf("CalcHarmonics should have returned an error for a harmoncNr that is too large")
	}
	if len(result) > 0 {
		t.Errorf("CalcHarmonics should have returned an empty result for a harmonicNr that is too large")
	}
}

func TestCalcHarmonicsHarmNrTooSmall(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 4, Position: 2.0},
		{Id: 8, Position: 3.3},
	}
	harmNr := 0.5
	hCalc := NewHarmonicsCalculation()
	result, err := hCalc.CalcHarmonics(positions, harmNr)
	if err == nil {
		t.Errorf("CalcHarmonics should have returned an error for a harmonicNr that is too small")
	}
	if len(result) > 0 {
		t.Errorf("CalcHarmonics should have returned an empty result for a harmonicNr that is too small,  was %d", len(result))
	}
}

func TestCalcHarmonicsEmptyInput(t *testing.T) {
	var positions []domain.SinglePosition
	hCalc := NewHarmonicsCalculation()
	result, err := hCalc.CalcHarmonics(positions, 2.0)

	if err == nil {
		t.Error("Expected error for empty input, got nil")
	}
	if len(result) != 0 {
		t.Errorf("Expected empty result, got slice with length %d", len(result))
	}
}

func TestCalcHarmonicsPositionTooLarge(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 4, Position: 380.0},
		{Id: 8, Position: 3.3},
	}
	harmNr := 2.0
	hCalc := NewHarmonicsCalculation()
	result, err := hCalc.CalcHarmonics(positions, harmNr)
	if err == nil {
		t.Errorf("CalcHarmonics should have returned an error for a position that is too large")
	}
	if len(result) > 0 {
		t.Errorf("CalcHarmonics should have returned an empty result for a position that is too large, was %d", len(result))
	}
}

func TestCalcHarmonicsPositionTooSmall(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: -100.0},
		{Id: 4, Position: 80.0},
		{Id: 8, Position: 3.3},
	}
	harmNr := 2.0
	hCalc := NewHarmonicsCalculation()
	result, err := hCalc.CalcHarmonics(positions, harmNr)
	if err == nil {
		t.Errorf("CalcHarmonics should have returned an error for a position that is too small")
	}
	if len(result) > 0 {
		t.Errorf("CalcHarmonics should have returned an empty result for a position that is too small")
	}
}
