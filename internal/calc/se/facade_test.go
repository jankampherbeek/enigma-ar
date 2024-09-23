/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package se

import (
	"enigma-ar/internal/domain"
	"math"
	"testing"
)

const DELTA = 1e-8

func TestJulDay(t *testing.T) {
	result := SeJulDayCalculation{}.SeCalcJd(2024, 5, 6, 20.5, 1)
	expected := 2460437.3541666665
	difference := math.Abs(result - expected)
	if difference > 0.000001 {
		t.Errorf("Julday(2024,5,6,20.5, true) = %f; want %f", result, expected)
	}
}

func TestPointPositions(t *testing.T) {
	ephePath := "..\\..\\..\\sedata" // path is relative from current package
	SetEphePath(ephePath)
	julDay := 2_470_000.0 // 2050/7/12 12:00
	body := domain.SE_MERCURY
	flags := domain.SEFLG_SWIEPH + domain.SEFLG_SPEED
	// TODO check all 6 values
	expected := []float64{132.309351305555, 1.309320472222, 1.106102572, 1.572654666667}
	result, err := SePointPosCalculation{}.SeCalcPointPos(julDay, body, flags)
	if err != nil {
		t.Errorf("PointPositions(2_470_000, SE_MERCURY, 256) returns error %s", err)
	} else {
		for i := 0; i <= 3; i++ {
			if math.Abs(result[i]-expected[i]) > DELTA {
				t.Errorf("PointPositionsJ(2_470_000, SE_MERCURY, 256) = %f; want %f", result[i], expected[i])
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
	result := SeHorPosCalculation{}.CalcHorPos(jdUt, geoLong, geoLat, geoHeight, pointRa, pointDecl, flags)
	for i := 0; i <= 2; i++ {
		if math.Abs(result[i]-expected[i]) > DELTA {
			t.Errorf("HorizontalPosition(2_434_406.8177, 6.9, 52.2166, 0.0, 0.0, 317.1878, -16.4229, 2048) = %f; want %f", result[i], expected[i])
		}
	}
}
