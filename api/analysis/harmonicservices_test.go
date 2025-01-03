/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package apianalysis

import (
	"enigma-ar/domain"
	"testing"
)

func TestHarmonicsHarmNrTooLarge(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 4, Position: 2.0},
		{Id: 8, Position: 3.3},
	}
	harmNr := 220_000.0
	hs := NewHarmonicService()
	result, err := hs.Harmonics(positions, harmNr)
	if err == nil {
		t.Errorf("Harmonics should have returned an error for a harmoncNr that is too large")
	}
	if result != nil {
		t.Errorf("Harmonics should have returned nil for a harmonicNr that is too large")
	}
}

func TestHarmonicsHarmNrTooSmall(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 4, Position: 2.0},
		{Id: 8, Position: 3.3},
	}
	harmNr := 0.5
	hs := NewHarmonicService()
	result, err := hs.Harmonics(positions, harmNr)
	if err == nil {
		t.Errorf("Harmonics should have returned an error for a harmonicNr that is too small")
	}
	if result != nil {
		t.Errorf("Harmonics should have returned nil for a harmonicNr that is too small")
	}
}

func TestHarmonicsEmptyInput(t *testing.T) {
	var positions []domain.SinglePosition
	hs := NewHarmonicService()
	result, err := hs.Harmonics(positions, 2.0)
	if err == nil {
		t.Error("Expected error for empty input, got nil")
	}
	if result != nil {
		t.Errorf("Expected nil for empty input")
	}
}

func TestHarmonicsPositionTooLarge(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 4, Position: 380.0},
		{Id: 8, Position: 3.3},
	}
	harmNr := 2.0
	hs := NewHarmonicService()
	result, err := hs.Harmonics(positions, harmNr)
	if err == nil {
		t.Errorf("CalcHarmonics should have returned an error for a position that is too large")
	}
	if result != nil {
		t.Errorf("Harmonics should have returned nil for a position that is too large")
	}
}

func TestHarmonicsPositionTooSmall(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: -100.0},
		{Id: 4, Position: 80.0},
		{Id: 8, Position: 3.3},
	}
	harmNr := 2.0
	hs := NewHarmonicService()
	result, err := hs.Harmonics(positions, harmNr)
	if err == nil {
		t.Errorf("Harmonics should have returned an error for a position that is too small")
	}
	if result != nil {
		t.Errorf("Harmonics should have returned nil for a position that is too small")
	}
}
