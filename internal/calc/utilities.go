/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	domain "enigma-ar/domain"
	"enigma-ar/domain/references"
)

// SeFlags calculates the total of all flags for the SE.
func SeFlags(coord references.CoordinateSystem, obsPos references.ObserverPosition, tropical bool) int {
	flags := domain.SeflgSwieph + domain.SeflgSpeed // always use SE + speed
	if coord == references.CoordEquatorial {
		flags += domain.SeflgEquatorial
	}
	if obsPos == references.ObsPosTopocentric {
		flags += domain.SeflgTopoc
	}
	if obsPos == references.ObsPosHeliocentric {
		flags += domain.SeflgHelioc
	}
	if coord == references.CoordEcliptical && !tropical {
		flags += domain.SeflgSidereal
	}
	return flags
}
