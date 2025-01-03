/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package analysis

import (
	"enigma-ar/domain"
	"math"
)

// DeclMidpointsCalculator calculates midpoints in declination
type DeclMidpointsCalculator interface {
	CalcDeclMidpoints(positions []domain.SinglePosition, orb float64) ([]domain.OccupiedMidpoint, error)
}

type DeclMidpointsCalculation struct{}

func NewDeclMidpointsCalculation() DeclMidpointsCalculator {
	return DeclMidpointsCalculation{}
}

// CalcDeclMidpoints calculates midpoints in declination

func (dmc DeclMidpointsCalculation) CalcDeclMidpoints(positions []domain.SinglePosition, orb float64) ([]domain.OccupiedMidpoint, error) {
	occMidpoints := make([]domain.OccupiedMidpoint, 0)

	var actualMP float64
	var actualOrb float64
	var exactness float64
	for i := 0; i < len(positions); i++ { // first point
		for j := i + 1; j < len(positions); j++ { // second point
			for k := 0; k < len(positions); k++ { // candidate for midpoint position
				actualMP = (positions[i].Position + positions[j].Position) / 2.0
				actualOrb = math.Abs(actualMP - positions[k].Position)
				if actualOrb <= orb { // match
					exactness = 1 - actualOrb/orb
					occMidpoints = append(occMidpoints, domain.OccupiedMidpoint{
						BaseMidpointPos1: positions[i],
						BaseMidpointPos2: positions[j],
						FocusPoint:       positions[k],
						ActualOrb:        actualOrb,
						Exactness:        exactness,
					})
				}
			}
		}
	}
	return occMidpoints, nil
}
