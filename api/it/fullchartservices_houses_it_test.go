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

func TestFullChartHousesPlacidus(t *testing.T) {
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
		HouseSys:  domain.HousesPlacidus,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Jd:        jd,
		GeoLong:   geoLong,
		GeoLat:    geoLat,
	})
	if err != nil {
		t.Errorf("Integration test for full chart houses Placidus failed with error: %v", err)
	}

	expCusp3 := 263.6838888888888
	resCusp3 := result.Cusps[3].LonPos
	match, errorTxt := housesEqual("Cusp 3", expCusp3, resCusp3)
	if !match {
		t.Errorf(errorTxt)
	}
	expCusp5 := 332.983333333333
	resCusp5 := result.Cusps[5].LonPos
	match, errorTxt = housesEqual("Cusp 5", expCusp5, resCusp5)
	if !match {
		t.Errorf(errorTxt)
	}
	expCusp12 := 181.240833333333
	resCusp12 := result.Cusps[12].LonPos
	match, errorTxt = housesEqual("Cusp 12", expCusp12, resCusp12)
	if !match {
		t.Errorf(errorTxt)
	}
}

func TestFullChartHousesRegiomontanus(t *testing.T) {
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
		t.Errorf("Integration test for full chart houses Regiomontanus failed with error: %v", err)
	}

	expCusp2 := 228.514444444445
	resCusp2 := result.Cusps[2].LonPos
	match, errorTxt := housesEqual("Cusp 2", expCusp2, resCusp2)
	if !match {
		t.Errorf(errorTxt)
	}
	expCusp9 := 80.235555555555
	resCusp9 := result.Cusps[9].LonPos
	match, errorTxt = housesEqual("Cusp 9", expCusp9, resCusp9)
	if !match {
		t.Errorf(errorTxt)
	}
}

func TestFullChartHousesPlacidusSouthernHemisphere(t *testing.T) {
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
	geoLat := -44.5
	calcService := apicalc.NewFullChartService()
	result, err := calcService.CalcFullChart(domain.FullChartRequest{
		Points: []domain.ChartPoint{
			domain.Sun,
			domain.Moon,
			domain.Jupiter,
			domain.Pluto,
		},
		HouseSys:  domain.HousesPlacidus,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Jd:        jd,
		GeoLong:   geoLong,
		GeoLat:    geoLat,
	})
	if err != nil {
		t.Errorf("Integration test for full chart houses Placidus failed with error: %v", err)
	}
	expCusp6 := 2.225
	resCusp6 := result.Cusps[6].LonPos
	match, errorTxt := housesEqual("Cusp 6", expCusp6, resCusp6)
	if !match {
		t.Errorf(errorTxt)
	}
	expCusp10 := 119.309722222222
	resCusp10 := result.Cusps[10].LonPos
	match, errorTxt = housesEqual("Cusp 10", expCusp10, resCusp10)
	if !match {
		t.Errorf(errorTxt)
	}
}

func TestFullChartHousesCampanusSouthernHemisphere(t *testing.T) {

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
	geoLat := -44.5
	calcService := apicalc.NewFullChartService()
	result, err := calcService.CalcFullChart(domain.FullChartRequest{
		Points: []domain.ChartPoint{
			domain.Sun,
			domain.Moon,
			domain.Jupiter,
			domain.Pluto,
		},
		HouseSys:  domain.HousesCampanus,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Jd:        jd,
		GeoLong:   geoLong,
		GeoLat:    geoLat,
	})
	if err != nil {
		t.Errorf("Integration test for full chart houses Placidus failed with error: %v", err)
	}
	expCusp1 := 233.110833333333
	resCusp1 := result.Cusps[1].LonPos
	match, errorTxt := housesEqual("Cusp 1", expCusp1, resCusp1)
	if !match {
		t.Errorf(errorTxt)
	}
	expCusp8 := 89.806388888888
	resCusp8 := result.Cusps[8].LonPos
	match, errorTxt = housesEqual("Cusp 8", expCusp8, resCusp8)
	if !match {
		t.Errorf(errorTxt)
	}

}

func housesEqual(cusp string, longExp, longRes float64) (bool, string) {
	const DELTA = 0.001
	noErrors := true
	errorText := ""
	if math.Abs(longExp-longRes) > DELTA {
		noErrors = false
		errorText = fmt.Sprintf("Longitude for %s does not match, want %f, got %f \n", cusp, longExp, longRes)
	}
	return noErrors, errorText
}
