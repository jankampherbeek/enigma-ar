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

// HarmonicServer provides services for the calculation of harmonics
type HarmonicServer interface {
	Harmonics(actPositions []domain.SinglePosition, harmonicNr float64) ([]domain.SinglePosition, error)
}

type HarmonicService struct {
	hrmCalc analysis.HarmonicsCalculator
}

func NewHarmonicService() *HarmonicService {
	hrmCalculator := analysis.NewHarmonicsCalculation()
	return &HarmonicService{
		hrmCalc: hrmCalculator,
	}
}

// Harmonics handles the calculation of harmonics
// PRE: 1 <= harmonicNr <= 1000
// PRE: for all values for position in actPositions: 0.0 <= value < 360.0
// PRE: length actPostions > 0
// POST: no errors -> returns calculated harmonics
// POST: contains errors -> returns nil and error
func (hs HarmonicService) Harmonics(actPositions []domain.SinglePosition, harmonicNr float64) ([]domain.SinglePosition, error) {
	const (
		MinHarmonic = 1
		MaxHarmonic = 100_000
	)
	slog.Info("Starting calculation of harmonics")

	if harmonicNr < MinHarmonic || harmonicNr > MaxHarmonic {
		slog.Error("Harmonic number out of range")
		return nil, fmt.Errorf("harmonics failed, harmonicNr should be > 0.0 and <= 1000, but was %f", harmonicNr)
	}
	if len(actPositions) < 1 {
		slog.Error("no data found")
		return nil, errors.New("harmonics failed, no data found")
	}
	for _, pos := range actPositions {
		if pos.Position < 0.0 || pos.Position >= 360.0 {
			slog.Error("position out of range")
			return nil, fmt.Errorf("harmonics failed, encountered position %f, this is outside range: >= 0.0 and < 360.0", pos.Position)
		}
	}
	slog.Info("Completed calculation of harmonics")
	return hs.hrmCalc.CalcHarmonics(actPositions, harmonicNr)
}
