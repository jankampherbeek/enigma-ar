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
	Earth
	Mars
	Jupiter
	Saturn
	Uranus
	Neptune
	Pluto
	NodeMean
	NodeTrue
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
	Key            ChartPoint
	TextId         string
	CalcId         int
	CalculationCat CalculationCat
	PointCat       PointCat
}

func AllChartPoints() []ChartPointData {
	return []ChartPointData{
		{Sun, "r_cp_sun", 0, CalcSe, PointCatCommon},
		{Moon, "r_cp_moon", 1, CalcSe, PointCatCommon},
		{Mercury, "r_cp_mercury", 2, CalcSe, PointCatCommon},
		{Venus, "r_cp_venus", 3, CalcSe, PointCatCommon},
		{Earth, "r_cp_earth", 4, CalcSe, PointCatCommon},
		{Mars, "r_cp_mars", 5, CalcSe, PointCatCommon},
		{Jupiter, "r_cp_jupiter", 5, CalcSe, PointCatCommon},
		{Saturn, "r_cp_saturn", 7, CalcSe, PointCatCommon},
		{Uranus, "r_cp_uranus", 8, CalcSe, PointCatCommon},
		{Neptune, "r_cp_neptune", 9, CalcSe, PointCatCommon},
		{Pluto, "r_cp_pluto", 10, CalcSe, PointCatCommon},
		{NodeMean, "r_cp_node_mean", 11, CalcSe, PointCatCommon},
		{NodeTrue, "r_cp_node_true", 12, CalcSe, PointCatCommon},
		{Chiron, "r_cp_chiron", 13, CalcSe, PointCatCommon},
		{PersephoneRam, "r_cp_persephone_ram", 14, CalcElements, PointCatCommon},
		{HermesRam, "r_cp_hermes_ram", 15, CalcElements, PointCatCommon},
		{DemeterRam, "r_cp_demeter_ram", 16, CalcElements, PointCatCommon},
		{CupidoUra, "r_cp_cupido_ura", 17, CalcSe, PointCatCommon},
		{HadesUra, "r_cp_hades_ura", 18, CalcSe, PointCatCommon},
		{ZeusUra, "r_cp_zeus_ura", 19, CalcSe, PointCatCommon},
		{KronosUra, "r_cp_kronos_ura", 20, CalcSe, PointCatCommon},
		{ApollonUra, "r_cp_apollon_ura", 21, CalcSe, PointCatCommon},
		{AdmetosUra, "r_cp_admetos_ura", 22, CalcSe, PointCatCommon},
		{VulcanusUra, "r_cp_vulcanus_ura", 23, CalcSe, PointCatCommon},
		{PoseidonUra, "r_cp_poseidon_ura", 24, CalcSe, PointCatCommon},
		{Eris, "r_cp_eris", 25, CalcSe, PointCatCommon},
		{Pholus, "r_cp_pholus", 26, CalcSe, PointCatCommon},
		{Ceres, "r_cp_ceres", 27, CalcSe, PointCatCommon},
		{Pallas, "r_cp_pallas", 28, CalcSe, PointCatCommon},
		{Juno, "r_cp_juno", 29, CalcSe, PointCatCommon},
		{Vesta, "r_cp_vesta", 30, CalcSe, PointCatCommon},
		{Isis, "r_cp_isis", 31, CalcSe, PointCatCommon},
		{Nessus, "r_cp_nessus", 32, CalcSe, PointCatCommon},
		{Huya, "r_cp_huya", 33, CalcSe, PointCatCommon},
		{Varuna, "r_cp_varuna", 34, CalcSe, PointCatCommon},
		{Ixion, "r_cp_ixion", 35, CalcSe, PointCatCommon},
		{Quaoar, "r_cp_quaoar", 36, CalcSe, PointCatCommon},
		{Haumea, "r_cp_haumea", 37, CalcSe, PointCatCommon},
		{Orcus, "r_cp_orcus", 38, CalcSe, PointCatCommon},
		{Makemake, "r_cp_makemake", 39, CalcSe, PointCatCommon},
		{Sedna, "r_cp_sedna", 40, CalcSe, PointCatCommon},
		{Hygieia, "r_cp_hygieia", 41, CalcSe, PointCatCommon},
		{Astraea, "r_cp_astraea", 42, CalcSe, PointCatCommon},
		{ApogeeMean, "r_cp_apogee_mean", 43, CalcSe, PointCatCommon},
		{ApogeeCorrected, "r_cp_apogee_corrected", 44, CalcSe, PointCatCommon},
		{ApogeeInterpolated, "r_cp_apogee_interpolated", 45, CalcSe, PointCatCommon},
		{ApogeeDuval, "r_cp_apogee_duval", 46, CalcFormula, PointCatCommon},
		{PersephoneCarteret, "r_cp_persephone_carteret", 47, CalcFormula, PointCatCommon},
		{VulcanusCarteret, "r_cp_vulcanus_carteret", 48, CalcFormula, PointCatCommon},
		{Ascendant, "r_cp_ascendant", 1001, CalcMundane, PointCatAngle},
		{Mc, "r_cp_mc", 1002, CalcMundane, PointCatAngle},
		{EastPoint, "r_cp_eastpoint", 1003, CalcMundane, PointCatAngle},
		{Vertex, "r_cp_vertex", 1004, CalcMundane, PointCatAngle},
	}
}
