/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

// References for calculations

// CoordinateSystem defines the set of coordinates that are used.
type CoordinateSystem int

const (
	Ecliptical CoordinateSystem = iota
	Equatorial
	Horizontal
)

// ObserverPosition defines the central position for the calculations.
type ObserverPosition int

const (
	ObsPosGeocentric ObserverPosition = iota
	ObsPosTopocentric
	ObsPosHeliocentric
)

// Projectiontype, standard 2D, oblique is for oblique longitude as used in the School of Ram.
type ProjectionType int

const (
	ProjType2D ObserverPosition = iota
	ProjTypeOblique
)
