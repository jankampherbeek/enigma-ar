/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	"enigma-ar/domain"
	"math"
	"testing"
)

func TestCalculatePersephoneRam(t *testing.T) {
	jdUt := 2434406.817711
	expected := 326.6011343685
	c := NewPointsElementsCalculation()
	result := c.Calculate(domain.AllChartPoints()[domain.PersephoneRam].CalcId, jdUt, domain.ObsPosGeocentric)
	if math.Abs(result[0]-expected) > 1e-8 {
		t.Errorf("Calculation of Persephone (Ram) failed. Expected %f, got %f", expected, result)
	}
}

func TestCalculateHermesRam(t *testing.T) {
	jdUt := 2434406.817711
	expected := 161.6211128197
	c := NewPointsElementsCalculation()
	result := c.Calculate(domain.AllChartPoints()[domain.HermesRam].CalcId, jdUt, domain.ObsPosGeocentric)
	if math.Abs(result[0]-expected) > 1e-8 {
		t.Errorf("Calculation of Hermes (Ram) failed. Expected %f, got %f", expected, result)
	}
}

func TestCalculateDemeterRam(t *testing.T) {
	jdUt := 2434406.817711
	expected := 261.4081200589
	c := NewPointsElementsCalculation()
	result := c.Calculate(domain.AllChartPoints()[domain.DemeterRam].CalcId, jdUt, domain.ObsPosGeocentric)
	if math.Abs(result[0]-expected) > 1e-8 {
		t.Errorf("Calculation of Demeter (Ram) failed. Expected %f, got %f", expected, result)
	}
}
