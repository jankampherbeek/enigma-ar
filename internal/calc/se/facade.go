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
	"fmt"
	"unsafe"
)

// SetEphePath initializes the SE and defines the location for the ephemeris files.
func SetEphePath(path string) {
	var _path *C.char = C.CString(path)
	defer C.free(unsafe.Pointer(_path))
	C.swe_set_ephe_path(_path)
}

// JulDay accesses the SE to calculate the Julian Day Number, given the values for the date, time and calendar.
func JulDay(year int, month int, day int, hour float64, gregCal bool) float64 {
	cYear := C.int(year)
	cMonth := C.int(month)
	cDay := C.int(day)
	cHour := C.double(hour)
	var gregFlag int32 = 1
	if !gregCal {
		gregFlag = 0
	}
	cGregFlag := C.int(gregFlag)
	result := float64(C.swe_julday(cYear, cMonth, cDay, cHour, cGregFlag))
	return result
}

// PointPositions accesses the SE to calculate positions for celestial points
func PointPositions(jdUt float64, body int, flags int) ([6]float64, error) {
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