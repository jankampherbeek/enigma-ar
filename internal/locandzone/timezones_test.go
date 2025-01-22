/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package locandzone

import (
	"enigma-ar/domain"
	"testing"
)

func TestActualTimeZone(t *testing.T) {
	tzHandler := NewTimeZoneHandling()
	dateTime := domain.DateTimeHms{
		Year:  1953,
		Month: 1,
		Day:   29,
		Hour:  8,
		Min:   37,
		Sec:   30,
		Greg:  true,
		Dst:   0,
		TZone: 0,
	}
	tzIndication := "Europe/Amsterdam"
	expectedName := "CET"
	expectedOffset := 3600
	resultName, resultOffset, err := tzHandler.ActualTimeZone(dateTime, tzIndication)
	if err != nil {
		t.Errorf("ActualTimeZone returned an unexpected error: %s", err.Error())
	}
	if resultName != expectedName {
		t.Errorf("Actual name for time zone is %s, expected %s", resultName, expectedName)
	}
	if expectedOffset != resultOffset {
		t.Errorf("Actual offset for time zone is %d, expected %d", resultOffset, expectedOffset)
	}
}

func TestActualTimeZoneForMAT(t *testing.T) {
	tzHandler := NewTimeZoneHandling()
	dateTime := domain.DateTimeHms{
		Year:  1908,
		Month: 1,
		Day:   29,
		Hour:  8,
		Min:   37,
		Sec:   30,
		Greg:  true,
		Dst:   0,
		TZone: 0,
	}
	tzIndication := "Europe/Amsterdam"
	expectedName := "AMT"
	expectedOffset := 19*60 + 32 // 0:19:32
	resultName, resultOffset, err := tzHandler.ActualTimeZone(dateTime, tzIndication)
	if err != nil {
		t.Errorf("ActualTimeZone for AMT returned an unexpected error: %s", err.Error())
	}
	if resultName != expectedName {
		t.Errorf("Actual name for time zone is %s, expected %s", resultName, expectedName)
	}
	if expectedOffset != resultOffset {
		t.Errorf("Actual offset for AMT is %d, expected %d", resultOffset, expectedOffset)
	}
}

func TestActualTimeZoneForBerlinLMT(t *testing.T) {
	tzHandler := NewTimeZoneHandling()
	dateTime := domain.DateTimeHms{
		Year:  1892,
		Month: 4,
		Day:   2,
		Hour:  12,
		Min:   0,
		Sec:   0,
		Greg:  true,
		Dst:   0,
		TZone: 0,
	}
	tzIndication := "Europe/Berlin"
	expectedName := "LMT"
	expectedOffset := 3208 // 0:53:28
	resultName, resultOffset, err := tzHandler.ActualTimeZone(dateTime, tzIndication)
	if err != nil {
		t.Errorf("ActualTimeZone for LMT returned an unexpected error: %s", err.Error())
	}
	if resultName != expectedName {
		t.Errorf("Actual name for time zone is %s, expected %s", resultName, expectedName)
	}
	if expectedOffset != resultOffset {
		t.Errorf("Actual offset for LMT is %d, expected %d", resultOffset, expectedOffset)
	}
}
