/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package locandzone

import (
	"enigma-ar/internal/calc"
	"enigma-ar/internal/se"
	"errors"
	"log/slog"
	"strconv"
	"strings"
)

const (
	pfLast = "last"
	pfGE1  = ">=1"
	pfGE2  = ">=2"
)

// DayDefHandler handles defining a date based on a definition.
type DayDefHandler interface {
	DayFromDefinition(year, month int, def string) (int, error)
}

type DayDefHandling struct {
	jdCalc  calc.JulDayCalculator
	dowCalc se.SwephDayOfWeekCalculator
}

func NewDayDefHandling() DayDefHandler {
	julDayCalculation := calc.NewJulDayCalculation()
	dowCalc := se.NewSwephDayOfWeekCalculation()
	return DayDefHandling{
		jdCalc:  julDayCalculation,
		dowCalc: dowCalc,
	}
}

// DayFromDefinition calculates the day number for a given definition in a given year and mont.
func (dd DayDefHandling) DayFromDefinition(year, month int, def string) (int, error) {
	var defDay, defType string
	if len(def) <= 2 {
		preDefinedDay, err := strconv.Atoi(def)
		if err != nil {
			return -1, err
		}
		return preDefinedDay, nil
	}
	if strings.Contains(def, pfLast) {
		defDay = def[4:]
		defType = pfLast
	} else if strings.Contains(def, pfGE1) {
		index := strings.Index(def, pfGE1)
		defDay = def[index-1 : index]
		defType = pfGE1
	} else if strings.Contains(def, pfGE2) {
		index := strings.Index(def, pfGE2)
		defDay = def[index-1 : index]
		defType = pfGE2
	} else {
		defType = "Unknown defDay"
	}
	switchDay, err := strconv.Atoi(defDay)
	if err != nil {
		slog.Error("could not parse defDay: " + def)
		return -1, errors.New("could not parse DefDay: " + def)
	}
	jd := dd.jdCalc.CalcJd(year, month, 1, 12.0, true) // jd for first day of month
	var firstDOW = dd.dowCalc.DayOfWeek(jd)            // index for first day of month, Mon=0...Sun=7
	var actualDay int
	switch defType {
	case pfLast:
		m31 := []int{1, 3, 5, 7, 8, 10, 12}
		if dd.contains(m31, month) {
			lastDayOfMonth := firstDOW + 30
			diff := lastDayOfMonth%7 - switchDay
			if diff < 0 {
				diff += 7
			}
			actualDay = 31 - diff
		} else { // assuming the last days of February are never used for a DST switch
			lastDayOfMonth := firstDOW + 29
			diff := lastDayOfMonth%7 - switchDay
			if diff < 0 {
				diff += 7
			}
			actualDay = 30 - diff
		}
	case pfGE1:
		diff := switchDay - firstDOW
		actualDay = 1 + diff
	case pfGE2:
		diff := switchDay - firstDOW
		actualDay = 8 + diff
	default:
		slog.Error("unknown def type: " + def)
		return -1, errors.New("unknown def type: " + def)
	}
	return actualDay, nil
}

// contains is a helper function for dayFromdefinition()
func (dd DayDefHandling) contains(numbers []int, num int) bool {
	for _, n := range numbers {
		if n == num {
			return true
		}
	}
	return false
}
