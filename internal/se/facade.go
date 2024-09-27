/*
 * Enigma Astrology Research.
 * Copyright (c) Jan Kampherbeek.
 * Enigma is open source.
 * Please check the file copyright.txt in the root of the source for further details.
 */

package se

/*
#include "swephexp.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"math"
	"unsafe"
)

type SeJulDayCalculator interface {
	SeCalcJd(year int, month int, day int, hour float64, gregFlag int) float64
}

type SePointPosCalculator interface {
	SeCalcPointPos(jdUt float64, body int, flags int) ([6]float64, error)
}

type SeHorizontalPosCalculator interface {
	CalcHorPos(jdUt float64, geoLong float64, geoLat float64, geoHeight float64, pointRa float64, pointDecl float64, flags int) [3]float64
}

type SeHousePosCalculator interface {
	CalcHousePos(houseSys rune, jdUt float64, geoLat float64, geoLong float64, flags int32) ([]float64, []float64, error)
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

// CalculateJd accesses the SE to calculate the Julian Day Number, given the values for the date, time and calendar.
func (jdc SeJulDayCalculation) SeCalcJd(year int, month int, day int, hour float64, gregFlag int) float64 {
	cYear := C.int(year)
	cMonth := C.int(month)
	cDay := C.int(day)
	cHour := C.double(hour)
	cGregFlag := C.int(gregFlag)
	result := float64(C.swe_julday(cYear, cMonth, cDay, cHour, cGregFlag))
	return result
}

type SePointPosCalculation struct{}

func NewSePointPosCalculation() SePointPosCalculator {
	return SePointPosCalculation{}
}

// CalculatePointPos accesses the SE to calculate positions for celestial points.
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
		return emptyArray, fmt.Errorf("PointPositions error: %s", err)
	}
	pos := make([]float64, 6)
	for i := 0; i < 6; i++ {
		pos[i] = float64(cPos[i])
	}
	return [6]float64(pos), nil
}

type SeHorPosCalculation struct{}

func NewSeHorPosCalculation() SeHorPosCalculation {
	return SeHorPosCalculation{}
}

// CalculateHorPos converts equatorial coordinates to azimuth, true altitude and apparent altitude. The SE does not return a result code.
func (hpc SeHorPosCalculation) CalcHorPos(jdUt float64, geoLong float64, geoLat float64, geoHeight float64, pointRa float64, pointDecl float64, flags int) [3]float64 {
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

type SeHousePos struct{}

func NewSeHousePos() *SeHousePos {
	return &SeHousePos{}
}

// CalcHousePos calculates mc, asc, and cusps for a given house system, jd, and location.
// Depending on the value of flags, tropical (0) positions or sidereal (65536) positions are returned.
// Values returned: array with cusps, starting at index 1, array with positions of asc, mc, armc, vertex, eq asc,
// co-asc Koch, co-asc Munkasey and three empty values.
func (hp *SeHousePos) CalcHousePos(houseSys rune, jdUt float64, geoLong float64, geoLat float64, flags int) ([]float64, []float64, error) {
	cJdUt := C.double(jdUt)
	cGeolat := C.double(geoLat)
	cGeolong := C.double(geoLong)
	cHouseSys := C.int(houseSys)
	cFlags := C.int(flags)

	var cCusps [13]C.double
	var cAscMc [10]C.double

	result := C.swe_houses_ex(cJdUt, cFlags, cGeolat, cGeolong, cHouseSys, &cCusps[0], &cAscMc[0])
	if result < 0 {
		fmt.Printf("Error in HousePositions: %d", result)
		err := errors.New("Error in HousePositions")
		return make([]float64, 13), make([]float64, 10), err
	}
	cusps := make([]float64, 13)
	for i := 0; i < int(13); i++ {
		cusps[i] = float64(cCusps[i])
	}
	ascMc := make([]float64, 10)
	for i := 0; i < int(10); i++ {
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
	for i := 0; i < int(3); i++ {
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
