/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package references

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
