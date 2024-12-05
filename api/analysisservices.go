/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package api

import (
	"enigma-ar/domain"
	"enigma-ar/internal/analysis"
)

// HarmonicsServer provides services for the calculation of harmonics.
type HarmonicsServer interface {
	Harmonics(actPositions []domain.SinglePosition, harmNr float64) ([]domain.SinglePosition, error)
}

type HarmonicsService struct {
	harmAnalysis analysis.HarmonicsCalculator
}

func NewHarmonicsService() *HarmonicsService {
	return &HarmonicsService{}
}

func (hs HarmonicsService) Harmonics(actPositions []domain.SinglePosition, harmNr float64) ([]domain.SinglePosition, error) {
	result, err := hs.harmAnalysis.CalcHarmonics(actPositions, harmNr)
	if err != nil {
		// TODO log error
	}
	return result, err
}
