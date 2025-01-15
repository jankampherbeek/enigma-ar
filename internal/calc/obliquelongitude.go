/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	"enigma-ar/domain"
	mathextra "enigma-ar/internal/calc/mathextra"
	"math"
)

// ObliqueLongCalculator provides calculations for oblique longitude ("true place") according to the School of Ram.
type ObliqueLongCalculator interface {
	calcObliqueLongitudes(points []domain.PointPosResult, armc, obliquity, geoLat, ayanOffset float64) ([]domain.PointPosResult, error)
	calculateSouthPoint(armc, obliquity, geoLat float64) (eclLong, eclLat float64, err error)
}

type ObliqueLongCalculation struct{}

func NewObliqueLongCalculation() ObliqueLongCalculator {
	return ObliqueLongCalculation{}
}

// calcObliqueLongitudes calculates oblique longitudes for the given request
func (olc ObliqueLongCalculation) calcObliqueLongitudes(points []domain.PointPosResult, armc, obliquity, geoLat, ayanOffset float64) ([]domain.PointPosResult, error) {

	emptyResult := make([]domain.PointPosResult, 0)
	result := make([]domain.PointPosResult, 0, len(points))
	spLong, spLat, err := olc.calculateSouthPoint(armc, obliquity, geoLat)
	if err != nil {
		return emptyResult, err
	}
	var oblLong float64
	for i := 0; i < len(points); i++ {
		oblLong, err = olc.oblLongForCelPoint(points[i].LonPos, points[i].LatPos, ayanOffset, spLong, spLat)
		if err != nil {
			return emptyResult, err
		}
		oblLong -= ayanOffset
		resultItem := points[i]
		resultItem.LonPos = oblLong
		result = append(result, resultItem)
	}

	return result, nil
}

// oblLongForCelPoint calculates oblique longitude for a celestial point
func (olc ObliqueLongCalculation) oblLongForCelPoint(eclLon, eclLat, ayanamshaOffset, longSp, latSp float64) (float64, error) {
	absLatSp := math.Abs(latSp)
	longPl := eclLon + ayanamshaOffset
	latPl := eclLat
	longSouthPMinusPlanet := math.Abs(longSp - longPl)
	longPlanetMinusSouthP := math.Abs(longPl - longSp)
	latSouthPMinusPlanet := absLatSp - latPl
	latSouthPPlusPlanet := absLatSp + latPl
	s := math.Min(longSouthPMinusPlanet, longPlanetMinusSouthP) / 2.0
	tanSRad := math.Tan(mathextra.DegToRad(s))
	qRad := math.Sin(mathextra.DegToRad(latSouthPMinusPlanet)) / math.Sin(mathextra.DegToRad(latSouthPPlusPlanet))
	v := mathextra.RadToDeg(math.Atan(tanSRad*qRad)) - s
	absoluteV, err := ValueToRange(math.Abs(v), -90.0, 90.0)
	if err != nil {
		return 0.0, err
	}
	absoluteV = math.Abs(absoluteV)

	var correctedV float64
	if olc.isRising(longSp, longPl) {
		if latPl < 0.0 {
			correctedV = absoluteV
		} else {
			correctedV = -absoluteV
		}
	} else {
		if latPl > 0.0 {
			correctedV = absoluteV
		} else {
			correctedV = -absoluteV
		}
	}
	//fmt.Printf("----------- DEBUG INFO ----------------------------\n")
	//fmt.Printf("eclLon %v and eclLat %v\n", eclLon, eclLat)
	//fmt.Printf("Southpoint longitude %v and latitude %v\n", longSp, latSp)
	//fmt.Printf("Ayanamsha offset  %v\n", ayanamshaOffset)
	//fmt.Printf("s: %v\n", s)
	//fmt.Printf("v: %v\n", v)
	//fmt.Printf("correctedV: %v\n", correctedV)
	//fmt.Printf("----------- END DEBUG INFO ------------------------\n")

	return ValueToRange(longPl+correctedV, 0.0, 360.0)
}

// isRising determines if a celestial point is rising based on longitude differences
func (olc ObliqueLongCalculation) isRising(longSp, longPl float64) bool {
	diff := longPl - longSp
	if diff < 0.0 {
		diff += 360.0
	}
	if diff >= 360.0 {
		diff -= 360.0
	}
	return diff < 180.0
}

// calculateSouthPoint calculates the ecliptic coordinates of the south point
func (olc ObliqueLongCalculation) calculateSouthPoint(armc, obliquity, geoLat float64) (eclLong, eclLat float64, err error) {
	declSp := -(90.0 - geoLat)
	arsp := armc
	if geoLat < 0.0 {
		arsp, err = ValueToRange(armc+180.0, 0.0, 360.0)
		if err != nil {
			return 0.0, 0.0, err
		}
		declSp = -90.0 - geoLat
	}
	sinSp := math.Sin(mathextra.DegToRad(arsp))
	cosEps := math.Cos(mathextra.DegToRad(obliquity))
	tanDecl := math.Tan(mathextra.DegToRad(declSp)) // error
	sinEps := math.Sin(mathextra.DegToRad(obliquity))
	cosArsp := math.Cos(mathextra.DegToRad(arsp))
	sinDecl := math.Sin(mathextra.DegToRad(declSp)) // error
	cosDecl := math.Cos(mathextra.DegToRad(declSp)) // error
	longSp := mathextra.RadToDeg(math.Atan2((sinSp*cosEps)+(tanDecl*sinEps), cosArsp))
	longSp, err = ValueToRange(longSp, 0.0, 360.0)
	latSp := mathextra.RadToDeg(math.Asin((sinDecl * cosEps) - (cosDecl * sinEps * sinSp)))
	return longSp, latSp, nil
}
