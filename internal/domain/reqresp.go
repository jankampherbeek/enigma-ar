/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package domain

// FloatResponse Response for a floating point value
type FloatResponse struct {
	Value  float64 `json:"value"`
	Result int     `json:"result"`
}

// JulDayRequest Request for the calculation of a Julian Day number
type JulDayRequest struct {
	Year  int     `json:"year"`
	Month int     `json:"month"`
	Day   int     `json:"day"`
	Ut    float64 `json:"ut"`
	Greg  bool    `json:"greg"`
}
