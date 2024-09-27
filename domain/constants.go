/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

// Constants that can be use globally

// Celestial points as used by the SE.
// _RAM = School of Ram, _URA = Uranian.
const (
	SeAdmetosUra    = 45
	SeApollonUra    = 44
	SeAstraea       = 10005
	SeCeres         = 17
	SeChiron        = 15
	SeCupidoUra     = 40
	SeDemeterRam    = 51
	SeEarth         = 14
	SeEclNut        = -1 // Obliquity/nutation
	SeEris          = 1009001
	SeHadesUra      = 41
	SeHaumea        = 146108
	SeHermesRam     = 50
	SeHuya          = 48628
	SeHygieia       = 10010
	SeIntpApog      = 21
	SeIsis          = 48
	SeIxion         = 38978
	SeJuno          = 19
	SeJupiter       = 5
	SeKronosUra     = 43
	SeMakemake      = 146472
	SeMars          = 4
	SeMeanApogee    = 12
	SeMeanNode      = 10
	SeMercury       = 2
	SeMoon          = 1
	SeNeptune       = 8
	SeNessus        = 17066
	SeOrcus         = 100482
	SeOscuApog      = 13
	SePallas        = 18
	SePersephoneRam = 49
	SePholus        = 16
	SePluto         = 9
	SePoseidonUra   = 47
	SeQuaoar        = 60000
	SeSaturn        = 6
	SeSedna         = 100377
	SeSun           = 0
	SeTrueNode      = 11
	SeUranus        = 7
	SeVaruna        = 30000
	SeVenus         = 3
	SeVesta         = 20
	SeVulcanusUra   = 46
	SeZeusUra       = 42
)

// SE flags
const (
	SeflgSwieph     = 2 // use Swiss Eph
	SeflgHelioc     = 8
	SeflgSpeed      = 256
	SeflgEquatorial = 2048
	SeflgTopoc      = 32768 // 32 * 1024
	SeflgSidereal   = 65536 // 64 * 1024
)

// House systems
const (
	NoHouses        = 'W'
	Placidus        = 'P'
	Koch            = 'K'
	Porphyri        = 'O'
	Regiomontanus   = 'R'
	Campanus        = 'C'
	Alcabitius      = 'B'
	TopoCentric     = 'T'
	Krusinski       = 'U'
	Apc             = 'Y'
	Morin           = 'M'
	WholeSign       = 'W'
	EqualAsc        = 'A'
	EqualMc         = 'D'
	EqualAries      = 'N'
	Vehlow          = 'V'
	Axial           = 'X'
	Horizon         = 'H'
	Carter          = 'F'
	Gauquelin       = 'G'
	SunShine        = 'i'
	SunShineTreindl = 'I'
	PullenSd        = 'L'
	PullenSr        = 'Q'
	Sripati         = 'S'
)

// Ayanamshas
// the range 0..40 is checked with a guard clause in api's for clculation.
const (
	None                  = 0
	Fagan                 = 1
	Lahiri                = 2
	DeLuce                = 3
	Raman                 = 4
	UshaShashi            = 5
	Krishnamurti          = 6
	DjwhalKhul            = 7
	Yukteshwar            = 8
	Bhasin                = 9
	Kugler1               = 10
	Kugler2               = 11
	Kugler3               = 12
	Huber                 = 13
	EtaPiscium            = 14
	Aldebaran15Tau        = 15
	Hipparchus            = 16
	Sassanian             = 17
	GalactCtr0Sag         = 18
	J2000                 = 19
	J1900                 = 20
	B1950                 = 21
	SuryaSiddhanta        = 22
	SuryaSiddhantaMeanSun = 23
	Aryabhata             = 24
	AryabhataMeanSun      = 25
	SsRevati              = 26
	SsCitra               = 27
	TrueCitra             = 28
	TrueRevati            = 29
	TruePushya            = 30
	GalacticCtrBrand      = 31
	GalacticEqIau1958     = 32
	GalacticEq            = 33
	GalacticEqMidMula     = 34
	Skydram               = 35
	TrueMula              = 36
	Dhruva                = 37
	Aryabhata522          = 38
	Britton               = 39
	GalacticCtrOCap       = 40
)
