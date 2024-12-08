/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package conversions

import "math"

func DegToRad(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func RadToDeg(radians float64) float64 {
	return radians * 180.0 / math.Pi
}
