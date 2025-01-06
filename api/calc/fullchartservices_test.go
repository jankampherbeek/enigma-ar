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

func TestCalcFullChartInsufficientPoints(T *testing.T) {
	request := domain.FullChartRequest{
		Points:    []domain.ChartPoint{},
		HouseSys:  domain.HousesAlcabitius,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Jd:        123456.789,
		GeoLong:   0.0,
		GeoLat:    0.0,
	}
	fcc := NewFullChartService()
	_, err := fcc.CalcFullChart(request)
	if err == nil {
		T.Errorf("Expected error for insufficient points, got nil")
	}
}

func TestCalcFullChartJdTooEarly(T *testing.T) {
	request := domain.FullChartRequest{
		Points: []domain.ChartPoint{
			domain.Sun,
			domain.Moon,
		},
		HouseSys:  domain.HousesAlcabitius,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Jd:        -2946708.5,
		GeoLong:   0.0,
		GeoLat:    0.0,
	}
	fcc := NewFullChartService()
	_, err := fcc.CalcFullChart(request)
	if err == nil {
		T.Errorf("Expected error for jd too early, got nil")
	}
}

func TestCalcFullChartJdTooLate(T *testing.T) {
	request := domain.FullChartRequest{
		Points: []domain.ChartPoint{
			domain.Sun,
			domain.Moon,
		},
		HouseSys:  domain.HousesAlcabitius,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Jd:        7865294.5,
		GeoLong:   0.0,
		GeoLat:    0.0,
	}
	fcc := NewFullChartService()
	_, err := fcc.CalcFullChart(request)
	if err == nil {
		T.Errorf("Expected error for jd too late, got nil")
	}
}

func TestCalcFullChartGeoLongTooSmall(T *testing.T) {
	request := domain.FullChartRequest{
		Points: []domain.ChartPoint{
			domain.Sun,
			domain.Moon,
		},
		HouseSys:  domain.HousesAlcabitius,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Jd:        123456.789,
		GeoLong:   -181.0,
		GeoLat:    0.0,
	}
	fcc := NewFullChartService()
	_, err := fcc.CalcFullChart(request)
	if err == nil {
		T.Errorf("Expected error for geolong too small, got nil")
	}
}

func TestCalcFullChartGeoLongTooLarge(T *testing.T) {
	request := domain.FullChartRequest{
		Points: []domain.ChartPoint{
			domain.Sun,
			domain.Moon,
		},
		HouseSys:  domain.HousesAlcabitius,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Jd:        123456.789,
		GeoLong:   181.0,
		GeoLat:    0.0,
	}
	fcc := NewFullChartService()
	_, err := fcc.CalcFullChart(request)
	if err == nil {
		T.Errorf("Expected error for geolong too large, got nil")
	}
}

func TestCalcFullChartGeoLatTooSmall(T *testing.T) {
	request := domain.FullChartRequest{
		Points: []domain.ChartPoint{
			domain.Sun,
			domain.Moon,
		},
		HouseSys:  domain.HousesAlcabitius,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Jd:        123456.789,
		GeoLong:   81.0,
		GeoLat:    -91.0,
	}
	fcc := NewFullChartService()
	_, err := fcc.CalcFullChart(request)
	if err == nil {
		T.Errorf("Expected error for geolat too small, got nil")
	}
}

func TestCalcFullChartGeoLatTooLarge(T *testing.T) {
	request := domain.FullChartRequest{
		Points: []domain.ChartPoint{
			domain.Sun,
			domain.Moon,
		},
		HouseSys:  domain.HousesAlcabitius,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Jd:        123456.789,
		GeoLong:   81.0,
		GeoLat:    91.0,
	}
	fcc := NewFullChartService()
	_, err := fcc.CalcFullChart(request)
	if err == nil {
		T.Errorf("Expected error for geolat too large, got nil")
	}
}
