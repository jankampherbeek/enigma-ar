/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package apicalc

import (
	"enigma-ar/domain"
	"enigma-ar/internal/calc"
	"errors"
	"log/slog"
)

type FullChartServer interface {
	CalcFullChart(request domain.FullChartRequest) (domain.FullChartResponse, error)
}

type FullChartService struct {
	fcc calc.FullChartCalculator
}

func NewFullChartService() FullChartServer {
	return FullChartService{calc.NewFullChartCalculation()}
}

// CalcFullChart provides services for the calculation of a fully defined chart
// PRE Size of FullChartsRequest.Points > 0
// PRE MinJdGeneral < FullChartRequest.Jd < NaxJdGeneral (jd between -2946707.5 and 7865293.5)
// PRE MinGeoLong <= FullChartRequest.Geolong < MaxGeoLong (geolong between -180.0 and 180.0)
// PRE MinGeoLat < = FullChartRequest.GeoLat < MaxGeoLat (geolat between -90.0 and 90.0)
// POST No errors: returns calculated full chart response, otherwise returns empty full chart response
func (fcs FullChartService) CalcFullChart(request domain.FullChartRequest) (domain.FullChartResponse, error) {
	slog.Info("Start calculation of full chart")
	emptyResponse := domain.FullChartResponse{
		Points:    nil,
		Mc:        domain.HousePosResult{},
		Asc:       domain.HousePosResult{},
		Vertex:    domain.HousePosResult{},
		EastPoint: domain.HousePosResult{},
		Cusps:     nil,
	}
	if len(request.Points) <= 0 {
		slog.Error("points is empty")
		return emptyResponse, errors.New("points is empty")
	}
	if request.Jd < domain.MinJdGeneral || request.Jd > domain.MaxJdGeneral {
		slog.Error("jd is out of range")
		return emptyResponse, errors.New("jd is out of range")
	}
	if request.GeoLong < domain.MinGeoLong || request.GeoLong > domain.MaxGeoLong {
		slog.Error("geoLong is out of range")
		return emptyResponse, errors.New("geoLong is out of range")
	}
	if request.GeoLat < domain.MinGeoLat || request.GeoLat > domain.MaxGeoLat {
		slog.Error("geoLat is out of range")
		return emptyResponse, errors.New("geoLat is out of range")
	}
	slog.Info("Completed calculation of full chart")
	result, err := fcs.fcc.CalcFullChart(request)
	return result, err
}
