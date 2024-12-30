/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package mathextra

import (
	"math"
	"testing"
)

func TestDegToRad(t *testing.T) {
	degrees := 100.0
	expectedRadians := 1.74532925199
	result := DegToRad(degrees)
	if math.Abs(result-expectedRadians) > 1e-8 {
		t.Errorf("DegToRad() returned %v, want %v", result, expectedRadians)
	}
}

func TestRadToDeg(t *testing.T) {
	radians := 2.0
	expectedDegrees := 114.591559026
	result := RadToDeg(radians)
	if math.Abs(result-expectedDegrees) > 1e-8 {
		t.Errorf("RadToDeg() returned %v, want %v", result, expectedDegrees)
	}
}

func TestRectangular2Polar(t *testing.T) {
	rectAngValues := RectAngCoordinates{
		XCoord: 1.0,
		YCoord: 2.0,
		ZCoord: 3.0,
	}
	expected := PolarCoordinates{
		PhiCoord:   1.107148717794,
		ThetaCoord: 0.930274014115,
		RCoord:     3.741657386774,
	}
	result := Rectangular2Polar(rectAngValues)
	if math.Abs(result.RCoord-expected.RCoord) > 1e-3 {
		t.Errorf("Rectangular2Polar() returned RCoord %v, want %v", result.RCoord, expected.RCoord)
	}
	if math.Abs(result.PhiCoord-expected.PhiCoord) > 1e-3 {
		t.Errorf("Rectangular2Polar() returned PhiCoord %v, want %v", result.PhiCoord, expected.PhiCoord)
	}
	if math.Abs(result.ThetaCoord-expected.ThetaCoord) > 1e-3 {
		t.Errorf("Rectangular2Polar() returned ThetaCoord %v, want %v", result.ThetaCoord, expected.ThetaCoord)
	}
}

func TestPolar2Rectangular(t *testing.T) {
	polarValues := PolarCoordinates{
		PhiCoord:   1.107148717794,
		ThetaCoord: 0.930274014115,
		RCoord:     3.741657386774,
	}
	expected := RectAngCoordinates{
		XCoord: 1.0,
		YCoord: 2.0,
		ZCoord: 3.0,
	}
	result := Polar2Rectangular(polarValues)
	if math.Abs(result.XCoord-expected.XCoord) > 1e-3 {
		t.Errorf("Polar2Rectangular() returned XCoord %v, want %v", result.XCoord, expected.XCoord)
	}
	if math.Abs(result.YCoord-expected.YCoord) > 1e-3 {
		t.Errorf("Polar2Rectangular() returned YCoord %v, want %v", result.YCoord, expected.YCoord)
	}
	if math.Abs(result.XCoord-expected.XCoord) > 1e-3 {
		t.Errorf("Polar2Rectangular() returned ZCoord %v, want %v", result.ZCoord, expected.ZCoord)
	}

}
