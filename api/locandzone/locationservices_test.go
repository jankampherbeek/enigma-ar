/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package apilocandzone

import "testing"

func TestCitiesCountryCodeWrongSize(t *testing.T) {
	countryCode := "ABC"
	locService := NewLocationService()
	_, err := locService.Cities(countryCode)
	if err == nil {
		t.Error("Expected error for wrong countryCode")
	}
}
