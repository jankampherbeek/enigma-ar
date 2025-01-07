/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	"math"
	"testing"
)

func TestCalcJd(t *testing.T) {
	fc := FakeSeJulDayCalculation{}
	jdc := JulDayCalculation{}
	jdc.seCalc = fc
	result := jdc.CalcJd(1, 2, 3, 0.0, true)
	expected := 123.456
	difference := math.Abs(result - expected)
	if difference > 0.001 {
		t.Errorf("Julday with fake = %f; want %f", result, expected)
	}
}

type FakeSeJulDayCalculation struct{}

func (fake FakeSeJulDayCalculation) CalcJd(year int, month int, day int, hour float64, gregFlag int) float64 {
	return 123.456
}

func TestCalcRevJd(t *testing.T) {
	fc := FakeSeRevJulDayCalculation{}
	rjdc := RevJulDayCalculation{}
	rjdc.seRevCalc = fc
	resultYear, resultMonth, resultDay, resultUt := rjdc.CalcRevJd(123.456, true)
	expectedYear := 2000
	expectedMonth := 1
	expectedDay := 4
	expectedUt := 0.5
	if expectedYear != resultYear || expectedMonth != resultMonth || expectedDay != resultDay {
		t.Errorf("CalcRevJd with fake year, month, day = %d, %d, %d; want %d, %d, %d", resultYear, resultMonth, resultDay, expectedYear, expectedMonth, expectedDay)
	}
	if math.Abs(expectedUt-resultUt) > 0.00001 {
		t.Errorf("CalcRevJd with fake ut = %f; want %f", resultUt, expectedUt)
	}
}

type FakeSeRevJulDayCalculation struct{}

func (fake FakeSeRevJulDayCalculation) RevCalcJd(jd float64, gregFlag int) (int, int, int, float64) {
	return 2000, 1, 4, 0.5
}
