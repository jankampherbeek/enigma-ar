/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package conversion

import (
	"enigma-ar/internal/calc/mathextra"
	"math"
)

// DeclinationToLongitude converts declination to longitude
func DeclinationToLongitude(obliquity, declination float64) float64 {
	sinDeclRad := math.Sin(mathextra.DegToRad(declination))
	sinOblRad := math.Sin(mathextra.DegToRad(obliquity))
	result := mathextra.RadToDeg(math.Asin(sinDeclRad / sinOblRad))
	if result > 360.0 {
		result -= 360.0
	}
	if result < 0.0 {
		result += 360.0
	}
	return result
}
