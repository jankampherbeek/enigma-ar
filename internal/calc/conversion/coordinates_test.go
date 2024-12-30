/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package conversion

import (
	"math"
	"testing"
)

func TestDeclinationToLongitudeHappyFlow(t *testing.T) {
	//longitude := 223.0
	obliquity := 23.437101628
	declination := -15.09002104
	expected1 := 220.884444444444
	expected2 := 319.115555555556
	result := DeclinationToLongitude(obliquity, declination)
	if math.Abs(result-expected1) > 1e-8 && math.Abs(result-expected2) > 1e-8 {
		t.Fatalf("Expected longitude %v or %v, got %v", expected1, expected2, result)
	}

}
