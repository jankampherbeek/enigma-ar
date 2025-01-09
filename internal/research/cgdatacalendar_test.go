/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package research

import "testing"

func TestDayFitsInMonthHappyFlow(t *testing.T) {
	c := NewCGDataHandling()
	year := 2025
	month := 1
	day := 9
	if !c.DayFitsInMonth(day, month, year) {
		t.Error("DayFitsInMonth returned false, should be true")
	}
}

func TestDayFitsInMonth2001Febr29(t *testing.T) {
	c := NewCGDataHandling()
	year := 2001
	month := 2
	day := 29
	if c.DayFitsInMonth(year, month, day) {
		t.Error("Expected false but got true for Febr 29 in non leap year")
	}
}

func TestDayFitsInMonth2004Febr29(t *testing.T) {
	c := NewCGDataHandling()
	year := 2004 // leap year
	month := 2
	day := 29
	if !c.DayFitsInMonth(day, month, year) {
		t.Error("DayFitsInMonth returned false for Febr 29 in leap year, should be true")
	}
}

func TestDayFitsInMonthApril31(t *testing.T) {
	c := NewCGDataHandling()
	year := 2001
	month := 4
	day := 31
	if c.DayFitsInMonth(year, month, day) {
		t.Error("Expected false but got true for April 31")
	}
}
