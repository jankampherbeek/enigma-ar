/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	"enigma-ar/internal/calc/se"
	"enigma-ar/internal/domain"
)

// JulianDay handles the calculation of a Julian day number.
func JulianDay(year int, month int, day int, ut float64, greg bool) float64 {
	return se.JulDay(year, month, day, ut, greg)
}

// FullPositions calculates fully defined positions for one or more celestial points
func FullPositions(request domain.PointPositionsRequest) ([]domain.PointPosResult, error) {
	positions := make([]domain.PointPosResult, 0)
	eclFlags := SeFlags(domain.Ecliptical, request.ObsPos, request.Tropical)
	equFlags := SeFlags(domain.Equatorial, request.ObsPos, request.Tropical)
	for i := 0; i < len(request.Points); i++ {
		var point = request.Points[i]
		posEcl, errEcl := se.PointPositions(request.JdUt, point, eclFlags)
		if errEcl != nil {
			return positions, errEcl
		}
		posEqu, errEqu := se.PointPositions(request.JdUt, point, equFlags)
		if errEqu != nil {
			return positions, errEqu
		}
		// todo calculate horizontal coordinates
		positions = append(positions, domain.PointPosResult{
			Point:     point,
			LonPos:    posEcl[0],
			LonSpeed:  posEcl[3],
			LatPos:    posEcl[1],
			LatSpeed:  posEcl[4],
			RaPos:     posEqu[0],
			RaSpeed:   posEqu[3],
			DeclPos:   posEqu[1],
			DeclSpeed: posEqu[4],
			RadvPos:   posEcl[2],
			RadvSpeed: posEcl[5],
			AzimPos:   0.0,
			AltitPos:  0.0,
		})
	}
	return positions, nil
}

// SeFlags calculates the total of all flags for the SE.
func SeFlags(coord domain.CoordinateSystem, obsPos domain.ObserverPosition, tropical bool) int {
	flags := domain.SEFLG_SWIEPH + domain.SEFLG_SPEED // always use SE + speed
	if coord == domain.Equatorial {
		flags += domain.SEFLG_EQUATORIAL
	}
	if obsPos == domain.Topocentric {
		flags += domain.SEFLG_TOPOCTR
	}
	if obsPos == domain.Heliocentric {
		flags += domain.SEFLG_HELIOC
	}
	if coord == domain.Equatorial && !tropical {
		flags += domain.SEFLG_SIDEREAL
	}
	return flags
}
