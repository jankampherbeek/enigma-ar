/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	"enigma-ar/domain"
	"enigma-ar/internal/calc/mathextra"
	"fmt"
	"math"
)

// OrbitDefinition represents orbital parameters
type OrbitDefinition struct {
	MeanAnomaly        []float64
	EccentricAnomaly   []float64
	SemiMajorAxis      float64
	ArgumentPerihelion []float64
	AscNode            []float64
	Inclination        []float64
}

// PointsElementsCalculator calculates positions of a celestial body based on elements.
type PointsElementsCalculator interface {
	Calculate(planetId int, jdUt float64, observerPosition domain.ObserverPosition) []float64
}

type PointsElementsCalculation struct{}

func NewPointsElementsCalculation() PointsElementsCalculator {
	return PointsElementsCalculation{}
}

// Calculate performs the calculation. Returns longitude, latitude and distance in that sequence.
func (c PointsElementsCalculation) Calculate(planetId int, jdUt float64, observerPosition domain.ObserverPosition) []float64 {
	centuryFractionT := c.factorT(jdUt)
	// no difference between geocentric and topocentric becasue of the distances involved
	if observerPosition == domain.ObsPosGeocentric || observerPosition == domain.ObsPosTopocentric {
		polarPlanetGeo := c.calcGeoPolarCoord(planetId, centuryFractionT)
		return c.definePosition(polarPlanetGeo)
	}
	// heliocentric
	polarPlanetHelio := c.calcHelioPolarCoord(planetId, centuryFractionT)
	return c.definePosition(polarPlanetHelio)
}

func (c PointsElementsCalculation) calcGeoPolarCoord(planetId int, centuryFractionT float64) mathextra.PolarCoordinates {
	rectAngEarthHelio := c.CalcEclipticHelioPosition(centuryFractionT, c.defineOrbitDefinition(domain.AllChartPoints()[domain.Earth].CalcId))
	rectAngPlanetHelio := c.CalcEclipticHelioPosition(centuryFractionT, c.defineOrbitDefinition(planetId))
	rectAngPlanetGeo := mathextra.RectAngCoordinates{
		XCoord: rectAngPlanetHelio.XCoord - rectAngEarthHelio.XCoord,
		YCoord: rectAngPlanetHelio.YCoord - rectAngEarthHelio.YCoord,
		ZCoord: rectAngPlanetHelio.ZCoord - rectAngEarthHelio.ZCoord,
	}
	return mathextra.Rectangular2Polar(rectAngPlanetGeo)
}

func (c PointsElementsCalculation) calcHelioPolarCoord(planetId int, centuryFractionT float64) mathextra.PolarCoordinates {
	rectAngPlanetHelio := c.CalcEclipticHelioPosition(centuryFractionT, c.defineOrbitDefinition(planetId))
	return mathextra.Rectangular2Polar(rectAngPlanetHelio)
}

func (c PointsElementsCalculation) definePosition(polarPlanetGeo mathextra.PolarCoordinates) []float64 {
	posLong := mathextra.RadToDeg(polarPlanetGeo.PhiCoord)
	if posLong < 0.0 {
		posLong += 360.0
	}
	posLat := mathextra.RadToDeg(polarPlanetGeo.ThetaCoord)
	posDist := mathextra.RadToDeg(polarPlanetGeo.RCoord)
	return []float64{posLong, posLat, posDist}
}

func (c PointsElementsCalculation) factorT(jdUt float64) float64 {
	return (jdUt - 2415020.5) / 36525
}

func (c PointsElementsCalculation) defineOrbitDefinition(planetId int) OrbitDefinition {
	var meanAnomaly, eccentricAnomaly, argumentPerihelion, ascNode, inclination []float64
	var semiMajorAxis float64

	eccentricAnomaly = []float64{0, 0, 0}
	argumentPerihelion = []float64{0, 0, 0}
	ascNode = []float64{0, 0, 0}
	inclination = []float64{0, 0, 0}

	switch planetId {
	case domain.AllChartPoints()[domain.Earth].CalcId:
		meanAnomaly = []float64{358.47584, 35999.0498, -.00015}
		eccentricAnomaly = []float64{.016751, -.41e-4, 0}
		semiMajorAxis = 1.00000013
		argumentPerihelion = []float64{101.22083, 1.71918, .00045}
	case domain.AllChartPoints()[domain.PersephoneRam].CalcId:
		meanAnomaly = []float64{295.0, 60, 0}
		semiMajorAxis = 71.137866
	case domain.AllChartPoints()[domain.HermesRam].CalcId:
		meanAnomaly = []float64{134.7, 50.0, 0}
		semiMajorAxis = 80.331954
	case domain.AllChartPoints()[domain.DemeterRam].CalcId:
		meanAnomaly = []float64{114.6, 40, 0}
		semiMajorAxis = 93.216975
		ascNode = []float64{125, 0, 0}
		inclination = []float64{5.5, 0, 0}
	default:
		panic(fmt.Sprintf("Unrecognized planet for OrbitDefinition: %d", planetId))
	}

	return OrbitDefinition{
		MeanAnomaly:        meanAnomaly,
		EccentricAnomaly:   eccentricAnomaly,
		SemiMajorAxis:      semiMajorAxis,
		ArgumentPerihelion: argumentPerihelion,
		AscNode:            ascNode,
		Inclination:        inclination,
	}
}

type OrbitCalculator struct {
	meanAnomaly2 float64
	semiAxis     float64
	inclination  float64
	eccentricity float64
	eccAnomaly   float64
	factorT      float64
}

// CalcEclipticHelioPosition calculates the ecliptic heliocentric position
func (c PointsElementsCalculation) CalcEclipticHelioPosition(factorT float64, orbitDefinition OrbitDefinition) mathextra.RectAngCoordinates {
	meanAnomaly1 := mathextra.DegToRad(c.ProcessTermsForFractionT(factorT, orbitDefinition.MeanAnomaly))
	if meanAnomaly1 < 0.0 {
		meanAnomaly1 += math.Pi * 2
	}
	eccentricity := c.ProcessTermsForFractionT(factorT, orbitDefinition.EccentricAnomaly)
	eccAnomaly := c.EccAnomalyFromKeplerEquation(meanAnomaly1, eccentricity)
	trueAnomalyPol := c.CalcPolarTrueAnomaly(orbitDefinition, eccAnomaly, eccentricity)
	semiAxis, inclination, meanAnomaly2 := c.ReduceToEcliptic(trueAnomalyPol, orbitDefinition, factorT)

	return c.CalcRectAngHelioCoordinates(
		semiAxis,
		inclination,
		eccAnomaly,
		eccentricity,
		meanAnomaly2,
		orbitDefinition,
	)
}

func (c PointsElementsCalculation) EccAnomalyFromKeplerEquation(meanAnomaly, eccentricity float64) float64 {
	eccAnomaly := meanAnomaly
	for count := 1; count < 6; count++ {
		eccAnomaly = meanAnomaly + (eccentricity * math.Sin(eccAnomaly))
	}
	return eccAnomaly
}

func (c PointsElementsCalculation) CalcPolarTrueAnomaly(orbitDefinition OrbitDefinition, eccAnomaly, eccentricity float64) mathextra.PolarCoordinates {
	xCoord := orbitDefinition.SemiMajorAxis * (math.Cos(eccAnomaly) - eccentricity)
	yCoord := orbitDefinition.SemiMajorAxis * math.Sin(eccAnomaly) * math.Sqrt(1-(eccentricity*eccentricity))
	const zCoord = 0.0
	anomalyVec := mathextra.RectAngCoordinates{XCoord: xCoord, YCoord: yCoord, ZCoord: zCoord}
	return mathextra.Rectangular2Polar(anomalyVec)
}

func (c PointsElementsCalculation) CalcRectAngHelioCoordinates(semiAxis, inclination, eccAnomaly, eccentricity, meanAnomaly float64, orbitDefinition OrbitDefinition) mathextra.RectAngCoordinates {
	phiCoord := mathextra.DegToRad(semiAxis)
	if phiCoord < 0.0 {
		phiCoord += math.Pi * 2
	}

	thetaCoord := math.Atan(math.Sin(phiCoord-meanAnomaly) * math.Tan(inclination))
	rCoord := mathextra.DegToRad(orbitDefinition.SemiMajorAxis) * (1 - (eccentricity * math.Cos(eccAnomaly)))
	helioPol := mathextra.PolarCoordinates{
		PhiCoord:   phiCoord,
		ThetaCoord: thetaCoord,
		RCoord:     rCoord,
	}
	return mathextra.Polar2Rectangular(helioPol)
}

func (c PointsElementsCalculation) ReduceToEcliptic(trueAnomalyPol mathextra.PolarCoordinates,
	orbitDefinition OrbitDefinition, factorT float64) (semiAxis, inclination, meanAnomaly2 float64) {
	semiAxis = mathextra.RadToDeg(trueAnomalyPol.PhiCoord) + c.ProcessTermsForFractionT(factorT, orbitDefinition.ArgumentPerihelion)
	meanAnomaly2 = mathextra.DegToRad(c.ProcessTermsForFractionT(factorT, orbitDefinition.AscNode))
	factorVDeg := semiAxis + mathextra.RadToDeg(meanAnomaly2)
	if factorVDeg < 0.0 {
		factorVDeg += 360.0
	}
	factorVRad := mathextra.DegToRad(factorVDeg)

	inclination = mathextra.DegToRad(c.ProcessTermsForFractionT(factorT, orbitDefinition.Inclination))
	semiAxis = math.Atan(math.Cos(inclination) * math.Tan(factorVRad-meanAnomaly2))
	if semiAxis < math.Pi {
		semiAxis += math.Pi
	}
	semiAxis = mathextra.RadToDeg(semiAxis + meanAnomaly2)
	if math.Abs(factorVDeg-semiAxis) > 10.0 {
		semiAxis -= 180.0
	}
	return semiAxis, inclination, meanAnomaly2
}

func (c PointsElementsCalculation) ProcessTermsForFractionT(fractionT float64, elements []float64) float64 {
	return elements[0] + (elements[1] * fractionT) + (elements[2] * fractionT * fractionT)
}
