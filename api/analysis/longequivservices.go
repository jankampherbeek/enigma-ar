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
	"fmt"
	"log/slog"
)

// LongequivServer provides services for the calculation of longitude equivalents
type LongequivServer interface {
	LongEquivs(positions []domain.DoublePosition, obliquity float64) ([]domain.SinglePosition, error)
}

type LongEquivService struct {
	leCalc analysis.LongEquivCalculator
}

func NewLongEquivService() *LongEquivService {
	leCalculator := analysis.NewLongEquivCalculation()
	return &LongEquivService{
		leCalc: leCalculator,
	}
}

const (
	MinObliquity = 22.0
	MaxObliquity = 25.0
)

// LongEquivs provides services for the calculation of longitude equivalents
// Positions should contain the index of the chartpoint, the longitude (Position1) and the declination (Position2)
// PRE positions contains >= 1 items
// PRE 22 < obliquity < 25
// PRE for all longitudes in positions (Position1): 0.0 <= longitude < 360.0
// PRE for all declinations in positions (Position2) : -180.0 < declonation < 180.0
// POST if no errors: returns slice with calculated longitude equivalents
// POST if error(s): returns empty slice with error
func (les LongEquivService) LongEquivs(positions []domain.DoublePosition, obliquity float64) ([]domain.SinglePosition, error) {
	slog.Info("Started calculation of longitude equivalents")
	if len(positions) < 1 {
		slog.Error("No positions found")
		return nil, errors.New("LongEquivs could not proceed: received empty slice of positions")
	}
	if obliquity < MinObliquity || obliquity > MaxObliquity {
		slog.Error("Obliquity out of range")
		return nil, fmt.Errorf("LongEquivs could not proceed: obliquity is out of range, value was %f and should be between %f and %f", obliquity, MinObliquity, MaxObliquity)
	}
	for _, longDeclPos := range positions {
		declination := longDeclPos.Position2
		longitude := longDeclPos.Position1
		if declination < domain.MinDeclination || declination > domain.MaxDeclination {
			slog.Error("Declination out of range")
			return nil, fmt.Errorf("LongEquivs could not proceed: found declination that is out of range, value was %f and should be between %f and %f", declination, domain.MinDeclination, domain.MaxDeclination)
		}
		if longitude < domain.MinLongitude || longitude > domain.MaxLongitude {
			slog.Error("Longitude out of range")
			return nil, fmt.Errorf("LongEquivs could not proceed: found longitude that is out of range, value was %f and should be between %f and %f", longitude, domain.MinLongitude, domain.MaxLongitude)
		}
	}
	slog.Info("Completed calculation of longitude equivalents")
	return les.leCalc.CalcEquivalents(positions, obliquity)
}
