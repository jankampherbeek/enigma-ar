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

// Integration tests for the calculation of hypothetical planets

func TestFullChartHypotheticals(t *testing.T) {
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
			domain.PersephoneRam,
			domain.HermesRam,
			domain.DemeterRam,
			domain.PersephoneCarteret,
			domain.VulcanusCarteret,
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
		t.Errorf("Integration test for hypotheticals failed with error: %v", err)
	}
	expPersephoneRamLong := 1.894376341404726
	expPersephoneRamLongSpeed := 0.009516589830895317
	resPersephoneRamLong := result.Points[1].LonPos
	resPersephoneRamLongSpeed := result.Points[1].LonSpeed
	match, errorTxt := hypoPosEqual("PersephoneRam", expPersephoneRamLong, resPersephoneRamLong, expPersephoneRamLongSpeed, resPersephoneRamLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

	expHermesRamLong := 189.35368397117426
	expHermesRamLongSpeed := -0.006953598733389299
	resHermesRamLong := result.Points[2].LonPos
	resHermesRamLongSpeed := result.Points[2].LonSpeed
	match, errorTxt = hypoPosEqual("HermesRam", expHermesRamLong, resHermesRamLong, expHermesRamLongSpeed, resHermesRamLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

	expDemeterRamLong := 284.29568101429845
	expDemeterRamLongSpeed := -0.005967496286586993
	resDemeterRamLong := result.Points[3].LonPos
	resDemeterRamLongSpeed := result.Points[3].LonSpeed
	match, errorTxt = hypoPosEqual("Demeter", expDemeterRamLong, resDemeterRamLong, expDemeterRamLongSpeed, resDemeterRamLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

	expPersephoneCarteretLong := 322.37784225246526
	expPersephoneCarteretLongSpeed := 0.002738
	resPersephoneCarteretLong := result.Points[4].LonPos
	resPersephoneCarteretLongSpeed := result.Points[4].LonSpeed
	match, errorTxt = hypoPosEqual("PersephoneCarteret", expPersephoneCarteretLong, resPersephoneCarteretLong, expPersephoneCarteretLongSpeed, resPersephoneCarteretLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}

	expVulcanusCarteretLong := 76.40781323885591
	expVulcanusCarteretLongSpeed := 0.001506
	resVulcanusCarteretLong := result.Points[5].LonPos
	resVulcanusCarteretLongSpeed := result.Points[5].LonSpeed
	match, errorTxt = hypoPosEqual("VulcanusCarteret", expVulcanusCarteretLong, resVulcanusCarteretLong, expVulcanusCarteretLongSpeed, resVulcanusCarteretLongSpeed)
	if !match {
		t.Errorf(errorTxt)
	}
}

func hypoPosEqual(point string, longExp, longRes, speedExp, speedRes float64) (bool, string) {
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
