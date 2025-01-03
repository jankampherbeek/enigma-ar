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

// ParallelsCalculator calculates parallels and contra-parallels.
type ParallelsCalculator interface {
	CalcParallels(actPositions []domain.SinglePosition, orb float64) ([]domain.MatchedParallel, error)
}

type ParallelsCalculation struct{}

func NewParallelsCalculation() ParallelsCalculator {
	return ParallelsCalculation{}
}

// CalcParallels calculates parallels and contraparallels (2 to n) of positions.
func (pc ParallelsCalculation) CalcParallels(actPositions []domain.SinglePosition, orb float64) ([]domain.MatchedParallel, error) {
	result := make([]domain.MatchedParallel, 0)

	for i := 0; i < len(actPositions); i++ {
		for j := i + 1; j < len(actPositions); j++ {
			pos1 := actPositions[i].Position
			pos2 := actPositions[j].Position
			distance := math.Abs(math.Abs(pos1) - math.Abs(pos2))
			if distance <= orb {
				notContra := (pos1 >= 0.0 && pos2 >= 0.0) || (pos1 <= 0.0 && pos2 <= 0.0)
				result = append(result, domain.MatchedParallel{Pos1: actPositions[i], Pos2: actPositions[j], Orb: distance, Parallel: notContra})
			}
		}
	}

	return result, nil
}
