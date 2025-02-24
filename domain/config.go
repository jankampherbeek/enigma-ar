/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

import "image/color"

type ConfigBasic = struct {
	Houses   HouseSystem
	Ayan     Ayanamsha
	ObsPos   ObserverPosition
	ProjType ProjectionType
	Wheel    WheelType
}

type ConfigOrbs = struct {
	BaseOrbAspects   float64
	BaseOrbMidpoints float64
	OrbDeclMidpoints float64
	OrbParallels     float64
	OrbTransits      float64
	OrbSecDir        float64
	OrbSymDir        float64
	OrbPrimDir       float64
}

type ConfigAspect = struct {
	ActualAspect Aspect
	IsUsed       bool
	ShowInChart  bool
	OrbFactor    float64
	Glyph        rune
	Color        color.NRGBA
}

type ConfigPoint = struct {
	ActualPoint ChartPoint
	IsUsed      bool
	ShowInChart bool
	OrbFactor   float64
	Glyph       rune
}

type ConfigProg = struct {
	TransitPoints  []ChartPoint
	SecDirPoints   []ChartPoint
	SymDirPoints   []ChartPoint
	SymDirTimeKey  SymDirKey
	PrimDirSign    []ChartPoint
	PrimDirProm    []ChartPoint
	PrimDirMethod  PrimDirMethods
	PrimDirTimeKey PrimDirKey
	PrimDirMundane bool
	SolarRelocate  bool
}

type Config = struct {
	Basic   ConfigBasic
	Orbs    ConfigOrbs
	Aspects []ConfigAspect
	Points  []ConfigPoint
	Prog    ConfigProg
}
