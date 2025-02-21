/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package conversion

import (
	"enigma-ar/domain"
	"errors"
	"math"
	"strconv"
)

// functions for conversion between decimal and sexagesimal values

// ParseDateTimeFromText converts a slice of strings into an instance of domain.SimpleDateTime.
// Should contain at least 3 items for year, month, date and optionally items for hour, minute, second.
// Optional items default to 0.
func ParseDateTimeFromText(items []string) (domain.SimpleDateTime, error) {
	emptySdt := domain.SimpleDateTime{
		Year:  0,
		Month: 0,
		Day:   0,
		Hour:  0,
		Min:   0,
		Sec:   0,
	}
	if len(items) < 3 {
		return emptySdt, errors.New("not enough items to define a date")
	}

	y, mo, d, h, mi, s := 0, 0, 0, 0, 0, 0
	tempY, err := strconv.Atoi(items[0])
	if err != nil {
		return emptySdt, err
	}
	y = tempY
	tempMo, err := strconv.Atoi(items[1])
	if err != nil {
		return emptySdt, err
	}
	mo = tempMo
	tempD, err := strconv.Atoi(items[2])
	if err != nil {
		return emptySdt, err
	}
	d = tempD
	// items for time are optional
	if len(items) > 3 {
		tempH, err := strconv.Atoi(items[3])
		if err != nil {
			return emptySdt, err
		}
		h = tempH
	}
	if len(items) > 4 {
		tempMi, err := strconv.Atoi(items[4])
		if err != nil {
			return emptySdt, err
		}
		mi = tempMi
	}
	if len(items) > 5 {
		tempS, err := strconv.Atoi(items[5])
		if err != nil {
			return emptySdt, err
		}
		s = tempS
	}
	sdt := domain.SimpleDateTime{
		Year:  y,
		Month: mo,
		Day:   d,
		Hour:  h,
		Min:   mi,
		Sec:   s,
	}
	return sdt, nil
}

// ParseHmsFromText define time as a float64 from strings for hour, minutes, and seconds.
// If items are invalid or missing, a vlaue of 0 is used.
func ParseHmsFromText(hTxt, mTxt, sTxt string) float64 {
	h, m, s := 0, 0, 0
	if hTxt != "" {
		tempH, err := strconv.Atoi(hTxt)
		if err != nil {
			h = 0
		} else {
			h = tempH
		}
	}
	if mTxt != "" {
		tempM, err := strconv.Atoi(mTxt)
		if err != nil {
			m = 0
		} else {
			m = tempM
		}
	}
	if sTxt != "" {
		tempS, err := strconv.Atoi(sTxt)
		if err != nil {
			s = 0
		} else {
			s = tempS
		}
	}
	return float64(h) + float64(m)/60.0 + float64(s)/3600.0
}

// ParseSexTextFromFloat returns a text with a sexagesimal value
func ParseSexTextFromFloat(value float64) string {
	hd := math.Trunc(value)
	remaining := value - hd + 1e-12 // add minor amount to prevent rounding problems
	mFrac := remaining * 60.0
	m := math.Trunc(mFrac)
	remaining = mFrac - m
	s := math.Trunc(remaining * 60.0)
	mPre := ""
	if m < 10.0 {
		mPre = "0"
	}
	sPre := ""
	if s < 10.0 {
		sPre = "0"
	}
	sep := ":"
	hdTxt := strconv.FormatFloat(hd, 'f', 0, 64)
	mTxt := strconv.FormatFloat(m, 'f', 0, 64)
	sTxt := strconv.FormatFloat(s, 'f', 0, 64)
	parsedText := hdTxt + sep + mPre + mTxt + sep + sPre + sTxt
	return parsedText
}
