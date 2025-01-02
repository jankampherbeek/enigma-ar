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

// SeJulDayCalculator retrieves the julian day number for ephemeris time from the SE.
type SeJulDayCalculator interface {
	SeCalcJd(year int, month int, day int, hour float64, gregFlag int) float64
}

// SeRevJulDayCalculator retrieves the date and time for a given jd nr from the SE.
type SeRevJulDayCalculator interface {
	SeRevCalcJd(jd float64, gragFlag int) (int, int, int, float64)
}

// SePointPosCalculator retrieves the positions and speed for ecliptical or equatorial coordinates.
type SePointPosCalculator interface {
	SeCalcPointPos(jdUt float64, body int, flags int) ([6]float64, error)
}

// SeEpsilonCalculator retrieves the value for the obliquity of the earths axis, either true (corrected for nutation) or mean.
type SeEpsilonCalculator interface {
	SeCalcEpsilon(jdUt float64, trueEps bool) (float64, error)
}

// SeHorPosCalculator retrieves the horizontal positions (azimuth and altitude) from the SE.
type SeHorPosCalculator interface {
	SeCalcHorPos(jdUt float64, geoLong float64, geoLat float64, geoHeight float64, pointRa float64, pointDecl float64, flags int) [3]float64
}

// SeHousePosCalculator retrieves the housepositions and several other mundane points from the SE.
type SeHousePosCalculator interface {
	SeCalcHousePos(houseSys rune, jdUt float64, geoLong float64, geoLat float64, flags int) ([]float64, []float64, error)
}

// SetEphePath initializes the SE and defines the location for the ephemeris files.
func SetEphePath(path string) {
	var _path *C.char = C.CString(path)
	defer C.free(unsafe.Pointer(_path))
	C.swe_set_ephe_path(_path)
}

type SeJulDayCalculation struct{}

func NewSeJulDayCalculation() SeJulDayCalculator {
	return SeJulDayCalculation{}
}

// SeJulDayCalculation accesses the SE to calculate the Julian Day Number, given the values for the date, time and calendar.
func (jdc SeJulDayCalculation) SeCalcJd(year int, month int, day int, hour float64, gregFlag int) float64 {
	cYear := C.int(year)
	cMonth := C.int(month)
	cDay := C.int(day)
	cHour := C.double(hour)
	cGregFlag := C.int(gregFlag)
	result := float64(C.swe_julday(cYear, cMonth, cDay, cHour, cGregFlag))
	return result
}

type SeRevJulDayCalculation struct{}

func NewSeRevJulDayCalculation() SeRevJulDayCalculator {
	return SeRevJulDayCalculation{}
}

// SeRevJulDayCalculation accesses the SE to calculate date and time from a julian day number. The return values are year,month,day and ut.
func (rjdc SeRevJulDayCalculation) SeRevCalcJd(jd float64, gragFlag int) (int, int, int, float64) {
	var cYear C.int
	var cMonth C.int
	var cDay C.int
	var cHour C.double
	C.swe_revjul(C.double(jd), C.int(gragFlag), &cYear, &cMonth, &cDay, &cHour)
	return int(cYear), int(cMonth), int(cDay), float64(cHour)
}

type SePointPosCalculation struct{}

func NewSePointPosCalculation() SePointPosCalculator {
	return SePointPosCalculation{}
}

// SeCalcPointPos accesses the SE to calculate positions for celestial points.
// The results that are returned are subsequently: longitude or ra, latitude or declination, distance, speed in long. or ra, speed in lat. or decl, speed in dist.
func (ppc SePointPosCalculation) SeCalcPointPos(jdUt float64, body int, flags int) ([6]float64, error) {
	var cPos [6]C.double
	cSerr := make([]C.char, C.AS_MAXCH)
	cJdUt := C.double(jdUt)
	cBody := C.int(body)
	cFlags := C.int(flags)
	result := C.swe_calc_ut(cJdUt, cBody, cFlags, &cPos[0], &cSerr[0])
	err := C.GoString(&cSerr[0])
	if result < 0 {
		var emptyArray [6]float64
		return emptyArray, fmt.Errorf("SeCalcPointPos error: %v", err)
	}
	pos := make([]float64, 6)
	for i := 0; i < 6; i++ {
		pos[i] = float64(cPos[i])
	}
	return [6]float64(pos), nil
}

type SeEpsilonCalculation struct{} // TODO create test for SeEpsilonCalculation

func NewSeEpsilonCalculation() SeEpsilonCalculator {
	return SeEpsilonCalculation{}
}

func (ec SeEpsilonCalculation) SeCalcEpsilon(jdUt float64, trueEps bool) (float64, error) {
	var cPos [6]C.double
	cSerr := make([]C.char, C.AS_MAXCH)
	cJdUt := C.double(jdUt)
	cBody := C.int(domain.EclNut)
	cFlags := C.int(domain.SeflgSwieph)
	result := C.swe_calc_ut(cJdUt, cBody, cFlags, &cPos[0], &cSerr[0])
	err := C.GoString(&cSerr[0])
	if result < 0 {
		return 0.0, fmt.Errorf("SeCalcEpsilon error: %v", err)
	}
	if trueEps {
		return float64(cPos[0]), nil
	}
	return float64(cPos[1]), nil
}

type SeHorPosCalculation struct{}

func NewSeHorPosCalculation() SeHorPosCalculator {
	return SeHorPosCalculation{}
}

// SeCalcHorPos converts equatorial coordinates to azimuth, true altitude and apparent altitude. The SE does not return a result code.
func (hpc SeHorPosCalculation) SeCalcHorPos(jdUt float64, geoLong float64, geoLat float64, geoHeight float64, pointRa float64, pointDecl float64, flags int) [3]float64 {
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

type SeHousePosCalculation struct{}

func NewSeHousePosCalculation() SeHousePosCalculator {
	return SeHousePosCalculation{}
}

// SeCalcHousePos calculates mc, asc, and cusps for a given house system, jd, and location.
// Values returned: array with cusps, starting at index 1, array with positions of asc, mc, armc, vertex, eq asc,
// co-asc Koch, co-asc Munkasey and three empty values.
func (hp SeHousePosCalculation) SeCalcHousePos(houseSys rune, jdUt float64, geoLong float64, geoLat float64, flags int) ([]float64, []float64, error) {
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

type SeCoordinateTransformer interface {
	Transform(valuesIn *[3]float64, eps float64, ec2Equ bool) []float64
}

type SeCoordinateTransform struct{}

func NewSeCoordinateTransform() *SeCoordinateTransform {
	return &SeCoordinateTransform{}
}

// Transform ecliptic to equatorial or the other way around. Valuesin and valuesout contain resp. long, lat, distance or ra, decl, distance.
func (ct SeCoordinateTransform) Transform(valuesIn *[3]float64, eps float64, ec2Equ bool) []float64 {
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

/*func Swe_cotrans(xpo *[6]float64, xpn *[6]float64, eps float64) {
	_xpo := (*C.double)(&xpo[0])
	_xpn := (*C.double)(&xpn[0])
	_eps := C.double(eps)

	C.swe_cotrans(_xpo, _xpn, _eps)
}*/
