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

type JulDayCalculator interface {
	CalcJd(year int, month int, day int, ut float64, greg bool) float64
}

type PointPosCalculator interface {
	CalcPointPos(request domain.PointPositionsRequest) ([]domain.PointPosResult, error)
}

type PointRangeCalculator interface {
	CalcPointRange(request domain.PointRangeRequest) ([]domain.PointRangeResult, error)
}

type JulDayCalculation struct {
	seCalc se.SeJulDayCalculator
}

func NewJulDayCalculation() JulDayCalculator {
	sjc := se.NewSeJulDayCalculation()
	return JulDayCalculation{sjc}
}

// CalcJd handles the calculation of a Julian day number.
func (jdc JulDayCalculation) CalcJd(year int, month int, day int, ut float64, greg bool) float64 {
	var gregFlag = 1
	if !greg {
		gregFlag = 0
	}
	return jdc.seCalc.SeCalcJd(year, month, day, ut, gregFlag)
}

type PointPosCalculation struct {
	sePointCalc  se.SePointPosCalculator
	seHorPosCalc se.SeHorPosCalculation
}

func NewPointPosCalculation() PointPosCalculator {
	ppc := se.NewSePointPosCalculation()
	hpc := se.NewSeHorPosCalculation()
	return PointPosCalculation{ppc, hpc}
}

// CalcPointPos calculates fully defined positions for one or more celestial points
func (calc PointPosCalculation) CalcPointPos(request domain.PointPositionsRequest) ([]domain.PointPosResult, error) {
	positions := make([]domain.PointPosResult, 0)
	eclFlags := SeFlags(domain.Ecliptical, request.ObsPos, request.Tropical)
	equFlags := SeFlags(domain.Equatorial, request.ObsPos, request.Tropical)
	for i := 0; i < len(request.Points); i++ {
		var point = request.Points[i]
		posEcl, errEcl := calc.sePointCalc.SeCalcPointPos(request.JdUt, point, eclFlags)
		if errEcl != nil {
			return positions, errEcl
		}
		posEqu, errEqu := calc.sePointCalc.SeCalcPointPos(request.JdUt, point, equFlags)
		if errEqu != nil {
			return positions, errEqu
		}
		height := 0.0
		pointRa := posEqu[0]
		pointDecl := posEqu[1]
		horFlags := domain.SEFLG_EQUATORIAL
		posHor := calc.seHorPosCalc.CalcHorPos(request.JdUt, request.GeoLong, request.GeoLat, height, pointRa, pointDecl, horFlags)
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
			AzimPos:   posHor[0],
			AltitPos:  posHor[2],
		})
	}
	return positions, nil
}

type PointRangeCalculation struct {
	sePointCalc se.SePointPosCalculator
}

func NewPointRangeCalculation() PointRangeCalculator {
	ppc := se.NewSePointPosCalculation()
	return PointRangeCalculation{ppc}
}

func (prc PointRangeCalculation) CalcPointRange(request domain.PointRangeRequest) ([]domain.PointRangeResult, error) {
	point := request.Point
	flags := SeFlags(request.Coord, request.ObsPos, request.Ayanamsha == 0)
	// TODO handle topocentric
	// TODO handle sidereal
	var rangePositions []domain.PointRangeResult
	var resultIndex int
	if request.Position {
		if request.MainValue {
			resultIndex = 0
		} else {
			resultIndex = 4
		}
	} else {
		if request.MainValue {
			resultIndex = 1
		} else {
			resultIndex = 5
		}
	}
	// TODO handle RADV/Distance
	for i := request.JdStart; i <= request.JdEnd; i += request.Interval {
		sePos, err := prc.sePointCalc.SeCalcPointPos(i, point, flags)
		if err != nil {
			return rangePositions, err
		}
		calcValue := sePos[resultIndex]
		rangePositions = append(rangePositions, domain.PointRangeResult{i, calcValue}) // TODO improve appending
	}
	return rangePositions, nil
}

/*func HousePos(hsys rune, jdUt float64, geoLong float64, geoLat float64, tropical bool) ([]domain.HousePosResult, []domain.HousePosResult, error) {
	cuspPos := make([]domain.HousePosResult, 0)
	mcAscPos := make([]domain.HousePosResult, 0)
	flagsEcl := SeFlags(domain.Ecliptical, domain.Geocentric, tropical)
	flagsEqu := SeFlags(domain.Equatorial, domain.Geocentric, tropical)
	hp := se.NewHousePos()
	cuspsEcl, mcAscEcl, errEcl := hp.CalcHousePos(hsys, jdUt, geoLat, geoLong, flagsEcl)
	if errEcl != nil {
		return cuspPos, mcAscPos, errEcl
	}
	cuspsEqu, mcAscEqu, errEqu := hp.CalcHousePos(hsys, jdUt, geoLat, geoLong, flagsEqu)
	if errEqu != nil {
		return cuspPos, mcAscPos, errEqu
	}

	horFlags := domain.SEFLG_EQUATORIAL
	height := 0.0
	for i := 0; i < len(cuspPos); i++ {
		posHor := se.HorizontalPosition(jdUt, geoLong, geoLat, height, cuspsEqu[i], pointDecl, horFlags)

	}

	//houseSys rune, jdUt float64, geoLat float64, geoLong float64, flags int32) ([]float64, []float64,
}*/

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
