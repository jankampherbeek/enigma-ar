/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	"enigma-ar/internal/domain"
	"fmt"
	"math"
	"testing"
)

const delta = 0.00000001

func TestCalcJd(t *testing.T) {
	fc := FakeSeJulDayCalculation{}
	jdc := JulDayCalculation{}
	jdc.seCalc = fc
	result := jdc.CalcJd(1, 2, 3, 0.0, true)
	expected := 123.456
	difference := math.Abs(result - expected)
	if difference > 0.001 {
		t.Errorf("Julday with fake = %f; want %f", result, expected)
	}
}

type FakeSeJulDayCalculation struct{}

func (fake FakeSeJulDayCalculation) SeCalcJd(year int, month int, day int, hour float64, gregFlag int) float64 {
	return 123.456
}

func TestCalcPointPos(t *testing.T) {
	fc := FakeSePointPosCalculation{}
	ppc := PointPosCalculation{}
	ppc.sePointCalc = fc
	// test happy flow
	var phf []int
	phf = append(phf, 2)
	requestHF := domain.PointPositionsRequest{
		Points:   phf,
		JdUt:     123.456,
		GeoLong:  0.0,
		GeoLat:   0.0,
		Coord:    domain.Ecliptical,
		ObsPos:   domain.Geocentric,
		Tropical: true,
	}
	resultHF, errorHF := ppc.CalcPointPos(requestHF)
	if errorHF != nil {
		t.Errorf("CalcPointPos happy flow returns unexpected error %d", errorHF)
	}
	if math.Abs(resultHF[0].LonPos-1.0) > delta {
		t.Errorf("CalcPointPos happy flow returns wrong value for longitude %f; wanted %f", resultHF[0].LonPos, 1.0)
	}
	if math.Abs(resultHF[0].LatPos-2.0) > delta {
		t.Errorf("CalcPointPos happy flow returns wrong value for latitude %f; wanted %f", resultHF[0].LatPos, 2.0)
	}
	var pError []int
	pError = append(pError, -100)
	requestError := domain.PointPositionsRequest{
		Points:   pError,
		JdUt:     123.456,
		GeoLong:  0.0,
		GeoLat:   0.0,
		Coord:    domain.Ecliptical,
		ObsPos:   domain.Geocentric,
		Tropical: true,
	}
	_, errorErr := ppc.CalcPointPos(requestError)
	if errorErr == nil {
		t.Errorf("CalcPointPos did not return expected error")
	}
}

type FakeSePointPosCalculation struct{}

func (fake FakeSePointPosCalculation) SeCalcPointPos(jdUt float64, body int, flags int) ([6]float64, error) {
	if body == 2 {
		return [6]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}, nil
	}
	if body == -100 {
		var emptyArray [6]float64
		err := "Error"
		return emptyArray, fmt.Errorf("PointPositions error: %s", err)
	}
	var emptyArray [6]float64
	return emptyArray, nil
}

/*
type SePointPosCalculator interface {
	SeCalcPointPos(jdUt float64, body int, flags int) ([6]float64, error)
}


type SePointPosCalculation struct{}

func NewSePointPosCalculation() SePointPosCalculation {
	return SePointPosCalculation{}
}

// CalculatePointPos accesses the SE to calculate positions for celestial points
func (ppc SePointPosCalculation) SeCalcPointPos(jdUt float64, body int, flags int) ([6]float64, error) {
	var cPos [6]C.double
........




--------
type PointPosCalculator interface {
	CalcPointPos(request domain.PointPositionsRequest) ([]domain.PointPosResult, error)
}


type PointPosCalculation struct {
	sePointCalc  se.SePointPosCalculation
	seHorPosCalc se.SeHorPosCalculation
}

func NewPointPosCalculation() PointPosCalculation {
	ppc := se.NewSePointPosCalculation()
	hpc := se.NewSeHorPosCalculation()
	return PointPosCalculation{ppc, hpc}
}

// CalcPointPos calculates fully defined positions for one or more celestial points
func (calc PointPosCalculation) CalcPointPos(request domain.PointPositionsRequest) ([]domain.PointPosResult, error) {
	positions := make([]domain.PointPosResult, 0)
	eclFlags := SeFlags(domain.Ecliptical, request.ObsPos, request.Tropical)
........
	}
	return positions, nil
}


*/
