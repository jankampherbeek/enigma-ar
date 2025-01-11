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

const DELTA = 0.001 // small DELTA to prevent effects from different SE versions with different values for Delta T

func TestFullChartStandardPlanets(t *testing.T) {
	// initialize SE
	sep := string(filepath.Separator)
	ephePath := ".." + sep + ".." + sep + "sedata" // path is relative from current package
	sp := se.NewSwephPreparation()
	sp.SetEphePath(ephePath)
	// first calc JD
	dateTime := domain.DateTime{
		Year:  2025,
		Month: 1,
		Day:   10,
		Ut:    0.0,
		Greg:  true,
	}
	jdService := api.NewJulDayService()
	jd := jdService.JulDay(&dateTime)
	geoLong := 0.0
	geoLat := 0.0
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
	expSunLong := 289.987436055555
	expSunLat := 9.30277777778e-5
	expSunDist := 0.983399922
	expSunLongSpeed := 1.01881383333
	resSunLong := result.Points[0].LonPos
	resSunLat := result.Points[0].LatPos
	resSunDist := result.Points[0].RadvPos
	resSunLongSpeed := result.Points[0].LonSpeed
	match, errorTxt := isEqual("Sun", expSunLong, resSunLong, expSunLat, resSunLat, expSunDist, resSunDist, expSunLongSpeed, resSunLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

	expMoonLong := 59.346283972222
	expMoonLat := 4.328678055555
	expMoonDist := 0.002482022
	expMoonLongSpeed := 14.131726583333
	resMoonLong := result.Points[1].LonPos
	resMoonLat := result.Points[1].LatPos
	resMoonDist := result.Points[1].RadvPos
	resMoonLongSpeed := result.Points[1].LonSpeed
	match, errorTxt = isEqual("Moon", expMoonLong, resMoonLong, expMoonLat, resMoonLat, expMoonDist, resMoonDist, expMoonLongSpeed, resMoonLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

	expJupLong := 72.359964833333
	expJupLat := -0.570847611112
	expJupDist := 4.271226405
	expJupLongSpeed := -0.0827720277778
	resJupLong := result.Points[2].LonPos
	resJupLat := result.Points[2].LatPos
	resJupDist := result.Points[2].RadvPos
	resJupLongSpeed := result.Points[2].LonSpeed
	match, errorTxt = isEqual("Jupiter", expJupLong, resJupLong, expJupLat, resJupLat, expJupDist, resJupDist, expJupLongSpeed, resJupLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

	expPluLong := 301.346121833333
	expPluLat := -3.282625916667
	expPluDist := 36.139161562
	expPluLongSpeed := 0.0317959722222
	resPluLong := result.Points[3].LonPos
	resPluLat := result.Points[3].LatPos
	resPluDist := result.Points[3].RadvPos
	resPluLongSpeed := result.Points[3].LonSpeed
	match, errorTxt = isEqual("Pluto", expPluLong, resPluLong, expPluLat, resPluLat, expPluDist, resPluDist, expPluLongSpeed, resPluLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

}

func isEqual(point string, longExp, longRes, latExp, latRes, distExp, distRes, speedExp, speedRes float64) (bool, string) {
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
