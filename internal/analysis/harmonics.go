/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package analysis

import (
	"enigma-ar/domain"
	"enigma-ar/internal/calc"
)

// HarmonicsCalculator calculates harmonics.
type HarmonicsCalculator interface {
	CalcHarmonics(actPositions []domain.SinglePosition, harmonicNr float64) ([]domain.SinglePosition, error)
}

type HarmonicsCalculation struct{}

func NewHarmonicsCalculation() HarmonicsCalculator {
	return HarmonicsCalculation{}
}

// CalcHarmonics calculates harmonics for a range (1 to n) of positions.
func (h HarmonicsCalculation) CalcHarmonics(actPositions []domain.SinglePosition, harmonicNr float64) ([]domain.SinglePosition, error) {
	result := make([]domain.SinglePosition, len(actPositions))
	for i, pos := range actPositions {
		harmonicPos := pos.Position * harmonicNr
		harmPosInRange, err := calc.ValueToRange(harmonicPos, 0.0, 360.0)
		if err != nil {
			return nil, err
		}
		result[i] = domain.SinglePosition{Id: pos.Id, Position: harmPosInRange}
	}
	return result, nil
}
