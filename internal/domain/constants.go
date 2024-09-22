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
	SE_ADMETOS_URA    = 45
	SE_APOLLON_URA    = 44
	SE_ASTRAEA        = 10005
	SE_CERES          = 17
	SE_CHIRON         = 15
	SE_CUPIDO_URA     = 40
	SE_DEMETER_RAM    = 51
	SE_EARTH          = 14
	SE_ECL_NUT        = -1 // Obliquity/nutation
	SE_ERIS           = 1009001
	SE_HADES_URA      = 41
	SE_HAUMEA         = 146108
	SE_HERMES_RAM     = 50
	SE_HUYA           = 48628
	SE_HYGIEIA        = 10010
	SE_INTP_APOG      = 21
	SE_ISIS           = 48
	SE_IXION          = 38978
	SE_JUNO           = 19
	SE_JUPITER        = 5
	SE_KRONOS_URA     = 43
	SE_MAKEMAKE       = 146472
	SE_MARS           = 4
	SE_MEAN_APOGEE    = 12
	SE_MEAN_NODE      = 10
	SE_MERCURY        = 2
	SE_MOON           = 1
	SE_NEPTUNE        = 8
	SE_NESSUS         = 17066
	SE_ORCUS          = 100482
	SE_OSCU_APOG      = 13
	SE_PALLAS         = 18
	SE_PERSEPHONE_RAM = 49
	SE_PHOLUS         = 16
	SE_PLUTO          = 9
	SE_POSEIDON_URA   = 47
	SE_QUAOAR         = 60000
	SE_SATURN         = 6
	SE_SEDNA          = 100377
	SE_SUN            = 0
	SE_TRUE_NODE      = 11
	SE_URANUS         = 7
	SE_VARUNA         = 30000
	SE_VENUS          = 3
	SE_VESTA          = 20
	SE_VULCANUS_URA   = 46
	SE_ZEUS_URA       = 42
)

// SE flags
const (
	SEFLG_SWIEPH     = 2 // use Swiss Eph
	SEFLG_HELIOC     = 8
	SEFLG_SPEED      = 256
	SEFLG_EQUATORIAL = 2048
	SEFLG_TOPOCTR    = 32768 // 32 * 1024
	SEFLG_SIDEREAL   = 65536 // 64 * 1024
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
