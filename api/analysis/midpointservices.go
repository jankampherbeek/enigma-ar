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
)

// MidpointServer provides services for the calculation of midpoints
type MidpointsServer interface {
	Midpoints(points []domain.SinglePosition) ([]domain.Midpoint, error)
	OccupiedMidpoints(points []domain.SinglePosition, dial domain.MpDial, orb float64) ([]domain.OccupiedMidpoint, error)
}

type MidpointService struct {
	mpCalc analysis.MidpointsCalculator
}

func NewMidpointService() *MidpointService {
	mpCalculator := analysis.NewMidpointsCalculation()
	return &MidpointService{
		mpCalc: mpCalculator,
	}
}

const (
	MinItemsForMP     = 2
	MinItemsForCalcMP = 3
	MinOrbForMP       = 0.0
	MaxOrbForMP       = 10.0
	MinPosForMP       = 0.0
	MaxPosForMP       = 360.0
)

// Midpoints handles the calculation of midpoints.
// PRE length points >= 2
// PRE for all positions : 0.0 <= position < 360.0
// POST no errors -> returns slice of midpoints
// POST errors: returns nil and error
func (mps MidpointService) Midpoints(points []domain.SinglePosition) ([]domain.Midpoint, error) {
	if len(points) < MinItemsForMP {
		return nil, errors.New("not enough points")
	}
	for i := 0; i < len(points); i++ {
		if points[i].Position < MinPosForMP || points[i].Position >= MaxPosForMP {
			return nil, errors.New("position must be between 0.0 and <360.0")
		}
	}
	return mps.mpCalc.CalcMidpoints(points)
}

// OccupiedMidpoints handles the calculation of occupied midpoints.
// PRE length points >= 3
// PRE for all positions : 0.0 <= position < 360.0
// PRE 0.0 < orb <= 10.0
// POST no errors -> returns nil
// POST errors: returns nil and error
func (mps MidpointService) OccupiedMidpoints(points []domain.SinglePosition, dial domain.MpDial, orb float64) ([]domain.OccupiedMidpoint, error) {
	if len(points) < MinItemsForCalcMP {
		return nil, errors.New("not enough points")
	}
	if orb <= MinOrbForMP || orb > MaxOrbForMP {
		return nil, errors.New("orb must be between 0.0 and 10.0")
	}
	for i := 0; i < len(points); i++ {
		if points[i].Position < MinPosForMP || points[i].Position >= MaxPosForMP {
			return nil, errors.New("positions must be between 0.0 and <360.0")
		}
	}
	return mps.mpCalc.CalcOccupiedMidpoints(points, dial, orb)
}
