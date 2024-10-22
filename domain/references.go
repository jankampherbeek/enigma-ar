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
	RatingUnknown Rating = iota
	RatingAA
	RatingA
	RatingB
	RatingC
	RatingDD
	RatingX
	RatingXX
)

func (r Rating) MetaData() string {
	switch r {
	case RatingAA:
		return "r_rr_aa"
	case RatingA:
		return "r_rr_a"
	case RatingB:
		return "r_rr_b"
	case RatingC:
		return "r_rr_c"
	case RatingDD:
		return "r_rr_dd"
	case RatingX:
		return "r_rr_x"
	case RatingXX:
		return "r_rr_xx"
	case RatingUnknown:
		return "r_rr_unknown"
	}
	return ""
}

type ChartCat int

const (
	CatFemale ChartCat = iota
	CatMale
	CatEvent
	CatHorary
	CatElection
	CatOther
	CatUnknown
)

func (cc ChartCat) MetaData() string {
	switch cc {
	case CatFemale:
		return "r_cc_female"
	case CatMale:
		return "r_cc_male"
	case CatEvent:
		return "r_cc_event"
	case CatHorary:
		return "r_cc_horary"
	case CatElection:
		return "r_cc_election"
	case CatOther:
		return "r_cc_other"
	case CatUnknown:
		return "r_cc_unknown"
	}
	return "r_cc_unknown"
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

func (tz TimeZone) MetaData() (string, float64) {
	switch tz {
	case TzUt:
		return "r_tz_ut", 0.0
	case TzCet:
		return "r_tz_cet", 1.0
	case TzEet:
		return "r_tz_eet", 2.0
	case TzEat:
		return "r_tz_eat", 3.0
	case TzIrst:
		return "r_tz_irst", 3.5
	case TzAmt:
		return "r_tz_amt", 4.0
	case TzAft:
		return "r_tz_aft", 4.5
	case TzPkt:
		return "r_tz_pkt", 5.0
	case TzIst:
		return "r_tz_ist", 5.5
	case TzIot:
		return "r_tz_iot", 6.0
	case TzMmt:
		return "r_tz_mmt", 6.5
	case IzIct:
		return "r_tz_ict", 7.0
	case TzWst:
		return "r_tz_wst", 8.0
	case TzJst:
		return "r_tz_jst", 9.0
	case TzAcst:
		return "r_tz_acst", 9.5
	case TzAest:
		return "r_tz_aest", 10.0
	case TzLhst:
		return "r_tz_lhst", 10.5
	case TzNct:
		return "r_tz_nct", 11.0
	case TzNzst:
		return "r_tz_nzst", 12.0
	case TzSst:
		return "r_tz_sst", -11.0
	case TzHast:
		return "r_tz_hast", -10.0
	case TzMart:
		return "r_tz_mart", -9.5
	case TzAkst:
		return "r_tz_akst", -9.0
	case TzPst:
		return "r_tz_pst", -8.0
	case TzMst:
		return "r_tz_mst", -7.0
	case TzCst:
		return "r_tz_cst", -6.0
	case TzEst:
		return "r_tz_est", -5.0
	case TzAst:
		return "r_tz_ast", -4.0
	case TzNst:
		return "r_tz_nst", -3.5
	case TzBrt:
		return "r_tz_brt", -3.0
	case TzGst:
		return "r_tz_gst", -2.0
	case TzAzot:
		return "r_tz_azot", -1.0
	case TzLmt:
		return "r_tz_lmt", 0.0

	}
	return "", 0.0
}
