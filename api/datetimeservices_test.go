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

// todo use fakes for testing jd calculations in datetimeservices
func TestJulDay(t *testing.T) {
	request := domain.DateTime{Year: 2024, Month: 5, Day: 6, Ut: 20.5, Greg: true}
	jdCalc := calc.NewJulDayCalculation()
	result := jdCalc.CalcJd(request.Year, request.Month, request.Day, request.Ut, request.Greg)
	expected := 2460437.3541666665
	difference := math.Abs(result - expected)
	if difference > 0.000001 {
		t.Errorf("Julday(2024,5,6,20.5, true) = %f; want %f", result, expected)
	}
}

func TestRevJulDay(t *testing.T) {
	jd := 2460437.3541666665
	revJdCalc := calc.NewRevJulDayCalculation()
	resultYear, resultMonth, resultDay, resultUt := revJdCalc.CalcRevJd(jd, true)
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
