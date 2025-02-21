/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package conversion

import (
	"enigma-ar/domain"
	"math"
	"testing"
)

func TestParseDateTimeFromTextHappyFlow(t *testing.T) {
	items := []string{"1953", "1", "29", "8", "37", "30"}
	expected := domain.SimpleDateTime{
		Year:  1953,
		Month: 1,
		Day:   29,
		Hour:  8,
		Min:   37,
		Sec:   30,
	}
	result, err := ParseDateTimeFromText(items)
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Fatalf("expected %v but got %v", expected, result)
	}
}

func TestParseDateTimeFromTextNoTime(t *testing.T) {
	items := []string{"2025", "2", "19"}
	expected := domain.SimpleDateTime{
		Year:  2025,
		Month: 2,
		Day:   19,
		Hour:  0,
		Min:   0,
		Sec:   0,
	}
	result, err := ParseDateTimeFromText(items)
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Fatalf("expected %v but got %v", expected, result)
	}
}

func TestParseDateTimeFromTextNoMinutesSeconds(t *testing.T) {
	items := []string{"2025", "2", "19", "14"}
	expected := domain.SimpleDateTime{
		Year:  2025,
		Month: 2,
		Day:   19,
		Hour:  14,
		Min:   0,
		Sec:   0,
	}
	result, err := ParseDateTimeFromText(items)
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Fatalf("expected %v but got %v", expected, result)
	}
}

func TestParseDateTimeFromTextNoSeconds(t *testing.T) {
	items := []string{"2025", "2", "19", "14", "33"}
	expected := domain.SimpleDateTime{
		Year:  2025,
		Month: 2,
		Day:   19,
		Hour:  14,
		Min:   33,
		Sec:   0,
	}
	result, err := ParseDateTimeFromText(items)
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Fatalf("expected %v but got %v", expected, result)
	}
}

func TestParseDateTimeFromTextIncomplete(t *testing.T) {
	items := []string{"2025", "2"}
	expected := domain.SimpleDateTime{
		Year:  0,
		Month: 0,
		Day:   0,
		Hour:  0,
		Min:   0,
		Sec:   0,
	}
	result, err := ParseDateTimeFromText(items)
	if err == nil {
		t.Fatal("expected error")
	}
	if result != expected {
		t.Fatalf("expected %v for incomplete input, but got %v", expected, result)
	}
}

func TestParseDateTimeFromTextInvalidInput(t *testing.T) {
	items := []string{"2025", "2", "19", "14", "33", "ab"}
	expected := domain.SimpleDateTime{
		Year:  0,
		Month: 0,
		Day:   0,
		Hour:  0,
		Min:   0,
		Sec:   0,
	}
	result, err := ParseDateTimeFromText(items)
	if err == nil {
		t.Fatal("expected error")
	}
	if result != expected {
		t.Fatalf("expected %v for invalid input, but got %v", expected, result)
	}
}

func TestParseHmsFromTextHappyFlow(t *testing.T) {
	hTxt := "14"
	mTxt := "50"
	sTxt := "15"
	expected := 14.0 + 50.0/60.0 + 15.0/3600.0
	result := ParseHmsFromText(hTxt, mTxt, sTxt)
	if math.Abs(result-expected) > 1e-8 {
		t.Fatalf("expected %v but got %v", expected, result)
	}
}

func TestParseHmsFromTextEmptyItems(t *testing.T) {
	hTxt := ""
	mTxt := ""
	sTxt := ""
	expected := 0.0
	result := ParseHmsFromText(hTxt, mTxt, sTxt)
	if math.Abs(result-expected) > 1e-8 {
		t.Fatalf("expected %v but got %v", expected, result)
	}
}

func TestParseHmsFromTextInvalidItems(t *testing.T) {
	hTxt := "22"
	mTxt := "xyz"
	sTxt := "0"
	expected := 22.0
	result := ParseHmsFromText(hTxt, mTxt, sTxt)
	if math.Abs(result-expected) > 1e-8 {
		t.Fatalf("expected %v but got %v", expected, result)
	}
}

func TestParseSexTextFromFloat(t *testing.T) {
	value := 1.5
	expected := "1:30:00"
	result := ParseSexTextFromFloat(value)
	if result != expected {
		t.Fatalf("expected %v but got %v", expected, result)
	}
}
