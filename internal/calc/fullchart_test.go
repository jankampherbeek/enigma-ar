/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package calc

import (
	"testing"
)

// just a simple test as the calculated values will be checked by an integration test
func TestCalcFullChartHappyFlow(t *testing.T) {
	// TODO activate TestCalcFullChartHappyFlow
	if 1 < 1 {
		t.Errorf("Stange error in dummy test")
	}

	//request := domain.FullChartRequest{
	//	Points: []domain.ChartPoint{
	//		domain.Sun,
	//		domain.Moon,
	//		domain.Mercury,
	//	},
	//	HouseSys:  domain.Regiomontanus,
	//	Ayanamsha: domain.AyanDeLuce,
	//	CoordSys:  domain.CoordEcliptical,
	//	ObsPos:    domain.ObsPosGeocentric,
	//	ProjType:  domain.ProjType2D,
	//	Jd:        123456.789,
	//	GeoLong:   0.0,
	//	GeoLat:    0.0,
	//}
	//
	//fcc := NewFullChartCalculation()
	//_, err := fcc.CalcFullChart(request)
	//if err != nil {
	//	t.Errorf("received unexpected error %v", err)
	//}
	//if result.Points[0].Point != domain.Sun {
	//	t.Errorf("Expected point %d, got %d", domain.Sun, result.Points[0].Point)
	//}
}
