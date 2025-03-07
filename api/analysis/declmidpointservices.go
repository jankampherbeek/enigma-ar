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
	"log/slog"
)

// DeclinationMidpointServer provides services for the calculation of midpoints in declination
type DeclinationMidpointServer interface {
	DeclinationMidpoints(positions []domain.SinglePosition, orb float64) ([]domain.OccupiedMidpoint, error)
}

type DeclinationMidpointService struct {
	dmpCalc analysis.DeclMidpointsCalculator
}

func NewDeclinationMidpointService() *DeclinationMidpointService {
	dmpCalculator := analysis.NewDeclMidpointsCalculation()
	return &DeclinationMidpointService{
		dmpCalc: dmpCalculator,
	}
}

const (
	MinOrbForDMP   = 0.0
	MaxOrbForDMP   = 10.0
	MinItemsForDMP = 3
	MinDeclForDMP  = -180.0
	MaxDeclForDMP  = 180.0
)

// DeclinationMidpoints handles the calculation of midpoints in declination
// PRE length points >= 3
// PRE for all points : -180.0 < position < 180.0
// PRE 0 < orb < 10.0
// POST no errors -> returns slice of occupied midpoints
// POST errors: returns empty slice and error
func (dmps DeclinationMidpointService) DeclinationMidpoints(positions []domain.SinglePosition, orb float64) ([]domain.OccupiedMidpoint, error) {
	slog.Info("Starting calculation of declination midpoints")
	if len(positions) < MinItemsForDMP {
		slog.Error("not enough positions")
		return nil, errors.New("not enough positions")
	}
	if orb <= MinOrbForDMP || orb > MaxOrbForDMP {
		slog.Error("Orb is out of range")
		return nil, errors.New("orb id out of range, must be between 0.0 and 10.0")
	}
	for i := 0; i < len(positions); i++ {
		if positions[i].Position <= MinDeclForDMP || positions[i].Position >= MaxDeclForDMP {
			slog.Error("Declination out of range")
			return nil, errors.New("declination out of range, must be between -180.0 and 180.0 (exclusive)")
		}
	}
	slog.Info("DeclinationMidpoints calculated")
	return dmps.dmpCalc.CalcDeclMidpoints(positions, orb)
}
