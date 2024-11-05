/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"enigma-ar/domain"
	"math"
	"testing"
)

// Tests for validators for dates are integration tests. The values are calculated in the backend and no fakes are used.

func TestCheckDateHappyFlow(t *testing.T) {
	validator := NewDateValidator()
	dateText := "1953/1/29"
	cal := domain.Calendar(domain.CalGregorian)
	ok, resultYear, resultMonth, resultDay := validator.CheckDate(dateText, cal)
	expYear := 1953
	expMonth := 1
	expDay := 29
	if ok != true {
		t.Errorf("Date check failed. Expected ok=%t, actual=%t", true, ok)
	}
	if resultYear != expYear || resultMonth != expMonth || resultDay != expDay {
		t.Errorf("DateValidator returns a wrong result for year, month, day: %d, %d, %d", resultYear, resultMonth, resultDay)
	}
}

func TestCheckDateLeapDay(t *testing.T) {
	validator := NewDateValidator()
	dateText := "2004/2/29"
	cal := domain.Calendar(domain.CalGregorian)
	ok, resultYear, resultMonth, resultDay := validator.CheckDate(dateText, cal)
	expYear := 2004
	expMonth := 2
	expDay := 29
	if ok != true {
		t.Errorf("Date check for leapday failed. Expected ok=%t, actual=%t", true, ok)
	}
	if resultYear != expYear || resultMonth != expMonth || resultDay != expDay {
		t.Errorf("DateValidator for leapday returns a wrong result for year, month, day: %d, %d, %d", resultYear, resultMonth, resultDay)
	}
}

func TestCheckDateError(t *testing.T) {
	validator := NewDateValidator()
	dateText := "dummy"
	cal := domain.Calendar(domain.CalGregorian)
	ok, _, _, _ := validator.CheckDate(dateText, cal)
	if ok == true {
		t.Errorf("Date check with error as input failed. Expected ok=%t, actual=%t", false, ok)
	}
}

func TestCheckDateNonLeapYear(t *testing.T) {
	validator := NewDateValidator()
	dateText := "2022/2/29"
	cal := domain.Calendar(domain.CalGregorian)
	ok, _, _, _ := validator.CheckDate(dateText, cal)
	if ok == true {
		t.Errorf("Date check with non-leapyear failed. Expected ok=%t, actual=%t", false, ok)
	}
}

func TestCheckDateJulianYear(t *testing.T) {
	validator := NewDateValidator()
	dateText := "1423/11/13"
	cal := domain.Calendar(domain.CalJulianCE)
	ok, resultYear, resultMonth, resultDay := validator.CheckDate(dateText, cal)
	expYear := 1423
	expMonth := 11
	expDay := 13
	if ok != true {
		t.Errorf("Date check with Julian Day failed. Expected ok=%t, actual=%t", true, ok)
	}
	if resultYear != expYear || resultMonth != expMonth || resultDay != expDay {
		t.Errorf("DateValidator for Julian Day returns a wrong result for year, month, day: %d, %d, %d", resultYear, resultMonth, resultDay)
	}
}

func TestCheckTimeHappyFlow(t *testing.T) {
	validator := NewTimeValidator()
	timeTxt := "5:13:47"
	ok, resultHour, resultMinute, resultSecond := validator.CheckTime(timeTxt)
	expHour := 5
	expMinute := 13
	expSecond := 47
	if ok != true {
		t.Errorf("Time check failed. Expected ok=%t, actual=%t", true, ok)
	}
	if resultHour != expHour || resultMinute != expMinute || resultSecond != expSecond {
		t.Errorf("TimeValidator returns a wrong result for hour, minute, second: %d, %d, %d", resultHour, resultMinute, resultSecond)
	}
}

func TestCheckNoSeconds(t *testing.T) {
	validator := NewTimeValidator()
	timeTxt := "5:13"
	ok, resultHour, resultMinute, resultSecond := validator.CheckTime(timeTxt)
	expHour := 5
	expMinute := 13
	expSecond := 0
	if ok != true {
		t.Errorf("Time check failed. Expected ok=%t, actual=%t", true, ok)
	}
	if resultHour != expHour || resultMinute != expMinute || resultSecond != expSecond {
		t.Errorf("TimeValidator returns a wrong result for hour, minute, second: %d, %d, %d", resultHour, resultMinute, resultSecond)
	}
}

func TestCheckTimeErrorInput(t *testing.T) {
	validator := NewTimeValidator()
	timeTxt := "dummy"
	ok, _, _, _ := validator.CheckTime(timeTxt)
	if ok == true {
		t.Errorf("Time check with erroneous input failed. Expected ok=%t, actual=%t", false, ok)
	}
}

func TestCheckTimeHourTooLarge(t *testing.T) {
	validator := NewTimeValidator()
	timeTxt := "26:1:2"
	ok, _, _, _ := validator.CheckTime(timeTxt)
	if ok == true {
		t.Errorf("Time check with hour too large failed. Expected ok=%t, actual=%t", false, ok)
	}
}

func TestCheckTimeMinuteTooSmall(t *testing.T) {
	validator := NewTimeValidator()
	timeTxt := "22:-1:2"
	ok, _, _, _ := validator.CheckTime(timeTxt)
	if ok == true {
		t.Errorf("Time check with minute too small failed. Expected ok=%t, actual=%t", false, ok)
	}
}

func TestCheckTimeSecondTooLarge(t *testing.T) {
	validator := NewTimeValidator()
	timeTxt := "22:1:60"
	ok, _, _, _ := validator.CheckTime(timeTxt)
	if ok == true {
		t.Errorf("Time check with second too large failed. Expected ok=%t, actual=%t", false, ok)
	}
}

func TestCheckGeoLongHappyFlow(t *testing.T) {
	validator := NewGeoLongValidator()
	gLongTxt := "20:30:40 e"
	lang := "en"
	ok, result := validator.CheckGeoLong(gLongTxt, lang)
	if ok != true {
		t.Errorf("GeoLong check failed. Expected ok=%t, actual=%t", true, ok)
	}
	expected := 20.511111111111
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("GeoLong check failed. Expected geoLong=%f, actual=%f", expected, result)
	}
}

func TestCheckGeoLongNoSeconds(t *testing.T) {
	validator := NewGeoLongValidator()
	gLongTxt := "20:30 e"
	lang := "en"
	ok, result := validator.CheckGeoLong(gLongTxt, lang)
	if ok != true {
		t.Errorf("GeoLong check without seconds failed. Expected ok=%t, actual=%t", true, ok)
	}
	expected := 20.5
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("GeoLong check without seconds failed. Expected geoLong=%f, actual=%f", expected, result)
	}
}

func TestCheckGeoLongWest(t *testing.T) {
	validator := NewGeoLongValidator()
	gLongTxt := "20:30:40 w"
	lang := "en"
	ok, result := validator.CheckGeoLong(gLongTxt, lang)
	if ok != true {
		t.Errorf("GeoLong check failed. Expected ok=%t, actual=%t", true, ok)
	}
	expected := -20.511111111111
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("GeoLong check for west failed. Expected geoLong=%f, actual=%f", expected, result)
	}
}

func TestCheckGeoLongFrench(t *testing.T) {
	validator := NewGeoLongValidator()
	gLongTxt := "20:30:40 o"
	lang := "fr"
	ok, result := validator.CheckGeoLong(gLongTxt, lang)
	if ok != true {
		t.Errorf("GeoLong check for French failed. Expected ok=%t, actual=%t", true, ok)
	}
	expected := -20.511111111111
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("GeoLong check for French failed. Expected geoLong=%f, actual=%f", expected, result)
	}
}

func TestCheckGeoLongError(t *testing.T) {
	validator := NewGeoLongValidator()
	gLongTxt := "dummy"
	lang := "en"
	ok, _ := validator.CheckGeoLong(gLongTxt, lang)
	if ok == true {
		t.Errorf("GeoLong check with erroneous input failed. Expected ok=%t, actual=%t", false, ok)
	}
}

func TestCheckGeoLongDegreeTooLarge(t *testing.T) {
	validator := NewGeoLongValidator()
	gLongTxt := "180:01:01 e"
	lang := "en"
	ok, _ := validator.CheckGeoLong(gLongTxt, lang)
	if ok == true {
		t.Errorf("GeoLong check with degree too large failed. Expected ok=%t, actual=%t", false, ok)
	}
}

func TestCheckGeoLongMinuteTooLarge(t *testing.T) {
	validator := NewGeoLongValidator()
	gLongTxt := "170:60:01 e"
	lang := "en"
	ok, _ := validator.CheckGeoLong(gLongTxt, lang)
	if ok == true {
		t.Errorf("GeoLong check with minute too large failed. Expected ok=%t, actual=%t", false, ok)
	}
}

func TestCheckGeoLongSecondTooLarge(t *testing.T) {
	validator := NewGeoLongValidator()
	gLongTxt := "10:0:60 w"
	lang := "ge"
	ok, _ := validator.CheckGeoLong(gLongTxt, lang)
	if ok == true {
		t.Errorf("GeoLong check with minute too large failed. Expected ok=%t, actual=%t", false, ok)
	}
}

func TestCheckGeoLatHappyFlow(t *testing.T) {
	validator := NewGeoLatValidator()
	gLatTxt := "20:30:40 n"
	lang := "en"
	ok, result := validator.CheckGeoLat(gLatTxt, lang)
	if ok != true {
		t.Errorf("GeoLat check failed. Expected ok=%t, actual=%t", true, ok)
	}
	expected := 20.511111111111
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("GeoLat check failed. Expected geoLat=%f, actual=%f", expected, result)
	}
}

func TestCheckGeoLatNoSeconds(t *testing.T) {
	validator := NewGeoLatValidator()
	gLatTxt := "20:30 n"
	lang := "en"
	ok, result := validator.CheckGeoLat(gLatTxt, lang)
	if ok != true {
		t.Errorf("GeoLat check without seconds failed. Expected ok=%t, actual=%t", true, ok)
	}
	expected := 20.5
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("GeoLat check without seconds failed. Expected geoLat=%f, actual=%f", expected, result)
	}
}

func TestCheckGeoLatSouth(t *testing.T) {
	validator := NewGeoLatValidator()
	gLatTxt := "20:30:40 s"
	lang := "en"
	ok, result := validator.CheckGeoLat(gLatTxt, lang)
	if ok != true {
		t.Errorf("GeoLat check failed. Expected ok=%t, actual=%t", true, ok)
	}
	expected := -20.511111111111
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("GeoLat check for west failed. Expected geoLat=%f, actual=%f", expected, result)
	}
}

func TestCheckGeoLatDutch(t *testing.T) {
	validator := NewGeoLatValidator()
	gLatTxt := "20:30:40 z"
	lang := "nl"
	ok, result := validator.CheckGeoLat(gLatTxt, lang)
	if ok != true {
		t.Errorf("GeoLat check for Dutch failed. Expected ok=%t, actual=%t", true, ok)
	}
	expected := -20.511111111111
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("GeoLat check for Dutch failed. Expected geoLat=%f, actual=%f", expected, result)
	}
}

func TestCheckGeoLatError(t *testing.T) {
	validator := NewGeoLatValidator()
	gLatTxt := "dummy"
	lang := "en"
	ok, _ := validator.CheckGeoLat(gLatTxt, lang)
	if ok == true {
		t.Errorf("GeoLat check with erroneous input failed. Expected ok=%t, actual=%t", false, ok)
	}
}

func TestCheckGeoLatDegreeTooLarge(t *testing.T) {
	validator := NewGeoLatValidator()
	gLatTxt := "90:01:01 n"
	lang := "en"
	ok, _ := validator.CheckGeoLat(gLatTxt, lang)
	if ok == true {
		t.Errorf("GeoLat check with degree too large failed. Expected ok=%t, actual=%t", false, ok)
	}
}

func TestCheckGeoLatMinuteTooLarge(t *testing.T) {
	validator := NewGeoLatValidator()
	gLatTxt := "170:60:01 e"
	lang := "en"
	ok, _ := validator.CheckGeoLat(gLatTxt, lang)
	if ok == true {
		t.Errorf("GeoLat check with minute too large failed. Expected ok=%t, actual=%t", false, ok)
	}
}

func TestCheckGeoLatSecondTooLarge(t *testing.T) {
	validator := NewGeoLatValidator()
	gLatTxt := "10:0:60 w"
	lang := "ge"
	ok, _ := validator.CheckGeoLat(gLatTxt, lang)
	if ok == true {
		t.Errorf("GeoLat check with minute too large failed. Expected ok=%t, actual=%t", false, ok)
	}
}
