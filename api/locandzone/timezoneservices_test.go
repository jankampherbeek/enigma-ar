/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package apilocandzone

import (
	"enigma-ar/domain"
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
	name, offset, err := tzs.ActualTimeZone(dateTime, tzIndication)
	if err == nil {
		t.Errorf("ActualTimeZone should have returned an error for an invalid tzIndication")
	}
	if name != "" {
		t.Errorf("ActualTimeZone should have returned an empty string as name for an invalid tzIndication")
	}
	if offset != 0 {
		t.Errorf("ActualTimeZone should have returned zero offset for an invalid tzIndication")
	}

}
