/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

type MpDial int

const (
	Dial360 = iota
	Dial90
	Dial45
	Dial225
)

type MpDialData struct {
	Key      int
	DialSize float64
}

func AllMpDials() []MpDialData {
	return []MpDialData{
		{Dial360, 360},
		{Dial90, 90},
		{Dial45, 45},
		{Dial225, 80},
	}
}

// OccupiedMidpoint contains data for an occupied midpoint including orb and exactness.
type OccupiedMidpoint struct {
	BaseMidpointPos1 SinglePosition
	BaseMidpointPos2 SinglePosition
	FocusPoint       SinglePosition
	ActualOrb        float64
	Exactness        float64
}

// Midpoint containts data for a midpoint without checking if it is occupied.
type Midpoint struct {
	Point1      SinglePosition
	Point2      SinglePosition
	MidpointPos float64
}
