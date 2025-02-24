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
	compareBasics(newConfig.Basic, defaultConfig.Basic, allDeltas)
	compareOrbs(newConfig.Orbs, defaultConfig.Orbs, allDeltas)
	err := compareAspects(newConfig.Aspects, defaultConfig.Aspects, allDeltas)
	if err != nil {
		return nil, err
	}
	err = comparePoints(newConfig.Points, defaultConfig.Points, allDeltas)
	if err != nil {
		return nil, err
	}
	compareProg(newConfig.Prog, defaultConfig.Prog, allDeltas)
	return allDeltas, nil
}

func compareBasics(newCfgBasic, defaultCfgBasic domain.ConfigBasic, allDeltas []CfgDelta) {
	if newCfgBasic.ObsPos != defaultCfgBasic.ObsPos {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "ObserverPosition",
			newValue: strconv.Itoa(int(newCfgBasic.ObsPos)),
		})
	}
	if newCfgBasic.Ayan != defaultCfgBasic.Ayan {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "Ayanamsha",
			newValue: strconv.Itoa(int(newCfgBasic.Ayan)),
		})
	}
	if newCfgBasic.ProjType != defaultCfgBasic.ProjType {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "ProjectionType",
			newValue: strconv.Itoa(int(newCfgBasic.ProjType)),
		})
	}
	if newCfgBasic.Houses != defaultCfgBasic.Houses {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "HouseSystem",
			newValue: strconv.Itoa(int(newCfgBasic.Houses)),
		})
	}
	if newCfgBasic.Wheel != defaultCfgBasic.Wheel {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "WheelType",
			newValue: strconv.Itoa(int(newCfgBasic.Wheel)),
		})
	}
}

func compareOrbs(newCfgOrb, defaultCfgOrb domain.ConfigOrbs, allDeltas []CfgDelta) {
	if math.Abs(newCfgOrb.BaseOrbAspects-defaultCfgOrb.BaseOrbAspects) > 1e-8 {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "BaseOrbAspects",
			newValue: fmt.Sprintf("%f", newCfgOrb.BaseOrbAspects),
		})
	}
	if math.Abs(newCfgOrb.OrbDeclMidpoints-defaultCfgOrb.OrbDeclMidpoints) > 1e-8 {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "OrbDeclMidpoints",
			newValue: fmt.Sprintf("%f", newCfgOrb.OrbDeclMidpoints),
		})
	}
	if math.Abs(newCfgOrb.OrbParallels-defaultCfgOrb.OrbParallels) > 1e-8 {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "OrbParallels",
			newValue: fmt.Sprintf("%f", newCfgOrb.OrbParallels),
		})
	}
	if math.Abs(newCfgOrb.BaseOrbMidpoints-defaultCfgOrb.BaseOrbMidpoints) > 1e-8 {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "BaseOrbMidpoints",
			newValue: fmt.Sprintf("%f", newCfgOrb.BaseOrbMidpoints),
		})
	}
	if math.Abs(newCfgOrb.OrbTransits-defaultCfgOrb.OrbTransits) > 1e-8 {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "OrbTransits",
			newValue: fmt.Sprintf("%f", newCfgOrb.OrbTransits),
		})
	}
	if math.Abs(newCfgOrb.OrbPrimDir-defaultCfgOrb.OrbPrimDir) > 1e-8 {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "OrbPrimdir",
			newValue: fmt.Sprintf("%f", newCfgOrb.OrbPrimDir),
		})
	}
	if math.Abs(newCfgOrb.OrbSecDir-defaultCfgOrb.OrbSecDir) > 1e-8 {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "OrbSecDir",
			newValue: fmt.Sprintf("%f", newCfgOrb.OrbSecDir),
		})
	}
	if math.Abs(newCfgOrb.OrbSymDir-defaultCfgOrb.OrbSymDir) > 1e-8 {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "OrbSymDir",
			newValue: fmt.Sprintf("%f", newCfgOrb.OrbSymDir),
		})
	}
}

func compareAspects(newCfgAsp, defaultCfgAsp []domain.ConfigAspect, allDeltas []CfgDelta) error {
	for _, newAsp := range newCfgAsp {
		defAsp, err := findAspect(newAsp, defaultCfgAsp)
		if err != nil {
			return err
		}
		if (newAsp.IsUsed != defAsp.IsUsed) || (newAsp.ShowInChart != defAsp.ShowInChart) ||
			(newAsp.Glyph != defAsp.Glyph) || (newAsp.Color != defAsp.Color) ||
			(math.Abs(newAsp.OrbFactor-defAsp.OrbFactor) > 1e-8) {
			details := fmt.Sprintf("use:%t|show:%t|factor:%f|glyph:%v|color:%v",
				newAsp.IsUsed, newAsp.ShowInChart, newAsp.OrbFactor, newAsp.Glyph, newAsp.Color)
			allDeltas = append(allDeltas, CfgDelta{
				cfgItem:  "Aspect_" + strconv.Itoa(int(defAsp.ActualAspect)),
				newValue: details,
			})
		}
	}
	return nil
}

func comparePoints(newCfgPoints, defaultCfgPoints []domain.ConfigPoint, allDeltas []CfgDelta) error {
	for _, newPoint := range newCfgPoints {
		defPoint, err := findPoint(newPoint, defaultCfgPoints)
		if err != nil {
			return err
		}
		if (newPoint.IsUsed != defPoint.IsUsed) || (newPoint.ShowInChart != defPoint.ShowInChart) ||
			(newPoint.Glyph != defPoint.Glyph) || (newPoint.OrbFactor-defPoint.OrbFactor) > 1e-8 {
			details := fmt.Sprintf("use:%t|show:%t|factor:%f|glyph:%v", newPoint.IsUsed, newPoint.ShowInChart,
				newPoint.OrbFactor, newPoint.Glyph)
			allDeltas = append(allDeltas, CfgDelta{
				cfgItem:  "Point_" + strconv.Itoa(int(defPoint.ActualPoint)),
				newValue: details,
			})
		}
	}
	return nil
}

func compareProg(newCfgProg, defaultCfgProg domain.ConfigProg, allDeltas []CfgDelta) {
	if newCfgProg.PrimDirMethod != defaultCfgProg.PrimDirMethod {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "Prog_PrimdirMethod",
			newValue: strconv.Itoa(int(newCfgProg.PrimDirMethod)),
		})
	}
	if newCfgProg.PrimDirMundane != defaultCfgProg.PrimDirMundane {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "Prog_PrimdirMundane",
			newValue: strconv.FormatBool(newCfgProg.PrimDirMundane),
		})
	}
	if newCfgProg.PrimDirTimeKey != defaultCfgProg.PrimDirTimeKey {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "Prog_PrimdirTimeKey",
			newValue: strconv.Itoa(int(newCfgProg.PrimDirTimeKey)),
		})
	}
	if newCfgProg.SymDirTimeKey != defaultCfgProg.SymDirTimeKey {
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "Prog_SymDirTimeKey",
			newValue: strconv.Itoa(int(newCfgProg.SymDirTimeKey)),
		})
	}
	if newCfgProg.SolarRelocate != defaultCfgProg.SolarRelocate {
		allDeltas = append(allDeltas, CfgDelta{
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
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "Prog_TransitPoints",
			newValue: details,
		})
	}
	if !reflect.DeepEqual(newCfgProg.SecDirPoints, defaultCfgProg.SecDirPoints) {
		details := createDetailsForPoints(newCfgProg.SecDirPoints)
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "Prog_SecDirPoints",
			newValue: details,
		})
	}
	if !reflect.DeepEqual(newCfgProg.SymDirPoints, defaultCfgProg.SymDirPoints) {
		details := createDetailsForPoints(newCfgProg.SymDirPoints)
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "Prog_SymDirPoints",
			newValue: details,
		})
	}
	if !reflect.DeepEqual(newCfgProg.PrimDirProm, defaultCfgProg.PrimDirProm) {
		details := createDetailsForPoints(newCfgProg.PrimDirProm)
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "Prog_PrimDirProm",
			newValue: details,
		})
	}
	if !reflect.DeepEqual(newCfgProg.PrimDirSign, defaultCfgProg.PrimDirSign) {
		details := createDetailsForPoints(newCfgProg.PrimDirSign)
		allDeltas = append(allDeltas, CfgDelta{
			cfgItem:  "Prog_PrimDirSign",
			newValue: details,
		})
	}
}

func findAspect(newAsp domain.ConfigAspect, defAsp []domain.ConfigAspect) (domain.ConfigAspect, error) {
	for _, asp := range defAsp {
		if asp == newAsp {
			return asp, nil
		}
	}
	return newAsp, fmt.Errorf("aspect not found")
}

func findPoint(newPoint domain.ConfigPoint, defPoint []domain.ConfigPoint) (domain.ConfigPoint, error) {
	for _, point := range defPoint {
		if point == newPoint {
			return point, nil
		}
	}
	return newPoint, fmt.Errorf("point not found")
}

func createDetailsForPoints(points []domain.ChartPoint) string {
	nrsAsStrings := make([]string, len(points))
	for i, num := range points {
		nrsAsStrings[i] = fmt.Sprint(num)
	}
	details := strings.Join(nrsAsStrings, "|")
	return details
}
