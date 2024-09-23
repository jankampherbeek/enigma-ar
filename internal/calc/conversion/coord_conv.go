/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package conversion

import "enigma-ar/internal/calc/se"

func ChangeEclToEqu(longitude float64, latitude float64, eps float64) (ra float64, decl float64) {
	var coords = [3]float64{longitude, latitude, 1.0}
	ct := se.NewCoordinateTransform()
	result := ct.Transform(coords, eps, true)
	return result[0], result[1]
}
