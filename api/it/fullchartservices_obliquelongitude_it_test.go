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

// Integration tests for the calculation of positions in oblique longitude

func TestFullChartOblLonPositions(t *testing.T) {
	// initialize SE
	sep := string(filepath.Separator)
	ephePath := ".." + sep + ".." + sep + "sedata" // path is relative from current package
	sp := se.NewSwephPreparation()
	sp.SetEphePath(ephePath)
	// calc JD
	dateTime := domain.DateTime{
		Year:  2010,
		Month: 5,
		Day:   18,
		Ut:    15.5,
		Greg:  true,
	}
	jdService := api.NewJulDayService()
	jd := jdService.JulDay(&dateTime)
	// calc Obliquity
	oblCalc := se.NewSwephEpsilonCalculation()
	obliquity, err := oblCalc.CalcEpsilon(jd, true)
	if err != nil {
		t.Errorf("Unexpected errors when calculating obliquity: %v\n", err)
	}
	geoLong := 12.75
	geoLat := 44.5
	calcService := apicalc.NewFullChartService()
	result, err := calcService.CalcFullChart(domain.FullChartRequest{
		Points: []domain.ChartPoint{
			domain.Moon,
			domain.Pluto,
		},
		HouseSys:  domain.HousesRegiomontanus,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjTypeOblique,
		Jd:        jd,
		Obliquity: obliquity,
		GeoLong:   geoLong,
		GeoLat:    geoLat,
	})

	if err != nil {
		t.Errorf("Integration test for full chart standard planets failed with error: %v", err)
	}

	expMoonLong := 116.41329857479772
	resMoonLong := result.Points[0].LonPos
	match, errorTxt := OblLonEqual("Moon", expMoonLong, resMoonLong)
	if !match {
		t.Errorf(errorTxt)
	}

	expPluLong := 272.9626152067933
	resPluLong := result.Points[1].LonPos
	match, errorTxt = OblLonEqual("Pluto", expPluLong, resPluLong)
	if !match {
		t.Errorf(errorTxt)
	}

}

func OblLonEqual(point string, longExp, longRes float64) (bool, string) {
	const DELTA = 0.000001
	noErrors := true
	errorText := ""
	if math.Abs(longExp-longRes) > DELTA {
		noErrors = false
		errorText = fmt.Sprintf("Longitude for %s does not match, want %f, got %f \n", point, longExp, longRes)
	}
	return noErrors, errorText
}
