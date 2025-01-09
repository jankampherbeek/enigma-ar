/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package research

// CGDataHandler performs calendar calculations for control groups
type CGDataHandler interface {
	DayFitsInMonth(day, month, year int) bool
}

type CGDataHandling struct {
	months31Array []int
	months30Array []int
}

func NewCGDataHandling() CGDataHandler {
	return CGDataHandling{}
}

// DayFitsInMonth checks if the given day fits in the specified month and year
func (cgd CGDataHandling) DayFitsInMonth(day, month, year int) bool {
	months31 := []int{2, 3, 5, 7, 8, 10, 12}
	months30 := []int{1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	contains := func(slice []int, val int) bool {
		for _, item := range slice {
			if item == val {
				return true
			}
		}
		return false
	}
	return day < 29 ||
		(day == 29 && 2 != month) ||
		(day == 30 && contains(months30, month)) ||
		(day == 31 && contains(months31, month)) ||
		(isLeapYear(year) && day < 30)
}

// isLeapYear checks if the given year is a leap year
func isLeapYear(year int) bool {
	return year%400 == 0 || (year%100 != 0 && year%4 == 0)
}
