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

func TestMidpointPositionTooLarge(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 2, Position: 12.0},
		{Id: 3, Position: 400.0},
		{Id: 5, Position: 220.5},
	}
	mpCalc := NewMidpointService()
	result, err := mpCalc.Midpoints(positions)
	if err == nil {
		t.Errorf("MidpointList should have returned an error for a position that is too large")
	}
	if result != nil {
		t.Errorf("MidpointList should have returned a nil for a position that is too large")
	}
}

func TestMidpointPositionTooSmall(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 2, Position: 12.0},
		{Id: 3, Position: -4.0},
		{Id: 5, Position: 220.5},
	}
	mpCalc := NewMidpointService()
	result, err := mpCalc.Midpoints(positions)
	if err == nil {
		t.Errorf("Midpoints should have returned an error for a position that is too small")
	}
	if result != nil {
		t.Errorf("Midpoints should have returned a nil for a position that is too small")
	}
}

func TestMidpointTooFewItems(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 2, Position: 12.0},
	}
	mpCalc := NewMidpointService()
	result, err := mpCalc.Midpoints(positions)
	if err == nil {
		t.Errorf("Midpoints should have returned an error for too few items")
	}
	if result != nil {
		t.Errorf("Midpoints should have returned a nil for too few items")
	}
}

func TestOccupiedMidpointPositionTooLarge(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 2, Position: 50.5},
		{Id: 4, Position: 200.1},
		{Id: 7, Position: 525.0},
		{Id: 8, Position: 255.3},
	}
	orb := 1.0
	mpCalc := NewMidpointService()
	result, err := mpCalc.OccupiedMidpoints(positions, 360.0, orb)
	if err == nil {
		t.Errorf("OccupiedMidpoints should have returned an error for a position that is too large")
	}
	if result != nil {
		t.Errorf("OccupiedMidpoints should have returned a nil result for a position that is too large")
	}
}

func TestOccupiedMidpointPositionTooSmall(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 2, Position: 50.5},
		{Id: 4, Position: 200.1},
		{Id: 7, Position: -25.0},
		{Id: 8, Position: 255.3},
	}
	orb := 1.0
	mpCalc := MidpointService{}
	result, err := mpCalc.OccupiedMidpoints(positions, 360.0, orb)
	if err == nil {
		t.Errorf("OccupiedMidpoints should have returned an error for a position that is too small")
	}
	if result != nil {
		t.Errorf("OccupiedMidpoints should have returned a nil for a position that is too small")
	}
}

func TestOccupiedMidpointTooFewItems(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 2, Position: 50.5},
	}
	orb := 1.0
	mpCalc := MidpointService{}
	result, err := mpCalc.OccupiedMidpoints(positions, 360.0, orb)
	if err == nil {
		t.Errorf("OccupiedMidpoints should have returned an error for too few items")
	}
	if len(result) > 0 {
		t.Errorf("OccupiedMidpoints should have returned a nil for too few items")
	}
}

func TestOccupiedMidpointOrbTooSmall(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 2, Position: 50.5},
		{Id: 4, Position: 200.1},
		{Id: 7, Position: 225.0},
		{Id: 8, Position: 255.3},
	}
	orb := -1.0
	mpCalc := MidpointService{}
	result, err := mpCalc.OccupiedMidpoints(positions, 360.0, orb)
	if err == nil {
		t.Errorf("OccupiedMidpoints should have returned an error for an orb that is too small")
	}
	if len(result) > 0 {
		t.Errorf("OccupiedMidpoints should have returned a nil for an orb that is too small")
	}
}

func TestOccupiedMidpointOrbTooLarge(t *testing.T) {
	var positions = []domain.SinglePosition{
		{Id: 1, Position: 100.0},
		{Id: 2, Position: 50.5},
		{Id: 4, Position: 200.1},
		{Id: 7, Position: 225.0},
		{Id: 8, Position: 255.3},
	}
	orb := 12.0
	mpCalc := MidpointService{}
	result, err := mpCalc.OccupiedMidpoints(positions, 360.0, orb)
	if err == nil {
		t.Errorf("OccupiedMidpoints should have returned an error for an orb that is too large")
	}
	if len(result) > 0 {
		t.Errorf("OccupiedMidpoints should have returned a nil for an orb that is too large")
	}
}
