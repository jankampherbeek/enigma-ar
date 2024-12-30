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
	internal "enigma-ar/internal/calc/mathextra"
	"enigma-ar/internal/se"
	"fmt"
	"math"
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

// FullChartCalculator calculates a full chart with celestial points and houses.
type FullChartCalculator interface {
	CalcFullChart(request domain.FullChartRequest) (domain.FullChartResponse, error)
}

type PointPosCalculation struct {
	sePointCalc  se.SePointPosCalculator
	seHorPosCalc se.SeHorPosCalculator
}

func NewPointPosCalculation() PointPosCalculator {
	ppc := se.NewSePointPosCalculation()
	hpc := se.NewSeHorPosCalculation()
	return PointPosCalculation{ppc, hpc}
}

// CalcPointPos calculates fully defined positions for one or more celestial points
// PRE MinJdGeneral () <= request.JdUt <= MaxJdGeneral ()
// PRE MinGeoLong <= request.GeoLong < MaxGeoLong
// PRE MinGeoLat <= request.GeoLat < MaxGeoLat
// POST : if no error occurred returns positions for the given points, otherwise returns empty slice and error
func (calc PointPosCalculation) CalcPointPos(request domain.PointPositionsRequest) ([]domain.PointPosResult, error) {
	emptyPositions := make([]domain.PointPosResult, 0)
	jdUt := request.JdUt
	geoLong := request.GeoLong
	geoLat := request.GeoLat
	if jdUt < domain.MinJdGeneral || jdUt > domain.MaxJdGeneral {
		return emptyPositions, fmt.Errorf("JdUt %f is out of range", jdUt)
	}
	if geoLong < domain.MinGeoLong || geoLong > domain.MaxGeoLong {
		return emptyPositions, fmt.Errorf("GeoLong %f is out of range", geoLong)
	}
	if geoLat < domain.MinGeoLat || geoLat > domain.MaxGeoLat {
		return emptyPositions, fmt.Errorf("GeoLat %f is out of range", geoLat)
	}
	positions := make([]domain.PointPosResult, 0)
	eclFlags := SeFlags(domain.CoordEcliptical, request.ObsPos, request.Tropical)
	equFlags := SeFlags(domain.CoordEquatorial, request.ObsPos, request.Tropical)

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
				return emptyPositions, fmt.Errorf("calc point positions failed for %v", point)
			}
			positions = append(positions, position)
		case domain.CalcElements:
			// handle elements
		case domain.CalcFormula:
			position, err := calc.calcPointPosViaFormula(calcId, point, jdUt, eclFlags, equFlags)
			if err != nil {
				return emptyPositions, fmt.Errorf("calc point positions failed for %v", point)
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
	ayanOffset := 0.0
	// TODO correct ayanoffset based on selection of Ayanamsha

	if request.ProjType == domain.ProjTypeOblique { // handle oblique longitude
		olc := NewObliqueLongCalculation()
		newPositions, err := olc.calcObliqueLongitudes(positions, request.Armc, request.Obliquity, geoLat, ayanOffset)
		if err != nil {
			return emptyPositions, fmt.Errorf("calc oblique positions failed")
		}
		positions = newPositions
	}
	return positions, nil
}

func (calc PointPosCalculation) calcPointPosViaSe(index int, point domain.ChartPoint, jdUt float64,
	eclFlags, equFlags int, geoLong, geoLat float64) (domain.PointPosResult, error) {

	var position domain.PointPosResult
	posEcl, errEcl := calc.sePointCalc.SeCalcPointPos(jdUt, index, eclFlags)
	if errEcl != nil {
		return position, errEcl
	}
	posEqu, errEqu := calc.sePointCalc.SeCalcPointPos(jdUt, index, equFlags)
	if errEqu != nil {
		return position, errEqu
	}
	height := 0.0
	pointRa := posEqu[0]
	pointDecl := posEqu[1]
	horFlags := domain.SeflgEquatorial
	posHor := calc.seHorPosCalc.SeCalcHorPos(jdUt, geoLong, geoLat, height, pointRa, pointDecl, horFlags)
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
	fmt.Printf("point %v lonPos %f\n", point, posEcl[0])

	return position, nil
}

func (calc PointPosCalculation) calcPointPosViaFormula(index int, point domain.ChartPoint, jdUt float64,
	eclFlags, equFlags int) (domain.PointPosResult, error) {

	var emptyPosition domain.PointPosResult
	var position domain.PointPosResult
	var eclLong, eclSpeed float64
	switch point {
	case domain.PersephoneCarteret:
		eclLong = calc.calcCarteretHypPlanet(jdUt, StartPointPersephoneCart, YearlySpeedPersephoneCart)
		eclSpeed = YearlySpeedPersephoneCart / domain.TropicalYearInDays
	case domain.VulcanusCarteret:
		eclLong = calc.calcCarteretHypPlanet(jdUt, StartPointVulcanusCart, YearlySpeedVulcanusCart)
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
		eclLong = result
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
	indexSun := domain.AllChartPoints()[domain.SeSun].CalcId
	indexApogeeMean := domain.AllChartPoints()[domain.ApogeeMean].CalcId
	longSun, err := calc.calcPointPosViaSe(indexSun, domain.SeSun, jdUt, flagsEcl, equFlags, geoLong, geoLat)
	longApogeeMean, err := calc.calcPointPosViaSe(indexApogeeMean, domain.ApogeeMean, jdUt, flagsEcl, equFlags, geoLong, geoLat)
	fmt.Printf("indexApogeeMean %d, flagsEcl %d, jdUt %f\n", indexApogeeMean, flagsEcl, jdUt)

	diff, err := valueToRange(longSun.LonPos-longApogeeMean.LonPos, -180.0, 180.0)
	if err != nil {
		return Zero, fmt.Errorf("error in calculation %v", err)
	}
	sin2Diff := math.Sin(internal.DegToRad(2 * diff))
	factor2 := math.Sin(internal.DegToRad(2 * (diff - 11.726*sin2Diff)))
	sin6Diff := math.Sin(internal.DegToRad(6 * diff))
	factor3 := (8.8 / 60.0) * sin6Diff
	corrFactor := factor1*factor2 + factor3
	valueInRange, err := valueToRange(longApogeeMean.LonPos+corrFactor, 0.0, 360.0)
	if err != nil {
		return Zero, fmt.Errorf("error in calculation %v", err)
	}
	//fmt.Printf("longSun: %f, longApogeeMean: %f, corrFactor: %f, valueInRange: %f\n", longSun.LonPos, longApogeeMean.LonPos, corrFactor, valueInRange)

	return valueInRange, nil
}

type PointRangeCalculation struct {
	sePointCalc se.SePointPosCalculator
}

func NewPointRangeCalculation() PointRangeCalculator {
	ppc := se.NewSePointPosCalculation()
	return PointRangeCalculation{ppc}
}

func (prc PointRangeCalculation) CalcPointRange(request domain.PointRangeRequest) ([]domain.PointRangeResult, error) {

	reqPoint := request.Point
	allPoints := domain.AllChartPoints()
	index := allPoints[reqPoint].CalcId

	flags := SeFlags(request.Coord, request.ObsPos, request.Ayanamsha == 0)
	// TODO handle topocentric
	// TODO handle sidereal
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
		sePos, err := prc.sePointCalc.SeCalcPointPos(i, index, flags)
		if err != nil {
			return rangePositions, err
		}
		calcValue := sePos[resultIndex]
		rangePositions = append(rangePositions, domain.PointRangeResult{Jd: i, Value: calcValue}) // TODO improve appending
	}
	return rangePositions, nil
}

type HousePosCalculation struct {
	seHouseCalc se.SeHousePosCalculator
	seEpsCalc   se.SeEpsilonCalculator
	seHorCalc   se.SeHorPosCalculator
}

func NewHousePosCalculation() HousePosCalculator {
	shpc := se.NewSeHousePosCalculation()
	sec := se.NewSeEpsilonCalculation()
	shc := se.NewSeHorPosCalculation()
	return HousePosCalculation{shpc, sec, shc}
}

func (hpc HousePosCalculation) CalcHousePos(request domain.HousePosRequest) ([]domain.HousePosResult, []domain.HousePosResult, error) {

	allHouseSystems := domain.AllHouseSystems()
	currentSystem := allHouseSystems[request.HouseSys]
	hSysId := currentSystem.Code

	var cuspPos = make([]domain.HousePosResult, 37)
	var mcAscPos = make([]domain.HousePosResult, 10)
	eclFlags := domain.SeflgSwieph + domain.SeflgSpeed
	//	equFlags := domain.SeflgSwieph + domain.SeflgSpeed + domain.SeflgEquatorial
	cuspsEcl, otherPointsEcl, errEcl := hpc.seHouseCalc.SeCalcHousePos(hSysId, request.JdUt, request.GeoLong, request.GeoLat, eclFlags)
	if errEcl != nil {
		return cuspPos, mcAscPos, errEcl
	}
	/*	cuspsEqu, otherPointsEqu, errEqu := hpc.seHouseCalc.SeCalcHousePos(request.HouseSys, request.JdUt, request.GeoLong, request.GeoLat, equFlags)
		if errEqu != nil {
			return cuspPos, mcAscPos, errEqu
		}*/
	trueEps := true // use true obliquity (corrected for nutation)
	eps, errEps := hpc.seEpsCalc.SeCalcEpsilon(request.JdUt, trueEps)
	if errEps != nil {
		return cuspPos, mcAscPos, errEps
	}
	nrOfCuspValues := len(cuspsEcl)
	lat := 0.0
	// TODO combined method for cusps and other mundane points
	for i := 1; i < nrOfCuspValues; i++ { // start with index 1, as the SE does the same
		ra, decl := conversion.ChangeEclToEqu(cuspsEcl[i], lat, eps)
		height := 0.0
		horFlags := domain.SeflgEquatorial
		posHor := hpc.seHorCalc.SeCalcHorPos(request.JdUt, request.GeoLong, request.GeoLat, height, ra, decl, horFlags)
		cuspPos[i] = domain.HousePosResult{
			LonPos:   cuspsEcl[i],
			RaPos:    ra,
			DeclPos:  decl,
			AzimPos:  posHor[0],
			AltitPos: posHor[1],
		}
	}
	for i := 0; i < len(mcAscPos); i++ {
		ra, decl := conversion.ChangeEclToEqu(otherPointsEcl[i], lat, eps)
		height := 0.0
		horFlags := domain.SeflgEquatorial
		posHor := hpc.seHorCalc.SeCalcHorPos(request.JdUt, request.GeoLong, request.GeoLat, height, ra, decl, horFlags)
		mcAscPos[i] = domain.HousePosResult{
			LonPos:   cuspsEcl[i],
			RaPos:    ra,
			DeclPos:  decl,
			AzimPos:  posHor[0],
			AltitPos: posHor[1],
		}
	}
	return cuspPos, mcAscPos, nil
}

type FullChartCalculation struct {
	ppc PointPosCalculator
	hpc HousePosCalculator
}

func NewFullChartCalculator() FullChartCalculator {
	ppc := NewPointPosCalculation()
	hpc := NewHousePosCalculation()
	return FullChartCalculation{ppc, hpc}
}

func (fcc FullChartCalculation) CalcFullChart(request domain.FullChartRequest) (domain.FullChartResponse, error) {

	var response domain.FullChartResponse
	pointsRequest := domain.PointPositionsRequest{
		Points:   request.Points,
		JdUt:     request.Jd,
		GeoLong:  request.GeoLong,
		GeoLat:   request.GeoLat,
		Coord:    request.CoordSys,
		ObsPos:   request.ObsPos,
		Tropical: request.Ayanamsha == 0,
	}
	pointsResult, pointsErr := fcc.ppc.CalcPointPos(pointsRequest)
	if pointsErr != nil {
		return response, pointsErr
	}
	houseRequest := domain.HousePosRequest{
		HouseSys: request.HouseSys,
		JdUt:     request.Jd,
		GeoLong:  request.GeoLong,
		GeoLat:   request.GeoLat,
	}
	housesResult, mundaneResult, mundaneErr := fcc.hpc.CalcHousePos(houseRequest)
	if mundaneErr != nil {
		return response, mundaneErr
	}
	// create response
	response = domain.FullChartResponse{
		Points:    pointsResult,
		Mc:        mundaneResult[1],
		Asc:       mundaneResult[0],
		Vertex:    mundaneResult[3],
		EastPoint: mundaneResult[4],
		Cusps:     housesResult,
	}

	return response, nil
}
