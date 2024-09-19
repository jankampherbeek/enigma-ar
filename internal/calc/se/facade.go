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
