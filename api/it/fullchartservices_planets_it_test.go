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

// Integration tests for the calculation of a full chart

func TestFullChartStandardPlanets(t *testing.T) {
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
			domain.Jupiter,
			domain.Pluto,
		},
		HouseSys:  domain.HousesRegiomontanus,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Jd:        jd,
		GeoLong:   geoLong,
		GeoLat:    geoLat,
	})
	if err != nil {
		t.Errorf("Integration test for full chart standard planets failed with error: %v", err)
	}
	expSunLong := 57.590494694445 // 27 ta 35'25.7809
	expSunLat := -0.000067944444  // -0° 0' 0.2446
	expSunDist := 1.011558980
	expSunLongSpeed := 0.963364 //0°57'48.1104
	resSunLong := result.Points[0].LonPos
	resSunLat := result.Points[0].LatPos
	resSunDist := result.Points[0].RadvPos
	resSunLongSpeed := result.Points[0].LonSpeed
	match, errorTxt := pointPosEqual("Sun", expSunLong, resSunLong, expSunLat, resSunLat, expSunDist, resSunDist, expSunLongSpeed, resSunLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

	expMoonLong := 116.706899111112 // 26 cn 42'24.8368
	expMoonLat := -1.248205222222   // -1°14'53.5388
	expMoonDist := 0.002477204
	expMoonLongSpeed := 14.095998472222 // 14° 5'45.5945
	resMoonLong := result.Points[1].LonPos
	resMoonLat := result.Points[1].LatPos
	resMoonDist := result.Points[1].RadvPos
	resMoonLongSpeed := result.Points[1].LonSpeed
	match, errorTxt = pointPosEqual("Moon", expMoonLong, resMoonLong, expMoonLat, resMoonLat, expMoonDist, resMoonDist, expMoonLongSpeed, resMoonLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

	expJupLong := 357.103241833333 // 27 pi  6'11.6706
	expJupLat := -1.100705416667   // -1° 6' 2.5395
	expJupDist := 5.394679142
	expJupLongSpeed := 0.1736915 // 0°10'25.2894
	resJupLong := result.Points[2].LonPos
	resJupLat := result.Points[2].LatPos
	resJupDist := result.Points[2].RadvPos
	resJupLongSpeed := result.Points[2].LonSpeed
	match, errorTxt = pointPosEqual("Jupiter", expJupLong, resJupLong, expJupLat, resJupLat, expJupDist, resJupDist, expJupLongSpeed, resJupLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

	expPluLong := 275.000652611111 // 5 cp  0' 2.3494
	expPluLat := 5.155871111111    // 5° 9'21.1360
	expPluDist := 31.028422017
	expPluLongSpeed := -0.01886775 // -0° 1' 7.9239
	resPluLong := result.Points[3].LonPos
	resPluLat := result.Points[3].LatPos
	resPluDist := result.Points[3].RadvPos
	resPluLongSpeed := result.Points[3].LonSpeed
	match, errorTxt = pointPosEqual("Pluto", expPluLong, resPluLong, expPluLat, resPluLat, expPluDist, resPluDist, expPluLongSpeed, resPluLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

}

func pointPosEqual(point string, longExp, longRes, latExp, latRes, distExp, distRes, speedExp, speedRes float64) (bool, string) {
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
