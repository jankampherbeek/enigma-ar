/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package analysis

import (
	"enigma-ar/domain"
	"enigma-ar/internal/conversions"
	"errors"
	"fmt"
	"math"
)

const (
	MinObliquity = 22.0
	MaxObliquity = 25.0
)

// LongEquivCalculator calculates longitude equivalents.
type LongEquivCalculator interface {
	CalcEquivalents(positions []domain.DoublePosition, obliquity float64) ([]domain.SinglePosition, error)
}

type LongEquivCalculation struct{}

func NewLongEquivCalculation() LongEquivCalculator {
	return LongEquivCalculation{}
}

// CalcEquivalents calculates the longitude equivalents for one or more cvhartpoints
// Positions shopuld contain the index of the chartpoint, the longitude (Position1) and the declination (Position2)
// PRE positions contains >= 1 items
// PRE 22 < obliquity < 25
// PRE for all longitudes in positions (Postion1): 0.0 <= longitude < 360.0
// PRE for all declinations in positions (Position2) : -180.0 < declonation < 180.0
// POST if no errors: returns slice with calculated longitude equivalents
// POST if error(s): returns empty slice with error
func (lec LongEquivCalculation) CalcEquivalents(positions []domain.DoublePosition, obliquity float64) ([]domain.SinglePosition, error) {
	longitudeEquivalents := make([]domain.SinglePosition, 0)
	emptyResult := make([]domain.SinglePosition, 0)
	if len(positions) < 1 {
		return emptyResult, errors.New("longequiv calculation failed, received empty slice of positions")
	}
	if obliquity < MinObliquity || obliquity > MaxObliquity {
		return emptyResult, fmt.Errorf("obliquity is out of range, value was %f and should be between %f and %f", obliquity, MinObliquity, MaxObliquity)
	}

	for _, longDeclPos := range positions {
		radixDeclination := longDeclPos.Position2
		declination := radixDeclination
		longitude := longDeclPos.Position1
		if declination < domain.MinDeclination || declination > domain.MaxDeclination {
			return emptyResult, fmt.Errorf("found declination that is out of range, value was %f and should be between %f and %f", declination, domain.MinDeclination, domain.MaxDeclination)
		}
		if longitude < domain.MinLongitude || longitude > domain.MaxLongitude {
			return emptyResult, fmt.Errorf("found longitude that is out of range, value was %f and should be between %f and %f", longitude, domain.MinLongitude, domain.MaxLongitude)
		}

		if math.Abs(radixDeclination) > obliquity { // OOB
			oobPart := math.Abs(radixDeclination) - obliquity
			if radixDeclination > 0 {
				declination = obliquity - oobPart
			} else {
				declination = oobPart - obliquity
			}
		}

		candidate1 := conversions.DeclinationToLongitude(obliquity, declination)
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
