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
)

// JulDayServer provides services for julian day related functionality.
type JulDayServer interface {
	JulDay(request *domain.DateTime) float64
}

type JulDayService struct {
	jdCalc calc.JulDayCalculator
}

func NewJulDayService(jdCalc calc.JulDayCalculator) *JulDayService {
	return &JulDayService{
		calc.NewJulDayCalculation(),
	}
}

// JulDay returns the calculated juilian day number for ephemeris time.
func (jds JulDayService) JulDay(request *domain.DateTime) float64 {
	jd := jds.jdCalc.CalcJd(request.Year, request.Month, request.Day, request.Ut, request.Greg)
	return jd
}
