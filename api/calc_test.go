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
	"math"
	"testing"
)

func TestCalc(t *testing.T) {
	request := domain.JulDayRequest{2024, 5, 6, 20.5, true}

	result := calc.CalcJd(request.Year, request.Month, request.Day, request.Ut, request.Greg)
	expected := 2460437.3541666665
	difference := math.Abs(result - expected)
	if difference > 0.000001 {
		t.Errorf("Julday(2024,5,6,20.5, true) = %f; want %f", result, expected)
	}

}
