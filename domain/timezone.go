/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

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
	TzIct
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
	Key    TimeZone
	TextId string
	Offset float64
}

func AllTimeZones() []TimeZoneData {
	return []TimeZoneData{
		{TzUt, "r_tz_ut", 0.0},
		{TzCet, "r_tz_cet", 1.0},
		{TzEet, "r_tz_eet", 2.0},
		{TzEat, "r_tz_eat", 3.0},
		{TzIrst, "r_tz_irst", 3.0},
		{TzAmt, "r_tz_amt", 4.0},
		{TzAft, "r_tz_aft", 4.0},
		{TzPkt, "r_tz_pkt", 5.0},
		{TzIst, "r_tz_ist", 5.0},
		{TzIot, "r_tz_iot", 6.0},
		{TzMmt, "r_tz_mmt", 6.0},
		{TzIct, "r_tz_ict", 7.0},
		{TzWst, "r_tz_wst", 8.0},
		{TzJst, "r_tz_jst", 9.0},
		{TzAcst, "r_tz_acst", 9.0},
		{TzAest, "r_tz_aest", 10.0},
		{TzLhst, "r_tz_lhst", 10.0},
		{TzNct, "r_tz_nct", 11.0},
		{TzNzst, "r_tz_nzst", 12.0},
		{TzSst, "r_tz_sst", -11.0},
		{TzHast, "r_tz_hast", -10.0},
		{TzMart, "r_tz_mart", -9.0},
		{TzAkst, "r_tz_akst", -9.0},
		{TzPst, "r_tz_pst", -8.0},
		{TzMst, "r_tz_mst", -7.0},
		{TzCst, "r_tz_cst", -6.0},
		{TzEst, "r_tz_est", -5.0},
		{TzAst, "r_tz_ast", -4.0},
		{TzNst, "r_tz_nst", -3.0},
		{TzBrt, "r_tz_brt", -3.0},
		{TzGst, "r_tz_gst", -2.0},
		{TzAzot, "r_tz_azot", -1.0},
		{TzLmt, "r_tz_lmt", 0.0},
	}
}
