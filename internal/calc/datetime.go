/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import "enigma-ar/internal/se"

// JulDayCalculator calculates the Julian Day for Epehemeris Time.
type JulDayCalculator interface {
	CalcJd(year int, month int, day int, ut float64, greg bool) float64
}

type JulDayCalculation struct {
	seCalc se.SeJulDayCalculator
}

func NewJulDayCalculation() JulDayCalculator {
	sjc := se.NewSeJulDayCalculation()
	return JulDayCalculation{sjc}
}

// CalcJd handles the calculation of a Julian day number.
func (jdc JulDayCalculation) CalcJd(year int, month int, day int, ut float64, greg bool) float64 {
	var gregFlag = 1
	if !greg {
		gregFlag = 0
	}
	return jdc.seCalc.SeCalcJd(year, month, day, ut, gregFlag)
}