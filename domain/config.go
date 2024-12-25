/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

type ConfigAspect = struct {
	ActualAspect Aspect
	OrbFactor    float64
	Glyph        rune
}

type ConfigPoint = struct {
	ActualPoint ChartPoint
	OrbFactor   float64
	Glyph       rune
}
