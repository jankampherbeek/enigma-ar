/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import domain2 "enigma-ar/domain"

// SeFlags calculates the total of all flags for the SE.
func SeFlags(coord domain2.CoordinateSystem, obsPos domain2.ObserverPosition, tropical bool) int {
	flags := domain2.SeflgSwieph + domain2.SeflgSpeed // always use SE + speed
	if coord == domain2.Equatorial {
		flags += domain2.SeflgEquatorial
	}
	if obsPos == domain2.ObsPosTopocentric {
		flags += domain2.SeflgTopoc
	}
	if obsPos == domain2.ObsPosHeliocentric {
		flags += domain2.SeflgHelioc
	}
	if coord == domain2.Ecliptical && !tropical {
		flags += domain2.SeflgSidereal
	}
	return flags
}
