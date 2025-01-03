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

func TestMidpointsOrbTooLarge(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 2, Position: 4.0},
		{Id: 4, Position: 8.1},
		{Id: 7, Position: -6.0},
		{Id: 8, Position: 3.3},
	}
	orb := 12.0
	dmS := NewDeclinationMidpointService()
	result, err := dmS.DeclinationMidpoints(positions, orb)
	if err == nil {
		t.Errorf("DeclMidpoints should have returned an error for an orb that is too large")
	}
	if result != nil {
		t.Errorf("DeclMidpoints should have returned nil for an orb that is too large")
	}
}

func TestMidpointsOrbTooSmall(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 2, Position: 4.0},
		{Id: 4, Position: 8.1},
		{Id: 7, Position: -6.0},
		{Id: 8, Position: 3.3},
	}
	orb := 0.0
	dmS := NewDeclinationMidpointService()
	result, err := dmS.DeclinationMidpoints(positions, orb)
	if err == nil {
		t.Errorf("DeclMidpoints should have returned an error for an orb that is too small")
	}
	if result != nil {
		t.Errorf("DeclMidpoints should have returned nil for an orb that is too small")
	}
}

func TestMidpointsEmptyInput(t *testing.T) {
	var positions []domain.SinglePosition
	dmS := NewDeclinationMidpointService()
	orb := 0.5
	result, err := dmS.DeclinationMidpoints(positions, orb)

	if err == nil {
		t.Error("DeclMidpoints should have returned an error for empty input")
	}
	if result != nil {
		t.Errorf("DeclMidpoints should have returned nil for empty input")
	}
}

func TestMidpointsPositionTooLarge(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 2, Position: 4.0},
		{Id: 4, Position: 190.1},
		{Id: 7, Position: -6.0},
		{Id: 8, Position: 3.3},
	}
	orb := 0.5
	dmS := NewDeclinationMidpointService()
	result, err := dmS.DeclinationMidpoints(positions, orb)
	if err == nil {
		t.Errorf("DeclMidpoints should have returned an error for a position that is too large")
	}
	if result != nil {
		t.Errorf("DeclMidpoints should have returned nil for a position that is too large")
	}
}

func TestMidpointsPositionTooSmall(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 12.0},
		{Id: 2, Position: -184.0},
		{Id: 4, Position: 10.1},
		{Id: 7, Position: -6.0},
		{Id: 8, Position: 3.3},
	}
	orb := 0.9
	dmS := NewDeclinationMidpointService()
	result, err := dmS.DeclinationMidpoints(positions, orb)
	if err == nil {
		t.Errorf("DeclMidpoints should have returned an error for a position that is too small")
	}
	if result != nil {
		t.Errorf("DeclMidpoints should have returned nil for a position that is too small")
	}
}
