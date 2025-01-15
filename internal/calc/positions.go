/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	"enigma-ar/domain"
	"enigma-ar/internal/calc/conversion"
	"enigma-ar/internal/calc/mathextra"
	"enigma-ar/internal/se"
	"fmt"
	"math"
	"path/filepath"
)

const (
	Zero                      = 0.0
	JD1900                    = 2415020.5 // Julian day for 1900/1/1 0:00:00 UT
	StartPointPersephoneCart  = 212.0
	StartPointVulcanusCart    = 15.7
	YearlySpeedPersephoneCart = 1.0
	YearlySpeedVulcanusCart   = 0.55
)

// PointPosCalculator calculates a fully defined set of positions and speeds, in ecliptical, equatorial and horizontal coordinates.
type PointPosCalculator interface {
	CalcPointPos(request domain.PointPositionsRequest) ([]domain.PointPosResult, error)
}

// PointRangeCalculator calculates the positions or speeds for a range of subsequent julian day numbers.
type PointRangeCalculator interface {
	CalcPointRange(request domain.PointRangeRequest) ([]domain.PointRangeResult, error)
}

// HousePosCalculator calculates the positions of houses and other mundane points.
type HousePosCalculator interface {
	CalcHousePos(request domain.HousePosRequest) ([]domain.HousePosResult, []domain.HousePosResult, error)
}

type PointPosCalculation struct {
	sePointCalc   se.SwephPointPosCalculator
	seHorPosCalc  se.SwephHorPosCalculator
	elementsCalc  PointsElementsCalculator
	seEpsilonCalc se.SwephEpsilonCalculator
	sePrep        se.SwephPreparator
}

func NewPointPosCalculation() PointPosCalculator {
	ppc := se.NewSwephPointPosCalculation()
	hpc := se.NewSwephHorPosCalculation()
	elc := NewPointsElementsCalculation()
	ec := se.NewSwephEpsilonCalculation()
	prep := se.NewSwephPreparation()
	return PointPosCalculation{ppc, hpc, elc, ec, prep}
}

// CalcPointPos calculates fully defined positions for one or more celestial points
// PRE MinJdGeneral () <= request.JdUt <= MaxJdGeneral ()
// PRE MinGeoLong <= request.GeoLong < MaxGeoLong
// PRE MinGeoLat <= request.GeoLat < MaxGeoLat
// POST : if no error occurred returns positions for the given points, otherwise returns empty slice and error
func (calc PointPosCalculation) CalcPointPos(request domain.PointPositionsRequest) ([]domain.PointPosResult, error) {

	jdUt := request.JdUt
	geoLong := request.GeoLong
	geoLat := request.GeoLat
	altitude := 0.0 // altitude in meters

	positions := make([]domain.PointPosResult, 0)
	eclFlags := SeFlags(domain.CoordEcliptical, request.ObsPos, request.Ayanamsha)
	equFlags := SeFlags(domain.CoordEquatorial, request.ObsPos, request.Ayanamsha)
	if request.ObsPos == domain.ObsPosTopocentric {
		calc.sePrep.SetTopo(geoLong, geoLat, altitude)
	}
	var ayanOffset float64
	var err error
	if request.Ayanamsha != domain.AyanNone {
		calc.sePrep.SetSidereal(request.Ayanamsha)
		ayanOffset, err = calc.sePrep.AyanOffset(request.JdUt)
		if err != nil {
			return nil, fmt.Errorf("error when defining offset for ayanamsha: %v", err)
		}
	}

	var allPoints = domain.AllChartPoints()
	var calcCat domain.CalculationCat

	for i := 0; i < len(request.Points); i++ {
		point := request.Points[i]
		calcId := allPoints[point].CalcId
		calcCat = domain.AllChartPoints()[point].CalcCat
		switch calcCat {
		case domain.CalcSe:
			position, err := calc.calcPointPosViaSe(calcId, point, jdUt, eclFlags, equFlags, geoLong, geoLat)
			if err != nil {
				return nil, fmt.Errorf("calc point positions failed for %v", point)
			}
			positions = append(positions, position)
		case domain.CalcElements:
			position, err := calc.calcElements(point, jdUt, ayanOffset, request.ObsPos)
			if err != nil {
				return nil, fmt.Errorf("calc point positions failed for %v", point)
			}
			positions = append(positions, position)
		case domain.CalcFormula:
			position, err := calc.calcPointPosViaFormula(calcId, point, jdUt, eclFlags, equFlags, ayanOffset)
			if err != nil {
				return nil, fmt.Errorf("calc point positions failed for %v", point)
			}
			positions = append(positions, position)
		case domain.CalcMundane:
			// handle mundane
		case domain.CalcZodiacFixed:
			// handle zodiac fixed
		case domain.CalcLots:
			// handle calc lots
		}
	}
	if request.ProjType == domain.ProjTypeOblique { // handle oblique longitude

		olc := NewObliqueLongCalculation()
		newPositions, err := olc.calcObliqueLongitudes(positions, request.Armc, request.Obliquity, geoLat, ayanOffset)
		if err != nil {
			return nil, fmt.Errorf("calc oblique positions failed")
		}
		positions = newPositions
	}
	return positions, nil
}

func (calc PointPosCalculation) calcPointPosViaSe(index int, point domain.ChartPoint, jdUt float64,
	eclFlags, equFlags int, geoLong, geoLat float64) (domain.PointPosResult, error) {

	// TODO check if this initalisation is required
	sep := string(filepath.Separator)
	ephePath := ".." + sep + ".." + sep + "sedata" // path is relative from current package
	sp := se.NewSwephPreparation()
	sp.SetEphePath(ephePath)

	var position domain.PointPosResult
	//posEcl, errEcl := calc.sePointCalc.CalcPointPos(jdUt, index, eclFlags)
	posEcl, errEcl := calc.sePointCalc.CalcPointPos(jdUt, int(point), eclFlags)
	if errEcl != nil {
		return position, errEcl
	}
	//posEqu, errEqu := calc.sePointCalc.CalcPointPos(jdUt, index, equFlags)
	posEqu, errEqu := calc.sePointCalc.CalcPointPos(jdUt, int(point), equFlags)
	if errEqu != nil {
		return position, errEqu
	}
	height := 0.0
	pointRa := posEqu[0]
	pointDecl := posEqu[1]
	horFlags := domain.SeflgEquatorial
	posHor := calc.seHorPosCalc.CalcHorPos(jdUt, geoLong, geoLat, height, pointRa, pointDecl, horFlags)
	position = domain.PointPosResult{
		Point:     point,
		LonPos:    posEcl[0],
		LonSpeed:  posEcl[3],
		LatPos:    posEcl[1],
		LatSpeed:  posEcl[4],
		RaPos:     posEqu[0],
		RaSpeed:   posEqu[3],
		DeclPos:   posEqu[1],
		DeclSpeed: posEqu[4],
		RadvPos:   posEcl[2],
		RadvSpeed: posEcl[5],
		AzimPos:   posHor[0],
		AltitPos:  posHor[2],
	}
	return position, nil
}

func (calc PointPosCalculation) calcElements(point domain.ChartPoint, jdUt float64,
	ayanOffset float64, obsPos domain.ObserverPosition) (domain.PointPosResult, error) {
	var position domain.PointPosResult
	pointId := domain.AllChartPoints()[point].CalcId

	// TODO add error checks
	positions := calc.elementsCalc.Calculate(pointId, jdUt, obsPos)
	lonPos := positions[0]
	latPos := positions[1]
	distance := positions[2]
	posBefore := calc.elementsCalc.Calculate(pointId, jdUt-0.5, obsPos)
	posAfter := calc.elementsCalc.Calculate(pointId, jdUt+0.5, obsPos)
	lonSpeed := posAfter[0] - posBefore[0]
	latSpeed := posAfter[1] - posBefore[1]
	distanceSpeed := posAfter[2] - posBefore[2]
	position = domain.PointPosResult{
		Point:     point,
		LonPos:    lonPos + ayanOffset,
		LonSpeed:  lonSpeed,
		LatPos:    latPos,
		LatSpeed:  latSpeed,
		RadvPos:   distance,
		RadvSpeed: distanceSpeed,
	}
	return position, nil
}

func (calc PointPosCalculation) calcPointPosViaFormula(index int, point domain.ChartPoint, jdUt float64,
	eclFlags, equFlags int, ayanOffset float64) (domain.PointPosResult, error) {

	var emptyPosition domain.PointPosResult
	var position domain.PointPosResult
	var eclLong, eclSpeed float64
	switch point {
	case domain.PersephoneCarteret:
		eclLong = calc.calcCarteretHypPlanet(jdUt, StartPointPersephoneCart, YearlySpeedPersephoneCart) + ayanOffset
		eclSpeed = YearlySpeedPersephoneCart / domain.TropicalYearInDays
	case domain.VulcanusCarteret:
		eclLong = calc.calcCarteretHypPlanet(jdUt, StartPointVulcanusCart, YearlySpeedVulcanusCart) + ayanOffset
		eclSpeed = YearlySpeedVulcanusCart / domain.TropicalYearInDays
	case domain.ApogeeDuval:
		result, err := calc.calcApogeeDuval(jdUt, eclFlags, equFlags)
		if err != nil {
			return emptyPosition, fmt.Errorf("error in calcPointPosViaFormula %v", err)
		}
		resultBefore, err := calc.calcApogeeDuval(jdUt-0.5, eclFlags, equFlags)
		if err != nil {
			return emptyPosition, fmt.Errorf("error in calcPointPosViaFormula %v", err)
		}
		resultAfter, err := calc.calcApogeeDuval(jdUt+0.5, eclFlags, equFlags)
		if err != nil {
			return emptyPosition, fmt.Errorf("error in calcPointPosViaFormula %v", err)
		}
		eclLong = result + ayanOffset
		eclSpeed = resultAfter - resultBefore
	default:
		return emptyPosition, fmt.Errorf("calcPointPosViaFOrmula encountered unknown point %v", point)
	}
	position = domain.PointPosResult{
		Point:     point,
		LonPos:    eclLong,
		LonSpeed:  eclSpeed,
		LatPos:    Zero,
		LatSpeed:  Zero,
		RaPos:     Zero,
		RaSpeed:   Zero,
		DeclPos:   Zero,
		DeclSpeed: Zero,
		RadvPos:   Zero,
		RadvSpeed: Zero,
		AzimPos:   Zero,
		AltitPos:  Zero,
	}
	return position, nil
}

func (calc PointPosCalculation) calcCarteretHypPlanet(jdUt, startPoint, yearlySpeed float64) float64 {
	return startPoint + ((jdUt - JD1900) * (yearlySpeed / domain.TropicalYearInDays))
}

func (calc PointPosCalculation) calcApogeeDuval(jdUt float64, eclFlags, equFlags int) (float64, error) {
	flagsEcl := 2 + 256 // use SE + speed
	factor1 := 12.37
	geoLat := 0.0
	geoLong := 0.0
	indexSun := domain.AllChartPoints()[domain.Sun].CalcId
	indexApogeeMean := domain.AllChartPoints()[domain.ApogeeMean].CalcId
	longSun, err := calc.calcPointPosViaSe(indexSun, domain.Sun, jdUt, flagsEcl, equFlags, geoLong, geoLat)
	longApogeeMean, err := calc.calcPointPosViaSe(indexApogeeMean, domain.ApogeeMean, jdUt, flagsEcl, equFlags, geoLong, geoLat)
	fmt.Printf("indexApogeeMean %d, flagsEcl %d, jdUt %f\n", indexApogeeMean, flagsEcl, jdUt)

	diff, err := ValueToRange(longSun.LonPos-longApogeeMean.LonPos, -180.0, 180.0)
	if err != nil {
		return Zero, fmt.Errorf("error in calculation %v", err)
	}
	sin2Diff := math.Sin(mathextra.DegToRad(2 * diff))
	factor2 := math.Sin(mathextra.DegToRad(2 * (diff - 11.726*sin2Diff)))
	sin6Diff := math.Sin(mathextra.DegToRad(6 * diff))
	factor3 := (8.8 / 60.0) * sin6Diff
	corrFactor := factor1*factor2 + factor3
	valueInRange, err := ValueToRange(longApogeeMean.LonPos+corrFactor, 0.0, 360.0)
	if err != nil {
		return Zero, fmt.Errorf("error in calculation %v", err)
	}

	return valueInRange, nil
}

type PointRangeCalculation struct {
	sePointCalc se.SwephPointPosCalculator
}

func NewPointRangeCalculation() PointRangeCalculator {
	ppc := se.NewSwephPointPosCalculation()
	return PointRangeCalculation{ppc}
}

func (prc PointRangeCalculation) CalcPointRange(request domain.PointRangeRequest) ([]domain.PointRangeResult, error) {

	reqPoint := request.Point
	allPoints := domain.AllChartPoints()
	index := allPoints[reqPoint].CalcId

	flags := SeFlags(request.Coord, request.ObsPos, request.Ayanamsha)
	// TODO handle topocentric
	var rangePositions []domain.PointRangeResult
	var resultIndex int
	if request.Position {
		if request.MainValue {
			resultIndex = 0
		} else {
			resultIndex = 4
		}
	} else {
		if request.MainValue {
			resultIndex = 1
		} else {
			resultIndex = 5
		}
	}
	// TODO handle RADV/Distance
	for i := request.JdStart; i <= request.JdEnd; i += request.Interval {
		sePos, err := prc.sePointCalc.CalcPointPos(i, index, flags)
		if err != nil {
			return rangePositions, err
		}
		calcValue := sePos[resultIndex]
		rangePositions = append(rangePositions, domain.PointRangeResult{Jd: i, Value: calcValue})
	}
	return rangePositions, nil
}

type HousePosCalculation struct {
	seHouseCalc se.SwephHousePosCalculator
	seEpsCalc   se.SwephEpsilonCalculator
	seHorCalc   se.SwephHorPosCalculator
}

func NewHousePosCalculation() HousePosCalculator {
	shpc := se.NewSwephHousePosCalculation()
	sec := se.NewSwephEpsilonCalculation()
	shc := se.NewSwephHorPosCalculation()
	return HousePosCalculation{shpc, sec, shc}
}

func (hpc HousePosCalculation) CalcHousePos(request domain.HousePosRequest) ([]domain.HousePosResult, []domain.HousePosResult, error) {

	allHouseSystems := domain.AllHouseSystems()
	currentSystem := allHouseSystems[request.HouseSys]
	hSysId := currentSystem.Code

	var cuspPos = make([]domain.HousePosResult, 37)
	var mcAscPos = make([]domain.HousePosResult, 4)
	eclFlags := domain.SeflgSwieph + domain.SeflgSpeed
	//	equFlags := domain.SeflgSwieph + domain.SeflgSpeed + domain.SeflgEquatorial
	cuspsEcl, otherPointsEcl, errEcl := hpc.seHouseCalc.CalcHousePos(hSysId, request.JdUt, request.GeoLong, request.GeoLat, eclFlags)
	if errEcl != nil {
		return cuspPos, mcAscPos, errEcl
	}
	/*	cuspsEqu, otherPointsEqu, errEqu := hpc.seHouseCalc.CalcHousePos(request.HouseSys, request.JdUt, request.GeoLong, request.GeoLat, equFlags)
		if errEqu != nil {
			return cuspPos, mcAscPos, errEqu
		}*/
	trueObliquity := true // use true obliquity (corrected for nutation)
	obliquity, errObl := hpc.seEpsCalc.CalcEpsilon(request.JdUt, trueObliquity)
	if errObl != nil {
		return cuspPos, mcAscPos, errObl
	}
	nrOfCuspValues := len(cuspsEcl)
	lat := 0.0

	for i := 1; i < nrOfCuspValues; i++ { // start with index 1, as the SE does the same
		ra, decl := conversion.ChangeEclToEqu(cuspsEcl[i], lat, obliquity)
		height := 0.0
		horFlags := domain.SeflgEquatorial
		posHor := hpc.seHorCalc.CalcHorPos(request.JdUt, request.GeoLong, request.GeoLat, height, ra, decl, horFlags)
		cuspPos[i] = domain.HousePosResult{
			LonPos:   cuspsEcl[i],
			RaPos:    ra,
			DeclPos:  decl,
			AzimPos:  posHor[0],
			AltitPos: posHor[1],
		}
	}

	mcAscPos[0] = hpc.createHousePosResult(otherPointsEcl[0], lat, obliquity, request.JdUt, request.GeoLong, request.GeoLat) // Ascendant
	mcAscPos[1] = hpc.createHousePosResult(otherPointsEcl[1], lat, obliquity, request.JdUt, request.GeoLong, request.GeoLat) // MC
	mcAscPos[2] = hpc.createHousePosResult(otherPointsEcl[3], lat, obliquity, request.JdUt, request.GeoLong, request.GeoLat) // Vertex
	mcAscPos[3] = hpc.createHousePosResult(otherPointsEcl[4], lat, obliquity, request.JdUt, request.GeoLong, request.GeoLat) // East point

	return cuspPos, mcAscPos, nil
}

func (hpc HousePosCalculation) createHousePosResult(position, lat, obliquity, jd, geoLong, geoLat float64) domain.HousePosResult {
	ra, decl := conversion.ChangeEclToEqu(position, lat, obliquity)
	height := 0.0
	horFlags := domain.SeflgEquatorial
	posHor := hpc.seHorCalc.CalcHorPos(jd, geoLong, geoLat, height, ra, decl, horFlags)
	return domain.HousePosResult{
		LonPos:   position,
		RaPos:    ra,
		DeclPos:  decl,
		AzimPos:  posHor[0],
		AltitPos: posHor[1],
	}
}
