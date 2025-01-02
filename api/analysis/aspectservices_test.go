/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package apianalysis

import (
	"enigma-ar/domain"
	"testing"
)

func TestAspectsNotEnoughPoints(t *testing.T) {
	baseOrb := 10.0

	var points = []domain.SinglePosition{
		{Id: 0, Position: 100.0},
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

	aspS := NewAspectService()
	result, err := aspS.Aspects(points, aspects, cfgPoints, cfgAspects, baseOrb)

	if err == nil {
		t.Errorf("Aspectsshould have returned an error for not enough points")
	}
	if len(result) > 0 {
		t.Errorf("Aspectsshould have returned an empty result for not enough points")
	}
}

func TestAspectsNoAspects(t *testing.T) {
	baseOrb := 10.0

	var points = []domain.SinglePosition{
		{Id: 0, Position: 100.0},
		{Id: 1, Position: 164.0},
		{Id: 2, Position: 110.0},
		{Id: 3, Position: 90.0},
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

	var aspects []domain.Aspect

	var cfgAspects = []domain.ConfigAspect{
		{ActualAspect: 0, OrbFactor: 100, Glyph: '\uE700'}, // conjunction
		{ActualAspect: 1, OrbFactor: 100, Glyph: '\uE710'}, // opposition
		{ActualAspect: 2, OrbFactor: 80, Glyph: '\uE720'},  // trine
		{ActualAspect: 3, OrbFactor: 80, Glyph: '\uE730'},  // square
		{ActualAspect: 4, OrbFactor: 20, Glyph: '\uE810'},  // septile (not used in aspects)
		{ActualAspect: 5, OrbFactor: 60, Glyph: '\uE700'},  // sextile
	}

	aspS := NewAspectService()
	result, err := aspS.Aspects(points, aspects, cfgPoints, cfgAspects, baseOrb)

	if err == nil {
		t.Errorf("Aspectsshould have returned an error for not enough aspects")
	}
	if len(result) > 0 {
		t.Errorf("Aspectsshould have returned an empty result for not enough aspects")
	}
}

func TestAspectsNotEnoughConfigPoints(t *testing.T) {
	baseOrb := 10.0

	var points = []domain.SinglePosition{
		{Id: 0, Position: 100.0},
		{Id: 1, Position: 164.0},
		{Id: 2, Position: 110.0},
		{Id: 3, Position: 90.0},
		{Id: 4, Position: 269.0},
		{Id: 5, Position: 175.0},
		{Id: 6, Position: 339.5},
	}

	var cfgPoints = []domain.ConfigPoint{
		{ActualPoint: 0, OrbFactor: 100, Glyph: '\uE200'}, // Sun
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

	aspS := NewAspectService()
	result, err := aspS.Aspects(points, aspects, cfgPoints, cfgAspects, baseOrb)

	if err == nil {
		t.Errorf("Aspectsshould have returned an error for not enough config points")
	}
	if len(result) > 0 {
		t.Errorf("Aspectsshould have returned an empty result for not enough config points")
	}
}

func TestAspectsNoConfigAspects(t *testing.T) {
	baseOrb := 10.0

	var points = []domain.SinglePosition{
		{Id: 0, Position: 100.0},
		{Id: 1, Position: 164.0},
		{Id: 2, Position: 110.0},
		{Id: 3, Position: 90.0},
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

	var cfgAspects []domain.ConfigAspect

	aspS := NewAspectService()
	result, err := aspS.Aspects(points, aspects, cfgPoints, cfgAspects, baseOrb)

	if err == nil {
		t.Errorf("Aspectsshould have returned an error for not enough configaspects")
	}
	if len(result) > 0 {
		t.Errorf("Aspectsshould have returned an empty result for not enough configaspects")
	}
}

func TestAspectsMissingConfigPoint(t *testing.T) {
	baseOrb := 10.0

	var points = []domain.SinglePosition{
		{Id: 0, Position: 100.0},
		{Id: 1, Position: 164.0},
		{Id: 2, Position: 110.0},
		{Id: 3, Position: 90.0},
		{Id: 4, Position: 269.0},
		{Id: 5, Position: 175.0},
		{Id: 6, Position: 339.5},
	}

	var cfgPoints = []domain.ConfigPoint{ // Mercury is missing, id 2
		{ActualPoint: 0, OrbFactor: 100, Glyph: '\uE200'}, // Sun
		{ActualPoint: 1, OrbFactor: 100, Glyph: '\uE201'}, // Moon
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

	aspS := NewAspectService()
	result, err := aspS.Aspects(points, aspects, cfgPoints, cfgAspects, baseOrb)

	if err == nil {
		t.Errorf("Aspectsshould have returned an error for missing configpoints")
	}
	if len(result) > 0 {
		t.Errorf("Aspectsshould have returned an empty result for missing configpoints")
	}
}

func TestAspectsMissingConfigAspect(t *testing.T) {
	baseOrb := 10.0

	var points = []domain.SinglePosition{
		{Id: 0, Position: 100.0},
		{Id: 1, Position: 164.0},
		{Id: 2, Position: 110.0},
		{Id: 3, Position: 90.0},
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

	var cfgAspects = []domain.ConfigAspect{ // trine is missing, id 2
		{ActualAspect: 0, OrbFactor: 100, Glyph: '\uE700'}, // conjunction
		{ActualAspect: 1, OrbFactor: 100, Glyph: '\uE710'}, // opposition
		{ActualAspect: 3, OrbFactor: 80, Glyph: '\uE730'},  // square
		{ActualAspect: 4, OrbFactor: 20, Glyph: '\uE810'},  // septile (not used in aspects)
		{ActualAspect: 5, OrbFactor: 60, Glyph: '\uE700'},  // sextile
	}

	aspS := NewAspectService()
	result, err := aspS.Aspects(points, aspects, cfgPoints, cfgAspects, baseOrb)

	if err == nil {
		t.Errorf("Aspectsshould have returned an error for missing config aspects")
	}
	if len(result) > 0 {
		t.Errorf("Aspectsshould have returned an empty result for missing config aspects")
	}
}

func TestAspectsPositionTooLarge(t *testing.T) {
	baseOrb := 10.0

	var points = []domain.SinglePosition{
		{Id: 0, Position: 100.0},
		{Id: 1, Position: 164.0},
		{Id: 2, Position: 410.0},
		{Id: 3, Position: 90.0},
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

	aspS := NewAspectService()
	result, err := aspS.Aspects(points, aspects, cfgPoints, cfgAspects, baseOrb)

	if err == nil {
		t.Errorf("Aspectsshould have returned an error for a position that is too large")
	}
	if len(result) > 0 {
		t.Errorf("Aspectsshould have returned an empty result for a position that is too large")
	}
}

func TestAspectsPositionTooSmall(t *testing.T) {
	baseOrb := 10.0

	var points = []domain.SinglePosition{
		{Id: 0, Position: 100.0},
		{Id: 1, Position: 164.0},
		{Id: 2, Position: 110.0},
		{Id: 3, Position: 90.0},
		{Id: 4, Position: -269.0},
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

	aspS := NewAspectService()
	result, err := aspS.Aspects(points, aspects, cfgPoints, cfgAspects, baseOrb)

	if err == nil {
		t.Errorf("Aspectsshould have returned an error for a position that is too small")
	}
	if len(result) > 0 {
		t.Errorf("Aspectsshould have returned an empty result for a position that is too small")
	}
}
