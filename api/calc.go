/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package api

import (
	"enigma-ar/internal/calc"
	"enigma-ar/internal/domain"
)

func JulDay(request domain.JulDayRequest) domain.FloatResponse {
	jd := calc.CalcJd(request.Year, request.Month, request.Day, request.Ut, request.Greg)
	err := 0
	return domain.FloatResponse{jd, err}
}

func FullPositions(request domain.PointPositionsRequest) ([]domain.PointPosResult, error) {
	positions, err := calc.FullPositions(request)
	// TODO log if error occurs
	return positions, err
}
