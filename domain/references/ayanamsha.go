/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package references

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
