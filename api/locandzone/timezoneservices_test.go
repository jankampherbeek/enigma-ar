/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package apilocandzone

import (
	"enigma-ar/domain"
	"math"
	"testing"
)

func TestActualTimeZone(t *testing.T) {
	tzs := NewTimeZoneService()
	tzIndication := "abc"
	dateTime := domain.DateTimeHms{
		Year:  2024,
		Month: 1,
		Day:   23,
		Hour:  20,
		Min:   4,
		Sec:   50,
		Greg:  true,
		Dst:   0,
		TZone: 0,
	}
	zoneInfo, err := tzs.ActualTimeZone(dateTime, tzIndication)
	if err == nil {
		t.Errorf("ActualTimezone should have returned an error for an invalid tzIndication")
	}
	if zoneInfo.ZoneName != "" {
		t.Errorf("ActualTimezone should have returned an empty string as name in the zoneInfo for an invalid tzIndication")
	}
	if math.Abs(zoneInfo.Offset) > 1e-8 {
		t.Errorf("ActualTimezone should have returned zero zoneInfo for an invalid tzIndication")
	}

}
