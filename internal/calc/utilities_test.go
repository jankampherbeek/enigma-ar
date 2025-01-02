/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	"enigma-ar/domain"
	"math"
	"testing"
)

func TestSeFlagsEcliptical(t *testing.T) {
	result := SeFlags(domain.CoordEcliptical, domain.ObsPosGeocentric, true)
	expected := domain.SeflgSwieph + domain.SeflgSpeed
	if result != expected {
		t.Errorf("SeFlags() for ecliptical = %v, want %v", result, expected)
	}
}

func TestSeFlagsEquatorial(t *testing.T) {
	result := SeFlags(domain.CoordEquatorial, domain.ObsPosGeocentric, true)
	expected := domain.SeflgSwieph + domain.SeflgSpeed + domain.SeflgEquatorial
	if result != expected {
		t.Errorf("SeFlags() for equatorial = %v, want %v", result, expected)
	}
}

func TestSeFlagsTopcentric(t *testing.T) {
	result := SeFlags(domain.CoordEcliptical, domain.ObsPosTopocentric, true)
	expected := domain.SeflgSwieph + domain.SeflgSpeed + domain.SeflgTopoc
	if result != expected {
		t.Errorf("SeFlags() for topocentric = %v, want %v", result, expected)
	}
}

func TestSeFlagsHeliocentric(t *testing.T) {
	result := SeFlags(domain.CoordEcliptical, domain.ObsPosHeliocentric, true)
	expected := domain.SeflgSwieph + domain.SeflgSpeed + domain.SeflgHelioc
	if result != expected {
		t.Errorf("SeFlags() for heliocentric = %v, want %v", result, expected)
	}
}

func TestSeFlagsSidereal(t *testing.T) {
	result := SeFlags(domain.CoordEcliptical, domain.ObsPosGeocentric, false)
	expected := domain.SeflgSwieph + domain.SeflgSpeed + domain.SeflgSidereal
	if result != expected {
		t.Errorf("SeFlags() for sidereal = %v, want %v", result, expected)
	}
}

func TestSeFlagsTopocEquatCombined(t *testing.T) {
	result := SeFlags(domain.CoordEquatorial, domain.ObsPosTopocentric, true)
	expected := domain.SeflgSwieph + domain.SeflgSpeed + domain.SeflgEquatorial + domain.SeflgTopoc
	if result != expected {
		t.Errorf("SeFlags() for equatorial/topocentric combined = %v, want %v", result, expected)
	}
}

func TestValueToRangeHappyFlow(t *testing.T) {
	testValue := 400.0
	lowerLimit := 0.0
	upperLimit := 360.0
	expected := 40.0
	result, err := ValueToRange(testValue, lowerLimit, upperLimit)
	if err != nil {
		t.Errorf("ValueToRange() returned unexpected error %v", err)
	}
	if math.Abs(result-expected) > 1e-8 {
		t.Errorf("ValueToRange() returned %v, want %v", result, expected)
	}
}

func TestValueToRangeLowerLimit(t *testing.T) {
	testValue := 0.0
	lowerLimit := 0.0
	upperLimit := 360.0
	expected := 0.0
	result, err := ValueToRange(testValue, lowerLimit, upperLimit)
	if err != nil {
		t.Errorf("ValueToRange() returned unexpected error %v", err)
	}
	if math.Abs(result-expected) > 1e-8 {
		t.Errorf("ValueToRange() returned %v, want %v", result, expected)
	}
}

func TestValueToRangeUpperLimit(t *testing.T) {
	testValue := 360.0
	lowerLimit := 0.0
	upperLimit := 360.0
	expected := 0.0
	result, err := ValueToRange(testValue, lowerLimit, upperLimit)
	if err != nil {
		t.Errorf("ValueToRange() returned unexpected error %v", err)
	}
	if math.Abs(result-expected) > 1e-8 {
		t.Errorf("ValueToRange() returned %v, want %v", result, expected)
	}
}

func TestValueToRangeNegativeValue(t *testing.T) {
	testValue := -100.0
	lowerLimit := 0.0
	upperLimit := 360.0
	expected := 260.0
	result, err := ValueToRange(testValue, lowerLimit, upperLimit)
	if err != nil {
		t.Errorf("ValueToRange() returned unexpected error %v", err)
	}
	if math.Abs(result-expected) > 1e-8 {
		t.Errorf("ValueToRange() returned %v, want %v", result, expected)
	}
}

func TestValueToRangeUpperLowerWrongSequence(t *testing.T) {
	testValue := 360.0
	lowerLimit := 180.0
	upperLimit := 90.0
	_, err := ValueToRange(testValue, lowerLimit, upperLimit)
	if err == nil {
		t.Error("ValueToRange() expected errorr for wrong sequence upper and lower did not occur")
	}
}
