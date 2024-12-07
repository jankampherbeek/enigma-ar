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

func TestCalcParallelsHappyFlow(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 3, Position: -12.8},
		{Id: 4, Position: 11.4},
		{Id: 7, Position: 23.5},
		{Id: 8, Position: 23.0},
	}

	orb := 1.0
	expected := []domain.MatchedParallel{
		{Pos1: domain.SinglePosition{Id: 1, Position: 12.0}, Pos2: domain.SinglePosition{Id: 3, Position: -12.8}, Orb: 0.8, Parallel: false},
		{Pos1: domain.SinglePosition{Id: 1, Position: 12.0}, Pos2: domain.SinglePosition{Id: 4, Position: 11.4}, Orb: 0.6, Parallel: true},
		{Pos1: domain.SinglePosition{Id: 7, Position: 23.5}, Pos2: domain.SinglePosition{Id: 8, Position: 23.0}, Orb: 0.5, Parallel: true},
	}

	pCalc := NewParallelsCalculation()
	result, err := pCalc.CalcParallels(positions, orb)
	if err != nil {
		t.Fatalf("Parallels calculation failed returned unexpected error %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("Expected %d results, got %d", len(expected), len(result))
	}

	for i := range result {
		diffOrb := math.Abs(result[i].Orb - expected[i].Orb)
		if diffOrb > 1e-8 {
			t.Errorf("Mismatch for orbat index %d: got %f, want %f", i, result[i].Orb, expected[i].Orb)
		}
		if result[i].Pos1.Id != expected[i].Pos1.Id {
			t.Errorf("ID mismatch at index %d: got %d, want %d", i, result[i].Pos1.Id, expected[i].Pos1.Id)
		}
		if result[i].Pos2.Id != expected[i].Pos2.Id {
			t.Errorf("ID mismatch at index %d: got %f, want %f", i, result[i].Pos2.Position, expected[i].Pos2.Position)
		}
		if math.Abs(result[i].Pos1.Position-expected[i].Pos1.Position) > 1e-8 {
			t.Errorf("Position mismatch at index %d: got %d, want %d", i, result[i].Pos1.Id, expected[i].Pos1.Id)
		}
		if math.Abs(result[i].Pos2.Position-expected[i].Pos2.Position) > 1e-8 {
			t.Errorf("Position mismatch at index %d: got %f, want %f", i, result[i].Pos2.Position, expected[i].Pos2.Position)
		}
	}

}

func TestCalcParallelsOrbTooSmall(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 3, Position: -12.8},
		{Id: 4, Position: 11.4},
		{Id: 7, Position: 23.5},
		{Id: 8, Position: 23.0},
	}
	orb := 0.0
	pCalc := NewParallelsCalculation()
	result, err := pCalc.CalcParallels(positions, orb)
	if err == nil {
		t.Errorf("Expected eror for orb that was too small, but no error was returned")
	}
	if len(result) != 0 {
		t.Errorf("Expected empty slice for orb that was too small, got %d", len(result))
	}
}

func TestCalcParallelsOrbTooLarge(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 3, Position: -12.8},
		{Id: 4, Position: 11.4},
		{Id: 7, Position: 23.5},
		{Id: 8, Position: 23.0},
	}
	orb := 11.0
	pCalc := NewParallelsCalculation()
	result, err := pCalc.CalcParallels(positions, orb)
	if err == nil {
		t.Errorf("Expected eror for orb that was too large, but no error was returned")
	}
	if len(result) != 0 {
		t.Errorf("Expected empty slice for orb that was too large, got %d", len(result))
	}
}

func TestCalcParallelsINsufficientData(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
	}
	orb := 1.0
	pCalc := NewParallelsCalculation()
	result, err := pCalc.CalcParallels(positions, orb)
	if err == nil {
		t.Errorf("Expected eror for insufficient data, but no error was returned")
	}
	if len(result) != 0 {
		t.Errorf("Expected empty slice for insufficient data, got %d", len(result))
	}
}

func TestCalcParallelsDataOutOfRange(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 3, Position: -12.8},
		{Id: 4, Position: 200.4},
		{Id: 7, Position: 23.5},
		{Id: 8, Position: 23.0},
	}
	orb := 1.0
	pCalc := NewParallelsCalculation()
	result, err := pCalc.CalcParallels(positions, orb)
	if err == nil {
		t.Errorf("Expected eror for declination that was too large, but no error was returned")
	}
	if len(result) != 0 {
		t.Errorf("Expected empty slice for declination that was too large, got %d", len(result))
	}
}
