/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import domain "enigma-ar/domain"

// SeFlags calculates the total of all flags for the SE.
func SeFlags(coord domain.CoordinateSystem, obsPos domain.ObserverPosition, tropical bool) int {
	flags := domain.SeflgSwieph + domain.SeflgSpeed // always use SE + speed
	if coord == domain.Equatorial {
		flags += domain.SeflgEquatorial
	}
	if obsPos == domain.ObsPosTopocentric {
		flags += domain.SeflgTopoc
	}
	if obsPos == domain.ObsPosHeliocentric {
		flags += domain.SeflgHelioc
	}
	if coord == domain.Ecliptical && !tropical {
		flags += domain.SeflgSidereal
	}
	return flags
}
