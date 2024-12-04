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
	CoordEcliptical CoordinateSystem = iota
	CoordEquatorial
	CoordHorizontal
)

type CoordinateSystemText struct {
	Key    CoordinateSystem
	TextId string
}

func AllCoordinateSystems() []CoordinateSystemText {
	return []CoordinateSystemText{
		{CoordEcliptical, "r_cs_ecliptical"},
		{CoordEquatorial, "r_cs_equatorial"},
		{CoordHorizontal, "r_cs_horizontal"},
	}
}

// ObserverPosition defines the central position for the calculations.
type ObserverPosition int

const (
	ObsPosGeocentric = iota
	ObsPosTopocentric
	ObsPosHeliocentric
)

type ObserverPosText struct {
	Key    ObserverPosition
	TextId string
}

func AllObserverPositions() []ObserverPosText {
	return []ObserverPosText{
		{ObsPosGeocentric, "r_op_geocentric"},
		{ObsPosTopocentric, "r_op_topocentric"},
		{ObsPosHeliocentric, "r_op_heliocentric"},
	}
}

// Projectiontype, standard 2D, oblique is for oblique longitude as used in the School of Ram.
type ProjectionType int

const (
	ProjType2D ProjectionType = iota
	ProjTypeOblique
)

type ProjectionTypeText struct {
	Key    ProjectionType
	TextId string
}

func AllProjectionTypes() []ProjectionTypeText {
	return []ProjectionTypeText{
		{ProjType2D, "r_pt_2d"},
		{ProjTypeOblique, "r_pt_oblique"},
	}
}

type Rating int

const (
	RatingUnknown Rating = iota
	RatingAA
	RatingA
	RatingB
	RatingC
	RatingDD
	RatingX
	RatingXX
)

type RatingText struct {
	Key    Rating
	TextId string
}

func AllRatings() []RatingText {
	return []RatingText{
		{RatingUnknown, "r_rr_unknown"},
		{RatingAA, "r_rr_aa"},
		{RatingA, "r_rr_a"},
		{RatingB, "r_rr_b"},
		{RatingC, "r_rr_c"},
		{RatingDD, "r_rr_dd"},
		{RatingX, "r_rr_x"},
		{RatingXX, "r_rr_xx"},
	}
}

type ChartCat int

const (
	CatUnknown ChartCat = iota
	CatFemale
	CatMale
	CatEvent
	CatHorary
	CatElection
	CatOther
)

type ChartCatText struct {
	Key    ChartCat
	TextId string
}

func AllChartCats() []ChartCatText {
	return []ChartCatText{
		{CatUnknown, "r_cc_unknown"},
		{CatFemale, "r_cc_female"},
		{CatMale, "r_cc_male"},
		{CatEvent, "r_cc_event"},
		{CatHorary, "r_cc_horary"},
		{CatElection, "r_cc_election"},
		{CatOther, "r_cc_other"},
	}
}

type Calendar int

const (
	CalGregorian Calendar = iota
	CalJulianCE
	CalJulianBCE
	CalAstronomical
)

type CalendarText struct {
	Key    Calendar
	TextId string
}

func AllCalendars() []CalendarText {
	return []CalendarText{
		{CalGregorian, "r_cal_gregorian"},
		{CalJulianCE, "r_cal_julian_ce"},
		{CalJulianBCE, "r_cal_julian_bce"},
		{CalAstronomical, "r_cal_astronomical"},
	}
}

type CalculationCat int

const (
	CalcSe CalculationCat = iota
	CalcElements
	CalcFormula
	CalcMundane
	CalcLots
	CalcZodiacFixed
)

type PointCat int

const (
	PointCatCommon PointCat = iota
	PointCatAngle
	PointCatCusp
	PointCatZodiac
	PointCatLot
	PointCatFixStar
)

type WheelType int

const (
	WheelTypeSignsEqual WheelType = iota
	WheelTypeHousesEqual
	WheelTypePlanetsOutside
	WheelTypeSimpleCircle
)

type WheelTypeText struct {
	Key    WheelType
	TextId string
}

func AllWheelTypes() []WheelTypeText {
	return []WheelTypeText{
		{WheelTypeSignsEqual, "r_wh_signs_equal"},
		{WheelTypeHousesEqual, "r_wh_houses_equal"},
		{WheelTypePlanetsOutside, "r_wh_planets_outside"},
		{WheelTypeSimpleCircle, "r_wh_simple_circle"},
	}

}
