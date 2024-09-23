/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package api

import (
	"enigma-ar/internal/calc"
	"enigma-ar/internal/domain"
)

type JulDayServer interface {
	JulDay(request domain.DateTime) float64
}

type FullPointServer interface {
	FullPositions(request domain.PointPositionsRequest) ([]domain.PointPosResult, error)
}

type JulDayService struct {
	jdCalc calc.JulDayCalculator
}

func NewJulDayService(jdCalc calc.JulDayCalculator) JulDayService {
	return JulDayService{
		calc.NewJulDayCalculation(),
	}
}

func (jds JulDayService) JulDay(request domain.DateTime) float64 {
	jd := jds.jdCalc.CalcJd(request.Year, request.Month, request.Day, request.Ut, request.Greg)
	return jd
}

type FullPointService struct {
	fpCalc calc.PointPosCalculator
}

func newFullPointService() FullPointService {
	return FullPointService{
		calc.NewPointPosCalculation(),
	}
}

func (fps FullPointService) FullPositions(request domain.PointPositionsRequest) ([]domain.PointPosResult, error) {
	positions, err := fps.fpCalc.CalcPointPos(request)
	// TODO log if error occurs
	return positions, err
}
