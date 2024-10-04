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

// FullPointServer returns all positions and speeds (ecliptical, equatorial and horizontal) for a given point.
type FullPointServer interface {
	FullPositions(request domain.PointPositionsRequest) ([]domain.PointPosResult, error)
}

// PointRangeServer returns the positions for a range of julian day numbers.
type PointRangeServer interface {
	DefinePointRange(request domain.PointRangeRequest) ([]domain.PointRangeResult, error)
}

type FullChartServer interface {
	DefineFullChart(request domain.FullChartRequest) (domain.FullChartResponse, error)
}

type FullPointService struct {
	fpCalc calc.PointPosCalculator
}

func NewFullPointService() FullPointService {
	return FullPointService{
		calc.NewPointPosCalculation(),
	}
}

func (fps FullPointService) FullPositions(request domain.PointPositionsRequest) ([]domain.PointPosResult, error) {
	positions, err := fps.fpCalc.CalcPointPos(request)
	// TODO log if error occurs
	return positions, err
}

type PointRangeService struct {
	prCalc calc.PointRangeCalculator
}

func NewPointRangeService() PointRangeService {
	return PointRangeService{
		calc.NewPointRangeCalculation(),
	}
}

func (prs PointRangeService) DefinePointRange(request domain.PointRangeRequest) ([]domain.PointRangeResult, error) {
	// TODO check validness of request:
	// existing id for point, jdEnd after jdStart, existing value for Ayanamsha, interval positive

	return prs.prCalc.CalcPointRange(request)
}

type FullChartService struct {
	fcc calc.FullChartCalculator
}

func NewFullChartServer() FullChartServer {
	return FullChartService{calc.NewFullChartCalculator()}
}

func (fcs FullChartService) DefineFullChart(request domain.FullChartRequest) (domain.FullChartResponse, error) {

	return fcs.fcc.CalcFullChart(request)

}
