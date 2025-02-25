/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package meta

import (
	"enigma-ar/domain"
	"image/color"
	"sort"
	"strconv"
	"testing"
)

func TestConfigDeltaHappyFlow(t *testing.T) {
	defaultConfig := DefaultConfig()
	newConfig := defaultConfig
	newConfig.Basic.Houses = domain.HousesCampanus
	expected := CfgDelta{
		domain.CfgHouseSystem,
		"5",
	}
	result, err := ConfigDelta(newConfig)
	if err != nil {
		t.Error(err)
	}
	if len(result) < 1 {
		t.Error("result is empty")
	}
	if expected != result[0] {
		t.Errorf("expected: %v, got: %v", expected, result[0])
	}
}

func TestConfigDeltaNoChanges(t *testing.T) {
	newConfig := DefaultConfig()
	result, err := ConfigDelta(newConfig)
	if err != nil {
		t.Error(err)
	}
	if len(result) > 0 {
		t.Error("result should be empty")
	}
}

func TestConfigDeltaAllBasics(t *testing.T) {
	defaultConfig := DefaultConfig()
	newConfig := defaultConfig
	newConfig.Basic.Houses = domain.HousesApc
	newConfig.Basic.Ayan = domain.AyanFagan
	newConfig.Basic.Wheel = domain.WheelTypePlanetsOutside
	newConfig.Basic.ObsPos = domain.ObsPosHeliocentric
	newConfig.Basic.ProjType = domain.ProjTypeOblique
	expected := []CfgDelta{
		{cfgItem: domain.CfgHouseSystem,
			newValue: "9",
		},
		{cfgItem: domain.CfgAyanamsha,
			newValue: "1",
		},
		{cfgItem: domain.CfgWheelType,
			newValue: "2",
		},
		{cfgItem: domain.CfgObspos,
			newValue: "2",
		},
		{cfgItem: domain.CfgProjType,
			newValue: "1",
		},
	}

	result, err := ConfigDelta(newConfig)
	if err != nil {
		t.Error(err)
	}
	if len(result) != len(expected) {
		t.Error("expected more results")
	}
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].cfgItem < expected[j].cfgItem
	})
	sort.Slice(result, func(i, j int) bool {
		return result[i].cfgItem < result[j].cfgItem
	})
	for i := 0; i < len(expected); i++ {
		if expected[i].cfgItem != result[i].cfgItem {
			t.Errorf("expected: %v, got: %v", expected[i], result[i])
		}
	}
}

func TestConfigDeltaOrbs(t *testing.T) {
	defaultConfig := DefaultConfig()
	newConfig := defaultConfig
	newConfig.Orbs.BaseOrbMidpoints = 2.0
	result, err := ConfigDelta(newConfig)
	if err != nil {
		t.Error(err)
	}
	if result[0].cfgItem != domain.CfgBaseOrbMidpoints {
		t.Errorf("expected: %v, got: %v", domain.CfgBaseOrbMidpoints, result[0].cfgItem)
	}
	if result[0].newValue != "2.000000" {
		t.Errorf("expected: %v, got: %v", "2.000000", result[0].newValue)
	}
}

func TestConfigDeltaAspects(t *testing.T) {
	defaultConfig := DefaultConfig()
	newConfig := defaultConfig
	newConfig.Aspects[1].OrbFactor = 90.0
	newConfig.Aspects[2].Color = color.NRGBA{R: 128, G: 128, B: 128, A: 255}
	result, err := ConfigDelta(newConfig)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 2 {
		t.Error("expected 2 results")
	}
	expCfgItem := domain.CfgAspectX + "1"
	expValue := "use:true|show:true|factor:90.000000|glyph:59152|color:{255 0 0 255}"
	if result[0].cfgItem != expCfgItem {
		t.Errorf("expected: %v, got: %v", expCfgItem, result[0].cfgItem)
	}
	if result[0].newValue != expValue {
		t.Errorf("expected: %v, got: %v", expValue, result[0].newValue)
	}
	expCfgItem = domain.CfgAspectX + "2"
	expValue = "use:true|show:true|factor:80.000000|glyph:59168|color:{128 128 128 255}"
	if result[1].cfgItem != expCfgItem {
		t.Errorf("expected: %v, got: %v", expCfgItem, result[1].cfgItem)
	}
	if result[1].newValue != expValue {
		t.Errorf("expected: %v, got: %v", expValue, result[1].newValue)
	}
}

func TestConfigDeltaPoints(t *testing.T) {
	defaultConfig := DefaultConfig()
	newConfig := defaultConfig
	newConfig.Points[6].OrbFactor = 3.0
	newConfig.Points[7].IsUsed = false
	result, err := ConfigDelta(newConfig)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 2 {
		t.Error("expected 2 results")
	}
	expCfgItem := domain.CfgPointX + "5"
	expValue := "use:true|show:true|factor:3.000000|glyph:57862"
	if result[0].cfgItem != expCfgItem {
		t.Errorf("expected: %v, got: %v", expCfgItem, result[0].cfgItem)
	}
	if result[0].newValue != expValue {
		t.Errorf("expected: %v, got: %v", expValue, result[0].newValue)
	}
	expCfgItem = domain.CfgPointX + "6"
	expValue = "use:false|show:true|factor:60.000000|glyph:57863"
	if result[1].cfgItem != expCfgItem {
		t.Errorf("expected: %v, got: %v", expCfgItem, result[1].cfgItem)
	}
	if result[1].newValue != expValue {
		t.Errorf("expected: %v, got: %v", expValue, result[1].newValue)
	}
}

func TestConfigDeltaProgBasics(t *testing.T) {
	defaultConfig := DefaultConfig()
	newConfig := defaultConfig
	newConfig.Prog.PrimDirMethod = domain.MethodRegiomontanus
	result, err := ConfigDelta(newConfig)
	expected := "1"
	if err != nil {
		t.Error(err)
	}
	if len(result) != 1 {
		t.Error("expected 1 result")
	}
	if result[0].cfgItem != domain.CfgProgPrimDirMethod {
		t.Errorf("expected: %v, got: %v", domain.CfgProgPrimDirMethod, result[0].cfgItem)
	}
	if result[0].newValue != strconv.Itoa(int(newConfig.Prog.PrimDirMethod)) {
		t.Errorf("expected: %v, got: %v", expected, result[0].newValue)
	}
}

func TestConfigDeltaProgSignificators(t *testing.T) {
	defaultConfig := DefaultConfig()
	newConfig := defaultConfig
	newConfig.Prog.PrimDirSign = []domain.ChartPoint{
		domain.Sun, domain.Moon, domain.Mercury,
	}
	result, err := ConfigDelta(newConfig)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 1 {
		t.Error("expected 1 result")
	}
	if result[0].cfgItem != domain.CfgProgPrimDirSign {
		t.Errorf("expected: %v, got: %v", domain.CfgProgPrimDirSign, result[0].cfgItem)
	}
	details := "0|1|2"
	if result[0].newValue != details {
		t.Errorf("expected: %v, got: %v", details, result[0].newValue)
	}
}

func TestConfigDeltaProgPromissors(t *testing.T) {
	defaultConfig := DefaultConfig()
	newConfig := defaultConfig
	newConfig.Prog.PrimDirProm = []domain.ChartPoint{
		domain.Ascendant, domain.Mc,
	}
	result, err := ConfigDelta(newConfig)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 1 {
		t.Error("expected 1 result")
	}
	if result[0].cfgItem != domain.CfgProgPrimDirProm {
		t.Errorf("expected: %v, got: %v", domain.CfgProgPrimDirProm, result[0].cfgItem)
	}
	details := "49|50"
	if result[0].newValue != details {
		t.Errorf("expected: %v, got: %v", details, result[0].newValue)
	}
}
