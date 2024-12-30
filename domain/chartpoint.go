/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

type ChartPoint int

const (
	Sun ChartPoint = iota
	Moon
	Mercury
	Venus
	Mars
	Jupiter
	Saturn
	Uranus
	Neptune
	Pluto
	NodeMean
	NodeTrue
	ApogeeMean
	ApogeeCorrected
	Earth
	Chiron
	Pholus
	Ceres
	Pallas
	Juno
	Vesta
	ApogeeInterpolated
	// placeholder for perigee interpolated
	CupidoUra
	HadesUra
	ZeusUra
	KronosUra
	ApollonUra
	AdmetosUra
	VulcanusUra
	PoseidonUra
	Isis

	PersephoneRam
	HermesRam
	DemeterRam
	Eris

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

	ApogeeDuval
	PersephoneCarteret
	VulcanusCarteret
	Ascendant
	Mc
	EastPoint
	Vertex
)

type ChartPointData struct {
	Key       ChartPoint
	TextId    string
	CalcId    int
	CalcCat   CalculationCat
	PointCat  PointCat
	Glyph     rune
	AltGlyphs []rune
}

func AllChartPoints() []ChartPointData {
	return []ChartPointData{
		{Sun, "r_cp_sun", 0, CalcSe, PointCatCommon, '\uE200', []rune{'\uE300'}},
		{Moon, "r_cp_moon", 1, CalcSe, PointCatCommon, '\uE201', []rune{}},
		{Mercury, "r_cp_mercury", 2, CalcSe, PointCatCommon, '\uE202', []rune{'\uE301'}},
		{Venus, "r_cp_venus", 3, CalcSe, PointCatCommon, '\uE203', []rune{}},
		{Mars, "r_cp_mars", 4, CalcSe, PointCatCommon, '\uE205', []rune{'\uE302'}},
		{Jupiter, "r_cp_jupiter", 5, CalcSe, PointCatCommon, '\uE206', []rune{'\uE303'}},
		{Saturn, "r_cp_saturn", 6, CalcSe, PointCatCommon, '\uE207', []rune{'\uE304'}},
		{Uranus, "r_cp_uranus", 7, CalcSe, PointCatCommon, '\uE208', []rune{'\uE305', '\uE306'}},
		{Neptune, "r_cp_neptune", 8, CalcSe, PointCatCommon, '\uE209', []rune{'\uE307'}},
		{Pluto, "r_cp_pluto", 9, CalcSe, PointCatCommon, '\uE210', []rune{'\uE308', '\uE309', '\uE310', '\uE311', '\uE312'}},
		{NodeMean, "r_cp_node_mean", 10, CalcSe, PointCatCommon, '\uE523', []rune{'\uE520'}},
		{NodeTrue, "r_cp_node_true", 11, CalcSe, PointCatCommon, '\uE525', []rune{'\uE520'}},
		{ApogeeMean, "r_cp_apogee_mean", 12, CalcSe, PointCatCommon, '\uE530', []rune{}},
		{ApogeeCorrected, "r_cp_apogee_corrected", 13, CalcSe, PointCatCommon, '\uE531', []rune{}},
		{Earth, "r_cp_earth", 14, CalcSe, PointCatCommon, '\uE204', []rune{}},
		{Chiron, "r_cp_chiron", 15, CalcSe, PointCatCommon, '\uE400', []rune{'\uE450'}},
		{Pholus, "r_cp_pholus", 16, CalcSe, PointCatCommon, '\uE402', []rune{}},
		{Ceres, "r_cp_ceres", 17, CalcSe, PointCatCommon, '\uE411', []rune{}},
		{Pallas, "r_cp_pallas", 18, CalcSe, PointCatCommon, '\uE412', []rune{}},
		{Juno, "r_cp_juno", 19, CalcSe, PointCatCommon, '\uE413', []rune{}},
		{Vesta, "r_cp_vesta", 20, CalcSe, PointCatCommon, '\uE414', []rune{}},
		{ApogeeInterpolated, "r_cp_apogee_interpolated", 21, CalcSe, PointCatCommon, '\uE530', []rune{}},
		// placeholder for periogee interpolated calcid 22
		{CupidoUra, "r_cp_cupido_ura", 40, CalcSe, PointCatCommon, '\uE600', []rune{}},
		{HadesUra, "r_cp_hades_ura", 41, CalcSe, PointCatCommon, '\uE601', []rune{}},
		{ZeusUra, "r_cp_zeus_ura", 42, CalcSe, PointCatCommon, '\uE602', []rune{}},
		{KronosUra, "r_cp_kronos_ura", 43, CalcSe, PointCatCommon, '\uE603', []rune{}},
		{ApollonUra, "r_cp_apollon_ura", 44, CalcSe, PointCatCommon, '\uE604', []rune{}},
		{AdmetosUra, "r_cp_admetos_ura", 45, CalcSe, PointCatCommon, '\uE605', []rune{}},
		{VulcanusUra, "r_cp_vulcanus_ura", 46, CalcSe, PointCatCommon, '\uE606', []rune{}},
		{PoseidonUra, "r_cp_poseidon_ura", 47, CalcSe, PointCatCommon, '\uE607', []rune{}},
		{Isis, "r_cp_isis", 48, CalcSe, PointCatCommon, '\uE611', []rune{}},
		{PersephoneRam, "r_cp_persephone_ram", 2000, CalcElements, PointCatCommon, '\uE608', []rune{}},
		{HermesRam, "r_cp_hermes_ram", 2001, CalcElements, PointCatCommon, '\uE609', []rune{}},
		{DemeterRam, "r_cp_demeter_ram", 2002, CalcElements, PointCatCommon, '\uE610', []rune{}},
		{Eris, "r_cp_eris", 2003, CalcSe, PointCatCommon, '\uE507', []rune{'\uE451', '\uE452', '\uE453', '\uE454', '\uE455', '\uE456'}},
		{Nessus, "r_cp_nessus", 2004, CalcSe, PointCatCommon, '\uE401', []rune{}},
		{Huya, "r_cp_huya", 2005, CalcSe, PointCatCommon, '\uE417', []rune{}},
		{Varuna, "r_cp_varuna", 2006, CalcSe, PointCatCommon, '\uE403', []rune{}},
		{Ixion, "r_cp_ixion", 2007, CalcSe, PointCatCommon, '\uE404', []rune{}},
		{Quaoar, "r_cp_quaoar", 2008, CalcSe, PointCatCommon, '\uE405', []rune{}},
		{Haumea, "r_cp_haumea", 2009, CalcSe, PointCatCommon, '\uE406', []rune{}},
		{Orcus, "r_cp_orcus", 2010, CalcSe, PointCatCommon, '\uE409', []rune{}},
		{Makemake, "r_cp_makemake", 2011, CalcSe, PointCatCommon, '\uE410', []rune{}},
		{Sedna, "r_cp_sedna", 2012, CalcSe, PointCatCommon, '\uE408', []rune{}},
		{Hygieia, "r_cp_hygieia", 2013, CalcSe, PointCatCommon, '\uE415', []rune{'\uE457'}},
		{Astraea, "r_cp_astraea", 2014, CalcSe, PointCatCommon, '\uE416', []rune{}},
		{ApogeeDuval, "r_cp_apogee_duval", 2015, CalcFormula, PointCatCommon, '\uE530', []rune{}},
		{PersephoneCarteret, "r_cp_persephone_carteret", 2016, CalcFormula, PointCatCommon, '\uE612', []rune{}},
		{VulcanusCarteret, "r_cp_vulcanus_carteret", 2017, CalcFormula, PointCatCommon, '\uE613', []rune{}},
		{Ascendant, "r_cp_ascendant", 1001, CalcMundane, PointCatAngle, '\uE600', []rune{'\uE550'}},
		{Mc, "r_cp_mc", 1002, CalcMundane, PointCatAngle, '\uE501', []rune{'\uE551'}},
		{EastPoint, "r_cp_eastpoint", 1003, CalcMundane, PointCatAngle, '\uE502', []rune{}},
		{Vertex, "r_cp_vertex", 1004, CalcMundane, PointCatAngle, '\uE503', []rune{}},
	} // TODO add node south, both mean and true
}
