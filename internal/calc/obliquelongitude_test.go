/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	"enigma-ar/domain"
	"math"
	"testing"
)

// calcObliqueLongitudes(points []domain.PointPosResult, armc, obliquity, geoLat, ayanOffset float64) ([]domain.PointPosResult, error)
func TestCalcObliqueLongitude(t *testing.T) {

	armc := 12.356358154336363
	ayanOffset := 0.0
	obliquity := 23.448018383804666
	geoLat := 51.5
	points := []domain.PointPosResult{
		{
			Point:     0,
			LonPos:    232.42300189310427,
			LonSpeed:  0,
			LatPos:    -9.592715239942409e-05,
			LatSpeed:  0,
			RaPos:     0,
			RaSpeed:   0,
			DeclPos:   0,
			DeclSpeed: 0,
			RadvPos:   0,
			RadvSpeed: 0,
			AzimPos:   0,
			AltitPos:  0,
		},
		{
			Point:     2,
			LonPos:    30.43675414259534,
			LonSpeed:  0,
			LatPos:    -0.4161709354253986,
			LatSpeed:  0,
			RaPos:     0,
			RaSpeed:   0,
			DeclPos:   0,
			DeclSpeed: 0,
			RadvPos:   0,
			RadvSpeed: 0,
			AzimPos:   0,
			AltitPos:  0,
		},
		{
			Point:     9,
			LonPos:    136.5627442049836,
			LonSpeed:  0,
			LatPos:    7.557009105978577,
			LatSpeed:  0,
			RaPos:     0,
			RaSpeed:   0,
			DeclPos:   0,
			DeclSpeed: 0,
			RadvPos:   0,
			RadvSpeed: 0,
			AzimPos:   0,
			AltitPos:  0,
		},
	}
	expected := []domain.PointPosResult{
		{
			Point:     0,
			LonPos:    232.42290213148797,
			LonSpeed:  0,
			LatPos:    -9.592715239942409e-05,
			LatSpeed:  0,
			RaPos:     0,
			RaSpeed:   0,
			DeclPos:   0,
			DeclSpeed: 0,
			RadvPos:   0,
			RadvSpeed: 0,
			AzimPos:   0,
			AltitPos:  0,
		},
		{
			Point:     1,
			LonPos:    30.744547179416834,
			LonSpeed:  0,
			LatPos:    -0.4161709354253986,
			LatSpeed:  0,
			RaPos:     0,
			RaSpeed:   0,
			DeclPos:   0,
			DeclSpeed: 0,
			RadvPos:   0,
			RadvSpeed: 0,
			AzimPos:   0,
			AltitPos:  0,
		},
		{
			Point:     9,
			LonPos:    130.33001057246562,
			LonSpeed:  0,
			LatPos:    7.557009105978577,
			LatSpeed:  0,
			RaPos:     0,
			RaSpeed:   0,
			DeclPos:   0,
			DeclSpeed: 0,
			RadvPos:   0,
			RadvSpeed: 0,
			AzimPos:   0,
			AltitPos:  0,
		},
	}
	olc := NewObliqueLongCalculation()
	result, err := olc.calcObliqueLongitudes(points, armc, obliquity, geoLat, ayanOffset)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 3; i++ {
		if math.Abs(result[i].LonPos-expected[i].LonPos) >= 1e-8 {
			t.Errorf("CalcOblique Longitude returned %f, expected %f", result[i].LonPos, expected[i].LonPos)
		}
		if math.Abs(result[i].LatPos-expected[i].LatPos) >= 1e-8 {
			t.Errorf("CalcOblique Latitude returned %f, expected %f", result[i].LonPos, expected[i].LonPos)
		}
	}

	/*
		input

		request = {ObliqueLongitudeRequest} ObliqueLongitudeRequest { Armc = 12,356358154336363, Obliquity = 23,448018383804666, GeoLat = 51,5, CelPointCoordinates = System.Collections.Generic.List`1[Enigma.Domain.Dtos.NamedEclipticCoordinates], AyanamshaOffset = 0 }
		 Armc = {double} 12.356358154336363
		 AyanamshaOffset = {double} 0
		 CelPointCoordinates = {List<NamedEclipticCoordinates>} Count = 12
		  [0] = ...{ CelPoint = Sun, EclipticCoordinate = EclipticCoordinates { Longitude = 232,42300189310427, Latitude = -9,592715239942409E-05 } }
		  [1] = ...{ CelPoint = Moon, EclipticCoordinate = EclipticCoordinates { Longitude = 30,43675414259534, Latitude = -0,4161709354253986 } }
		  [2] = {NamedEclipticCoordinates} NamedEclipticCoordinates { CelPoint = Mercury, EclipticCoordinate = EclipticCoordinates { Longitude = 216,9576865573118, Latitude = 1,780835037657735 } }
		  [3] = {NamedEclipticCoordinates} NamedEclipticCoordinates { CelPoint = Venus, EclipticCoordinate = EclipticCoordinates { Longitude = 196,3844683356489, Latitude = 1,825362496978752 } }
		  [4] = {NamedEclipticCoordinates} NamedEclipticCoordinates { CelPoint = Mars, EclipticCoordinate = EclipticCoordinates { Longitude = 260,9489659733187, Latitude = -0,9280297233715826 } }
		  [5] = {NamedEclipticCoordinates} NamedEclipticCoordinates { CelPoint = Jupiter, EclipticCoordinate = EclipticCoordinates { Longitude = 269,8856092667219, Latitude = 0,06872534014187279 } }
		  [6] = {NamedEclipticCoordinates} NamedEclipticCoordinates { CelPoint = Saturn, EclipticCoordinate = EclipticCoordinates { Longitude = 155,26743966241523, Latitude = 1,4391042093967628 } }
		  [7] = {NamedEclipticCoordinates} NamedEclipticCoordinates { CelPoint = Uranus, EclipticCoordinate = EclipticCoordinates { Longitude = 89,92909217025462, Latitude = 0,19970670994628562 } }
		  [8] = {NamedEclipticCoordinates} NamedEclipticCoordinates { CelPoint = Neptune, EclipticCoordinate = EclipticCoordinates { Longitude = 194,12908509921593, Latitude = 1,524560281275531 } }
		  [9] = ...{ CelPoint = Pluto, EclipticCoordinate = EclipticCoordinates { Longitude = 136,5627442049836, Latitude = 7,557009105978577 } }
		  [10] = {NamedEclipticCoordinates} NamedEclipticCoordinates { CelPoint = TrueNode, EclipticCoordinate = EclipticCoordinates { Longitude = 34,959957039085296, Latitude = 0 } }
		  [11] = {NamedEclipticCoordinates} NamedEclipticCoordinates { CelPoint = Chiron, EclipticCoordinate = EclipticCoordinates { Longitude = 238,22461117628117, Latitude = 3,159363255166229 } }
		 EqualityContract = {RuntimeType} Enigma.Domain.Requests.ObliqueLongitudeRequest
		 GeoLat = {double} 51.5
		 Obliquity = {double} 23.448018383804666
	*/

	/*
		output

			Sun,232.42290213148797,Enigma.Domain.Dtos.NamedEclipticLongitude
			Moon,30.744547179416834,Enigma.Domain.Dtos.NamedEclipticLongitude
			Mercury,218.49263250004265,Enigma.Domain.Dtos.NamedEclipticLongitude
			Venus,197.29228618319931,Enigma.Domain.Dtos.NamedEclipticLongitude
			Mars,259.82966113823733,Enigma.Domain.Dtos.NamedEclipticLongitude
			Jupiter,269.96799723186774,Enigma.Domain.Dtos.NamedEclipticLongitude
			Saturn,154.72264979643595,Enigma.Domain.Dtos.NamedEclipticLongitude
			Uranus,89.689542536385176,Enigma.Domain.Dtos.NamedEclipticLongitude
			Neptune,194.8139205547628,Enigma.Domain.Dtos.NamedEclipticLongitude
			Pluto,130.33001057246562,Enigma.Domain.Dtos.NamedEclipticLongitude
			TrueNode,34.959957039085296,Enigma.Domain.Dtos.NamedEclipticLongitude
			Chiron,241.78809945007299,Enigma.Domain.Dtos.NamedEclipticLongitude

	*/
}

func TestCalculateSouthPointHappyFlow(t *testing.T) {
	armc := 331.883333333333
	obliquity := 23.449614320676233 // mean obliquity
	geoLat := 48.8333333333333
	expectedLong := 318.50043580207006
	expectedLat := -27.562090280566338

	olc := NewObliqueLongCalculation()
	resultLong, resultLat, err := olc.calculateSouthPoint(armc, obliquity, geoLat)
	if err != nil {
		t.Error(err)
	}
	if math.Abs(resultLong-expectedLong) > 1e-3 {
		t.Errorf("Error in longitude of south point, expected %f, got %f", expectedLong, resultLong)
	}
	if math.Abs(resultLat-expectedLat) > 1e-3 {
		t.Errorf("Error in latitude of south point, expected %f, got %f", expectedLat, resultLat)
	}
}

func TestCalculateSouthPointSouthernHemisphere(t *testing.T) {
	armc := 331.883333333333
	obliquity := 23.449614320676233 // mean obliquity
	geoLat := -48.8333333333333
	expectedLong := 174.53494810489755
	expectedLat := -48.16467239725159

	olc := NewObliqueLongCalculation()
	resultLong, resultLat, err := olc.calculateSouthPoint(armc, obliquity, geoLat)
	if err != nil {
		t.Error(err)
	}
	if math.Abs(resultLong-expectedLong) > 1e-3 {
		t.Errorf("Error in longitude of south point for southern hemisphere, expected %f, got %f", expectedLong, resultLong)
	}
	if math.Abs(resultLat-expectedLat) > 1e-3 {
		t.Errorf("Error in latitude of south point for southern hemisphere, expected %f, got %f", expectedLat, resultLat)
	}

}
