/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	domain "enigma-ar/domain"
	"enigma-ar/domain/references"
	"enigma-ar/internal/calc/conversion"
	"enigma-ar/internal/se"
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
func (calc PointPosCalculation) CalcPointPos(request domain.PointPositionsRequest) ([]domain.PointPosResult, error) {
	positions := make([]domain.PointPosResult, 0)
	eclFlags := SeFlags(references.CoordEcliptical, request.ObsPos, request.Tropical)
	equFlags := SeFlags(references.CoordEquatorial, request.ObsPos, request.Tropical)

	var allPoints = references.AllChartPoints()

	for i := 0; i < len(request.Points); i++ {
		reqPoint := request.Points[i]
		index := allPoints[reqPoint].CalcId
		posEcl, errEcl := calc.sePointCalc.SeCalcPointPos(request.JdUt, index, eclFlags)
		if errEcl != nil {
			return positions, errEcl
		}
		posEqu, errEqu := calc.sePointCalc.SeCalcPointPos(request.JdUt, index, equFlags)
		if errEqu != nil {
			return positions, errEqu
		}
		height := 0.0
		pointRa := posEqu[0]
		pointDecl := posEqu[1]
		horFlags := domain.SeflgEquatorial
		posHor := calc.seHorPosCalc.SeCalcHorPos(request.JdUt, request.GeoLong, request.GeoLat, height, pointRa, pointDecl, horFlags)
		positions = append(positions, domain.PointPosResult{
			Point:     reqPoint,
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
		})
	}
	return positions, nil
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
	allPoints := references.AllChartPoints()
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

	allHouseSystems := references.AllHouseSystems()
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
