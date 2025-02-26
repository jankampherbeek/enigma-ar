/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package meta

import (
	"enigma-ar/domain"
	"math"
	"testing"
)

func TestActualConfigBase(t *testing.T) {
	deltas := []string{
		domain.CfgHouseSystem + "=9", // APC
		domain.CfgObspos + "=1",      // Topocentric
		domain.CfgProjType + "=1",    // Oblique longitude
	}
	actCfg, err := ActualConfig(deltas)
	if err != nil {
		t.Fatal(err)
	}
	if actCfg.Basic.Houses != domain.HousesApc {
		t.Errorf("expected: %v, got: %v", domain.HousesApc, actCfg.Basic.Houses)
	}
	if actCfg.Basic.ObsPos != domain.ObsPosTopocentric {
		t.Errorf("expected: %v, got: %v", domain.ObsPosTopocentric, actCfg.Basic.ObsPos)
	}
	if actCfg.Basic.ProjType != domain.ProjTypeOblique {
		t.Errorf("expected: %v, got: %v", domain.ProjTypeOblique, actCfg.Basic.ProjType)
	}
}

func TestActualConfigOrbs(t *testing.T) {
	deltas := []string{
		domain.CfgBaseOrbAspects + "=5",
		domain.CfgOrbDeclMidpoints + "=4",
	}
	actCfg, err := ActualConfig(deltas)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(actCfg.Orbs.BaseOrbAspects-5) > 1e-8 {
		t.Errorf("expected: %v, got: %v", 5, actCfg.Orbs.BaseOrbAspects)
	}
	if math.Abs(actCfg.Orbs.OrbDeclMidpoints-4) > 1e-8 {
		t.Errorf("expected: %v, got: %v", 4, actCfg.Orbs.OrbDeclMidpoints)
	}
}

func TestActualConfigAspects(t *testing.T) {
	deltas := []string{
		domain.CfgAspectX + "0=use:true|show:true|factor:66.000000|glyph:59152|color:{255 255 0 255}",
		domain.CfgAspectX + "1=use:false|show:true|factor:88.000000|glyph:59153|color:{0 0 0 255}",
	}
	actCfg, err := ActualConfig(deltas)
	if err != nil {
		t.Fatal(err)
	}
	var asp1 domain.ConfigAspect
	for _, act := range actCfg.Aspects {
		if int(act.ActualAspect) == 0 {
			asp1 = act
			break
		}
	}
	if !asp1.IsUsed {
		t.Errorf("expected: %v, got: %v", true, asp1.IsUsed)
	}
	if math.Abs(asp1.OrbFactor-66.0) > 1e-8 {
		t.Errorf("expected: %v, got: %v", 66.0, asp1.OrbFactor)
	}
	if asp1.Color.R != 255 || asp1.Color.G != 255 || asp1.Color.B != 0 {
		t.Error("wrong color")
	}
	var asp2 domain.ConfigAspect
	for _, act := range actCfg.Aspects {
		if int(act.ActualAspect) == 1 {
			asp2 = act
			break
		}
	}
	if asp2.IsUsed {
		t.Errorf("expected: %v, got: %v", true, asp2.IsUsed)
	}
	if math.Abs(asp2.OrbFactor-88.0) > 1e-8 {
		t.Errorf("expected: %v, got: %v", 88.0, asp2.OrbFactor)
	}
	if asp2.Color.R != 0 || asp2.Color.G != 0 || asp2.Color.B != 0 {
		t.Error("wrong color")
	}
	if asp2.Glyph != 59153 {
		t.Errorf("expected glyph: %v, got: %v", 59153, asp2.Glyph)
	}

}

func TestActualConfigPoints(t *testing.T) {
	deltas := []string{
		domain.CfgPointX + "0=use:true|show:false|factor:5.500000|glyph:57863",
		domain.CfgPointX + "1=use:false|show:true|factor:1.000000|glyph:57864",
	}
	actCfg, err := ActualConfig(deltas)
	if err != nil {
		t.Fatal(err)
	}
	var point1 domain.ConfigPoint
	for _, act := range actCfg.Points {
		if int(act.ActualPoint) == 0 {
			point1 = act
			break
		}
	}
	if !point1.IsUsed {
		t.Errorf("expected: %v, got: %v", true, point1.IsUsed)
	}
	if point1.ShowInChart {
		t.Errorf("expected: %v, got: %v", false, point1.ShowInChart)
	}
	if math.Abs(point1.OrbFactor-5.5) > 1e-8 {
		t.Errorf("expected: %v, got: %v", 5.5, point1.OrbFactor)
	}
	if point1.Glyph != 57863 {
		t.Errorf("expected glyph: %v, got: %v", 57863, point1.Glyph)
	}

	var point2 domain.ConfigPoint
	for _, act := range actCfg.Points {
		if int(act.ActualPoint) == 1 {
			point2 = act
			break
		}
	}
	if point2.IsUsed {
		t.Errorf("expected: %v, got: %v", true, point2.IsUsed)
	}
	if !point2.ShowInChart {
		t.Errorf("expected: %v, got: %v", false, point2.ShowInChart)
	}
	if math.Abs(point2.OrbFactor-1.0) > 1e-8 {
		t.Errorf("expected: %v, got: %v", 1.0, point2.OrbFactor)
	}
	if point2.Glyph != 57864 {
		t.Errorf("expected glyph: %v, got: %v", 57864, point1.Glyph)
	}

}

func TestActualConfigProgBase(t *testing.T) {
	deltas := []string{
		domain.CfgProgPrimDirMethod + "=1",    // Regiomontanus
		domain.CfgProgSymDirTimeKey + "=1",    // Mean Sun
		domain.CfgProgSolarRelocate + "=true", // Oblique longitude
	}
	actCfg, err := ActualConfig(deltas)
	if err != nil {
		t.Fatal(err)
	}
	if actCfg.Prog.PrimDirMethod != domain.MethodRegiomontanus {
		t.Errorf("expected: %v, got: %v", 1, int(actCfg.Prog.PrimDirMethod))
	}
	if actCfg.Prog.SymDirTimeKey != domain.SymKeyMeanSun {
		t.Errorf("expected: %v, got: %v", 1, int(actCfg.Prog.SymDirTimeKey))
	}
	if !actCfg.Prog.SolarRelocate {
		t.Errorf("expected: %v, got: %v", false, actCfg.Prog.SolarRelocate)
	}
}

func TestActualConfigProgPdProms(t *testing.T) {
	deltas := []string{
		domain.CfgProgPrimDirProm + "=0|1|2",
	}
	actCfg, err := ActualConfig(deltas)
	if err != nil {
		t.Fatal(err)
	}
	if len(actCfg.Prog.PrimDirProm) != 3 {
		t.Errorf("expected items for promissors: %v, got: %v", 3, len(actCfg.Prog.PrimDirProm))
	}
	if actCfg.Prog.PrimDirProm[0] != 0 || actCfg.Prog.PrimDirProm[1] != 1 || actCfg.Prog.PrimDirProm[2] != 2 {
		t.Errorf("expected promissors 0, 1 and 2 but got %v, %v and %v", actCfg.Prog.PrimDirProm[0], actCfg.Prog.PrimDirProm[1], actCfg.Prog.PrimDirProm[2])

	}
}

func TestActualConfigProgPdSign(t *testing.T) {
	deltas := []string{
		domain.CfgProgPrimDirSign + "=5|6",
	}
	actCfg, err := ActualConfig(deltas)
	if err != nil {
		t.Fatal(err)
	}
	if len(actCfg.Prog.PrimDirSign) != 2 {
		t.Errorf("expected items for significators: %v, got: %v", 2, len(actCfg.Prog.PrimDirSign))
	}
	if actCfg.Prog.PrimDirSign[0] != 5 || actCfg.Prog.PrimDirSign[1] != 6 {
		t.Errorf("expected significators 5 and 6 but got %vand %v", actCfg.Prog.PrimDirSign[0], actCfg.Prog.PrimDirSign[1])
	}
}

func TestActualConfigProgTrPoints(t *testing.T) {
	deltas := []string{
		domain.CfgProgTransitPoints + "=2|8",
	}
	actCfg, err := ActualConfig(deltas)
	if err != nil {
		t.Fatal(err)
	}
	if len(actCfg.Prog.TransitPoints) != 2 {
		t.Errorf("expected items for transit points: %v, got %v", 2, len(actCfg.Prog.TransitPoints))
	}
	if actCfg.Prog.TransitPoints[0] != 2 || actCfg.Prog.TransitPoints[1] != 8 {
		t.Errorf("expected transit points: %v and %v; got %v and %v", 2, 8, actCfg.Prog.TransitPoints[0], actCfg.Prog.TransitPoints[1])
	}
}

func TestActualConfigProgSecDPoints(t *testing.T) {
	deltas := []string{
		domain.CfgProgSecDirPoints + "=3|11|12",
	}
	actCfg, err := ActualConfig(deltas)
	if err != nil {
		t.Fatal(err)
	}
	if len(actCfg.Prog.SecDirPoints) != 3 {
		t.Errorf("expected items for sec points: %v, got %v", 3, len(actCfg.Prog.SecDirPoints))
	}
	if actCfg.Prog.SecDirPoints[0] != 3 || actCfg.Prog.SecDirPoints[1] != 11 || actCfg.Prog.SecDirPoints[2] != 12 {
		t.Errorf("expected sec points: %v, %v and %v, got %v, %v and %v", 3, 11, 12, actCfg.Prog.SecDirPoints[0], actCfg.Prog.SecDirPoints[1], actCfg.Prog.SecDirPoints[2])
	}
}

func TestActualConfigProgSymDirPoints(t *testing.T) {
	deltas := []string{
		domain.CfgProgSymDirPoints + "=4|12",
	}
	actCfg, err := ActualConfig(deltas)
	if err != nil {
		t.Fatal(err)
	}
	if len(actCfg.Prog.SymDirPoints) != 2 {
		t.Errorf("expected items for sym points: %v, got %v", 2, len(actCfg.Prog.SymDirPoints))
	}
	if actCfg.Prog.SymDirPoints[0] != 4 || actCfg.Prog.SymDirPoints[1] != 12 {
		t.Errorf("expected sym points: %v and %v, got %v and %v", 4, 12, actCfg.Prog.SymDirPoints[0], actCfg.Prog.SymDirPoints[1])
	}
}
