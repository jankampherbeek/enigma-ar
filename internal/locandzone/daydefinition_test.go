/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package locandzone

import "testing"

func TestDayFromDefinitionFirstSun(t *testing.T) {
	ddh := NewDayDefHandling()
	expected := 2
	result, err := ddh.DayFromDefinition(2025, 2, "6>=1")
	if err != nil {
		t.Errorf("DayFromDefinition returned error: %v", err)
	}
	if result != expected {
		t.Errorf("DayFromDefinition returned %d, expected %d", result, expected)
	}
}

func TestDayFromDefinitionSecondSat(t *testing.T) {
	ddh := NewDayDefHandling()
	expected := 9
	result, err := ddh.DayFromDefinition(2026, 5, "5>=2")
	if err != nil {
		t.Errorf("DayFromDefinition returned error: %v", err)
	}
	if result != expected {
		t.Errorf("DayFromDefinition returned %d, expected %d", result, expected)
	}
}

func TestDayFromDefinitionLastWed(t *testing.T) {
	ddh := NewDayDefHandling()
	expected := 26
	result, err := ddh.DayFromDefinition(2025, 2, "last2")
	if err != nil {
		t.Errorf("DayFromDefinition returned error: %v", err)
	}
	if result != expected {
		t.Errorf("DayFromDefinition returned %d, expected %d", result, expected)
	}
}
