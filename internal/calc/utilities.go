/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	domain "enigma-ar/domain"
	"fmt"
)

// SeFlags calculates the total of all flags for the SE.
func SeFlags(coord domain.CoordinateSystem, obsPos domain.ObserverPosition, tropical bool) int {
	flags := domain.SeflgSwieph + domain.SeflgSpeed // always use SE + speed
	if coord == domain.CoordEquatorial {
		flags += domain.SeflgEquatorial
	}
	if obsPos == domain.ObsPosTopocentric {
		flags += domain.SeflgTopoc
	}
	if obsPos == domain.ObsPosHeliocentric {
		flags += domain.SeflgHelioc
	}
	if coord == domain.CoordEcliptical && !tropical {
		flags += domain.SeflgSidereal
	}
	return flags
}

// ValueToRange converts a value to a limited range, using the difference between the max and the min value as steps to
// increase/decrease the size of the tested value. lowerLImit is inclusive, upperLimit is exclusive
func ValueToRange(testValue, lowerLimit, upperLimit float64) (float64, error) {
	if upperLimit < lowerLimit {
		return 0, fmt.Errorf("upper limit  %f cannot be less than lower limit %f", upperLimit, lowerLimit)
	}
	rangeSize := upperLimit - lowerLimit

	for testValue < lowerLimit {
		testValue += rangeSize
	}
	for testValue >= upperLimit {
		testValue -= rangeSize
	}
	return testValue, nil
}
