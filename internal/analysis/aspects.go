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
func (ac AspectsCalculation) CalcAspects(points []domain.SinglePosition,
	aspects []domain.Aspect,
	cfgPoints []domain.ConfigPoint,
	cfgAspects []domain.ConfigAspect,
	baseOrb float64) ([]domain.ActualAspect, error) {

	const (
		FullCircle = 360.0
		Zero       = 0.0
	)

	actualAspects := make([]domain.ActualAspect, 0)

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
