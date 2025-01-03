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

func TestLongEquivsEmptyInput(t *testing.T) {

	var positions []domain.DoublePosition
	obliquity := 23.437101628
	leS := NewLongEquivService()
	result, err := leS.LongEquivs(positions, obliquity)
	if err == nil {
		t.Fatalf("LongEquivs failed, expected error because of empty input, but no error was returned")
	}
	if result != nil {
		t.Fatalf("Expected nil, got %d", len(result))
	}
}

func TestLongEquivsDeclinationTooLarge(t *testing.T) {

	pos := domain.DoublePosition{
		Id:        3,
		Position1: 10.0,
		Position2: 184.0,
	}

	positions := []domain.DoublePosition{pos}
	obliquity := 23.437101628
	leS := NewLongEquivService()
	result, err := leS.LongEquivs(positions, obliquity)
	if err == nil {
		t.Fatalf("LongEquivs failed, expected error because of declination that was too large, but no error was returned")
	}
	if result != nil {
		t.Fatalf("Expected nil, got %d", len(result))
	}
}

func TestLongEquivsDeclinationTooSmall(t *testing.T) {

	pos := domain.DoublePosition{
		Id:        3,
		Position1: 10.0,
		Position2: -190.0,
	}

	positions := []domain.DoublePosition{pos}
	obliquity := 23.437101628
	leS := NewLongEquivService()
	result, err := leS.LongEquivs(positions, obliquity)
	if err == nil {
		t.Fatalf("LongEquivs failed, expected error because of declination that was too small, but no error was returned")
	}
	if result != nil {
		t.Fatalf("Expected nil, got %d", len(result))
	}
}

func TestLongEquivsLongitudeTooLarge(t *testing.T) {

	pos := domain.DoublePosition{
		Id:        3,
		Position1: 370.0,
		Position2: 184.0,
	}

	positions := []domain.DoublePosition{pos}
	obliquity := 23.437101628
	leS := NewLongEquivService()
	result, err := leS.LongEquivs(positions, obliquity)
	if err == nil {
		t.Fatalf("LongEquivs failed, expected error because of longitude that was too large, but no error was returned")
	}
	if result != nil {
		t.Fatalf("Expected nil, got %d", len(result))
	}
}

func TestLongEquivsObliquityTooSmall(t *testing.T) {

	pos := domain.DoublePosition{
		Id:        3,
		Position1: 170.0,
		Position2: 184.0,
	}

	positions := []domain.DoublePosition{pos}
	obliquity := 20.437101628
	leS := NewLongEquivService()
	result, err := leS.LongEquivs(positions, obliquity)
	if err == nil {
		t.Fatalf("LongEquivs failed, expected error because of obliquity that was too small, but no error was returned")
	}
	if result != nil {
		t.Fatalf("Expected nil, got %d", len(result))
	}
}

func TestLongEquivsObliquityTooLarge(t *testing.T) {

	pos := domain.DoublePosition{
		Id:        3,
		Position1: 170.0,
		Position2: 184.0,
	}

	positions := []domain.DoublePosition{pos}
	obliquity := 25.437101628
	leS := NewLongEquivService()
	result, err := leS.LongEquivs(positions, obliquity)
	if err == nil {
		t.Fatalf("LongEquivs failed, expected error because of obliquity that was too large, but no error was returned")
	}
	if result != nil {
		t.Fatalf("Expected nil, got %d", len(result))
	}
}
