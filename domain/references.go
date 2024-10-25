/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

type ReferenceText struct {
	Key    int
	TextId string
}

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
	ProjType2D ProjectionType = iota
	ProjTypeOblique
)

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

type TimeZone int

const (
	TzUt TimeZone = iota
	TzCet
	TzEet
	TzEat
	TzIrst
	TzAmt
	TzAft
	TzPkt
	TzIst
	TzIot
	TzMmt
	IzIct
	TzWst
	TzJst
	TzAcst
	TzAest
	TzLhst
	TzNct
	TzNzst
	TzSst
	TzHast
	TzMart
	TzAkst
	TzPst
	TzMst
	TzCst
	TzEst
	TzAst
	TzNst
	TzBrt
	TzGst
	TzAzot
	TzLmt
)

type TimeZoneData struct {
	TextId string
	Offset float64
}

func AllTimeZones() []TimeZoneData {
	return []TimeZoneData{
		{"r_tz_ut", 0.0},
		{"r_tz_cet", 1.0},
		{"r_tz_eet", 2.0},
		{"r_tz_eat", 3.0},
		{"r_tz_irst", 3.0},
		{"r_tz_amt", 4.0},
		{"r_tz_aft", 4.0},
		{"r_tz_pkt", 5.0},
		{"r_tz_ist", 5.0},
		{"r_tz_iot", 6.0},
		{"r_tz_mmt", 6.0},
		{"r_tz_ict", 7.0},
		{"r_tz_wst", 8.0},
		{"r_tz_jst", 9.0},
		{"r_tz_acst", 9.0},
		{"r_tz_aest", 10.0},
		{"r_tz_lhst", 10.0},
		{"r_tz_nct", 11.0},
		{"r_tz_nzst", 12.0},
		{"r_tz_sst", -11.0},
		{"r_tz_hast", -10.0},
		{"r_tz_mart", -9.0},
		{"r_tz_akst", -9.0},
		{"r_tz_pst", -8.0},
		{"r_tz_mst", -7.0},
		{"r_tz_cst", -6.0},
		{"r_tz_est", -5.0},
		{"r_tz_ast", -4.0},
		{"r_tz_nst", -3.0},
		{"r_tz_brt", -3.0},
		{"r_tz_gst", -2.0},
		{"r_tz_azot", -1.0},
		{"r_tz_lmt", 0.0},
	}

}
