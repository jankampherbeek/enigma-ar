/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

import "enigma-ar/domain/references"

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
	Points   []references.ChartPoint
	JdUt     float64
	GeoLong  float64
	GeoLat   float64
	Coord    references.CoordinateSystem
	ObsPos   references.ObserverPosition
	Tropical bool
}

// PointPosResult Calculated positions for a single point
type PointPosResult struct {
	Point     references.ChartPoint
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

// PointRangeRequest for the calculation of a range of positions or speeds for a given point.
// The Interval is in days and can be fractional.
// MainValue indicates if longitude or ra is used (true) or latitude or declination.
// It does not have effect if the coord is radv (distance).
// Position indicates that the position is used (true) or the speed (false).
// If the Ayanamsha is zero, a tropical zodiac is used, otherwise a sidereal zodiac with the given ayanamsha.
type PointRangeRequest struct {
	Point     references.ChartPoint
	JdStart   float64
	JdEnd     float64
	Interval  float64
	Coord     references.CoordinateSystem
	MainValue bool
	Position  bool
	ObsPos    references.ObserverPosition
	Ayanamsha int
}

// PointRangeResult calculated value for position or speed for a given date/time, to be used in a range of positions.
type PointRangeResult struct {
	Jd    float64
	Value float64
}

// HousePosRequest for the calculation of cusps and other mundane poiints.
type HousePosRequest struct {
	HouseSys references.HouseSystem
	JdUt     float64
	GeoLong  float64
	GeoLat   float64
}

// HousePosResult Calculated positions for a single cusp or other mundane point.
type HousePosResult struct {
	LonPos   float64
	RaPos    float64
	DeclPos  float64
	AzimPos  float64
	AltitPos float64
}

// FullChartRequest for the calculation of a complete chart with positions of points and mundane positions.
type FullChartRequest struct {
	Points    []references.ChartPoint
	HouseSys  references.HouseSystem
	Ayanamsha int
	CoordSys  references.CoordinateSystem
	ObsPos    references.ObserverPosition
	ProjType  references.ProjectionType
	Jd        float64
	GeoLong   float64
	GeoLat    float64
}

// FullChartResult contains the calculated positions for a complete chart. Use housecusps from index 1, zero is an empty placeholder.
type FullChartResponse struct {
	Points    []PointPosResult
	Mc        HousePosResult
	Asc       HousePosResult
	Vertex    HousePosResult
	EastPoint HousePosResult
	Cusps     []HousePosResult
}
