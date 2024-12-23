/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package analysis

import (
	"enigma-ar/domain"
	"errors"
	"math"
)

const (
	MinOrbForMP       = 0.0
	MaxOrbForMP       = 10.0
	MinItemsForMP     = 2
	MinItemsForCalcMP = 3
	MinPosForMP       = 0.0
	MaxPosForMP       = 360.0
)

// MidpointsCalculator calculates midpoints in longitude (or in ra)
type MidpointsCalculator interface {
	CalcMidpointList(points []domain.SinglePosition) ([]domain.Midpoint, error)
	CalcOccupiedMidpoints(points []domain.SinglePosition, dial domain.MpDial, orb float64) ([]domain.OccupiedMidpoint, error)
}

type MidpointsCalculation struct{}

func NewMidpointsCalculation() MidpointsCalculator {
	return MidpointsCalculation{}
}

// CalcMidpointList calculates midpoints
// PRE length points >= 2
// PRE for all positions : 0.0 <= position < 360.0
// POST no errors -> returns slice of midpoints
// POST errors: returns empty slice and error
func (mc MidpointsCalculation) CalcMidpointList(points []domain.SinglePosition) ([]domain.Midpoint, error) {
	dialSize := 360.0
	emptyResults := make([]domain.Midpoint, 0)
	midpoints := make([]domain.Midpoint, 0)
	var actualMp float64
	if len(points) < MinItemsForMP {
		return emptyResults, errors.New("not enough points")
	}
	for i := 0; i < len(points); i++ { // first point
		if points[i].Position < MinPosForMP || points[i].Position >= MaxPosForMP {
			return emptyResults, errors.New("position must be between 0.0 and <360.0")
		}
		for j := i + 1; j < len(points); j++ { // second point
			actualMp = constructEffectiveMidpoint(points[i], points[j], dialSize)
			midpoints = append(midpoints, domain.Midpoint{Point1: points[i], Point2: points[j], MidpointPos: actualMp})
		}
	}
	return midpoints, nil
}

// CalcOccupiedMidpoints calculates midpoints
// PRE length points >= 3
// PRE for all positions : 0.0 <= position < 360.0
// PRE 0.0 < orb <= 10.0
// POST no errors -> returns slice of occupied midpoints
// POST errors: returns empty slice and error
func (mc MidpointsCalculation) CalcOccupiedMidpoints(points []domain.SinglePosition, dial domain.MpDial, orb float64) ([]domain.OccupiedMidpoint, error) {
	emptyResults := make([]domain.OccupiedMidpoint, 0)
	occMidpoints := make([]domain.OccupiedMidpoint, 0)
	pointsInDial := make([]domain.SinglePosition, 0)
	if len(points) < MinItemsForCalcMP {
		return emptyResults, errors.New("not enough points")
	}
	if orb <= MinOrbForMP || orb > MaxOrbForMP {
		return emptyResults, errors.New("orb must be between 0.0 and 10.0")
	}
	for i := 0; i < len(points); i++ {
		if points[i].Position < MinPosForMP || points[i].Position >= MaxPosForMP {
			return emptyResults, errors.New("positions must be between 0.0 and <360.0")
		}
	}
	// reduce all positions to size of dial
	dialSize := 360.0
	for i := 0; i < len(domain.AllMpDials()); i++ {
		if domain.AllMpDials()[i].Key == int(dial) {
			dialSize = domain.AllMpDials()[i].DialSize
		}
	}
	var tempPos float64
	for i := 0; i < len(points); i++ {
		tempPos = points[i].Position
		for tempPos >= dialSize {
			tempPos -= dialSize
		}
		pointsInDial = append(pointsInDial, domain.SinglePosition{
			Id:       points[i].Id,
			Position: tempPos,
		})
	}
	var actOrb float64
	var exactness float64
	for i := 0; i < len(pointsInDial); i++ { // first point
		for j := i + 1; j < len(pointsInDial); j++ { // second point
			mp := constructEffectiveMidpoint(pointsInDial[i], pointsInDial[j], dialSize) // calc midpoint
			for k := 0; k < len(pointsInDial); k++ {
				mpCandidatePos1 := pointsInDial[k].Position
				mpCandidatePos2 := mpCandidatePos1 - (dialSize / 2.0)
				if mpCandidatePos2 < 0.0 {
					mpCandidatePos2 = mpCandidatePos1 + (dialSize / 2.0)
				}
				if math.Abs(mpCandidatePos1-mp) <= orb || math.Abs(mpCandidatePos2-mp) <= orb { // match
					actOrb = math.Abs(mpCandidatePos1 - mp)
					if math.Abs(mpCandidatePos2-mp) < actOrb {
						actOrb = math.Abs(mpCandidatePos2 - mp)
					}
					exactness = (1 - (actOrb / orb)) * 100.0
					occMidpoints = append(occMidpoints, domain.OccupiedMidpoint{
						BaseMidpointPos1: pointsInDial[i],
						BaseMidpointPos2: pointsInDial[j],
						FocusPoint:       pointsInDial[k],
						ActualOrb:        actOrb,
						Exactness:        exactness,
					})
				}
			}
		}
	}
	return occMidpoints, nil
}

func constructEffectiveMidpoint(point1, point2 domain.SinglePosition, dialSize float64) float64 {
	halfDial := dialSize / 2
	pos1 := point1.Position
	pos2 := point2.Position

	smallPos := pos1
	if pos2 < pos1 {
		smallPos = pos2
	}
	largePos := pos2
	if pos1 > pos2 {
		largePos = pos1
	}
	diff := largePos - smallPos
	firstPosShortestArc := smallPos
	if diff >= halfDial {
		firstPosShortestArc = largePos
	}
	lastPosShortestArc := largePos
	if diff >= halfDial {
		lastPosShortestArc = smallPos
	}
	diff = lastPosShortestArc - firstPosShortestArc
	if diff < 0.0 {
		diff += dialSize
	}
	mPos := (diff / 2) + firstPosShortestArc
	if mPos >= dialSize {
		mPos -= dialSize
	}
	return mPos
}
