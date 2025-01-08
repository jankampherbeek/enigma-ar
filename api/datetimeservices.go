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
	"log/slog"
)

// JulDayServer provides services for julian day related functionality.
type JulDayServer interface {
	JulDay(request *domain.DateTime) float64
}

type RevJulDayServer interface {
	RevJulDay(jd float64, cal domain.Calendar) (int, int, int, float64)
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
	slog.Info("Starting calculation of JD")
	jd := jds.jdCalc.CalcJd(request.Year, request.Month, request.Day, request.Ut, request.Greg)
	slog.Info("Completed calculation of JD")
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

// RevJulDay returns date and time for a given jd. The return values are year, month, day and ut.
func (rjds RevJulDayService) RevJulDay(jd float64, cal domain.Calendar) (int, int, int, float64) {
	slog.Info("Starting calculation of date/time from JD")
	y, m, d, t := rjds.revJdCalc.CalcRevJd(jd, cal == domain.CalGregorian)
	slog.Info("Completed calculation of date/time from JD")
	return y, m, d, t
}
