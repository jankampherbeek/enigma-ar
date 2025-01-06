/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import "enigma-ar/domain"

// FullChartCalculator calculates a full chart with celestial points and houses.
type FullChartCalculator interface {
	CalcFullChart(request domain.FullChartRequest) (domain.FullChartResponse, error)
}

type FullChartCalculation struct {
	ppc PointPosCalculator
	hpc HousePosCalculator
}

func NewFullChartCalculation() FullChartCalculator {
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
