/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package api

import (
	"enigma-ar/domain"
	"enigma-ar/internal/calc"
	"math"
	"testing"
)

func TestCalcInt(t *testing.T) {
	request := domain.DateTime{Year: 2024, Month: 5, Day: 6, Ut: 20.5, Greg: true}
	jdCalc := calc.NewJulDayCalculation()
	result := jdCalc.CalcJd(request.Year, request.Month, request.Day, request.Ut, request.Greg)
	expected := 2460437.3541666665
	difference := math.Abs(result - expected)
	if difference > 0.000001 {
		t.Errorf("Julday(2024,5,6,20.5, true) = %f; want %f", result, expected)
	}
}

func TestCalc(t *testing.T) {

}
