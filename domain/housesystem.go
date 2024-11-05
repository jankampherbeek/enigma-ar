/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

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
