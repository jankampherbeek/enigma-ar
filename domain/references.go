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

type HouseSystem int

const (
	HousesNone = iota
	HousesPlacidus
	HousesKoch
	HousesPorphyri
	HousesRegiomontanus
	HousesCampanus
	HousesAlcabitius
	HousesTopocentric
	HousesKrusinski
	HousesApc
	HousesMorin
	HousesWholeSign
	HousesEqualAsc
	HousesEqualMc
	HousesEqualAries
	HousesVehlow
	HousesAxial
	HousesHorizon
	HousesCarter
	HousesGauquelin
	HousesSunShine
	HousesSunShineTreindl
	HousesPullenSd
	HousesPullenSr
	HousesSripati
)

type HouseSystemData struct {
	TextId           string
	SeSupported      bool
	Code             rune
	Number           int
	CounterClockWise bool
	Quadrant         bool
}

func AllHouseSystems() []HouseSystemData {
	return []HouseSystemData{
		{"r_hs_none", false, 'W', 0, false, false},
		{"r_hs_placidus", true, 'P', 12, true, true},
		{"r_hs_koch", true, 'K', 12, true, true},
		{"r_hs_porphyri", true, 'O', 12, true, true},
		{"r_hs_regiomontanus", true, 'R', 12, true, true},
		{"r_hs_campanus", true, 'C', 12, true, true},
		{"r_hs_alcabitius", true, 'B', 12, true, true},
		{"r_hs_topocentric", true, 'T', 12, true, true},
		{"r_hs_krusinski", true, 'U', 12, true, true},
		{"r_hs_apc", true, 'Y', 12, true, true},
		{"r_hs_morin", true, 'M', 12, true, false},
		{"r_hs_wholesign", true, 'W', 12, true, true},
		{"r_hs_equalasc", true, 'A', 12, true, false},
		{"r_hs_equalmc", true, 'D', 12, true, false},
		{"r_hs_equalaries", true, 'N', 12, true, false},
		{"r_hs_vehlow", true, 'V', 12, true, false},
		{"r_hs_axial", true, 'X', 12, true, false},
		{"r_hs_horizon", true, 'H', 12, true, false},
		{"r_hs_carter", true, 'F', 12, true, false},
		{"r_hs_gauquelin", true, 'G', 36, true, false},
		{"r_hs_sunshine", true, 'i', 12, true, false},
		{"r_hs_sunshine_treindl", true, 'I', 12, true, false},
		{"r_hs_pullen_sd", true, 'L', 12, true, true},
		{"r_hs_pullen_sr", true, 'Q', 12, true, true},
		{"r_hs_sripati", true, 'S', 12, true, false},
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

type ChartPoint int

const (
	Sun = iota
	Moon
	Mercury
	Venus
	Earth
	Mars
	Jupiter
	Saturn
	Uranus
	Neptune
	Pluto
	MeanNode
	TrueNode
	Chiron
	PersephoneRam
	HermesRam
	DemeterRam
	CupidoUra
	HadesUra
	ZeusUra
	KronosUra
	ApollonUra
	AdmetosUra
	VulcanusUra
	PoseidonUra
	Eris
	Pholus
	Ceres
	Pallas
	Juno
	Vesta
	Isis
	Nessus
	Huya
	Varuna
	Ixion
	Quaoar
	Haumea
	Orcus
	Makemake
	Sedna
	Hygieia
	Astraea
	ApogeeMean
	ApogeeCorrected
	ApogeeInterpolated
	ApogeeDuval
	PersephoneCarteret
	VulcanusCarteret
	Ascendant
	Mc
	EastPoint
	Vertex
)

type ChartPointData struct {
	TextId         string
	CalcId         int
	CalculationCat CalculationCat
	PointCat       PointCats
}

func AllChartPoints() []ChartPointData {
	return []ChartPointData{
		{"r_cp_sun", 0, CalcSe, PointCatCommon},
		{"r_cp_moon", 1, CalcSe, PointCatCommon},
		{"r_cp_mercury", 2, CalcSe, PointCatCommon},
		{"r_cp_venus", 3, CalcSe, PointCatCommon},
		{"r_cp_earth", 4, CalcSe, PointCatCommon},
		{"r_cp_mars", 5, CalcSe, PointCatCommon},
		{"r_cp_jupiter", 5, CalcSe, PointCatCommon},
		{"r_cp_saturn", 7, CalcSe, PointCatCommon},
		{"r_cp_uranus", 8, CalcSe, PointCatCommon},
		{"r_cp_neptune", 9, CalcSe, PointCatCommon},
		{"r_cp_pluto", 10, CalcSe, PointCatCommon},
		{"r_cp_node_mean", 11, CalcSe, PointCatCommon},
		{"r_cp_node_true", 12, CalcSe, PointCatCommon},
		{"r_cp_chiron", 13, CalcSe, PointCatCommon},
		{"r_cp_persephone_ram", 14, CalcElements, PointCatCommon},
		{"r_cp_hermes_ram", 15, CalcElements, PointCatCommon},
		{"r_cp_demeter_ram", 16, CalcElements, PointCatCommon},
		{"r_cp_cupido_ura", 17, CalcSe, PointCatCommon},
		{"r_cp_hades_ura", 18, CalcSe, PointCatCommon},
		{"r_cp_zeus_ura", 19, CalcSe, PointCatCommon},
		{"r_cp_kronos_ura", 20, CalcSe, PointCatCommon},
		{"r_cp_apollon_ura", 21, CalcSe, PointCatCommon},
		{"r_cp_admetos_ura", 22, CalcSe, PointCatCommon},
		{"r_cp_vulcanus_ura", 23, CalcSe, PointCatCommon},
		{"r_cp_poseidon_ura", 24, CalcSe, PointCatCommon},
		{"r_cp_eris", 25, CalcSe, PointCatCommon},
		{"r_cp_pholus", 26, CalcSe, PointCatCommon},
		{"r_cp_ceres", 27, CalcSe, PointCatCommon},
		{"r_cp_pallas", 28, CalcSe, PointCatCommon},
		{"r_cp_juno", 29, CalcSe, PointCatCommon},
		{"r_cp_vesta", 30, CalcSe, PointCatCommon},
		{"r_cp_isis", 31, CalcSe, PointCatCommon},
		{"r_cp_nessus", 32, CalcSe, PointCatCommon},
		{"r_cp_huya", 33, CalcSe, PointCatCommon},
		{"r_cp_varuna", 34, CalcSe, PointCatCommon},
		{"r_cp_ixion", 35, CalcSe, PointCatCommon},
		{"r_cp_quaoar", 36, CalcSe, PointCatCommon},
		{"r_cp_haumea", 37, CalcSe, PointCatCommon},
		{"r_cp_orcus", 38, CalcSe, PointCatCommon},
		{"r_cp_makemake", 39, CalcSe, PointCatCommon},
		{"r_cp_sedna", 40, CalcSe, PointCatCommon},
		{"r_cp_hygieia", 41, CalcSe, PointCatCommon},
		{"r_cp_astraea", 42, CalcSe, PointCatCommon},
		{"r_cp_apogee_mean", 43, CalcSe, PointCatCommon},
		{"r_cp_apogee_corrected", 44, CalcSe, PointCatCommon},
		{"r_cp_apogee_interpolated", 45, CalcSe, PointCatCommon},
		{"r_cp_apogee_duval", 46, CalcFormula, PointCatCommon},
		{"r_cp_persephone_carteret", 47, CalcFormula, PointCatCommon},
		{"r_cp_vulcanus_carteret", 48, CalcFormula, PointCatCommon},
		{"r_cp_ascendant", 1001, CalcMundane, PointCatAngle},
		{"r_cp_mc", 1002, CalcMundane, PointCatAngle},
		{"r_cp_eastpoint", 1003, CalcMundane, PointCatAngle},
		{"r_cp_vertex", 1004, CalcMundane, PointCatAngle},
	}
}

type Ayanamsha int

const (
	AyanNone = iota
	AyanFagan
	AyanLahiri
	AyanDeLuce
	AyanRaman
	AyanUshaShashi
	AyanKrishnamurti
	AyanDjwhalKhul
	AyanYukteshwar
	AyanBhasin
	AyanKugler1
	AyanKugler2
	AyanKugler3
	AyanHuber
	AyanEtaPiscium
	AyanAldebaran15Tau
	AyanHipparchus
	AyanSassanian
	AyanGalactCtr0Sag
	AyanJ2000
	AyanJ1900
	AyanB1950
	AyanSuryaSiddhanta
	AyanSuryaSiddhantaMeanSun
	AyanAryabhata
	AyanAryabhataMeanSun
	AyanSsRevati
	AyanSsCitra
	AyanTrueCitra
	AyanTrueRevati
	AyanTruePushya
	AyanGalacticCtrBrand
	AyanGalacticEqIau1958
	AyanGalacticEq
	AyanGalacticEqMidMula
	AyanSkydram
	AyanTrueMula
	AyanDhruva
	AyanAryabhata522
	AyanBritton
	AyanGalacticCtrOCap = 40
)

type AyanamshaData struct {
	Key    int
	TextId string
	CalcId int
}

func AllAyanamshas() []AyanamshaData {
	return []AyanamshaData{
		{AyanNone, "r_ay_none", -1},
		{AyanFagan, "r_ay_fagan", 0},
		{AyanLahiri, "r_ay_lahiri", 1},
		{AyanDeLuce, "r_ay_deluce", 2},
		{AyanRaman, "r_ay_raman", 3},
		{AyanUshaShashi, "r_ay_shashi", 4},
		{AyanKrishnamurti, "r_ay_krishnamurti", 5},
		{AyanDjwhalKhul, "r_ay_djwhalkhul", 6},
		{AyanYukteshwar, "r_ay_yukteshwar", 7},
		{AyanBhasin, "r_ay_bhasin", 8},
		{AyanKugler1, "r_ay_kugler1", 9},
		{AyanKugler2, "r_ay_kugler2", 10},
		{AyanKugler3, "r_ay_kugler3", 11},
		{AyanHuber, "r_ay_huber", 12},
		{AyanEtaPiscium, "r_ay_etapiscium", 13},
		{AyanAldebaran15Tau, "r_ay_aldebaran15tau", 14},
		{AyanHipparchus, "r_ay_hipparchus", 15},
		{AyanSassanian, "r_ay_sassanian", 16},
		{AyanGalactCtr0Sag, "r_ay_galactctr0sag", 17},
		{AyanJ2000, "r_ay_j2000", 18},
		{AyanJ1900, "r_ay_j2000", 19},
		{AyanB1950, "r_ay_b1950", 20},
		{AyanSuryaSiddhanta, "r_ay_surya_siddhanta", 21},
		{AyanSuryaSiddhantaMeanSun, "r_ay_surya_siddhantameansun", 22},
		{AyanAryabhata, "r_ay_aryabhata", 23},
		{AyanAryabhataMeanSun, "r_ay_aryabhatameansun", 24},
		{AyanSsRevati, "r_ay_ssrevati", 25},
		{AyanSsCitra, "r_ay_sscitra", 26},
		{AyanTrueCitra, "r_ay_truecitra", 27},
		{AyanTrueRevati, "r_ay_truerevati", 28},
		{AyanTruePushya, "r_ay_truepushya", 29},
		{AyanGalacticCtrBrand, "r_ay_galactctr_brand", 30},
		{AyanGalacticEqIau1958, "r_ay_galactctr_eqiau1958", 31},
		{AyanGalacticEq, "r_ay_galactic_eq", 32},
		{AyanGalacticEqMidMula, "r_ay_galacticeqmidmula", 33},
		{AyanSkydram, "r_ay_skydram", 34},
		{AyanTrueMula, "r_ay_truemula", 35},
		{AyanDhruva, "r_ay_dhruva", 36},
		{AyanAryabhata522, "r_ay_aryabhata", 37},
		{AyanBritton, "r_ay_britton", 38},
		{AyanGalacticCtrOCap, "r_ay_galacticctr0cap", 39},
	}
}
