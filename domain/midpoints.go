/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

type OccupiedMidpoint struct {
	BaseMidpointPos1 SinglePosition
	BaseMidpointPos2 SinglePosition
	FocusPoint       SinglePosition
	ActualOrb        float64
	Exactness        float64
}
