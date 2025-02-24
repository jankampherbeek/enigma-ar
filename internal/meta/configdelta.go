/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package meta

import (
	"enigma-ar/domain"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

type CfgDelta struct {
	cfgItem  string
	newValue string
}

func ConfigDelta(newConfig domain.Config) ([]CfgDelta, error) {
	defaultConfig := DefaultConfig()
	var allDeltas []CfgDelta
	allDeltas = append(allDeltas, compareBasics(newConfig.Basic, defaultConfig.Basic)...)
	allDeltas = append(allDeltas, compareOrbs(newConfig.Orbs, defaultConfig.Orbs)...)
	newDeltas, err := compareAspects(newConfig.Aspects, defaultConfig.Aspects)
	if err != nil {
		return nil, err
	}
	allDeltas = append(allDeltas, newDeltas...)
	newDeltas, err = comparePoints(newConfig.Points, defaultConfig.Points)
	if err != nil {
		return nil, err
	}
	allDeltas = append(allDeltas, newDeltas...)
	allDeltas = append(allDeltas, compareProg(newConfig.Prog, defaultConfig.Prog)...)
	return allDeltas, nil
}

func compareBasics(newCfgBasic, defaultCfgBasic domain.ConfigBasic) []CfgDelta {
	var newDeltas []CfgDelta
	if newCfgBasic.ObsPos != defaultCfgBasic.ObsPos {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "ObserverPosition",
			newValue: strconv.Itoa(int(newCfgBasic.ObsPos)),
		})
	}
	if newCfgBasic.Ayan != defaultCfgBasic.Ayan {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "Ayanamsha",
			newValue: strconv.Itoa(int(newCfgBasic.Ayan)),
		})
	}
	if newCfgBasic.ProjType != defaultCfgBasic.ProjType {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "ProjectionType",
			newValue: strconv.Itoa(int(newCfgBasic.ProjType)),
		})
	}
	if newCfgBasic.Houses != defaultCfgBasic.Houses {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "HouseSystem",
			newValue: strconv.Itoa(int(newCfgBasic.Houses)),
		})
	}
	if newCfgBasic.Wheel != defaultCfgBasic.Wheel {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "WheelType",
			newValue: strconv.Itoa(int(newCfgBasic.Wheel)),
		})
	}
	return newDeltas
}

func compareOrbs(newCfgOrb, defaultCfgOrb domain.ConfigOrbs) []CfgDelta {
	var newDeltas []CfgDelta
	if math.Abs(newCfgOrb.BaseOrbAspects-defaultCfgOrb.BaseOrbAspects) > 1e-8 {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "BaseOrbAspects",
			newValue: fmt.Sprintf("%f", newCfgOrb.BaseOrbAspects),
		})
	}
	if math.Abs(newCfgOrb.OrbDeclMidpoints-defaultCfgOrb.OrbDeclMidpoints) > 1e-8 {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "OrbDeclMidpoints",
			newValue: fmt.Sprintf("%f", newCfgOrb.OrbDeclMidpoints),
		})
	}
	if math.Abs(newCfgOrb.OrbParallels-defaultCfgOrb.OrbParallels) > 1e-8 {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "OrbParallels",
			newValue: fmt.Sprintf("%f", newCfgOrb.OrbParallels),
		})
	}
	if math.Abs(newCfgOrb.BaseOrbMidpoints-defaultCfgOrb.BaseOrbMidpoints) > 1e-8 {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "BaseOrbMidpoints",
			newValue: fmt.Sprintf("%f", newCfgOrb.BaseOrbMidpoints),
		})
	}
	if math.Abs(newCfgOrb.OrbTransits-defaultCfgOrb.OrbTransits) > 1e-8 {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "OrbTransits",
			newValue: fmt.Sprintf("%f", newCfgOrb.OrbTransits),
		})
	}
	if math.Abs(newCfgOrb.OrbPrimDir-defaultCfgOrb.OrbPrimDir) > 1e-8 {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "OrbPrimdir",
			newValue: fmt.Sprintf("%f", newCfgOrb.OrbPrimDir),
		})
	}
	if math.Abs(newCfgOrb.OrbSecDir-defaultCfgOrb.OrbSecDir) > 1e-8 {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "OrbSecDir",
			newValue: fmt.Sprintf("%f", newCfgOrb.OrbSecDir),
		})
	}
	if math.Abs(newCfgOrb.OrbSymDir-defaultCfgOrb.OrbSymDir) > 1e-8 {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "OrbSymDir",
			newValue: fmt.Sprintf("%f", newCfgOrb.OrbSymDir),
		})
	}
	return newDeltas
}

func compareAspects(newCfgAsp, defaultCfgAsp []domain.ConfigAspect) ([]CfgDelta, error) {
	var newDeltas []CfgDelta
	if len(newCfgAsp) != len(defaultCfgAsp) {
		return nil, fmt.Errorf("nr of default aspects and new aspects must be equal")
	}
	for i := 0; i < len(defaultCfgAsp); i++ {
		newAsp := newCfgAsp[i]
		defAsp := defaultCfgAsp[i]
		if (newAsp.IsUsed != defAsp.IsUsed) || (newAsp.ShowInChart != defAsp.ShowInChart) ||
			(newAsp.Glyph != defAsp.Glyph) || (newAsp.Color != defAsp.Color) ||
			(math.Abs(newAsp.OrbFactor-defAsp.OrbFactor) > 1e-8) {
			details := fmt.Sprintf("use:%t|show:%t|factor:%f|glyph:%v|color:%v",
				newAsp.IsUsed, newAsp.ShowInChart, newAsp.OrbFactor, newAsp.Glyph, newAsp.Color)
			newDeltas = append(newDeltas, CfgDelta{
				cfgItem:  "Aspect_" + strconv.Itoa(int(defAsp.ActualAspect)),
				newValue: details,
			})
		}
	}
	return newDeltas, nil
}

func comparePoints(newCfgPoints, defaultCfgPoints []domain.ConfigPoint) ([]CfgDelta, error) {
	var newDeltas []CfgDelta
	if len(newCfgPoints) != len(defaultCfgPoints) {
		return nil, fmt.Errorf("nr of default points and new points must be equal")
	}
	for i := 0; i < len(defaultCfgPoints); i++ {
		newPoint := newCfgPoints[i]
		defPoint := defaultCfgPoints[i]
		if (newPoint.IsUsed != defPoint.IsUsed) || (newPoint.ShowInChart != defPoint.ShowInChart) ||
			(newPoint.Glyph != defPoint.Glyph) || (math.Abs(newPoint.OrbFactor-defPoint.OrbFactor) > 1e-8) {
			details := fmt.Sprintf("use:%t|show:%t|factor:%f|glyph:%v", newPoint.IsUsed, newPoint.ShowInChart,
				newPoint.OrbFactor, newPoint.Glyph)
			newDeltas = append(newDeltas, CfgDelta{
				cfgItem:  "Point_" + strconv.Itoa(int(defPoint.ActualPoint)),
				newValue: details,
			})
		}
	}
	return newDeltas, nil
}

func compareProg(newCfgProg, defaultCfgProg domain.ConfigProg) []CfgDelta {
	var newDeltas []CfgDelta
	if newCfgProg.PrimDirMethod != defaultCfgProg.PrimDirMethod {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "Prog_PrimdirMethod",
			newValue: strconv.Itoa(int(newCfgProg.PrimDirMethod)),
		})
	}
	if newCfgProg.PrimDirMundane != defaultCfgProg.PrimDirMundane {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "Prog_PrimdirMundane",
			newValue: strconv.FormatBool(newCfgProg.PrimDirMundane),
		})
	}
	if newCfgProg.PrimDirTimeKey != defaultCfgProg.PrimDirTimeKey {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "Prog_PrimdirTimeKey",
			newValue: strconv.Itoa(int(newCfgProg.PrimDirTimeKey)),
		})
	}
	if newCfgProg.SymDirTimeKey != defaultCfgProg.SymDirTimeKey {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "Prog_SymDirTimeKey",
			newValue: strconv.Itoa(int(newCfgProg.SymDirTimeKey)),
		})
	}
	if newCfgProg.SolarRelocate != defaultCfgProg.SolarRelocate {
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "Prog_SolarRelocate",
			newValue: strconv.FormatBool(newCfgProg.SolarRelocate),
		})
	}
	if !reflect.DeepEqual(newCfgProg.TransitPoints, defaultCfgProg.TransitPoints) {
		nrsAsStrings := make([]string, len(newCfgProg.TransitPoints))
		for i, num := range newCfgProg.TransitPoints {
			nrsAsStrings[i] = fmt.Sprint(num)
		}
		details := strings.Join(nrsAsStrings, "|")
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "Prog_TransitPoints",
			newValue: details,
		})
	}
	if !reflect.DeepEqual(newCfgProg.SecDirPoints, defaultCfgProg.SecDirPoints) {
		details := createDetailsForPoints(newCfgProg.SecDirPoints)
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "Prog_SecDirPoints",
			newValue: details,
		})
	}
	if !reflect.DeepEqual(newCfgProg.SymDirPoints, defaultCfgProg.SymDirPoints) {
		details := createDetailsForPoints(newCfgProg.SymDirPoints)
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "Prog_SymDirPoints",
			newValue: details,
		})
	}
	if !reflect.DeepEqual(newCfgProg.PrimDirProm, defaultCfgProg.PrimDirProm) {
		details := createDetailsForPoints(newCfgProg.PrimDirProm)
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "Prog_PrimDirProm",
			newValue: details,
		})
	}
	if !reflect.DeepEqual(newCfgProg.PrimDirSign, defaultCfgProg.PrimDirSign) {
		details := createDetailsForPoints(newCfgProg.PrimDirSign)
		newDeltas = append(newDeltas, CfgDelta{
			cfgItem:  "Prog_PrimDirSign",
			newValue: details,
		})
	}
	return newDeltas
}

func createDetailsForPoints(points []domain.ChartPoint) string {
	nrsAsStrings := make([]string, len(points))
	for i, num := range points {
		nrsAsStrings[i] = fmt.Sprint(num)
	}
	details := strings.Join(nrsAsStrings, "|")
	return details
}
