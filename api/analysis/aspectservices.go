/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package apianalysis

import (
	"enigma-ar/domain"
	"enigma-ar/internal/analysis"
	"errors"
	"fmt"
	"log/slog"
)

// AspectServer provides services for the calculation of aspects.
type AspectServer interface {
	Aspects(points []domain.SinglePosition,
		aspects []domain.Aspect,
		cfgPoints []domain.ConfigPoint,
		cfgAspects []domain.ConfigAspect,
		baseOrb float64) ([]domain.ActualAspect, error)
}

type AspectService struct {
	aspCalc analysis.AspectsCalculator
}

func NewAspectService() *AspectService {
	aspCalculator := analysis.NewAspectsCalculation()
	return &AspectService{
		aspCalc: aspCalculator,
	}
}

// Aspects handles the calculation of aspects
// PRE length points >= 2
// PRE length aspects >= 1
// PRE length cfgPoints >= 2
// PRE length cfgAspects >= 1
// PRE each point is represented in cfgPoints
// PRE each aspect is represented in cfgAspects
// PRE for all positions : 0.0 <= position < 360.0
// POST no errors -> returns slice of occupied midpoints
// POST errors: returns nil and error
func (as AspectService) Aspects(points []domain.SinglePosition,
	aspects []domain.Aspect,
	cfgPoints []domain.ConfigPoint,
	cfgAspects []domain.ConfigAspect,
	baseOrb float64) ([]domain.ActualAspect, error) {

	const (
		MinPointsForCalcAsp  = 2
		MinAspectsForCalcAsp = 1
	)
	slog.Info("received request")
	if len(points) < MinPointsForCalcAsp {
		slog.Error("not enough points")
		return nil, errors.New("not enough points")
	}
	if len(cfgPoints) < MinPointsForCalcAsp {
		slog.Error("not enough configured points")
		return nil, errors.New("not enough configured points")
	}
	if len(aspects) < MinAspectsForCalcAsp {
		slog.Error("nog enough aspects")
		return nil, errors.New("not enough aspects")
	}
	if len(cfgAspects) < MinAspectsForCalcAsp {
		slog.Error("not enough configured aspects")
		return nil, errors.New("not enough configured aspects")
	}
	var match bool
	// check if points are available as configured point
	for _, point := range points {
		match = false
		for _, cfgPoint := range cfgPoints {
			if point.Id == cfgPoint.ActualPoint {
				match = true
			}
		}
		if !match {
			slog.Error("point not found in configured points", "point", point.Id)
			return nil, fmt.Errorf("point %d not found in configured points", point.Id)
		}
	}
	// check if aspects are available as configured aspect
	for _, aspect := range aspects {
		match = false
		for _, cfgAspect := range cfgAspects {
			if aspect == cfgAspect.ActualAspect {
				match = true
			}
		}
		if !match {
			slog.Error("aspect not found in configured aspects", "aspect", aspect)
			return nil, fmt.Errorf("aspect %d not found in configured aspects", aspect)
		}
	}
	// check if positions are within range
	for _, point := range points {
		if point.Position > domain.MaxLongitude || point.Position < domain.MinLongitude {
			slog.Error("point is out of range", "longitude", fmt.Sprint(point.Position))
			return nil, fmt.Errorf("point %d is out of range, longitude is %f and should be >= %f and < %f",
				point.Id, point.Position, domain.MinLongitude, domain.MaxLongitude)
		}
	}
	// no errors in input, handle the calculation of aspects
	slog.Info("completed calculation of aspects")
	return as.aspCalc.CalcAspects(points, aspects, cfgPoints, cfgAspects, baseOrb)
}
