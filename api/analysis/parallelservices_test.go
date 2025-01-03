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

func TestParallelsOrbTooSmall(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 3, Position: -12.8},
		{Id: 4, Position: 11.4},
		{Id: 7, Position: 23.5},
		{Id: 8, Position: 23.0},
	}
	orb := 0.0
	pService := NewParallelService()
	result, err := pService.Parallels(positions, orb)
	if err == nil {
		t.Errorf("Expected error for orb that was too small, but no error was returned")
	}
	if result != nil {
		t.Errorf("Expected nil for orb that was too small")
	}
}

func TestParallelsOrbTooLarge(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 3, Position: -12.8},
		{Id: 4, Position: 11.4},
		{Id: 7, Position: 23.5},
		{Id: 8, Position: 23.0},
	}
	orb := 11.0
	pService := NewParallelService()
	result, err := pService.Parallels(positions, orb)
	if err == nil {
		t.Errorf("Expected error for orb that was too large, but no error was returned")
	}
	if result != nil {
		t.Errorf("Expected nil for orb that was too large")
	}
}

func TestParallelsInsufficientData(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
	}
	orb := 1.0
	pService := NewParallelService()
	result, err := pService.Parallels(positions, orb)
	if err == nil {
		t.Errorf("Expected error for insufficient data, but no error was returned")
	}
	if result != nil {
		t.Errorf("Expected nil for insufficient data")
	}
}

func TestParallelsDataOutOfRange(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 3, Position: -12.8},
		{Id: 4, Position: 200.4},
		{Id: 7, Position: 23.5},
		{Id: 8, Position: 23.0},
	}
	orb := 1.0
	pService := NewParallelService()
	result, err := pService.Parallels(positions, orb)
	if err == nil {
		t.Errorf("Expected error for declination that was too large, but no error was returned")
	}
	if result != nil {
		t.Errorf("Expected nil for declination that was too large")
	}
}
