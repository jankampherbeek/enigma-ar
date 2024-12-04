/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

// Sign is a zodiacal sign.
type Sign int

const (
	Aries Sign = iota
	Taurus
	Gemini
	Cancer
	Leo
	Virgo
	Libra
	Scorpio
	Sagittarius
	Capricorn
	Aquarius
	Pisces
)

// SignData contains presentation data for zodiacal signs.
type SignData struct {
	Key       Sign
	Index     int
	TextId    string
	Glyph     rune
	AltGlyphs []rune
}

// AllSigns returns all zodiac signs with their presentation data.
func AllSigns() []SignData {
	return []SignData{
		{Aries, 1, "r_si_aries", '\uE000', []rune{}},
		{Taurus, 2, "r_si_taurus", '\uE001', []rune{}},
		{Gemini, 3, "r_si_gemini", '\uE002', []rune{}},
		{Cancer, 4, "r_si_cancer", '\uE003', []rune{}},
		{Leo, 5, "r_si_leo", '\uE004', []rune{}},
		{Virgo, 6, "r_si_virgo", '\uE005', []rune{}},
		{Libra, 7, "r_si_libra", '\uE006', []rune{}},
		{Scorpio, 8, "r_si_scorpio", '\uE007', []rune{}},
		{Sagittarius, 9, "r_si_sagittarius", '\uE008', []rune{}},
		{Capricorn, 10, "r_si_capricorn", '\uE009', []rune{'\uE012'}},
		{Aquarius, 11, "r_si_aquarius", '\uE010', []rune{}},
		{Pisces, 12, "r_si_pisces", '\uE011', []rune{}},
	}
}
