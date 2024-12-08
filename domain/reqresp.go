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
	Points   []ChartPoint
	JdUt     float64
	GeoLong  float64
	GeoLat   float64
	Coord    CoordinateSystem
	ObsPos   ObserverPosition
	Tropical bool
}

// PointPosResult Calculated positions for a single point
type PointPosResult struct {
	Point     ChartPoint
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
	Point     ChartPoint
	JdStart   float64
	JdEnd     float64
	Interval  float64
	Coord     CoordinateSystem
	MainValue bool
	Position  bool
	ObsPos    ObserverPosition
	Ayanamsha int
}

// PointRangeResult calculated value for position or speed for a given date/time, to be used in a range of positions.
type PointRangeResult struct {
	Jd    float64
	Value float64
}

// HousePosRequest for the calculation of cusps and other mundane poiints.
type HousePosRequest struct {
	HouseSys HouseSystem
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
	Points    []ChartPoint
	HouseSys  HouseSystem
	Ayanamsha int
	CoordSys  CoordinateSystem
	ObsPos    ObserverPosition
	ProjType  ProjectionType
	Jd        float64
	GeoLong   float64
	GeoLat    float64
}

// FullChartResponse contains the calculated positions for a complete chart. Use housecusps from index 1, zero is an empty placeholder.
type FullChartResponse struct {
	Points    []PointPosResult
	Mc        HousePosResult
	Asc       HousePosResult
	Vertex    HousePosResult
	EastPoint HousePosResult
	Cusps     []HousePosResult
}

// FullChartMeta contains data that is not used by the backend but will be shown in the UI.
type FullChartMeta struct {
	Name         string
	Description  string
	Category     ChartCat
	Rating       Rating
	Source       string
	LocationName string
	GeoLat       string
	GeoLong      string
	Date         string
	Calendar     Calendar
	Time         string
	TimeZone     TimeZone
	Dst          bool
	GeoLongLmt   string
}

type PersistableChart = struct {
	Id          int
	Name        string
	Description string
	Category    string
}

type PersistableDateLocation = struct {
	Id           int
	ChartId      int
	Source       string
	NameLocation string
	Rating       string
	GeoLong      float64
	GeoLat       float64
	DateText     string
	TimeText     string
	Jd           float64
}

// SinglePosition contains a single value for a specific coordinate, and the id for the Chartpoint.
// It's main use is to support calculations like midpoints, harmonics etc.
type SinglePosition = struct {
	Id       ChartPoint
	Position float64
}

// DoublePosition contains two values for a chartPoioint, and the id for that chartpoint.
// It supports combinations like longitude/declination, ra/declination and azimuth/altitude.
type DoublePosition = struct {
	Id        ChartPoint
	Position1 float64
	Position2 float64
}

// MatchedParallel contains two single positions that form a (contra)parallel, an orb and an indication for
// parallel (true) or contraparallel (false).
type MatchedParallel = struct {
	Pos1     SinglePosition
	Pos2     SinglePosition
	Orb      float64
	Parallel bool
}
