/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

import "path/filepath"

// Constants that can be used globally

// Min and max values
const (
	// TODO check min and max values for JD
	MinJdGeneral             = -2946707.5 // -12999/08/02
	MaxJdGeneral             = 7865293.5  // 16799/12/30
	MinJdChiron              = 1967598.5  // 0675/01/01
	MaxJdChiron              = 3419437.5  // 4650/01/01
	MinJdPholus              = 641716.5   // -2958/01/01
	MaxJdPholus              = 4390615.5  // 7308/12/30
	MinJdCeresVesta          = -2946707.5 // -12999/08/02
	MaxJdCeresVesta          = 5224242.5  // 9591/05/23
	MinJdMinorPoints         = 626157.5   // -3000/03/18  Nessus, Huya, Ixion, ORcus, Varuna, MakeMake, Haumea, Quaoar, Eris, Sedna
	MaxJdMinorPoints         = 5224242.5  // 9591/05/23
	MinLongitude             = 0.0
	MaxLongitude             = 360.0
	MinArmc                  = 0.0
	MaxArmc                  = 360.0
	MinDeclination           = -180.0
	MaxDeclination           = 180.0
	MinObliquity             = 20.0
	MaxObliquity             = 30.0
	MinGeoLong               = -180.0
	MaxGeoLong               = 180.0
	MinGeoLat                = -90.0
	MaxGeoLat                = 90.0
	MinMultiplicationCGroups = 1
	MaxMultiplicationCGroups = 1000
	MinSizeCGroups           = 2
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

// general purpose constants
const (
	PathSep = string(filepath.Separator)
)
