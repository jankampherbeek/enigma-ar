/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package analysis

import (
	"enigma-ar/domain"
	"math"
	"testing"
)

func TestCalcAspectsHappyFlow(t *testing.T) {

	baseOrb := 10.0

	var points = []domain.SinglePosition{
		{Id: 0, Position: 100.0}, // 60 with 1, 0 with 2, 120 with 6,
		{Id: 1, Position: 164.0}, // 60 with 2    (border case)
		{Id: 2, Position: 110.0}, // 60 with 5
		{Id: 3, Position: 90.0},  // 180 with 4
		{Id: 4, Position: 269.0},
		{Id: 5, Position: 175.0},
		{Id: 6, Position: 339.5},
	}

	var cfgPoints = []domain.ConfigPoint{
		{ActualPoint: 0, OrbFactor: 100, Glyph: '\uE200'}, // Sun
		{ActualPoint: 1, OrbFactor: 100, Glyph: '\uE201'}, // Moon
		{ActualPoint: 2, OrbFactor: 80, Glyph: '\uE202'},  // Mercury
		{ActualPoint: 3, OrbFactor: 80, Glyph: '\uE203'},  // Venus
		{ActualPoint: 4, OrbFactor: 80, Glyph: '\uE204'},  // Mars
		{ActualPoint: 5, OrbFactor: 60, Glyph: '\uE205'},  // Jupiter
		{ActualPoint: 6, OrbFactor: 60, Glyph: '\uE206'},  // Saturn
	}

	var aspects = []domain.Aspect{
		0, // conjunction
		1, // opposition
		2, // trine
		3, // square
		5, // sextile
	}

	var cfgAspects = []domain.ConfigAspect{
		{ActualAspect: 0, OrbFactor: 100, Glyph: '\uE700'}, // conjunction
		{ActualAspect: 1, OrbFactor: 100, Glyph: '\uE710'}, // opposition
		{ActualAspect: 2, OrbFactor: 80, Glyph: '\uE720'},  // trine
		{ActualAspect: 3, OrbFactor: 80, Glyph: '\uE730'},  // square
		{ActualAspect: 4, OrbFactor: 20, Glyph: '\uE810'},  // septile (not used in aspects)
		{ActualAspect: 5, OrbFactor: 60, Glyph: '\uE700'},  // sextile
	}

	var expected = []domain.ActualAspect{
		{Pos1: domain.SinglePosition{Id: 0, Position: 100.0},
			Pos2:         domain.SinglePosition{Id: 1, Position: 164.0},
			ActualAspect: 5,
			ActualOrb:    4.0,
			Exactness:    34,
		},
		{Pos1: domain.SinglePosition{Id: 0, Position: 100.0},
			Pos2:         domain.SinglePosition{Id: 2, Position: 110.0},
			ActualAspect: 0,
			ActualOrb:    10.0,
			Exactness:    0,
		},
		{Pos1: domain.SinglePosition{Id: 0, Position: 100.0},
			Pos2:         domain.SinglePosition{Id: 3, Position: 90.0},
			ActualAspect: 0,
			ActualOrb:    10.0,
			Exactness:    0,
		},
		{Pos1: domain.SinglePosition{Id: 0, Position: 100.0},
			Pos2:         domain.SinglePosition{Id: 6, Position: 339.5},
			ActualAspect: 2,
			ActualOrb:    0.5,
			Exactness:    94,
		},
		{Pos1: domain.SinglePosition{Id: 1, Position: 164.0},
			Pos2:         domain.SinglePosition{Id: 2, Position: 110.0},
			ActualAspect: 5,
			ActualOrb:    6.0,
			Exactness:    0,
		},
		{Pos1: domain.SinglePosition{Id: 1, Position: 164.0},
			Pos2:         domain.SinglePosition{Id: 6, Position: 339.5},
			ActualAspect: 1,
			ActualOrb:    4.5,
			Exactness:    55,
		},
		{Pos1: domain.SinglePosition{Id: 3, Position: 90.0},
			Pos2:         domain.SinglePosition{Id: 4, Position: 269.0},
			ActualAspect: 1,
			ActualOrb:    1.0,
			Exactness:    88,
		},
		{Pos1: domain.SinglePosition{Id: 3, Position: 90.0},
			Pos2:         domain.SinglePosition{Id: 5, Position: 175.0},
			ActualAspect: 3,
			ActualOrb:    5.0,
			Exactness:    22,
		},
		{Pos1: domain.SinglePosition{Id: 4, Position: 269.0},
			Pos2:         domain.SinglePosition{Id: 5, Position: 175.0},
			ActualAspect: 3,
			ActualOrb:    4.0,
			Exactness:    38,
		},
	}

	aspCalc := AspectsCalculation{}
	result, err := aspCalc.CalcAspects(points, aspects, cfgPoints, cfgAspects, baseOrb)

	if err != nil {
		t.Fatalf("aspects calculation failed, returned unexpected error %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("Expected %d results, got %d, result was %v", len(expected), len(result), result)
	}

	var difference float64
	for i := range result {
		difference = math.Abs(result[i].ActualOrb - expected[i].ActualOrb)
		if difference > 1e-8 {
			t.Errorf("Wrong value for actualorb for points %d and %d at index %d: got %f, want %f",
				result[i].Pos1.Id, result[i].Pos2.Id, i, result[i].ActualOrb, expected[i].ActualOrb)
		}
		if result[i].Exactness != expected[i].Exactness {
			t.Errorf("Wrong value for exactness index %d: got %d, want %d",
				i, result[i].Exactness, expected[i].Exactness)
		}
		if result[i].Pos1.Id != expected[i].Pos1.Id {
			t.Errorf("ID mismatch for pos1 at index %d: got %d, want %d",
				i, result[i].Pos1.Id, expected[i].Pos1.Id)
		}
		if result[i].Pos2.Id != expected[i].Pos2.Id {
			t.Errorf("ID mismatch for pos2 at index %d: got %d, want %d",
				i, result[i].Pos2.Id, expected[i].Pos2.Id)
		}
		if result[i].ActualAspect != expected[i].ActualAspect {
			t.Errorf("Aspects mismatch at index %d: got %d, want %d",
				i, result[i].ActualAspect, expected[i].ActualAspect)
		}
		difference = math.Abs(result[i].Pos1.Position - expected[i].Pos1.Position)
		if difference > 1e-8 {
			t.Errorf("Wrong position for pos1 at index %d: got %f, want %f",
				i, result[i].Pos1.Position, expected[i].Pos1.Position)
		}
		difference = math.Abs(result[i].Pos2.Position - expected[i].Pos2.Position)
		if difference > 1e-8 {
			t.Errorf("Wrong position for pos2 at index %d: got %f, want %f",
				i, result[i].Pos2.Position, expected[i].Pos2.Position)
		}
	}
}
