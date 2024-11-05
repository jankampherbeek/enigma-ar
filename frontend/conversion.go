/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"fmt"
	"math"
	"strconv"
)

const (
	DEGREE_SIGN = "\u00B0" // °
	MINUTE_SIGN = "\u2032" // ′
	SECOND_SIGN = "\u2033" // ″
)

// IDoubleToDmsConversions is an interface for converting double values to presentable strings.
type IDoubleToDmsConversions interface {
	// ConvertDoubleToDmInSignNoGlyph converts longitude to degrees and minutes within a sign without a glyph.
	ConvertDoubleToDmInSignNoGlyph(position float64) string
	// ConvertDoubleToDmsWithGlyph converts longitude to degrees, minutes, and seconds within a sign with a glyph.
	ConvertDoubleToDmsWithGlyph(position float64) (string, rune)
	// ConvertDoubleToDmsInSignNoGlyph converts longitude to degrees, minutes, and seconds within a sign without a glyph.
	ConvertDoubleToDmsInSignNoGlyph(position float64) string
	// ConvertDoubleToPositionsDmsText converts a value to a sexagesimal text, indicating negative values with a minus sign.
	ConvertDoubleToPositionsDmsText(position float64) string
}

// DoubleToDmsConversions is a struct that implements IDoubleToDmsConversions.
type DoubleToDmsConversions struct{}

func NewDoubleToDmsConversions() DoubleToDmsConversions {
	return DoubleToDmsConversions{}
}

// ConvertDoubleToDmInSignNoGlyph converts longitude to degrees and minutes within a sign without a glyph.
func (d *DoubleToDmsConversions) ConvertDoubleToDmInSignNoGlyph(position float64) string {
	posInRange := valueToRange(position, 0.0, 360.0)
	remaining := posInRange
	degrees := int(posInRange)
	nrOfSigns := 1 + (degrees / 30)
	degreesInSign := degrees - ((nrOfSigns - 1) * 30)
	remaining = math.Abs(remaining - float64(degrees))
	minutes := int(remaining * 60.0)
	return createDmString(degreesInSign, minutes)
}

// ConvertDoubleToDmsWithGlyph converts longitude to degrees, minutes, and seconds within a sign with a glyph.
func (d *DoubleToDmsConversions) ConvertDoubleToDmsWithGlyph(position float64) (string, rune) {
	posInRange := valueToRange(position, 0.0, 360.0)
	remaining := posInRange
	degrees := int(posInRange)
	nrOfSigns := 1 + (degrees / 30)
	degreesInSign := degrees - ((nrOfSigns - 1) * 30)
	remaining = math.Abs(remaining - float64(degrees))
	minutes := int(remaining * 60.0)
	remaining -= float64(minutes) / 60.0
	seconds := int(remaining * 3600.0)
	longTxt := createDmsString(degreesInSign, minutes, seconds)
	glyph := defineGlyph(nrOfSigns)
	return longTxt, glyph
}

// ConvertDoubleToDmsInSignNoGlyph returns the longitude in degrees, minutes, and seconds within a sign without a glyph.
func (d *DoubleToDmsConversions) ConvertDoubleToDmsInSignNoGlyph(position float64) string {
	longTxt, _ := d.ConvertDoubleToDmsWithGlyph(position)
	return longTxt
}

// ConvertDoubleToPositionsDmsText converts a value to a sexagesimal text, indicating negative values with a minus sign.
func (d *DoubleToDmsConversions) ConvertDoubleToPositionsDmsText(position float64) string {
	minusSign := ""
	if position < 0.0 {
		minusSign = "-"
	}
	const correctionForDouble = 0.00000001 // Correction to prevent double values like 0.99999999999
	remaining := math.Abs(position) + correctionForDouble
	if remaining >= 360.0 {
		remaining -= 360.0
	}
	degrees := int(remaining)
	remaining -= float64(degrees)
	minutes := int(remaining * 60.0)
	remaining -= float64(minutes) / 60.0
	seconds := int(remaining * 3600.0)
	return minusSign + createDmsString(degrees, minutes, seconds)
}

// createDmsString creates a string representation of degrees, minutes, and seconds.
func createDmsString(degrees, minutes, seconds int) string {
	degreeText := strconv.Itoa(degrees)
	minuteText := fmt.Sprintf("%02d", minutes)
	secondText := fmt.Sprintf("%02d", seconds)
	return degreeText + DEGREE_SIGN + minuteText + MINUTE_SIGN + secondText + SECOND_SIGN
}

// createDmString creates a string representation of degrees and minutes.
func createDmString(degrees, minutes int) string {
	degreeText := strconv.Itoa(degrees)
	minuteText := fmt.Sprintf("%02d", minutes)
	return degreeText + DEGREE_SIGN + minuteText + MINUTE_SIGN
}

// defineGlyph defines a glyph based on the number of signs.
func defineGlyph(nrOfSigns int) rune {
	allGlyphs := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-', '='}
	if nrOfSigns-1 >= 0 && nrOfSigns-1 < len(allGlyphs) {
		return allGlyphs[nrOfSigns-1]
	}
	return '?'
}

// valueToRange adjusts a value to be within the range [min, max).
func valueToRange(value, min, max float64) float64 {
	rangeWidth := max - min
	adjustedValue := math.Mod(value-min, rangeWidth)
	if adjustedValue < 0 {
		adjustedValue += rangeWidth
	}
	return adjustedValue + min
}
