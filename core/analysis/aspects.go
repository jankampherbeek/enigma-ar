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
	"fmt"
	"math"
)

type AspectsCalculator interface {
	CalcAspects(points []domain.SinglePosition,
		aspects []domain.Aspect,
		cfgPoints []domain.ConfigPoint,
		cfgAspects []domain.ConfigAspect,
		baseOrb float64) ([]domain.ActualAspect, error)
}

type AspectsCalculation struct{}

func NewAspectsCalculation() AspectsCalculator {
	return AspectsCalculation{}
}

// CalcAspects returns the actual aspects
// PRE length points >= 2
// PRE length aspects >= 1
// PRE length cfgPoints >= 2
// PRE length cfgAspects >= 1
// PRE each point is represented in cfgPoints
// PRE each aspect is represented in cfgAspects
// PRE for all positions : 0.0 <= position < 360.0
// POST no errors -> returns slice of occupied midpoints
// POST errors: returns empty slice and error
func (ac AspectsCalculation) CalcAspects(points []domain.SinglePosition,
	aspects []domain.Aspect,
	cfgPoints []domain.ConfigPoint,
	cfgAspects []domain.ConfigAspect,
	baseOrb float64) ([]domain.ActualAspect, error) {

	const (
		MinPointsForCalcAsp  = 2
		MinAspectsForCalcAsp = 1
		FullCircle           = 360.0
		Zero                 = 0.0
	)

	emptyResults := make([]domain.ActualAspect, 0)
	actualAspects := make([]domain.ActualAspect, 0)

	if len(points) < MinPointsForCalcAsp {
		return emptyResults, errors.New("not enough points")
	}
	if len(cfgPoints) < MinPointsForCalcAsp {
		return emptyResults, errors.New("not enough configured points")
	}
	if len(aspects) < MinAspectsForCalcAsp {
		return emptyResults, errors.New("not enough aspects")
	}
	if len(cfgAspects) < MinAspectsForCalcAsp {
		return emptyResults, errors.New("not enough configured aspects")
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
			return emptyResults, fmt.Errorf("point %d not found in configured points", point.Id)
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
			return emptyResults, fmt.Errorf("aspect %d not found in configured aspects", aspect)
		}
	}
	// check if positions are within range
	for _, point := range points {
		if point.Position > domain.MaxLongitude || point.Position < domain.MinLongitude {
			return emptyResults, fmt.Errorf("point %d is out of range, longitude is %f and should be >= %f and < %f",
				point.Id, point.Position, domain.MinLongitude, domain.MaxLongitude)
		}
	}

	var currentCfgOrb, aspectDistance, delta, delta1, delta2 float64
	var exactness int
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			// calculate distances
			distance1 := math.Abs(points[i].Position - points[j].Position)
			if distance1 < Zero {
				distance1 += FullCircle
			}
			distance2 := FullCircle - distance1
			for _, aspect := range aspects {
				// define orb
				var factor1, factor2 float64
				var aspectFactor float64
				for _, cfgPoint := range cfgPoints {
					if cfgPoint.ActualPoint == points[i].Id {
						factor1 = cfgPoint.OrbFactor
					}
					if cfgPoint.ActualPoint == points[j].Id {
						factor2 = cfgPoint.OrbFactor
					}
				}
				for _, cfgAspect := range cfgAspects {
					if aspect == cfgAspect.ActualAspect {
						aspectFactor = cfgAspect.OrbFactor
						aspectDistance = domain.AllAspects()[aspect].Distance
					}
				}
				currentCfgOrb = ((math.Max(factor1, factor2) * aspectFactor) / 10000) * baseOrb
				// check for match
				delta1 = math.Abs(distance1 - aspectDistance)
				if delta1 < Zero {
					delta1 += FullCircle
				}
				delta2 = math.Abs(distance2 - aspectDistance)
				if delta2 < Zero {
					delta2 += FullCircle
				}
				delta = math.Min(delta1, delta2)

				if delta <= currentCfgOrb {
					// match
					exactness = 100 - int((delta/currentCfgOrb)*100)
					actualAspects = append(actualAspects, domain.ActualAspect{
						Pos1:         points[i],
						Pos2:         points[j],
						ActualAspect: aspect,
						ActualOrb:    delta,
						Exactness:    exactness,
					})
				}
			}
		}
	}
	return actualAspects, nil
}
