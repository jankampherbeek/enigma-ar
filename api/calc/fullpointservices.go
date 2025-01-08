/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package apicalc

import (
	"enigma-ar/domain"
	"enigma-ar/internal/calc"
	"errors"
	"fmt"
	"log/slog"
)

// FullPointServer returns all positions and speeds (ecliptical, equatorial and horizontal) for a given point.
type FullPointServer interface {
	FullPositions(request domain.PointPositionsRequest) ([]domain.PointPosResult, error)
}

// PointRangeServer returns the positions for a range of julian day numbers.
type PointRangeServer interface {
	DefinePointRange(request domain.PointRangeRequest) ([]domain.PointRangeResult, error)
}

type FullPointService struct {
	fpCalc calc.PointPosCalculator
}

type PointRangeService struct {
	prCalc calc.PointRangeCalculator
}

func NewFullPointService() FullPointService {
	return FullPointService{
		calc.NewPointPosCalculation(),
	}
}

func NewPointRangeService() PointRangeService {
	return PointRangeService{
		calc.NewPointRangeCalculation(),
	}
}

// FullPositions calculates all positions for a chart
// PRE request.Points contains at least 1 chartpoint
// PRE MinJdGeneral < request.JdUt < MaxJdGeneral
// PRE MinGeoLang <= request.GeoLang <= MaxGeoLang
// PRE MinGeoLat < request.GeoLat < MaxGeoLat
// PRE MinArmc <= request.Armc < MaxArmc
// PRE MinObliquity < request.Obliquity < MaxObliquity
// POST No errors -> returns calculated chart, otherwise returns nil
func (fps FullPointService) FullPositions(request domain.PointPositionsRequest) ([]domain.PointPosResult, error) {
	if len(request.Points) < 1 {
		slog.Error("No points found")
		return nil, errors.New("points must have at least one point")
	}
	if request.JdUt <= domain.MinJdGeneral || request.JdUt >= domain.MaxJdGeneral {
		slog.Error("Jd out of range")
		return nil, fmt.Errorf("jdUt %f is out of range", request.JdUt)
	}
	if request.GeoLong < domain.MinGeoLong || request.GeoLong > domain.MaxGeoLong {
		slog.Error("GeoLong out of range")
		return nil, fmt.Errorf("geoLong %f is out of range", request.GeoLong)
	}
	if request.GeoLat <= domain.MinGeoLat || request.GeoLat >= domain.MaxGeoLat {
		slog.Error("GeoLong out of range")
		return nil, fmt.Errorf("geoLat %f is out of range", request.GeoLat)
	}
	if request.Armc < domain.MinArmc || request.Armc >= domain.MaxArmc {
		slog.Error("Armc out of range")
		return nil, fmt.Errorf("armc %f is out of range", request.Armc)
	}
	if request.Obliquity < domain.MinObliquity || request.Obliquity >= domain.MaxObliquity {
		slog.Error("Obliquity out of range")
		return nil, fmt.Errorf("obliquity %f is out of range", request.Obliquity)
	}
	positions, err := fps.fpCalc.CalcPointPos(request)
	if err != nil {
		slog.Error("Error calculating full points", "error", err, "request", request)
	}
	slog.Info("Completed calculation of full points")
	return positions, err
}

func (prs PointRangeService) DefinePointRange(request domain.PointRangeRequest) ([]domain.PointRangeResult, error) {
	return prs.prCalc.CalcPointRange(request)
}
