/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	domain "enigma-ar/domain"
	"math"
	"testing"
)

const delta = 0.000001

func TestCalcPointPosViaFormulaPersephoneCarteret(t *testing.T) {
	jdUt := 2432870.3851736113
	expected := 260.871365956251
	c := NewPointPosCalculation()
	request := domain.PointPositionsRequest{
		Points: []domain.ChartPoint{
			domain.PersephoneCarteret,
		},
		JdUt:      jdUt,
		GeoLong:   0,
		GeoLat:    0,
		Armc:      0,
		Obliquity: 0,
		Coord:     domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Ayanamsha: domain.AyanNone,
	}
	result, err := c.CalcPointPos(request)
	if err != nil {
		t.Error(err)
	}
	if math.Abs(expected-result[0].LonPos) > delta {
		t.Errorf("Error in calculation of Persephone (Carteret), expected %f, got %f", expected, result[0].LonPos)
	}
}

func TestCalcPointPosViaFormulaVulcanusCarteret(t *testing.T) {
	jdUt := 2432870.3851736113
	expected := 42.579251275938063
	c := NewPointPosCalculation()
	request := domain.PointPositionsRequest{
		Points: []domain.ChartPoint{
			domain.VulcanusCarteret,
		},
		JdUt:      jdUt,
		GeoLong:   0,
		GeoLat:    0,
		Armc:      0,
		Obliquity: 0,
		Coord:     domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Ayanamsha: domain.AyanNone,
	}
	result, err := c.CalcPointPos(request)
	if err != nil {
		t.Error(err)
	}
	if math.Abs(expected-result[0].LonPos) > delta {
		t.Errorf("Error in calculation of Vulcanus (Carteret), expected %f, got %f", expected, result[0].LonPos)
	}
}

func TestCalcPointPosViaFormulaApogeeDuval(t *testing.T) {
	jdUt := 2432870.3851736113
	expected := 353.50375056094902
	c := NewPointPosCalculation()
	request := domain.PointPositionsRequest{
		Points: []domain.ChartPoint{
			domain.ApogeeDuval,
		},
		JdUt:      jdUt,
		GeoLong:   0,
		GeoLat:    0,
		Armc:      0,
		Obliquity: 0,
		Coord:     domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Ayanamsha: domain.AyanNone,
	}
	result, err := c.CalcPointPos(request)
	if err != nil {
		t.Error(err)
	}
	if math.Abs(expected-result[0].LonPos) > 1e-4 {
		t.Errorf("Error in calculation of Apogee (Duval), expected %f, got %f", expected, result[0].LonPos)
	}
}
