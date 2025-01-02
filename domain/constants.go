/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

// Constants that can be used globally

// Min and max values
const (
	// TODO check min and max values for JD
	MinJdGeneral     = -2946707.5 // -12999/08/02
	MaxJdGeneral     = 7865293.5  // 16799/12/30
	MinJdChiron      = 1967598.5  // 0675/01/01
	MaxJdChiron      = 3419437.5  // 4650/01/01
	MinJdPholus      = 641716.5   // -2958/01/01
	MaxJdPholus      = 4390615.5  // 7308/12/30
	MinJdCeresVesta  = -2946707.5 // -12999/08/02
	MaxJdCeresVesta  = 5224242.5  // 9591/05/23
	MinJdMinorPoints = 626157.5   // -3000/03/18  Nessus, Huya, Ixion, ORcus, Varuna, MakeMake, Haumea, Quaoar, Eris, Sedna
	MaxJdMinorPoints = 5224242.5  // 9591/05/23
	MinLongitude     = 0.0
	MaxLongitude     = 360.0
	MinDeclination   = -180.0
	MaxDeclination   = 180.0
	MinGeoLong       = 0.0
	MaxGeoLong       = 360.0
	MinGeoLat        = -180.0
	MaxGeoLat        = 180.0
)

// Astronomical constants
const (
	// Length of tropical year measured in tropical days. According to: NASA 365 days, 5 hours, 48 minutes, and 46 seconds,
	// https://www.grc.nasa.gov/www/k-12/Numbers/Math/Mathematical_Thinking/calendar_calculations.htm

	TropicalYearInDays = 365.242199074
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
