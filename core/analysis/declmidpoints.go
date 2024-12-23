/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package analysis

import (
	"enigma-ar/domain"
	"errors"
	"math"
)

const (
	MinOrbForDMP   = 0.0
	MaxOrbForDMP   = 10.0
	MinItemsForDMP = 3
	MinDeclForDMP  = -180.0
	MaxDeclForDMP  = 180.0
)

// DeclMidpointsCalculator calculates midpoints in declination
type DeclMidpointsCalculator interface {
	CalcDeclMidpoints(positions []domain.SinglePosition, obliquity float64) ([]domain.OccupiedMidpoint, error)
}

type DeclMidpointsCalculation struct{}

func NewDeclMidpointsCalculation() DeclMidpointsCalculator {
	return DeclMidpointsCalculation{}
}

// CalcDeclMidpoints calculates midpoints in declination
// PRE length points >= 3
// PRE for all points : -180.0 < position < 180.0
// PRE 0 < orb < 10.0
// POST no errors -> returns slice of occupied midpoints
// POST errors: returns empty slice and error
func (dmc DeclMidpointsCalculation) CalcDeclMidpoints(points []domain.SinglePosition, orb float64) ([]domain.OccupiedMidpoint, error) {
	emptyResults := make([]domain.OccupiedMidpoint, 0)
	occMidpoints := make([]domain.OccupiedMidpoint, 0)
	if len(points) < MinItemsForDMP {
		return emptyResults, errors.New("not enough points")
	}
	if orb <= MinOrbForDMP || orb > MaxOrbForDMP {
		return emptyResults, errors.New("orb must be between 0.0 and 10.0")
	}

	var actualMP float64
	var actualOrb float64
	var exactness float64
	for i := 0; i < len(points); i++ { // first point
		if points[i].Position <= MinDeclForDMP || points[i].Position >= MaxDeclForDMP {
			return emptyResults, errors.New("declination must be between -180.0 and 180.0 (exclusive)")
		}
		for j := i + 1; j < len(points); j++ { // second point
			for k := 0; k < len(points); k++ { // candidate for midpoint position
				actualMP = (points[i].Position + points[j].Position) / 2.0
				actualOrb = math.Abs(actualMP - points[k].Position)
				if actualOrb <= orb { // match
					exactness = 1 - actualOrb/orb
					occMidpoints = append(occMidpoints, domain.OccupiedMidpoint{
						BaseMidpointPos1: points[i],
						BaseMidpointPos2: points[j],
						FocusPoint:       points[k],
						ActualOrb:        actualOrb,
						Exactness:        exactness,
					})
				}
			}
		}
	}
	return occMidpoints, nil
}
