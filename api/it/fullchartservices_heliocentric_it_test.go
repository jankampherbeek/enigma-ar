/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package it

import (
	"enigma-ar/api"
	apicalc "enigma-ar/api/calc"
	"enigma-ar/domain"
	"enigma-ar/internal/se"
	"fmt"
	"math"
	"path/filepath"
	"testing"
)

// Integration tests for heliocentric positions

func TestFullChartHelioPositions(t *testing.T) {
	// initialize SE
	sep := string(filepath.Separator)
	ephePath := ".." + sep + ".." + sep + "sedata" // path is relative from current package
	sp := se.NewSwephPreparation()
	sp.SetEphePath(ephePath)
	// first calc JD
	dateTime := domain.DateTime{
		Year:  2010,
		Month: 5,
		Day:   18,
		Ut:    15.5,
		Greg:  true,
	}
	jdService := api.NewJulDayService()
	jd := jdService.JulDay(&dateTime)
	geoLong := 12.75
	geoLat := 44.5
	calcService := apicalc.NewFullChartService()
	result, err := calcService.CalcFullChart(domain.FullChartRequest{
		Points: []domain.ChartPoint{
			domain.Mercury,
			domain.Uranus,
			domain.Jupiter,
			domain.Pluto,
		},
		HouseSys:  domain.HousesRegiomontanus,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosHeliocentric,
		ProjType:  domain.ProjType2D,
		Jd:        jd,
		GeoLong:   geoLong,
		GeoLat:    geoLat,
	})
	if err != nil {
		t.Errorf("Integration test for heliocentric positions failed with error: %v", err)
	}
	expMercuryLong := 274.595112027778 // 4 cp 35'42.4033
	expMercuryLat := -5.0628306111112  // -5° 3'46.1902
	expMercuryDist := 0.461361774
	expMercuryLongSpeed := 2.820337638888 // 2°49'13.2155
	resMercuryLong := result.Points[0].LonPos
	resMercuryLat := result.Points[0].LatPos
	resMercuryDist := result.Points[0].RadvPos
	resMercuryLongSpeed := result.Points[0].LonSpeed
	match, errorTxt := helioPosEqual("Mercury", expMercuryLong, resMercuryLong, expMercuryLat, resMercuryLat, expMercuryDist, resMercuryDist, expMercuryLongSpeed, resMercuryLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

	expUranusLong := 357.247148444445 // 27 pi 14'49.7344
	expUranusLat := -0.751907527778   // -0°45' 6.8671
	expUranusDist := 20.094760671
	expUranusLongSpeed := 0.01076975 // 0° 0'38.7711
	resUranusLong := result.Points[1].LonPos
	resUranusLat := result.Points[1].LatPos
	resUranusDist := result.Points[1].RadvPos
	resUranusLongSpeed := result.Points[1].LonSpeed
	match, errorTxt = helioPosEqual("Uranus", expUranusLong, resUranusLong, expUranusLat, resUranusLat, expUranusDist, resUranusDist, expUranusLongSpeed, resUranusLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

}

func helioPosEqual(point string, longExp, longRes, latExp, latRes, distExp, distRes, speedExp, speedRes float64) (bool, string) {
	const DELTA = 0.000001
	noErrors := true
	errorText := ""
	if math.Abs(longExp-longRes) > DELTA {
		noErrors = false
		errorText = fmt.Sprintf("Longitude for %s does not match, want %f, got %f \n", point, longExp, longRes)
	}
	if math.Abs(latExp-latRes) > DELTA {
		noErrors = false
		errorText = fmt.Sprintf("Latitude for %s does not match, want %f, got %f \n", point, latExp, latRes)
	}
	if math.Abs(distExp-distRes) > DELTA {
		noErrors = false
		errorText = fmt.Sprintf("Distance for %s does not match, want %f, got %f \n", point, distExp, distRes)
	}
	if math.Abs(speedExp-speedRes) > DELTA {
		noErrors = false
		errorText = fmt.Sprintf("Speed for %s does not match, want %f, got %f \n", point, speedExp, speedRes)
	}
	return noErrors, errorText
}
