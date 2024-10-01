/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	domain2 "enigma-ar/domain"
	"fmt"
	"math"
	"testing"
)

const delta = 0.000001

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

// TODO reactivate TestCalcPointPos
/*
func TestCalcPointPos(t *testing.T) {
	fc := FakeSePointPosCalculation{}
	ppc := PointPosCalculation{}
	ppc.sePointCalc = fc
	// test happy flow
	phf := make([]int, 1)
	phf[0] = 2
	phf = append(phf, 2)
	requestHF := domain2.PointPositionsRequest{
		Points:   phf,
		JdUt:     123.456,
		GeoLong:  0.0,
		GeoLat:   0.0,
		Coord:    domain2.Ecliptical,
		ObsPos:   domain2.ObsPosGeocentric,
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
	requestError := domain2.PointPositionsRequest{
		Points:   pError,
		JdUt:     123.456,
		GeoLong:  0.0,
		GeoLat:   0.0,
		Coord:    domain2.Ecliptical,
		ObsPos:   domain2.ObsPosGeocentric,
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
}*/

func TestCalcPointRange(t *testing.T) {
	fcr := FakeSePointPosCalcForRange{}
	prc := PointRangeCalculation{}
	prc.sePointCalc = fcr
	request := domain2.PointRangeRequest{
		Point:     2,
		JdStart:   2_400_000.0,
		JdEnd:     2_400_004.0,
		Interval:  2.0,
		Coord:     domain2.Ecliptical,
		MainValue: true,
		Position:  true,
		ObsPos:    domain2.ObsPosGeocentric,
		Ayanamsha: 0,
	}
	result, err := prc.CalcPointRange(request)
	if err != nil {
		t.Errorf("CalcPointRange happy flow returns unexpected error %d", err)
	}
	// TODO use table test
	if math.Abs(result[0].Jd-2_400_000) > delta {
		t.Errorf("CalcPointRange returns wrong jd for first row: %f; wanted %f", result[0].Jd, 2_400_000.0)
	}
	if math.Abs(result[1].Jd-2_400_002) > delta {
		t.Errorf("CalcPointRange returns wrong jd for second row: %f; wanted %f", result[1].Jd, 2_400_002.0)
	}
	if math.Abs(result[2].Jd-2_400_004) > delta {
		t.Errorf("CalcPointRange returns wrong jd for third row: %f; wanted %f", result[2].Jd, 2_400_004.0)
	}
	if math.Abs(result[0].Value-100.0) > delta {
		t.Errorf("CalcPointRange returns wrong value for first row: %f; wanted %f", result[0].Value, 100.0)
	}
	if math.Abs(result[1].Value-103.0) > delta {
		t.Errorf("CalcPointRange returns wrong value for second row: %f; wanted %f", result[1].Value, 103.0)
	}
	if math.Abs(result[2].Value-106.0) > delta {
		t.Errorf("CalcPointRange returns wrong value for third row: %f; wanted %f", result[2].Value, 106.0)
	}

}

type FakeSePointPosCalcForRange struct{}

func (fake FakeSePointPosCalcForRange) SeCalcPointPos(jdUt float64, body int, flags int) ([6]float64, error) {
	if math.Abs(jdUt-2_400_000) < delta {
		return [6]float64{100.0, 1.0, 2.0, 3.0, 4.0, 5.0}, nil
	}
	if math.Abs(jdUt-2_400_002) < delta {
		return [6]float64{103.0, 1.0, 2.0, 3.0, 4.0, 5.0}, nil
	}
	if math.Abs(jdUt-2_400_004) < delta {
		return [6]float64{106.0, 1.0, 2.0, 3.0, 4.0, 5.0}, nil
	}
	if math.Abs(jdUt-2_400_006) < delta {
		return [6]float64{109.0, 1.0, 2.0, 3.0, 4.0, 5.0}, nil
	}
	var emptyArray [6]float64
	err := fmt.Errorf("FakeSePointPosCalcForRange: unexpected value for JD: %f ", jdUt)
	return emptyArray, fmt.Errorf("PointPosCalcForRange: %s", err)
}

/*
func TestCalcHousePos(t *testing.T) {
	fhc := FakeSeHouseCalculation{}

}

type FakeSeHouseCalculation struct{}

func (fake FakeSeHouseCalculation) SeCalcHousePos(jdUt float64, body int, flags int) ([]float64, []float64, error) {
	mcAscPos := []float64{100.0, 192.0, 0.0, 0.0, 0.0, 0.0}
	cuspPos := []float64{0.0, 100.0, 130.0, 160.0, 192.0, 220.0, 250.0, 280.0, 310.0, 340.0, 10.0, 40.0, 70.0}
	return cuspPos, mcAscPos, nil
}
*/
