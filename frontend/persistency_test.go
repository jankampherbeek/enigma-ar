/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"enigma-ar/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePersistableChartData(t *testing.T) {
	meta := createMeta()
	expected := domain.PersistableChart{
		Id:          -1,
		Name:        "Test name",
		Description: "Some description",
		Category:    "r_cc_male",
	}
	result := CreatePersistableChartData(meta)

	if result != expected {
		t.Errorf("CreatePersistableChartData() returned %+v, expected %+v", result, expected)
	}
}

func TestCreatePersistableDateLocation(t *testing.T) {
	meta := createMeta()
	request := createRequest()
	result := CreatePersistableDateLocation(meta, request)
	expected := domain.PersistableDateLocation{
		Id:           -1,
		ChartId:      -1,
		Source:       "Text for source",
		NameLocation: "Name of location +23:30 -45:45",
		Rating:       "r_rr_aa",
		GeoLong:      23.5,
		GeoLat:       -45.75,
		DateText:     "2024/11/17 r_cal_gregorian",
		TimeText:     "15:40:30 r_tz_cet",
		Jd:           12345.678,
	}
	assert.Equal(t, expected.Id, result.Id, 1e-9)
	assert.Equal(t, expected.ChartId, result.ChartId)
	assert.Equal(t, expected.Source, result.Source)
	assert.Equal(t, expected.NameLocation, result.NameLocation)
	assert.Equal(t, expected.Rating, result.Rating)
	assert.InDelta(t, expected.GeoLong, result.GeoLong, 1e-9)
	assert.InDelta(t, expected.GeoLat, result.GeoLat, 1e-9)
	assert.Equal(t, expected.DateText, result.DateText)
	assert.Equal(t, expected.TimeText, result.TimeText)
	assert.InDelta(t, expected.Jd, result.Jd, 1e-9)

}

func createMeta() domain.FullChartMeta {
	meta := domain.FullChartMeta{
		Name:         "Test name",
		Description:  "Some description",
		Category:     domain.CatMale,
		Rating:       domain.RatingAA,
		Source:       "Text for source",
		LocationName: "Name of location",
		GeoLat:       "45:45 z",
		GeoLong:      "23:30  e",
		Date:         "2024/11/17",
		Calendar:     domain.CalGregorian,
		Time:         "15:40:30",
		TimeZone:     domain.TzCet,
		Dst:          false,
		GeoLongLmt:   "",
	}
	return meta
}

func createRequest() domain.FullChartRequest {
	request := domain.FullChartRequest{
		Points:    nil,
		HouseSys:  domain.HousesAlcabitius,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Jd:        12345.678,
		GeoLong:   23.5,
		GeoLat:    -45.75,
	}
	return request
}
