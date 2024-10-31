/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package references

type ReferenceText struct {
	Key    int
	TextId string
}

// References for calculations

// CoordinateSystem defines the set of coordinates that are used.
type CoordinateSystem int

const (
	CoordEcliptical = iota
	CoordEquatorial
	CoordHorizontal
)

func AllCoordinateSystems() []ReferenceText {
	return []ReferenceText{
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

func AllObserverPositions() []ReferenceText {
	return []ReferenceText{
		{ObsPosGeocentric, "r_op_geocentric"},
		{ObsPosTopocentric, "r_op_topocentric"},
		{ObsPosHeliocentric, "r_op_heliocentric"},
	}
}

// Projectiontype, standard 2D, oblique is for oblique longitude as used in the School of Ram.
type ProjectionType int

const (
	ProjType2D = iota
	ProjTypeOblique
)

func AllProjectionTypes() []ReferenceText {
	return []ReferenceText{
		{ProjType2D, "r_pt_2d"},
		{ProjTypeOblique, "r_pt_oblique"},
	}
}

type Rating int

const (
	RatingUnknown = iota
	RatingAA
	RatingA
	RatingB
	RatingC
	RatingDD
	RatingX
	RatingXX
)

func AllRatings() []ReferenceText {
	return []ReferenceText{
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
	CatUnknown = iota
	CatFemale
	CatMale
	CatEvent
	CatHorary
	CatElection
	CatOther
)

func AllChartCats() []ReferenceText {
	return []ReferenceText{
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
	CalGregorian = iota
	CalJulianCE
	CalJulianBCE
	CalAstronomical
)

func AllCalendars() []ReferenceText {
	return []ReferenceText{
		{CalGregorian, "r_cal_gregorian"},
		{CalJulianCE, "r_cal_julian_ce"},
		{CalJulianBCE, "r_cal_julian_bce"},
		{CalAstronomical, "r_cal_astronomical"},
	}
}

type CalculationCat int

const (
	CalcSe = iota
	CalcElements
	CalcFormula
	CalcMundane
	CalcLots
	CalcZodiacFixed
)

type PointCats int

const (
	PointCatCommon = iota
	PointCatAngle
	PointCatCusp
	PointCatZodiac
	PointCatLot
	PointCatFixStar
)
