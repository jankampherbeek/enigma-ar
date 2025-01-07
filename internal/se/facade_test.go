/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package se

import (
	"enigma-ar/domain"
	"math"
	"path/filepath"
	"testing"
)

const DELTA = 1e-7

func TestJulDay(t *testing.T) {
	result := SwephJulDayCalculation{}.CalcJd(2024, 5, 6, 20.5, 1)
	expected := 2460437.3541666665
	difference := math.Abs(result - expected)
	if difference > 0.000001 {
		t.Errorf("Julday(2024,5,6,20.5, true) = %f; want %f", result, expected)
	}
}

func TestRevJulDay(t *testing.T) {
	jd := 2460437.3541666665
	resultYear, resultMonth, resultDay, resultUt := SwephRevJulDayCalculation{}.RevCalcJd(jd, 1)
	expectedYear := 2024
	expectedMonth := 5
	expectedDay := 6
	expectedUt := 20.5
	if resultYear != expectedYear || resultMonth != expectedMonth || resultDay != expectedDay {
		t.Errorf("RevJulDay returns a wrong result for year, month, day: %d, %d, %d", resultYear, resultMonth, resultDay)
	}
	if math.Abs(expectedUt-resultUt) > 0.00001 {
		t.Errorf("RevJulDay returns a wrong result for ut: %f", resultUt)
	}
}

func TestPointPositions(t *testing.T) {
	sep := string(filepath.Separator)
	ephePath := ".." + sep + ".." + sep + "sedata" // path is relative from current package
	sp := NewSwephPreparation()
	sp.SetEphePath(ephePath)
	julDay := 2_470_000.0 // 2050/7/12 12:00
	body := domain.AllChartPoints()[domain.Mercury].CalcId
	flags := domain.SeflgSwieph + domain.SeflgSpeed
	// TODO check all 6 values
	expected := []float64{132.309351305555, 1.309320472222, 1.106102572, 1.572654666667}
	result, err := SwephPointPosCalculation{}.CalcPointPos(julDay, body, flags)
	if err != nil {
		t.Errorf("PointPositions(2_470_000, SeMercury, 256) returns error %s", err)
	} else {
		for i := 0; i <= 3; i++ {
			if math.Abs(result[i]-expected[i]) > DELTA {
				t.Errorf("PointPositionsJ(2_470_000, SeMercury, 256) = %f; want %f", result[i], expected[i])
			}
		}
	}
}

func TestHorizontalPosition(t *testing.T) {
	jdUt := 2_434_406.8177083335
	geoLong := 6.9
	geoLat := 52.216666666666669
	geoHeight := 0.0
	pointRa := 317.18784726228648
	pointDecl := -16.422932391786961
	flags := 2048
	expected := []float64{297.4812938568067, 0.0, 0.50662370470219853}
	result := SwephHorPosCalculation{}.CalcHorPos(jdUt, geoLong, geoLat, geoHeight, pointRa, pointDecl, flags)
	for i := 0; i <= 2; i++ {
		if math.Abs(result[i]-expected[i]) > DELTA {
			t.Errorf("HorizontalPosition(2_434_406.8177, 6.9, 52.2166, 0.0, 0.0, 317.1878, -16.4229, 2048) = %f; want %f", result[i], expected[i])
		}
	}
}

func TestHousePositionEcliptical(t *testing.T) {
	julDay := 2_470_000.0 // 2050/7/12 12:00
	geoLong := 0.0
	geoLat := 51.5
	flags := domain.SeflgSwieph + domain.SeflgSpeed
	var houseSys rune = 'P'
	cuspResult, mcAscResult, err := SwephHousePosCalculation{}.CalcHousePos(houseSys, julDay, geoLong, geoLat, flags)
	expectedMc := 109.0150128333
	expectedAsc := 194.5072978611
	expectedCusp2 := 220.0397175278
	expectedCusp9 := 71.8124521388888
	if err != nil {
		t.Errorf("HousePosition ecliptical returned error = %v", err)
	}
	if math.Abs(mcAscResult[0]-expectedAsc) > DELTA {
		t.Errorf("HousePosition ecliptical for Asc = %.8f; want %.8f", mcAscResult[0], expectedAsc)
	}
	if math.Abs(mcAscResult[1]-expectedMc) > DELTA {
		t.Errorf("HousePosition ecliptical for Mc = %.8f; want %.8f", mcAscResult[1], expectedMc)
	}
	if math.Abs(cuspResult[2]-expectedCusp2) > DELTA {
		t.Errorf("HousePosition ecliptical for cusp 2 = %.8f; want %.8f", cuspResult[2], expectedCusp2)
	}
	if math.Abs(cuspResult[9]-expectedCusp9) > DELTA {
		t.Errorf("HousePosition ecliptical for cusp 9 = %.8f; want %.8f", cuspResult[9], expectedCusp9)
	}
}
