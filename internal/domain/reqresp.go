/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

// DateTime Representation of a date and time in UT
type DateTime struct {
	Year  int
	Month int
	Day   int
	Ut    float64
	Greg  bool
}

// PointPositionsRequest Request for the calculation of all positions for one or more points
type PointPositionsRequest struct {
	Points   []int
	JdUt     float64
	GeoLong  float64
	GeoLat   float64
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

// PointRangeResult calculated value for position or speed for a given date/time, to be used in a range of positions.
type PointRangeResult struct {
	Jd    float64
	Value float64
}

// PointRangeRequest for the calculation of a range of positions or speeds for a given point.
// The Interval is in days and can be fractional.
// MainValue indicates if longitude or ra is used (true) or latitude or declination.
// It does not have effect if the coord is radv (distance).
// Position indicates that the position is used (true) or the speed (false).
// If the Ayanamsha is zero, a tropical zodiac is used, otherwise a sidereal zodiac with the given ayanamsha.
type PointRangeRequest struct {
	Point     int
	JdStart   float64
	JdEnd     float64
	Interval  float64
	Coord     CoordinateSystem
	MainValue bool
	Position  bool
	ObsPos    ObserverPosition
	Ayanamsha int
}

// HousePosResult Calculated positions for a cusp or other mundane point
type HousePosResult struct {
	LonPos   float64
	RaPos    float64
	DeclPos  float64
	AzimPos  float64
	AltitPos float64
}
