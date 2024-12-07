/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	"enigma-ar/domain"
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
