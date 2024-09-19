/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import "enigma-ar/internal/calc/se"

// CalcJd handles the calculation of a Julian day number.
func CalcJd(year int, month int, day int, ut float64, greg bool) float64 {
	return se.JulDay(year, month, day, ut, greg)
}
