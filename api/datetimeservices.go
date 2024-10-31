/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package api

import (
	"enigma-ar/domain"
	"enigma-ar/domain/references"
	"enigma-ar/internal/calc"
)

// JulDayServer provides services for julian day related functionality.
type JulDayServer interface {
	JulDay(request *domain.DateTime) float64
}

type RevJulDayServer interface {
	RevJulDay(jd float64, cal references.Calendar) (int, int, int, float64)
}

type JulDayService struct {
	jdCalc calc.JulDayCalculator
}

func NewJulDayService() *JulDayService {
	jdCalc := calc.NewJulDayCalculation()
	return &JulDayService{
		jdCalc: jdCalc,
	}
}

// JulDay returns the calculated julian day number for ephemeris time.
func (jds JulDayService) JulDay(request *domain.DateTime) float64 {
	jd := jds.jdCalc.CalcJd(request.Year, request.Month, request.Day, request.Ut, request.Greg)
	return jd
}

type RevJulDayService struct {
	revJdCalc calc.RevJulDayCalculator
}

func NewRevJulDayService() *RevJulDayService {
	revJdCalc := calc.NewRevJulDayCalculation()
	return &RevJulDayService{
		revJdCalc: revJdCalc,
	}
}

// RevJulDay returns date and time for a given jd. The returnvalues are year, month, day and ut.
func (rjds RevJulDayService) RevJulDay(jd float64, cal references.Calendar) (int, int, int, float64) {
	return rjds.revJdCalc.CalcRevJd(jd, cal == references.CalGregorian)
}
