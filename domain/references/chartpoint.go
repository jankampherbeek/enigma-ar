/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package references

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
