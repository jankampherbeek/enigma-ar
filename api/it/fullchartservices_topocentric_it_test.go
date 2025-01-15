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

// Integration tests for the calculation of topocentric positions

func TestFullChartTopocentricPositions(t *testing.T) {
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
			domain.Sun,
			domain.Moon,
			domain.Mercury,
		},
		HouseSys:  domain.HousesRegiomontanus,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosTopocentric,
		ProjType:  domain.ProjType2D,
		Jd:        jd,
		GeoLong:   geoLong,
		GeoLat:    geoLat,
	})
	if err != nil {
		t.Errorf("Integration test for full chart topocentric podsitions failed with error: %v", err)
	}
	expSunLong := 57.588681277778  // 27 ta 35'19.2526
	expSunLat := -0.00101452777778 // -0° 0' 3.6523
	expSunDist := 1.011537390
	expSunLongSpeed := 0.958007805555 // 0°57'28.8281
	resSunLong := result.Points[0].LonPos
	resSunLat := result.Points[0].LatPos
	resSunDist := result.Points[0].RadvPos
	resSunLongSpeed := result.Points[0].LonSpeed
	match, errorTxt := pointTopocentricEqual("Sun", expSunLong, resSunLong, expSunLat, resSunLat, expSunDist, resSunDist, expSunLongSpeed, resSunLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

	expMoonLong := 116.74988125   // 26 cn 44'59.5725
	expMoonLat := -1.666205444445 // -1°39'58.3396
	expMoonDist := 0.002438640
	expMoonLongSpeed := 9.913286944445 // 9°54'47.8330
	resMoonLong := result.Points[1].LonPos
	resMoonLat := result.Points[1].LatPos
	resMoonDist := result.Points[1].RadvPos
	resMoonLongSpeed := result.Points[1].LonSpeed
	match, errorTxt = pointTopocentricEqual("Moon", expMoonLong, resMoonLong, expMoonLat, resMoonLat, expMoonDist, resMoonDist, expMoonLongSpeed, resMoonLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

	expMercuryLong := 34.365003361112 // 4 ta 21'54.0121
	expMercuryLat := -3.323620805555  //  -3°19'25.0349
	expMercuryDist := 0.702508140
	expMercuryLongSpeed := 0.497060722222 // 0°29'49.4186
	resMercuryLong := result.Points[2].LonPos
	resMercuryLat := result.Points[2].LatPos
	resMercuryDist := result.Points[2].RadvPos
	resMercuryLongSpeed := result.Points[2].LonSpeed
	match, errorTxt = pointTopocentricEqual("Mercury", expMercuryLong, resMercuryLong, expMercuryLat, resMercuryLat, expMercuryDist, resMercuryDist, expMercuryLongSpeed, resMercuryLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

}

func pointTopocentricEqual(point string, longExp, longRes, latExp, latRes, distExp, distRes, speedExp, speedRes float64) (bool, string) {
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
