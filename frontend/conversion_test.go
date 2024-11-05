/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import "testing"

func TestDoubleToDmsHappyFlow(t *testing.T) {
	converter := NewDoubleToDmsConversions()
	inputValue := 6.5
	expected := "6" + DEGREE_SIGN + "30" + MINUTE_SIGN + "00" + SECOND_SIGN
	result := converter.ConvertDoubleToPositionsDmsText(inputValue)
	if result != expected {
		t.Error("DoubleToDmsHappyFlow. Expected", expected, "got", result)
	}
}

func TestDoubleToDmsNegative(t *testing.T) {
	converter := NewDoubleToDmsConversions()
	inputValue := -6.5
	expected := "-6" + DEGREE_SIGN + "30" + MINUTE_SIGN + "00" + SECOND_SIGN
	result := converter.ConvertDoubleToPositionsDmsText(inputValue)
	if result != expected {
		t.Error("DoubleToDmsNegative. Expected", expected, "got", result)
	}
}

func TestDoubleToDmsSmallNegative(t *testing.T) {
	converter := NewDoubleToDmsConversions()
	inputValue := -0.03
	expected := "-0" + DEGREE_SIGN + "01" + MINUTE_SIGN + "48" + SECOND_SIGN
	result := converter.ConvertDoubleToPositionsDmsText(inputValue)
	if result != expected {
		t.Error("DoubleToDmsSmallNegative. Expected", expected, "got", result)
	}
}

func TestDoubleToDmsLargeValue(t *testing.T) {
	converter := NewDoubleToDmsConversions()
	inputValue := 342.5
	expected := "342" + DEGREE_SIGN + "30" + MINUTE_SIGN + "00" + SECOND_SIGN
	result := converter.ConvertDoubleToPositionsDmsText(inputValue)
	if result != expected {
		t.Error("DoubleToDmsLargeValue. Expected", expected, "got", result)
	}
}

func TestDoubleToDmsBorderValue(t *testing.T) {
	converter := NewDoubleToDmsConversions()
	inputValue := 42.999999
	expected := "42" + DEGREE_SIGN + "59" + MINUTE_SIGN + "59" + SECOND_SIGN
	result := converter.ConvertDoubleToPositionsDmsText(inputValue)
	if result != expected {
		t.Error("DoubleToDmsBorderValue. Expected", expected, "got", result)
	}
}

func TestDoubleToDmsHappyFlowWithGlyph(t *testing.T) {
	converter := NewDoubleToDmsConversions()
	inputValue := 36.5
	expectedTxt := "6" + DEGREE_SIGN + "30" + MINUTE_SIGN + "00" + SECOND_SIGN
	expectedGlyph := '2'
	resultTxt, resultGlyph := converter.ConvertDoubleToDmsWithGlyph(inputValue)
	if resultTxt != expectedTxt {
		t.Error("DoubleToDmsHappyFlowWithGlyph. Expected text", expectedTxt, "got", resultTxt)
	}
	if resultGlyph != expectedGlyph {
		t.Error("DoubleToDmsHappyFlowWithGlyph. Expected glyph", expectedGlyph, "got", resultGlyph)
	}
}

func TestDoubleToDmsZeroDegreeWithGlyph(t *testing.T) {
	converter := NewDoubleToDmsConversions()
	inputValue := 0.0
	expectedTxt := "0" + DEGREE_SIGN + "00" + MINUTE_SIGN + "00" + SECOND_SIGN
	expectedGlyph := '1'
	resultTxt, resultGlyph := converter.ConvertDoubleToDmsWithGlyph(inputValue)
	if resultTxt != expectedTxt {
		t.Error("DoubleToDmsZeroDegreeWithGlyph. Expected text", expectedTxt, "got", resultTxt)
	}
	if resultGlyph != expectedGlyph {
		t.Error("DoubleToDmsZeroDegreeWithGlyph. Expected glyph", expectedGlyph, "got", resultGlyph)
	}
}
