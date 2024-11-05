/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

type Aspect int

const (
	Conjunction = iota
	Opposition
	Trine
	Square
	Septile
	Sextile
	Quintile
	SemiSextile
	SemiSquare
	SemiQuintile
	BiQuintile
	Inconjunct
	SesquiQuadrate
	TriDecile
	BiSeptile
	TriSeptile
	Novile
	BiNovile
	QuadraNovile
	Undecile
	Centile
	Vigintile
)

type AspectData struct {
	Key      int
	TextId   string
	Distance float64
}

func AllAspects() []AspectData {
	return []AspectData{
		{Conjunction, "r_as_conjunction", 0.0},
		{Opposition, "r_as_opposition", 18.0},
		{Trine, "r_as_trine", 120.0},
		{Square, "r_as_square", 90.0},
		{Septile, "r_as_septile", 51.42857143},
		{Sextile, "r_as_sextile", 60.0},
		{Quintile, "r_as_quintile", 72.0},
		{SemiSextile, "r_as_semi_sextile", 30.0},
		{SemiSquare, "r_as_semi_square", 45.0},
		{SemiQuintile, "r_as_semi_quintile", 36.0},
		{BiQuintile, "r_as_biquintile", 144.0},
		{Inconjunct, "r_as_inconjunct", 150.0},
		{SesquiQuadrate, "r_as_sesquadrate", 135.0},
		{TriDecile, "r_as_tridecile", 108.0},
		{BiSeptile, "r_as_biseptile", 102.85714286},
		{TriSeptile, "r_as_triseptile", 154.28571429},
		{Novile, "r_as_novile", 40.0},
		{BiNovile, "r_as_binovile", 80.0},
		{QuadraNovile, "r_as_quadranovile", 160.0},
		{Undecile, "r_as_undecile", 33.0},
		{Centile, "r_as_centile", 100.0},
		{Vigintile, "r_as_vigintile", 18.0},
	}
}
