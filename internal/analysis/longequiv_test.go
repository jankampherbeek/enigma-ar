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

func TestLongEquivCalculationHappyFlow(t *testing.T) {

	longitude := 223.0
	obliquity := 23.437101628
	declination := -15.09002104
	expected := 220.884444444444

	pos := domain.DoublePosition{
		Id:        1,
		Position1: longitude,
		Position2: declination,
	}
	positions := []domain.DoublePosition{pos}

	leCalc := NewLongEquivCalculation()
	result, err := leCalc.CalcEquivalents(positions, obliquity)
	if err != nil {
		t.Fatalf("LongEquiv calculation failed, returned unexpected error %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("Expected 1 result, got %d", len(result))
	}

	if result[0].Id != 1 {
		t.Fatalf("Expected id 1, got %d", result[0].Id)
	}
	if math.Abs(result[0].Position-expected) > 1e-8 {
		t.Fatalf("Expected position %v, got %v", expected, result[0].Position)
	}

}

func TestLongEquivCalculationMultiplePositions(t *testing.T) {

	longitude := 223.0
	obliquity := 23.437101628
	declination := -15.09002104

	pos1 := domain.DoublePosition{
		Id:        1,
		Position1: longitude,
		Position2: declination,
	}
	pos2 := domain.DoublePosition{
		Id:        3,
		Position1: 10.0,
		Position2: 3.0,
	}

	positions := []domain.DoublePosition{pos1, pos2}

	leCalc := NewLongEquivCalculation()
	result, err := leCalc.CalcEquivalents(positions, obliquity)
	if err != nil {
		t.Fatalf("LongEquiv calculation failed, returned unexpected error %v", err)
	}
	if len(result) != 2 {
		t.Fatalf("Expected 2 resuls, got %d", len(result))
	}

	if result[1].Id != 3 {
		t.Fatalf("Expected id 1, got %d", result[0].Id)
	}
}

func TestLongEquivCalculation24Libra(t *testing.T) { // Test for fix 0001
	pos := domain.DoublePosition{
		Id:        10,
		Position1: 204.9008333333,
		Position2: 5.92,
	}
	positions := []domain.DoublePosition{pos}
	obliquity := 23.437101628
	expected := 164.970690431
	leCalc := NewLongEquivCalculation()
	result, err := leCalc.CalcEquivalents(positions, obliquity)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if math.Abs(result[0].Position-expected) > 1e-8 {
		t.Fatalf("Expected position %v, got %v", expected, result[0].Position)
	}
}
