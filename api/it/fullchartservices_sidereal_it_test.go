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

// Integration tests for the calculation of sidereal positions

func TestFullChartSiderealPositions(t *testing.T) {
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
			domain.Venus,
			domain.Saturn,
		},
		HouseSys:  domain.HousesRegiomontanus,
		Ayanamsha: domain.AyanLahiri,
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
	expSunLong := 33.584084777778 //  3 ta 35' 2.7052
	resSunLong := result.Points[0].LonPos
	match, errorTxt := pointSiderealEqual("Sun", expSunLong, resSunLong)
	if !match {
		t.Errorf(errorTxt)
	}

	expVenusLong := 64.316452583333 // 4 ge 18'59.2293
	resVenusLong := result.Points[1].LonPos
	match, errorTxt = pointSiderealEqual("Venus", expVenusLong, resVenusLong)
	if !match {
		t.Errorf(errorTxt)
	}

	expSaturnLong := 153.950636805555 // 3 vi 57' 2.2925
	resSaturnLong := result.Points[2].LonPos
	match, errorTxt = pointSiderealEqual("Saturn", expSaturnLong, resSaturnLong)
	if !match {
		t.Errorf(errorTxt)
	}

}

func pointSiderealEqual(point string, longExp, longRes float64) (bool, string) {
	const DELTA = 0.000001
	noErrors := true
	errorText := ""
	if math.Abs(longExp-longRes) > DELTA {
		noErrors = false
		errorText = fmt.Sprintf("Longitude for %s does not match, want %f, got %f \n", point, longExp, longRes)
	}
	return noErrors, errorText
}
