/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

// JulDayRequest Request for the calculation of a Julian Day number
type JulDayRequest struct { // TODO change into general struct for date/time, remove Json
	Year  int     `json:"year"`
	Month int     `json:"month"`
	Day   int     `json:"day"`
	Ut    float64 `json:"ut"`
	Greg  bool    `json:"greg"`
}

// PointPositionsRequest Request for the calculation of all positions for one or more points
type PointPositionsRequest struct {
	Points   []int
	JdUt     float64
	Coord    CoordinateSystem
	ObsPos   ObserverPosition
	Tropical bool
}

// PointPosResult Calculated positions for a single point
type PointPosResult struct {
	Point     int
	LonPos    float64
	LonSpeed  float64
	LatPos    float64
	LatSpeed  float64
	RaPos     float64
	RaSpeed   float64
	DeclPos   float64
	DeclSpeed float64
	RadvPos   float64
	RadvSpeed float64
	AzimPos   float64
	AltitPos  float64
}
