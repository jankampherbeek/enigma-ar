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
	"image/color"
	"strconv"
	"strings"
)

func ActualConfig(deltas []string) (domain.Config, error) {

	actConfig, err := createConfig(deltas)
	if err != nil {
		return domain.Config{}, err
	}
	return actConfig, nil
}

func createConfig(deltas []string) (domain.Config, error) {
	actCfg := DefaultConfig()
	newCfg := &actCfg

	for _, delta := range deltas {
		items := strings.Split(delta, "=")
		if len(items) != 2 {
			return domain.Config{}, fmt.Errorf("error in delta : %v", delta)
		}
		err := updateBasic(newCfg, items[0], items[1])
		if err != nil {
			return domain.Config{}, err
		}
		err = updateOrbs(newCfg, items[0], items[1])
		if err != nil {
			return domain.Config{}, err
		}
		err = updateAspects(newCfg, items[0], items[1])
		if err != nil {
			return domain.Config{}, err
		}
		err = updatePoints(newCfg, items[0], items[1])
		if err != nil {
			return domain.Config{}, err
		}
		err = updateProgBase(newCfg, items[0], items[1])
		if err != nil {
			return domain.Config{}, err
		}
		err = updateProgPoints(newCfg, items[0], items[1])
		if err != nil {
			return domain.Config{}, err
		}
	}
	return actCfg, nil
}

func updateBasic(c *domain.Config, item, value string) error {

	switch item {
	case domain.CfgObspos:
		newObsPos, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		c.Basic.ObsPos = domain.ObserverPosition(newObsPos)
	case domain.CfgProjType:
		newProjType, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		c.Basic.ProjType = domain.ProjectionType(newProjType)
	case domain.CfgHouseSystem:
		newHouseSystem, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		c.Basic.Houses = domain.HouseSystem(newHouseSystem)
	case domain.CfgAyanamsha:
		newAyanamsha, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		c.Basic.Ayan = domain.Ayanamsha(newAyanamsha)
	case domain.CfgWheelType:
		newWheelType, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		c.Basic.Wheel = domain.WheelType(newWheelType)
	}
	return nil
}

func updateOrbs(c *domain.Config, item, value string) error {
	switch item {
	case domain.CfgBaseOrbAspects:
		newBaseOrbAspects, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Orbs.BaseOrbAspects = newBaseOrbAspects
	case domain.CfgBaseOrbMidpoints:
		newBaseOrbMidpoints, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Orbs.BaseOrbMidpoints = newBaseOrbMidpoints
	case domain.CfgOrbDeclMidpoints:
		newBaseOrbDeclMidpoints, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Orbs.OrbDeclMidpoints = newBaseOrbDeclMidpoints
	case domain.CfgOrbParallels:
		newBaseOrbParallels, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Orbs.OrbParallels = newBaseOrbParallels
	case domain.CfgOrbTransits:
		newBaseOrbTransits, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Orbs.OrbTransits = newBaseOrbTransits
	case domain.CfgOrbSecDir:
		newBaseOrbSecDir, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Orbs.OrbSecDir = newBaseOrbSecDir
	case domain.CfgOrbSymDir:
		newBaseOrbSymDir, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Orbs.OrbSymDir = newBaseOrbSymDir
	case domain.CfgOrbPrimDir:
		newBaseOrbPrimDir, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Orbs.OrbPrimDir = newBaseOrbPrimDir
	}
	return nil
}

func updateAspects(c *domain.Config, item, value string) error {

	if strings.HasPrefix(item, domain.CfgAspectX) {
		index := len(domain.CfgAspectX)
		aspectNr, err := strconv.Atoi(item[index:])
		if err != nil {
			return err
		}
		for i, asp := range c.Aspects {
			if asp.ActualAspect == domain.Aspect(aspectNr) {
				c.Aspects[i], err = constructAspect(aspectNr, value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func updatePoints(c *domain.Config, item, value string) error {
	if strings.HasPrefix(item, domain.CfgPointX) {
		index := len(domain.CfgPointX)
		pointNr, err := strconv.Atoi(item[index:])
		if err != nil {
			return err
		}
		for i, point := range c.Points {
			if point.ActualPoint == domain.ChartPoint(pointNr) {
				c.Points[i], err = constructPoint(pointNr, value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func updateProgBase(c *domain.Config, item, value string) error {
	switch item {
	case domain.CfgProgSymDirTimeKey:
		newSymDirTK, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Prog.SymDirTimeKey = domain.SymDirKey(newSymDirTK)
	case domain.CfgProgPrimDirTimeKey:
		newPrimDirTK, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Prog.PrimDirTimeKey = domain.PrimDirKey(newPrimDirTK)
	case domain.CfgProgPrimDirMundane:
		newPrimDirMundane, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		c.Prog.PrimDirMundane = newPrimDirMundane
	case domain.CfgProgPrimDirMethod:
		newPrimDirMethod, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		c.Prog.PrimDirMethod = domain.PrimDirMethods(newPrimDirMethod)
	case domain.CfgProgSolarRelocate:
		newSolarRelocate, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		c.Prog.SolarRelocate = newSolarRelocate
	}
	return nil
}

func updateProgPoints(c *domain.Config, item, value string) error {
	switch item {
	case domain.CfgProgTransitPoints:
		tPts, err := createPointList(value)
		if err != nil {
			return err
		}
		c.Prog.TransitPoints = tPts
	case domain.CfgProgSecDirPoints:
		scdPts, err := createPointList(value)
		if err != nil {
			return err
		}
		c.Prog.SecDirPoints = scdPts
	case domain.CfgProgSymDirPoints:
		smdPts, err := createPointList(value)
		if err != nil {
			return err
		}
		c.Prog.SymDirPoints = smdPts
	case domain.CfgProgPrimDirProm:
		promtPts, err := createPointList(value)
		if err != nil {
			return err
		}
		c.Prog.PrimDirProm = promtPts
	case domain.CfgProgPrimDirSign:
		signPts, err := createPointList(value)
		if err != nil {
			return err
		}
		c.Prog.PrimDirSign = signPts
	}
	return nil
}

func constructAspect(aspectNr int, value string) (domain.ConfigAspect, error) {
	items := strings.Split(value, "|")
	if len(items) != 5 {
		return domain.ConfigAspect{}, fmt.Errorf("wrong nr of items for aspect")
	}
	usedItems := strings.Split(items[0], ":")
	if len(usedItems) != 2 {
		return domain.ConfigAspect{}, fmt.Errorf("wrong nr of items for isUsed")
	}
	used, err := strconv.ParseBool(usedItems[1])
	if err != nil {
		return domain.ConfigAspect{}, err
	}
	usedItems = strings.Split(items[1], ":")
	show, err := strconv.ParseBool(usedItems[1])
	if err != nil {
		return domain.ConfigAspect{}, err
	}
	usedItems = strings.Split(items[2], ":")
	ofact, err := strconv.ParseFloat(usedItems[1], 64)
	if err != nil {
		return domain.ConfigAspect{}, err
	}
	usedItems = strings.Split(items[3], ":")
	runeStr := usedItems[1]
	runeInt, _ := strconv.Atoi(runeStr)
	gl := rune(runeInt)

	usedItems = strings.Split(items[4], ":")
	col, err := rgbToNRGBA(usedItems[1])
	if err != nil {
		return domain.ConfigAspect{}, err
	}
	ca := domain.ConfigAspect{
		ActualAspect: domain.Aspect(aspectNr),
		IsUsed:       used,
		ShowInChart:  show,
		OrbFactor:    ofact,
		Glyph:        gl,
		Color:        col,
	}
	return ca, nil
}

func constructPoint(pointNr int, value string) (domain.ConfigPoint, error) {
	items := strings.Split(value, "|")
	if len(items) != 4 {
		return domain.ConfigPoint{}, fmt.Errorf("wrong nr of items for point")
	}
	usedItems := strings.Split(items[0], ":")
	if len(usedItems) != 2 {
		return domain.ConfigPoint{}, fmt.Errorf("wrong nr of items for isUsed")
	}
	used, err := strconv.ParseBool(usedItems[1])
	if err != nil {
		return domain.ConfigPoint{}, err
	}
	usedItems = strings.Split(items[1], ":")
	show, err := strconv.ParseBool(usedItems[1])
	if err != nil {
		return domain.ConfigPoint{}, err
	}
	usedItems = strings.Split(items[2], ":")
	ofact, err := strconv.ParseFloat(usedItems[1], 64)
	if err != nil {
		return domain.ConfigPoint{}, err
	}
	usedItems = strings.Split(items[3], ":")
	runeStr := usedItems[1]
	runeInt, _ := strconv.Atoi(runeStr)
	gl := rune(runeInt)

	cp := domain.ConfigPoint{
		ActualPoint: domain.ChartPoint(pointNr),
		IsUsed:      used,
		ShowInChart: show,
		OrbFactor:   ofact,
		Glyph:       gl,
	}
	return cp, nil
}

func rgbToNRGBA(rgb string) (color.NRGBA, error) {
	// Remove "rgb(" and ")" and split by comma
	rgb = strings.TrimPrefix(rgb, "{")
	rgb = strings.TrimSuffix(rgb, "}")
	parts := strings.Split(rgb, " ")

	if len(parts) != 4 {
		return color.NRGBA{}, fmt.Errorf("invalid RGB format")
	}

	r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return color.NRGBA{}, err
	}

	g, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return color.NRGBA{}, err
	}

	b, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		return color.NRGBA{}, err
	}

	return color.NRGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}, nil
}

func createPointList(value string) ([]domain.ChartPoint, error) {
	pts := strings.Split(value, "|")
	cfgPoints := make([]domain.ChartPoint, len(pts))
	for i, tPt := range pts {
		ptNr, err := strconv.Atoi(tPt)
		if err != nil {
			return nil, err
		}
		cfgPoints[i] = domain.ChartPoint(ptNr)
	}
	return cfgPoints, nil
}
