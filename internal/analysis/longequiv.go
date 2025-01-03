/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package analysis

import (
	"enigma-ar/domain"
	"enigma-ar/internal/calc/conversion"
	"math"
)

// LongEquivCalculator calculates longitude equivalents.
type LongEquivCalculator interface {
	CalcEquivalents(positions []domain.DoublePosition, obliquity float64) ([]domain.SinglePosition, error)
}

type LongEquivCalculation struct{}

func NewLongEquivCalculation() LongEquivCalculator {
	return LongEquivCalculation{}
}

// CalcEquivalents calculates the longitude equivalents for one or more chartpoints

func (lec LongEquivCalculation) CalcEquivalents(positions []domain.DoublePosition, obliquity float64) ([]domain.SinglePosition, error) {
	longitudeEquivalents := make([]domain.SinglePosition, 0)

	for _, longDeclPos := range positions {
		radixDeclination := longDeclPos.Position2
		declination := radixDeclination
		longitude := longDeclPos.Position1

		if math.Abs(radixDeclination) > obliquity { // OOB
			oobPart := math.Abs(radixDeclination) - obliquity
			if radixDeclination > 0 {
				declination = obliquity - oobPart
			} else {
				declination = oobPart - obliquity
			}
		}

		candidate1 := conversion.DeclinationToLongitude(obliquity, declination)
		if candidate1 < 0.0 {
			candidate1 += 360.0
		}

		var candidate2 float64
		if longitude < 180.0 {
			candidate2 = 180.0 - candidate1
		} else {
			candidate2 = 540.0 - candidate1
		}
		if candidate2 > 360.0 {
			candidate2 -= 360.0
		}
		if candidate2 < 0.0 {
			candidate2 += 360.0
		}

		diff1 := math.Abs(candidate1 - longitude)
		diff2 := math.Abs(candidate2 - longitude)

		longitudeEquivalent := candidate1
		if diff2 < diff1 {
			longitudeEquivalent = candidate2
		}

		longitudeEquivalents = append(longitudeEquivalents, domain.SinglePosition{
			Id:       longDeclPos.Id,
			Position: longitudeEquivalent,
		})
	}

	return longitudeEquivalents, nil
}
