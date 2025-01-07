/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package apicalc

import (
	"enigma-ar/domain"
	"testing"
)

func TestFullPositionsInsufficientPoints(t *testing.T) {
	request := domain.PointPositionsRequest{
		Points:    []domain.ChartPoint{},
		JdUt:      12344.678,
		GeoLong:   0.0,
		GeoLat:    0.0,
		Armc:      0.0,
		Obliquity: 23.447,
		Coord:     0.0,
		ObsPos:    0.0,
		ProjType:  0.0,
		Ayanamsha: domain.AyanNone,
	}
	fps := NewFullPointService()
	result, err := fps.FullPositions(request)
	if err == nil {
		t.Errorf("full positions: expected error for insufficient points")
	}
	if result != nil {
		t.Errorf("full positions: expected nil for insufficient points")
	}
}

func TestFullPositionsJdTooEarly(t *testing.T) {
	request := domain.PointPositionsRequest{
		Points:    []domain.ChartPoint{domain.Sun, domain.Moon},
		JdUt:      domain.MinJdGeneral - 1.0,
		GeoLong:   0.0,
		GeoLat:    0.0,
		Armc:      0.0,
		Obliquity: 23.447,
		Coord:     0.0,
		ObsPos:    0.0,
		ProjType:  0.0,
		Ayanamsha: domain.AyanNone,
	}
	fps := NewFullPointService()
	result, err := fps.FullPositions(request)
	if err == nil {
		t.Errorf("full positions: expected error for jd that is too early")
	}
	if result != nil {
		t.Errorf("full positions: expected nil for jd that is too early")
	}
}

func TestFullPositionsJdTooLate(t *testing.T) {
	request := domain.PointPositionsRequest{
		Points:    []domain.ChartPoint{domain.Sun, domain.Moon},
		JdUt:      domain.MaxJdGeneral + 1.0,
		GeoLong:   0.0,
		GeoLat:    0.0,
		Armc:      0.0,
		Obliquity: 23.447,
		Coord:     0.0,
		ObsPos:    0.0,
		ProjType:  0.0,
		Ayanamsha: domain.AyanNone,
	}
	fps := NewFullPointService()
	result, err := fps.FullPositions(request)
	if err == nil {
		t.Errorf("full positions: expected error for jd that is too late")
	}
	if result != nil {
		t.Errorf("full positions: expected nil for jd that is too late")
	}
}

func TestFullPositionsGeoLongTooSmall(t *testing.T) {
	request := domain.PointPositionsRequest{
		Points:    []domain.ChartPoint{domain.Sun, domain.Moon},
		JdUt:      123456.789,
		GeoLong:   -190.0,
		GeoLat:    0.0,
		Armc:      0.0,
		Obliquity: 23.447,
		Coord:     0.0,
		ObsPos:    0.0,
		ProjType:  0.0,
		Ayanamsha: domain.AyanNone,
	}
	fps := NewFullPointService()
	result, err := fps.FullPositions(request)
	if err == nil {
		t.Errorf("full positions: expected error for geoLong that is too small")
	}
	if result != nil {
		t.Errorf("full positions: expected nil for geoLong that is too small")
	}
}

func TestFullPositionsGeoLongTooLarge(t *testing.T) {
	request := domain.PointPositionsRequest{
		Points:    []domain.ChartPoint{domain.Sun, domain.Moon},
		JdUt:      123456.789,
		GeoLong:   190.0,
		GeoLat:    0.0,
		Armc:      0.0,
		Obliquity: 23.447,
		Coord:     0.0,
		ObsPos:    0.0,
		ProjType:  0.0,
		Ayanamsha: domain.AyanNone,
	}
	fps := NewFullPointService()
	result, err := fps.FullPositions(request)
	if err == nil {
		t.Errorf("full positions: expected error for geoLong that is too large")
	}
	if result != nil {
		t.Errorf("full positions: expected nil for geoLong that is too large")
	}
}

func TestFullPositionsGeoLatTooSmall(t *testing.T) {
	request := domain.PointPositionsRequest{
		Points:    []domain.ChartPoint{domain.Sun, domain.Moon},
		JdUt:      123456.789,
		GeoLong:   90.0,
		GeoLat:    -91.0,
		Armc:      0.0,
		Obliquity: 23.447,
		Coord:     0.0,
		ObsPos:    0.0,
		ProjType:  0.0,
		Ayanamsha: domain.AyanNone,
	}
	fps := NewFullPointService()
	result, err := fps.FullPositions(request)
	if err == nil {
		t.Errorf("full positions: expected error for geoLat that is too small")
	}
	if result != nil {
		t.Errorf("full positions: expected nil for geoLat that is too small")
	}
}

func TestFullPositionsGeoLatTooLarge(t *testing.T) {
	request := domain.PointPositionsRequest{
		Points:    []domain.ChartPoint{domain.Sun, domain.Moon},
		JdUt:      123456.789,
		GeoLong:   90.0,
		GeoLat:    91.0,
		Armc:      0.0,
		Obliquity: 23.447,
		Coord:     0.0,
		ObsPos:    0.0,
		ProjType:  0.0,
		Ayanamsha: domain.AyanNone,
	}
	fps := NewFullPointService()
	result, err := fps.FullPositions(request)
	if err == nil {
		t.Errorf("full positions: expected error for geoLat that is too large")
	}
	if result != nil {
		t.Errorf("full positions: expected nil for geoLat that is too large")
	}
}

func TestFullPositionsArmcTooSmall(t *testing.T) {
	request := domain.PointPositionsRequest{
		Points:    []domain.ChartPoint{domain.Sun, domain.Moon},
		JdUt:      123456.789,
		GeoLong:   90.0,
		GeoLat:    11.0,
		Armc:      -1.0,
		Obliquity: 23.447,
		Coord:     0.0,
		ObsPos:    0.0,
		ProjType:  0.0,
		Ayanamsha: domain.AyanNone,
	}
	fps := NewFullPointService()
	result, err := fps.FullPositions(request)
	if err == nil {
		t.Errorf("full positions: expected error for armc that is too small")
	}
	if result != nil {
		t.Errorf("full positions: expected nil for armc that is too small")
	}
}

func TestFullPositionsArmcTooLarge(t *testing.T) {
	request := domain.PointPositionsRequest{
		Points:    []domain.ChartPoint{domain.Sun, domain.Moon},
		JdUt:      123456.789,
		GeoLong:   90.0,
		GeoLat:    11.0,
		Armc:      361.0,
		Obliquity: 23.447,
		Coord:     0.0,
		ObsPos:    0.0,
		ProjType:  0.0,
		Ayanamsha: domain.AyanNone,
	}
	fps := NewFullPointService()
	result, err := fps.FullPositions(request)
	if err == nil {
		t.Errorf("full positions: expected error for armc that is too large")
	}
	if result != nil {
		t.Errorf("full positions: expected nil for armc that is too large")
	}
}

func TestFullPositionsObliquityTooSmall(t *testing.T) {
	request := domain.PointPositionsRequest{
		Points:    []domain.ChartPoint{domain.Sun, domain.Moon},
		JdUt:      123456.789,
		GeoLong:   90.0,
		GeoLat:    11.0,
		Armc:      1.0,
		Obliquity: domain.MinObliquity - 1.0,
		Coord:     0.0,
		ObsPos:    0.0,
		ProjType:  0.0,
		Ayanamsha: domain.AyanNone,
	}
	fps := NewFullPointService()
	result, err := fps.FullPositions(request)
	if err == nil {
		t.Errorf("full positions: expected error for obliquity that is too small")
	}
	if result != nil {
		t.Errorf("full positions: expected nil for obliquity that is too small")
	}
}

func TestFullPositionsObliquityTooLarge(t *testing.T) {
	request := domain.PointPositionsRequest{
		Points:    []domain.ChartPoint{domain.Sun, domain.Moon},
		JdUt:      123456.789,
		GeoLong:   90.0,
		GeoLat:    11.0,
		Armc:      1.0,
		Obliquity: domain.MaxObliquity + 1.0,
		Coord:     0.0,
		ObsPos:    0.0,
		ProjType:  0.0,
		Ayanamsha: domain.AyanNone,
	}
	fps := NewFullPointService()
	result, err := fps.FullPositions(request)
	if err == nil {
		t.Errorf("full positions: expected error for obliquity that is too large")
	}
	if result != nil {
		t.Errorf("full positions: expected nil for obliquity that is too large")
	}
}
