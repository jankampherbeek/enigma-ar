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
)

const (
	MIN_HARMONIC = 1
	MAX_HARMONIC = 100_000
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
// PRE: 1 <= harmonicNr <= 1000
// PRE: for all values for position in actPositions: 0.0 <= value < 360.0
// PRE: length actPostions > 0
// POST: no errors -> returns calculated harmonics
// POST: contains errors -> returns slice of zero positions and error
func (h HarmonicsCalculation) CalcHarmonics(actPositions []domain.SinglePosition, harmonicNr float64) ([]domain.SinglePosition, error) {
	result := make([]domain.SinglePosition, len(actPositions))
	emptyResult := make([]domain.SinglePosition, 0)
	if harmonicNr < MIN_HARMONIC || harmonicNr > MAX_HARMONIC {
		return emptyResult, fmt.Errorf("harmonics calculation failed, harmonicNr should be > 0.0 and <= 1000, but was %f", harmonicNr)
	}
	if len(actPositions) < 1 {
		return emptyResult, errors.New("harmonics calculation failed, no data found")
	}
	for i, pos := range actPositions {
		if pos.Position < 0.0 || pos.Position >= 360.0 {
			return emptyResult, fmt.Errorf("harmonics calculation failed, encountered position %f, this is outside range: >= 0.0 and < 360.0", pos.Position)
		}
		result[i] = domain.SinglePosition{Id: pos.Id, Position: inRange360(pos.Position * harmonicNr)}
	}
	return result, nil
}

func inRange360(originalValue float64) float64 {
	inRangeValue := originalValue
	for inRangeValue >= 360.0 {
		inRangeValue -= 360.0
	}
	return inRangeValue
}
