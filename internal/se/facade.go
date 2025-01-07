/*
 * Enigma Astrology Research.
 * Copyright (c) Jan Kampherbeek.
 * Enigma is open source.
 * Please check the file copyright.txt in the root of the source for further details.
 */

package se

/*
#cgo android CFLAGS: -D_FILE_OFFSET_BITS=64 -D__ANDROID_API__=24 -DANDROID
#cgo android LDFLAGS: -lm -llog
#include <stdio.h>
#include "swe_android.h"
#include "swephexp.h"
*/
import "C"
import (
	"enigma-ar/domain"
	"errors"
	"fmt"
	"math"
	"unsafe"
)

// SwephPreparator handles the initialization of the SE
type SwephPreparator interface {
	SetEphePath(path string)
	SetTopo(geoLong, geoLat, altitudeMtrs float64)
	SetSidereal(ayanamsha domain.Ayanamsha)
	AyanOffset(jdUt float64) (float64, error)
}

// SwephJulDayCalculator retrieves the julian day number for ephemeris time from the SE.
type SwephJulDayCalculator interface {
	CalcJd(year int, month int, day int, hour float64, gregFlag int) float64
}

// SwephRevJulDayCalculator retrieves the date and time for a given jd nr from the SE.
type SwephRevJulDayCalculator interface {
	RevCalcJd(jd float64, gragFlag int) (int, int, int, float64)
}

// SwephPointPosCalculator retrieves the positions and speed for ecliptical or equatorial coordinates.
type SwephPointPosCalculator interface {
	CalcPointPos(jdUt float64, body int, flags int) ([6]float64, error)
}

// SwephEpsilonCalculator retrieves the value for the obliquity of the earths axis, either true (corrected for nutation) or mean.
type SwephEpsilonCalculator interface {
	CalcEpsilon(jdUt float64, trueEps bool) (float64, error)
}

// SwephHorPosCalculator retrieves the horizontal positions (azimuth and altitude) from the SE.
type SwephHorPosCalculator interface {
	CalcHorPos(jdUt float64, geoLong float64, geoLat float64, geoHeight float64, pointRa float64, pointDecl float64, flags int) [3]float64
}

// SwephHousePosCalculator retrieves the housepositions and several other mundane points from the SE.
type SwephHousePosCalculator interface {
	CalcHousePos(houseSys rune, jdUt float64, geoLong float64, geoLat float64, flags int) ([]float64, []float64, error)
}

type SwephPreparation struct{}

func NewSwephPreparation() SwephPreparation {
	return SwephPreparation{}
}

// SetEphePath initializes the SE and defines the location for the ephemeris files.
func (sp SwephPreparation) SetEphePath(path string) {
	var _path *C.char = C.CString(path)
	defer C.free(unsafe.Pointer(_path))
	C.swe_set_ephe_path(_path)
}

// SetTopo initializes the SE for topocentric calculations
func (sp SwephPreparation) SetTopo(geoLong, geoLat, altitudeMtrs float64) {
	gLongC := C.double(geoLong)
	gLatC := C.double(geoLat)
	altitudeC := C.double(altitudeMtrs)
	C.swe_set_topo(gLongC, gLatC, altitudeC)
}

// SetSidereal prepares the Se for sidereal calculations and defines the ayanamsha to be used
func (sp SwephPreparation) SetSidereal(ayanamsha domain.Ayanamsha) {
	seIdAyan := domain.AllAyanamshas()[ayanamsha].CalcId
	ayan := C.int32(seIdAyan)
	C.swe_set_sid_mode(ayan, 0.0, 0.0)
}

func (sp SwephPreparation) AyanOffset(jdUt float64) (float64, error) {
	epheFlag := domain.SeflgSwieph // Enigma always uses this flag
	cSerr := make([]C.char, C.AS_MAXCH)
	cAyanValue := C.double(0.0)
	cJd := C.double(jdUt)
	cFlag := C.int(epheFlag)
	C.swe_get_ayanamsa_ex_ut(cJd, cFlag, &cAyanValue, &cSerr[0])
	errTxt := C.GoString(&cSerr[0])
	if errTxt != "" {
		return 0.0, errors.New(errTxt)
	}
	result := float64(cAyanValue)
	return result, nil
}

type SwephJulDayCalculation struct{}

func NewSwephJulDayCalculation() SwephJulDayCalculator {
	return SwephJulDayCalculation{}
}

// CalcJd accesses the SE to calculate the Julian Day Number, given the values for the date, time and calendar.
func (jdc SwephJulDayCalculation) CalcJd(year int, month int, day int, hour float64, gregFlag int) float64 {
	cYear := C.int(year)
	cMonth := C.int(month)
	cDay := C.int(day)
	cHour := C.double(hour)
	cGregFlag := C.int(gregFlag)
	result := float64(C.swe_julday(cYear, cMonth, cDay, cHour, cGregFlag))
	return result
}

type SwephRevJulDayCalculation struct{}

func NewSwephRevJulDayCalculation() SwephRevJulDayCalculator {
	return SwephRevJulDayCalculation{}
}

// RevCalcJd accesses the SE to calculate date and time from a julian day number. The return values are year,month,day and ut.
func (rjdc SwephRevJulDayCalculation) RevCalcJd(jd float64, gragFlag int) (int, int, int, float64) {
	var cYear C.int
	var cMonth C.int
	var cDay C.int
	var cHour C.double
	C.swe_revjul(C.double(jd), C.int(gragFlag), &cYear, &cMonth, &cDay, &cHour)
	return int(cYear), int(cMonth), int(cDay), float64(cHour)
}

type SwephPointPosCalculation struct{}

func NewSwephPointPosCalculation() SwephPointPosCalculator {
	return SwephPointPosCalculation{}
}

// CalcPointPos accesses the SE to calculate positions for celestial points.
// The results that are returned are subsequently: longitude or ra, latitude or declination, distance, speed in long. or ra, speed in lat. or decl, speed in dist.
func (ppc SwephPointPosCalculation) CalcPointPos(jdUt float64, body int, flags int) ([6]float64, error) {
	var cPos [6]C.double
	cSerr := make([]C.char, C.AS_MAXCH)
	cJdUt := C.double(jdUt)
	cBody := C.int(body)
	cFlags := C.int(flags)
	result := C.swe_calc_ut(cJdUt, cBody, cFlags, &cPos[0], &cSerr[0])
	if result < 0 {
		var emptyArray [6]float64
		return emptyArray, fmt.Errorf("CalcPointPos error: %v", cSerr)
	}

	pos := make([]float64, 6)
	for i := 0; i < 6; i++ {
		pos[i] = float64(cPos[i])
	}
	return [6]float64(pos), nil
}

type SwephEpsilonCalculation struct{} // TODO create test for SwephEpsilonCalculation

func NewSwephEpsilonCalculation() SwephEpsilonCalculator {
	return SwephEpsilonCalculation{}
}

func (ec SwephEpsilonCalculation) CalcEpsilon(jdUt float64, trueEps bool) (float64, error) {
	var cPos [6]C.double
	cSerr := make([]C.char, C.AS_MAXCH)
	cJdUt := C.double(jdUt)
	cBody := C.int(domain.EclNut)
	cFlags := C.int(domain.SeflgSwieph)
	result := C.swe_calc_ut(cJdUt, cBody, cFlags, &cPos[0], &cSerr[0])
	err := C.GoString(&cSerr[0])
	if result < 0 {
		return 0.0, fmt.Errorf("CalcEpsilon error: %v", err)
	}
	if trueEps {
		return float64(cPos[0]), nil
	}
	return float64(cPos[1]), nil
}

type SwephHorPosCalculation struct{}

func NewSwephHorPosCalculation() SwephHorPosCalculator {
	return SwephHorPosCalculation{}
}

// CalcHorPos converts equatorial coordinates to azimuth, true altitude and apparent altitude. The SE does not return a result code.
func (hpc SwephHorPosCalculation) CalcHorPos(jdUt float64, geoLong float64, geoLat float64, geoHeight float64, pointRa float64, pointDecl float64, flags int) [3]float64 {
	var cHorCoord [3]C.double
	cJdUt := C.double(jdUt)
	cFlags := C.int(flags)
	cAtPress := C.double(0.0)
	cAtTemp := C.double(0.0)
	geoCoord := []float64{geoLong, geoLat, geoHeight}
	pointCoord := []float64{pointRa, pointDecl}
	cGeoCoord := (*C.double)(&geoCoord[0])
	cPointCoord := (*C.double)(&pointCoord[0])
	_ = C.swe_azalt(cJdUt, cFlags, cGeoCoord, cAtPress, cAtTemp, cPointCoord, &cHorCoord[0])
	pos := make([]float64, 3)
	for i := 0; i < 3; i++ {
		pos[i] = float64(cHorCoord[i])
	}
	return [3]float64(pos)
}

type SwephHousePosCalculation struct{}

func NewSwephHousePosCalculation() SwephHousePosCalculator {
	return SwephHousePosCalculation{}
}

// CalcHousePos calculates mc, asc, and cusps for a given house system, jd, and location.
// Values returned: array with cusps, starting at index 1, array with positions of asc, mc, armc, vertex, eq asc,
// co-asc Koch, co-asc Munkasey and three empty values.
func (hp SwephHousePosCalculation) CalcHousePos(houseSys rune, jdUt float64, geoLong float64, geoLat float64, flags int) ([]float64, []float64, error) {
	cJdUt := C.double(jdUt)
	cGeolat := C.double(geoLat)
	cGeolong := C.double(geoLong)
	cHouseSys := C.int(houseSys)
	cFlags := C.int(flags)

	var cCusps [13]C.double
	var cAscMc [10]C.double

	result := C.swe_houses_ex(cJdUt, cFlags, cGeolat, cGeolong, cHouseSys, &cCusps[0], &cAscMc[0])
	if result < 0 {
		fmt.Printf("Error in HousePositions: %v", result)
		err := errors.New("error in housepositions")
		return make([]float64, 13), make([]float64, 10), err
	}
	cusps := make([]float64, 13)
	for i := 0; i < 13; i++ {
		cusps[i] = float64(cCusps[i])
	}
	ascMc := make([]float64, 10)
	for i := 0; i < 10; i++ {
		ascMc[i] = float64(cAscMc[i])
	}
	return cusps, ascMc, nil
}

// Coordinate transformation

type SwephCoordinateTransformer interface {
	Transform(valuesIn *[3]float64, eps float64, ec2Equ bool) []float64
}

type SwephCoordinateTransform struct{}

func NewSwephCoordinateTransform() *SwephCoordinateTransform {
	return &SwephCoordinateTransform{}
}

// Transform ecliptic to equatorial or the other way around. Valuesin and valuesout contain resp. long, lat, distance or ra, decl, distance.
func (ct SwephCoordinateTransform) Transform(valuesIn *[3]float64, eps float64, ec2Equ bool) []float64 {
	cValuesIn := (*C.double)(&valuesIn[0])
	var correctedEsp = math.Abs(eps) // SE expects positive epsilon for equatorial 2 ecliptical
	if ec2Equ {                      // and negatieve epsilon for ecliptical to equatorial
		correctedEsp *= -1
	}
	cEps := C.double(eps)
	var cValuesOut [3]C.double
	C.swe_cotrans(cValuesIn, &cValuesOut[0], cEps)
	valuesOut := make([]float64, 3)
	for i := 0; i < 3; i++ {
		valuesOut[i] = float64(cValuesOut[i])
	}
	return valuesOut
}
