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
	"fmt"
	"math"
	"testing"
)

// Integration tests for the apogee according to Max Duval

func TestFullChartApogeeDuval(t *testing.T) {
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
			domain.ApogeeDuval,
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
		t.Errorf("Integration test for apogeel according to Duval: %v", err)
	}
	expApogeeLong := 324.1637740856539
	expApogeeLongSpeed := -0.416664
	resApogeeLong := result.Points[1].LonPos
	resApogeeLongSpeed := result.Points[1].LonSpeed
	match, errorTxt := pointApogeeEqual("Apogee", expApogeeLong, resApogeeLong, expApogeeLongSpeed, resApogeeLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

}

func pointApogeeEqual(point string, longExp, longRes, speedExp, speedRes float64) (bool, string) {
	const DELTA = 0.000001
	noErrors := true
	errorText := ""
	if math.Abs(longExp-longRes) > DELTA {
		noErrors = false
		errorText = fmt.Sprintf("Longitude for %s does not match, want %f, got %f \n", point, longExp, longRes)
	}
	if math.Abs(speedExp-speedRes) > DELTA {
		noErrors = false
		errorText = fmt.Sprintf("Speed for %s does not match, want %f, got %f \n", point, speedExp, speedRes)
	}
	return noErrors, errorText
}
