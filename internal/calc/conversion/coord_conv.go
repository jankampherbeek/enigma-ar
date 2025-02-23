/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package conversion

import "enigma-ar/internal/se"

// TODO add error handling to conversion.ChangeeclToEqu

// ChangeEclToEqu converts ecliptic coordinates to equatorial coordinates.
func ChangeEclToEqu(longitude float64, latitude float64, eps float64) (ra float64, decl float64) {
	var coords = [3]float64{longitude, latitude, 1.0}
	ct := se.NewSwephCoordinateTransform()
	result := ct.Transform(&coords, eps, true)
	return result[0], result[1]
}
