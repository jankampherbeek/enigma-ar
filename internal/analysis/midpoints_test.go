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

func TestCalcMidpointListHappyFlow(t *testing.T) {

	var positions = []domain.SinglePosition{
		{Id: 2, Position: 12.0},
		{Id: 3, Position: 100.0},
		{Id: 5, Position: 220.5},
	}

	var expected = []domain.Midpoint{{
		Point1:      domain.SinglePosition{Id: 2, Position: 12.0},
		Point2:      domain.SinglePosition{Id: 3, Position: 100.0},
		MidpointPos: 56.0,
	}, {
		Point1:      domain.SinglePosition{Id: 2, Position: 12.0},
		Point2:      domain.SinglePosition{Id: 5, Position: 220.5},
		MidpointPos: 296.25,
	}, {
		Point1:      domain.SinglePosition{Id: 3, Position: 100.0},
		Point2:      domain.SinglePosition{Id: 5, Position: 220.5},
		MidpointPos: 160.25,
	},
	}

	mpCalc := MidpointsCalculation{}
	result, err := mpCalc.CalcMidpoints(positions)

	if err != nil {
		t.Fatalf("midpoints calculation failed, returned unexpected error %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("Expected %d results, got %d, result was %v", len(expected), len(result), result)
	}

	var difference float64
	for i := range result {
		difference = math.Abs(result[i].MidpointPos - expected[i].MidpointPos)
		if difference > 1e-8 {
			t.Errorf("Position mismatch for MidpointPos at index %d: got %f, want %f",
				i, result[i].MidpointPos, expected[i].MidpointPos)
		}
		if result[i].Point1.Id != expected[i].Point1.Id {
			t.Errorf("ID mismatch for point1 at index %d: got %d, want %d",
				i, result[i].Point1.Id, expected[i].Point1.Id)
		}
		if result[i].Point2.Id != expected[i].Point2.Id {
			t.Errorf("ID mismatch for point2 at index %d: got %d, want %d",
				i, result[i].Point1.Id, expected[i].Point1.Id)
		}
		difference = math.Abs(result[i].Point1.Position - expected[i].Point1.Position)
		if difference > 1e-8 {
			t.Errorf("Wrong position for point1 at index %d: got %f, want %f",
				i, result[i].Point1.Position, expected[i].Point1.Position)
		}
		difference = math.Abs(result[i].Point2.Position - expected[i].Point2.Position)
		if difference > 1e-8 {
			t.Errorf("Wrong position for point2 at index %d: got %f, want %f",
				i, result[i].Point2.Position, expected[i].Point2.Position)
		}

	}
}

func TestOccupiedMidpointsHappyFlow(t *testing.T) {

	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 2, Position: 50.5},
		{Id: 4, Position: 200.1},
		{Id: 7, Position: 125.0},
		{Id: 8, Position: 255.3},
	}
	orb := 1.0

	// expected midpoints:
	// 1 / 2 = 8
	// 2 / 4 = 7

	var expected = []domain.OccupiedMidpoint{{
		BaseMidpointPos1: domain.SinglePosition{Id: 1, Position: 100.0},
		BaseMidpointPos2: domain.SinglePosition{Id: 2, Position: 50.5},
		FocusPoint:       domain.SinglePosition{Id: 8, Position: 255.3},
		ActualOrb:        0.05,
		Exactness:        95.0,
	}, {
		BaseMidpointPos1: domain.SinglePosition{Id: 2, Position: 50.5},
		BaseMidpointPos2: domain.SinglePosition{Id: 4, Position: 200.1},
		FocusPoint:       domain.SinglePosition{Id: 7, Position: 125.0},
		ActualOrb:        0.3,
		Exactness:        70.0,
	},
	}

	mpCalc := MidpointsCalculation{}
	result, err := mpCalc.CalcOccupiedMidpoints(positions, domain.Dial360, orb)

	if err != nil {
		t.Fatalf("midpoints calculation failed, returned unexpected error %v", err)
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

func TestOccupiedMidpointDisc45(t *testing.T) {

	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 2, Position: 50.5},
		{Id: 4, Position: 200.1},
		{Id: 7, Position: 125.0},
		{Id: 8, Position: 255.3},
	}
	orb := 1.0

	var expected = []domain.OccupiedMidpoint{
		{
			BaseMidpointPos1: domain.SinglePosition{Id: 1, Position: 100.0},
			BaseMidpointPos2: domain.SinglePosition{Id: 2, Position: 50.5},
			FocusPoint:       domain.SinglePosition{Id: 8, Position: 255.3},
			ActualOrb:        0.05,
			Exactness:        95.0,
		}, {
			BaseMidpointPos1: domain.SinglePosition{Id: 2, Position: 50.5},
			BaseMidpointPos2: domain.SinglePosition{Id: 4, Position: 200.1},
			FocusPoint:       domain.SinglePosition{Id: 7, Position: 125.0},
			ActualOrb:        0.3,
			Exactness:        70.0,
		},
	}

	mpCalc := MidpointsCalculation{}
	result, err := mpCalc.CalcOccupiedMidpoints(positions, domain.Dial360, orb)

	if err != nil {
		t.Fatalf("midpoints calculation failed, returned unexpected error %v", err)
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
