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

// Integration tests for the calculation of a full chart

const DELTA = 0.001

func TestFullChart(t *testing.T) {
	// first calc JD
	dateTime := domain.DateTime{
		Year:  2025,
		Month: 1,
		Day:   10,
		//Ut:    10.933333333333333333333,
		Ut:   0.0,
		Greg: true,
	}
	jdService := api.NewJulDayService()
	jd := jdService.JulDay(&dateTime)
	fmt.Printf("=========== JulDay %v\n", jd)
	geoLong := 6.0 + 54.0/60.0
	geoLat := 52.0 + 13.0/60.0
	calcService := apicalc.NewFullChartService()
	result, err := calcService.CalcFullChart(domain.FullChartRequest{
		Points: []domain.ChartPoint{
			domain.Sun,
			domain.Moon,
			domain.Mercury,
			domain.Venus,
			domain.Mars,
			domain.Jupiter,
			domain.Saturn,
			domain.Uranus,
			domain.Neptune,
			domain.Pluto,
			domain.NodeMean,
			//domain.Chiron,
			domain.Vertex,
			domain.EastPoint,
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
		t.Errorf("Integration test for full charts failed with error: %v", err)
	}

	//expSunLong := 290.4515581013796
	//resSunLong := result.Points[domain.Sun].LonPos
	//if !isEqual(expSunLong, resSunLong) {
	//	t.Errorf("Mismatch longitude Sun, expected %v, got %v", expSunLong, resSunLong)
	//}
	expSunLong := 289.9874435693395
	resSunLong := result.Points[domain.Sun].LonPos
	if !isEqual(expSunLong, resSunLong) {
		t.Errorf("Mismatch longitude Sun, expected %v, got %v", expSunLong, resSunLong)
	}

	//expMoonLong := 65.7792265292281
	//resMoonLong := result.Points[domain.Moon].LonPos
	//if !isEqual(expMoonLong, resMoonLong) {
	//	t.Errorf("Mismatch longitude Moon, expected %v, got %v", expMoonLong, resMoonLong)
	//}
	expMoonLong := 59.34630699452788
	resMoonLong := result.Points[domain.Moon].LonPos
	if !isEqual(expMoonLong, resMoonLong) {
		t.Errorf("Mismatch longitude Moon, expected %v, got %v", expMoonLong, resMoonLong)
	}

	//expPlutoLong := 301.3606931011713
	//resPlutoLong := result.Points[domain.Pluto].LonPos
	//if !isEqual(expPlutoLong, resPlutoLong) {
	//	t.Errorf("Mismatch longitude Pluto, expected %v, got %v", expPlutoLong, resPlutoLong)
	//}

	expPlutoLong := 301.3461975241484
	resPlutoLong := result.Points[domain.Pluto].LonPos
	if !isEqual(expPlutoLong, resPlutoLong) {
		t.Errorf("Mismatch longitude Pluto, expected %v, got %v", expPlutoLong, resPlutoLong)
	}

}

func isEqual(pos1, pos2 float64) bool {
	return math.Abs(pos1-pos2) < DELTA
}
