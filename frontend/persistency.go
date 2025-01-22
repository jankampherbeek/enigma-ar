/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"enigma-ar/api"
	"enigma-ar/domain"
	"fmt"
	"strings"
)

// Prepare chart data to be persisted or read data from locandzone

func persistCurrentChart() {
	dv := GetDataVaultRadix()
	request := dv.Request
	meta := dv.Meta
	persChart := CreatePersistableChartData(meta)
	persDateLoc := CreatePersistableDateLocation(meta, request)
	pkCharts, pkDateLoc, err := api.WriteChart(persChart, persDateLoc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("PK's : %d %d", pkCharts, pkDateLoc)
	// TODO log pk's

}

func CreatePersistableChartData(meta domain.FullChartMeta) domain.PersistableChart {
	allCats := domain.AllChartCats()
	catText := allCats[meta.Category].TextId
	chart := domain.PersistableChart{
		Id:          -1,
		Name:        meta.Name,
		Description: meta.Description,
		Category:    catText,
	}
	return chart
}

func CreatePersistableDateLocation(meta domain.FullChartMeta, request domain.FullChartRequest) domain.PersistableDateLocation {
	allRatings := domain.AllRatings()
	ratingTxt := allRatings[meta.Rating].TextId
	geoLongSign := "+"
	if request.GeoLong < 0.0 {
		geoLongSign = "-"
	}
	geoLongTxt := geoLongSign + strings.Split(meta.GeoLong, " ")[0]
	geoLatSign := "+"
	if request.GeoLat < 0.0 {
		geoLatSign = "-"
	}
	geoLatTxt := geoLatSign + strings.Split(meta.GeoLat, " ")[0]

	allCalendars := domain.AllCalendars()
	calTxt := allCalendars[meta.Calendar].TextId
	allTimeZones := domain.AllTimeZones()
	var zoneTxt = allTimeZones[meta.TimeZone].TextId
	if meta.TimeZone == domain.TzLmt {
		zoneTxt += " " + meta.GeoLongLmt
	}
	if meta.Dst == true {
		zoneTxt += " " + "v_input_radix_dst"
	}
	dateLoc := domain.PersistableDateLocation{
		Id:           -1,
		ChartId:      -1,
		Source:       meta.Source,
		NameLocation: meta.LocationName + " " + geoLongTxt + " " + geoLatTxt,
		Rating:       ratingTxt,
		GeoLong:      request.GeoLong,
		GeoLat:       request.GeoLat,
		DateText:     meta.Date + " " + calTxt,
		TimeText:     meta.Time + " " + zoneTxt,
		Jd:           request.Jd,
	}
	return dateLoc
}
