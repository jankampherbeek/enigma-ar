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
	"enigma-ar/domain/references"
	"strconv"
	"strings"
)

type DateValidator interface {
	CheckDate(date string, cal references.Calendar) (bool, int, int, int)
}

type TimeValidator interface {
	CheckTime(time string) (bool, int, int, int)
}

type GeoLongValidator interface {
	CheckGeoLong(geoLong string, lang string) (bool, float64)
}

type GeoLatValidator interface {
	CheckGeoLat(geoLat string, lang string) (bool, float64)
}

// -------------------------------------------------------------------------

type DateValidation struct{}

func NewDateValidator() DateValidator {
	return DateValidation{}
}

func (v DateValidation) CheckDate(date string, cal references.Calendar) (bool, int, int, int) {
	var dateValid bool
	var year, month, day int
	jdApi := api.NewJulDayService()
	revJdApi := api.NewRevJulDayService()
	items, err := StringToInts(date, "/")
	if err == nil && len(items) == 3 {
		year = items[0]
		month = items[1]
		day = items[2]
		var ut float64
		greg := true
		if cal != references.CalGregorian {
			greg = false
		}
		if cal == references.CalJulianBCE {
			year--
		}
		jdReq := &domain.DateTime{Year: year, Month: month, Day: day, Ut: ut, Greg: greg}
		jd := jdApi.JulDay(jdReq)
		y, m, d, _ := revJdApi.RevJulDay(jd, cal) // ignore ut
		dateValid = year == y && month == m && day == d
	}
	return dateValid, year, month, day
}

// --------------------------------------------------------------------------
type TimeValidation struct{}

func NewTimeValidator() TimeValidator {
	return TimeValidation{}
}

func (v TimeValidation) CheckTime(time string) (bool, int, int, int) {
	timeValid := true
	var hour, minute, second int
	items, err := StringToInts(time, ":")
	if err == nil && (len(items) == 3 || len(items) == 2) {
		if len(items) == 2 { // add zero for seconds
			items = append(items, 0)
		}
		hour = items[0]
		minute = items[1]
		second = items[2]
		if hour < 0 || hour > 23 || minute < 0 || minute > 59 || second < 0 || second > 59 {
			timeValid = false
		}
	} else {
		timeValid = false
	}
	return timeValid, hour, minute, second
}

// --------------------------------------------------------------------------

type GeoLongValidation struct {
}

func NewGeoLongValidator() GeoLongValidator {
	return GeoLongValidation{}
}

func (v GeoLongValidation) CheckGeoLong(s string, lang string) (bool, float64) {
	gLongValid := true
	var geoLong float64
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		gLongValid = false
	}
	if gLongValid {
		valueTxt := parts[0]
		dir := strings.ToUpper(parts[1])
		var dirFactor = 1.0
		if (lang == "en" || lang == "nl" || lang == "ge") && dir == "W" {
			dirFactor = -1.0
		}
		if lang == "fr" && (dir == "o" || dir == "O") {
			dirFactor = -1.0
		}
		items, err := StringToInts(valueTxt, ":")
		if err == nil && (len(items) == 3 || len(items) == 2) {
			if len(items) == 2 {
				items = append(items, 0)
			}
			d := items[0]
			m := items[1]
			s := items[2]
			if d < 0 || d > 179 || m < 0 || m > 59 || s < 0 || s > 59 {
				gLongValid = false
			} else {
				geoLong = (float64(d) + float64(m)/60.0 + float64(s)/3600.0) * dirFactor
			}
		} else {
			gLongValid = false
		}
	}
	return gLongValid, geoLong
}

// --------------------------------------------------------------------------

type GeoLatValidation struct {
}

func NewGeoLatValidator() GeoLatValidator {
	return GeoLatValidation{}
}

func (v GeoLatValidation) CheckGeoLat(s string, lang string) (bool, float64) {
	gLatValid := true
	var geoLat float64
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		gLatValid = false
	}
	if gLatValid {
		valueTxt := parts[0]
		dir := strings.ToUpper(parts[1])
		var dirFactor = 1.0
		if dir != "N" { // 'N' is used for north in all supported languages
			dirFactor = -1.0
		}

		items, err := StringToInts(valueTxt, ":")
		if err == nil && (len(items) == 3 || len(items) == 2) {
			if len(items) == 2 {
				items = append(items, 0)
			}
			d := items[0]
			m := items[1]
			s := items[2]
			if d < 0 || d > 89 || m < 0 || m > 59 || s < 0 || s > 59 {
				gLatValid = false
			} else {
				geoLat = (float64(d) + float64(m)/60.0 + float64(s)/3600.0) * dirFactor
			}
		} else {
			gLatValid = false
		}
	}
	return gLatValid, geoLat
}

// --------------------------------------------------------------------------

func StringToInts(s string, sep string) ([]int, error) {
	txtItems := strings.Split(s, sep)
	var err error
	var valueItems []int
	for i, item := range txtItems {
		txtItems[i] = strings.TrimSpace(item)
		value, err := strconv.Atoi(txtItems[i])
		if err == nil {
			valueItems = append(valueItems, value)
		}
	}
	return valueItems, err
}
