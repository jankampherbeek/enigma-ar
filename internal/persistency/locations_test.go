/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package persistency

import (
	"testing"
)

func TestCountries(t *testing.T) {
	handler := NewLocationHandling()
	result, err := handler.Countries()
	if err != nil {
		t.Fatal(err)
	}
	if len(result) < 250 {
		t.Errorf("countries should have returned at least 250 entries")
	}
	for _, country := range result {
		if country.Code == "PT" {
			if country.Name != "Portugal" {
				t.Errorf("country should have country 'Portugal' but was '%s'", country.Name)
			}
		}
	}
}

func TestCities(t *testing.T) {
	handler := NewLocationHandling()
	countryCode := "NL"
	expectedGeoLat := "52.21833"
	expectedGeoLong := "6.89583"
	expectedElevation := "45"
	expectedTz := "Europe/Amsterdam"
	expectedRegion := "Overijssel"

	result, err := handler.Cities(countryCode)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) < 1000 {
		t.Errorf("cities should have returned at least 1000 entries")
	}
	for _, city := range result {
		if city.Country != countryCode {
			t.Errorf("Expected country code '%v' but got '%v'", countryCode, city.Country)
		}
		if city.Name == "Enschede" {
			if city.GeoLat != expectedGeoLat {
				t.Errorf("Expected latitude '%v' but got '%v'", expectedGeoLat, city.GeoLat)
			}
			if city.GeoLong != expectedGeoLong {
				t.Errorf("Expected longitude '%v' but got '%v'", expectedGeoLong, city.GeoLong)
			}
			if city.Elevation != expectedElevation {
				t.Errorf("Expected elevation '%v' but got '%v'", expectedElevation, city.Elevation)
			}
			if city.IndicationTz != expectedTz {
				t.Errorf("Expected timezone indication '%v' but got '%v'", expectedTz, city.IndicationTz)
			}
			if city.Region != expectedRegion {
				t.Errorf("Expected region '%v' but got '%v'", expectedRegion, city.Region)
			}
		}

	}
}
