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

func (fake FakeSeJulDayCalculation) SeCalcJd(year int, month int, day int, hour float64, gregFlag int) float64 {
	return 123.456
}
