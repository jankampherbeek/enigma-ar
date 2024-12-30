/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package mathextra

import (
	"fmt"
	"math"
)

// PolarCoordinates represents polar coordinate system
type PolarCoordinates struct {
	PhiCoord   float64
	ThetaCoord float64
	RCoord     float64
}

// RectAngCoordinates represents rectangular coordinates
type RectAngCoordinates struct {
	XCoord float64
	YCoord float64
	ZCoord float64
}

func DegToRad(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func RadToDeg(radians float64) float64 {
	return radians * 180.0 / math.Pi
}

// TODO check if this function can be removed

// Rectangular2PolarArray converts array with rectangular coordinates to array with polar coordinates
func Rectangular2PolarArray(rectangularValues []float64) ([]float64, error) {
	if len(rectangularValues) != 3 {
		return nil, fmt.Errorf("invalid input for rectangularValues: expected 3 values, got %d", len(rectangularValues))
	}
	x := rectangularValues[0]
	y := rectangularValues[1]
	z := rectangularValues[2]

	r := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2) + math.Pow(z, 2))
	if r == 0 {
		r = math.SmallestNonzeroFloat64
	}
	if x == 0 {
		x = math.SmallestNonzeroFloat64
	}

	phi := math.Atan2(y, x)
	theta := math.Asin(z / r)

	return []float64{phi, theta, r}, nil
}

// Rectangular2Polar converts RectAngCoordinates to PolarCoordinates
func Rectangular2Polar(rectAngCoordinates RectAngCoordinates) PolarCoordinates {
	xCoord := rectAngCoordinates.XCoord
	yCoord := rectAngCoordinates.YCoord
	zCoord := rectAngCoordinates.ZCoord

	polarRCoord := math.Sqrt(math.Pow(xCoord, 2) + math.Pow(yCoord, 2) + math.Pow(zCoord, 2))
	if polarRCoord == 0 {
		polarRCoord = math.SmallestNonzeroFloat64
	}
	if xCoord == 0 {
		xCoord = math.SmallestNonzeroFloat64
	}

	polarPhiCoord := math.Atan2(yCoord, xCoord)
	polarThetaCoord := math.Asin(zCoord / polarRCoord)

	return PolarCoordinates{
		PhiCoord:   polarPhiCoord,
		ThetaCoord: polarThetaCoord,
		RCoord:     polarRCoord,
	}
}

// TODO check if this function can be removed

// Polar2RectangularArray converts array with polar coordinates to array with rectangular coordinates
func Polar2RectangularArray(polarValues []float64) ([]float64, error) {
	if len(polarValues) != 3 {
		return nil, fmt.Errorf("invalid input for polar values: expected 3 values, got %d", len(polarValues))
	}

	phi := polarValues[0]
	theta := polarValues[1]
	r := polarValues[2]

	x := r * math.Cos(theta) * math.Cos(phi)
	y := r * math.Cos(theta) * math.Sin(phi)
	z := r * math.Sin(theta)

	return []float64{x, y, z}, nil
}

// Polar2Rectangular converts PolarCoordinates to RectAngCoordinates
func Polar2Rectangular(polarCoordinates PolarCoordinates) RectAngCoordinates {
	phiCoord := polarCoordinates.PhiCoord
	thetaCoord := polarCoordinates.ThetaCoord
	rCoord := polarCoordinates.RCoord

	rectAngXCoord := rCoord * math.Cos(thetaCoord) * math.Cos(phiCoord)
	rectAngYCoord := rCoord * math.Cos(thetaCoord) * math.Sin(phiCoord)
	rectAngZCoord := rCoord * math.Sin(thetaCoord)

	return RectAngCoordinates{
		XCoord: rectAngXCoord,
		YCoord: rectAngYCoord,
		ZCoord: rectAngZCoord,
	}
}
