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

func TestCalcMidpointsHappyFlow(t *testing.T) {

	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 2, Position: 4.0},
		{Id: 4, Position: 8.1},
		{Id: 7, Position: -6.0},
		{Id: 8, Position: 3.3},
	}
	orb := 1.0

	var expected = []domain.OccupiedMidpoint{{
		BaseMidpointPos1: domain.SinglePosition{Id: 1, Position: 12.0},
		BaseMidpointPos2: domain.SinglePosition{Id: 2, Position: 4.0},
		FocusPoint:       domain.SinglePosition{Id: 4, Position: 8.1},
		ActualOrb:        0.1,
		Exactness:        0.9,
	}, {
		BaseMidpointPos1: domain.SinglePosition{Id: 1, Position: 12.0},
		BaseMidpointPos2: domain.SinglePosition{Id: 7, Position: -6.0},
		FocusPoint:       domain.SinglePosition{Id: 2, Position: 4.0},
		ActualOrb:        1.0,
		Exactness:        0.0,
	}, {
		BaseMidpointPos1: domain.SinglePosition{Id: 1, Position: 12.0},
		BaseMidpointPos2: domain.SinglePosition{Id: 7, Position: -6.0},
		FocusPoint:       domain.SinglePosition{Id: 8, Position: 3.3},
		ActualOrb:        0.3,
		Exactness:        0.7,
	}, {
		BaseMidpointPos1: domain.SinglePosition{Id: 1, Position: 12.0},
		BaseMidpointPos2: domain.SinglePosition{Id: 8, Position: 3.3},
		FocusPoint:       domain.SinglePosition{Id: 4, Position: 8.1},
		ActualOrb:        0.45,
		Exactness:        0.55,
	}, {
		BaseMidpointPos1: domain.SinglePosition{Id: 2, Position: 4.0},
		BaseMidpointPos2: domain.SinglePosition{Id: 8, Position: 3.3},
		FocusPoint:       domain.SinglePosition{Id: 2, Position: 4.0},
		ActualOrb:        0.35,
		Exactness:        0.65,
	}, {
		BaseMidpointPos1: domain.SinglePosition{Id: 2, Position: 4.0},
		BaseMidpointPos2: domain.SinglePosition{Id: 8, Position: 3.3},
		FocusPoint:       domain.SinglePosition{Id: 8, Position: 3.3},
		ActualOrb:        0.35,
		Exactness:        0.65,
	},
	}

	dmCalc := DeclMidpointsCalculation{}
	result, err := dmCalc.CalcMidpoints(positions, orb)

	if err != nil {
		t.Fatalf("declination midpoints calculation failed, returned unexpected error %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("Expected %d results, got %d, result was %v", len(expected), len(result), result)
	}

	var difference float64
	for i := range result {
		difference = math.Abs(result[i].BaseMidpointPos1.Position - expected[i].BaseMidpointPos1.Position)
		if difference > 1e-8 {
			t.Errorf("Position mismatch for BaseMidpointPos1 at index %d: got %f, want %f",
				i, result[i].BaseMidpointPos1.Position, expected[i].BaseMidpointPos1.Position)
		}
		difference = math.Abs(result[i].BaseMidpointPos2.Position - expected[i].BaseMidpointPos2.Position)
		if difference > 1e-8 {
			t.Errorf("Position mismatch for BaseMidpointPos2 at index %d: got %f, want %f",
				i, result[i].BaseMidpointPos2.Position, expected[i].BaseMidpointPos2.Position)
		}
		difference = math.Abs(result[i].FocusPoint.Position - expected[i].FocusPoint.Position)
		if difference > 1e-8 {
			t.Errorf("Position mismatch for FocusPoint at index %d: got %f, want %f result: %v",
				i, result[i].FocusPoint.Position, expected[i].FocusPoint.Position, result)
		}
		if result[i].BaseMidpointPos1.Id != expected[i].BaseMidpointPos1.Id {
			t.Errorf("ID mismatch for BaseMidpointPos1 at index %d: got %d, want %d",
				i, result[i].BaseMidpointPos1.Id, expected[i].BaseMidpointPos1.Id)
		}
		if result[i].BaseMidpointPos2.Id != expected[i].BaseMidpointPos2.Id {
			t.Errorf("ID mismatch for BaseMidpointPos2 at index %d: got %d, want %d",
				i, result[i].BaseMidpointPos2.Id, expected[i].BaseMidpointPos2.Id)
		}
		if math.Abs(result[i].ActualOrb-expected[i].ActualOrb) > 1e-8 {
			t.Errorf("Actual orbs do not match at index %d, got %f, wanted %f", i, result[i].ActualOrb, expected[i].ActualOrb)
		}
		if math.Abs(result[i].Exactness-expected[i].Exactness) > 1e-8 {
			t.Errorf("Exactnesses do not match at index %d, got %f, wanted %f", i, result[i].Exactness, expected[i].Exactness)
		}
	}
}

func TestCalcMidpointsOrbTooLarge(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 2, Position: 4.0},
		{Id: 4, Position: 8.1},
		{Id: 7, Position: -6.0},
		{Id: 8, Position: 3.3},
	}
	orb := 12.0
	dmCalc := NewDeclMidpointsCalculation()
	result, err := dmCalc.CalcMidpoints(positions, orb)
	if err == nil {
		t.Errorf("CalcMidpoints should have returned an error for an orb that is too large")
	}
	if len(result) > 0 {
		t.Errorf("CalcMidpoints should have returned an empty result for an orb that is too large")
	}
}

func TestCalcMidpointsOrbTooSmall(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 2, Position: 4.0},
		{Id: 4, Position: 8.1},
		{Id: 7, Position: -6.0},
		{Id: 8, Position: 3.3},
	}
	orb := 0.0
	dmCalc := NewDeclMidpointsCalculation()
	result, err := dmCalc.CalcMidpoints(positions, orb)
	if err == nil {
		t.Errorf("CalcMidpoints should have returned an error for an orb that is too small")
	}
	if len(result) > 0 {
		t.Errorf("CalcMidpoints should have returned an empty result for an orb that is too small")
	}
}

func TestCalcMidpointsEmptyInput(t *testing.T) {
	var positions []domain.SinglePosition
	dmCalc := NewDeclMidpointsCalculation()
	result, err := dmCalc.CalcMidpoints(positions, 2.0)

	if err == nil {
		t.Error("Expected error for empty input, got nil")
	}
	if len(result) != 0 {
		t.Errorf("Expected empty result, got slice with length %d", len(result))
	}
}

func TestCalcMidpointsPositionTooLarge(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 2, Position: 4.0},
		{Id: 4, Position: 190.1},
		{Id: 7, Position: -6.0},
		{Id: 8, Position: 3.3},
	}
	orb := 0.5
	dmCalc := NewDeclMidpointsCalculation()
	result, err := dmCalc.CalcMidpoints(positions, orb)
	if err == nil {
		t.Errorf("CalcMidpoints should have returned an error for a position that is too large")
	}
	if len(result) > 0 {
		t.Errorf("CalcMidpoints should have returned an empty result for a position that is too large, was %d", len(result))
	}
}

func TestCalcMidpointsPositionTooSmall(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 2, Position: -184.0},
		{Id: 4, Position: 10.1},
		{Id: 7, Position: -6.0},
		{Id: 8, Position: 3.3},
	}
	orb := 0.9
	dmCalc := NewDeclMidpointsCalculation()
	result, err := dmCalc.CalcMidpoints(positions, orb)
	if err == nil {
		t.Errorf("CalcMidpoints should have returned an error for a position that is too small")
	}
	if len(result) > 0 {
		t.Errorf("CalcMidpoints should have returned an empty result for a position that is too small, was %d", len(result))
	}
}
