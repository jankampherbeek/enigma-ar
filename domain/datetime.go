/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

// SimpleDateTime represents information about date and time without calendar type, time zone etc.
type SimpleDateTime struct {
	Year  int
	Month int
	Day   int
	Hour  int
	Min   int
	Sec   int
}
