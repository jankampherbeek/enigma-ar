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
	"math"
)

// ParallelServer provides services for the calculation of parallels and contraparallels
type ParallelServer interface {
	Parallels(actPositions []domain.SinglePosition, orb float64) ([]domain.MatchedParallel, error)
}

type ParallelService struct {
	parCalc analysis.ParallelsCalculator
}

func NewParallelService() *ParallelService {
	parCalculator := analysis.NewParallelsCalculation()
	return &ParallelService{
		parCalc: parCalculator,
	}
}

// Parallels handles the calculation of parallels and contra parallels
// PRE: length actPostions >= 2
// PRE: 0 < orb < 10
// PRE: for all values for position in actPositions: 0.0 <= value < 180.0
// POST: no errors -> returns calculated parallels and contra parallels
// POST: contains errors -> returns nil and error
func (ps ParallelService) Parallels(actPositions []domain.SinglePosition, orb float64) ([]domain.MatchedParallel, error) {

	const MaxDecl = 180.0
	slog.Info("Started calculation of parallels")
	if len(actPositions) < 2 {
		slog.Error("Not enough positions")
		return nil, errors.New("parallels failed, not enough data")
	}
	if orb <= 0.0 || orb >= 10.0 {
		return nil, errors.New("parallels failed, orb not > 0.0 or not <= 10.0")
	}
	for i := 0; i < len(actPositions); i++ {
		for j := i + 1; j < len(actPositions); j++ {
			pos1 := actPositions[i].Position
			pos2 := actPositions[j].Position
			if math.Abs(pos1) >= MaxDecl || math.Abs(pos2) >= MaxDecl {
				slog.Error("Declination out of range")
				return nil, errors.New("parallels failed, found declination >= 180.0")
			}
		}
	}
	slog.Info("Completed calculation of parallels")
	return ps.parCalc.CalcParallels(actPositions, orb)
}
