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

// RevJulDayCalculator calculates the date and time for a given Julian Day number.
type RevJulDayCalculator interface {
	CalcRevJd(jd float64, greg bool) (int, int, int, float64)
}

type JulDayCalculation struct {
	seCalc se.SwephJulDayCalculator
}

func NewJulDayCalculation() JulDayCalculator {
	sjc := se.NewSwephJulDayCalculation()
	return JulDayCalculation{sjc}
}

// CalcJd handles the calculation of a Julian day number.
func (jdc JulDayCalculation) CalcJd(year int, month int, day int, ut float64, greg bool) float64 {
	var gregFlag = 1
	if !greg {
		gregFlag = 0
	}
	return jdc.seCalc.CalcJd(year, month, day, ut, gregFlag)
}

type RevJulDayCalculation struct {
	seRevCalc se.SwephRevJulDayCalculator
}

func NewRevJulDayCalculation() RevJulDayCalculator {
	srjc := se.NewSwephRevJulDayCalculation()
	return RevJulDayCalculation{srjc}
}

// CalcRevJd Calculates date and trme for a jd nr. Returns year, month, day and ut.
func (rjdc RevJulDayCalculation) CalcRevJd(jd float64, greg bool) (int, int, int, float64) {
	var gregFlag = 1
	if !greg {
		gregFlag = 0
	}
	return rjdc.seRevCalc.RevCalcJd(jd, gregFlag)
}
