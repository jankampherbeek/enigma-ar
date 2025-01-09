/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package research

import (
	"enigma-ar/domain"
	"fmt"
	"sort"
)

type StandardInputItem struct {
	ID           string
	Name         string
	GeoLongitude float64
	GeoLatitude  float64
	DateTime     domain.DateTimeHms
}

type ControlGroupCreator interface {
	CreateMultipleControlData(inputItems []StandardInputItem, multiplicity int) []StandardInputItem
}

type ControlGroupCreation struct {
	dataHandler       CGDataHandler
	cgRandomizer      CGroupRandomizer
	controlGroupItems []StandardInputItem
	years             []int
	months            []int
	days              []int
	hours             []int
	minutes           []int
	seconds           []int
	dsts              []float64
	zoneOffsets       []float64
	latitudes         []float64
	longitudes        []float64
}

func NewControlGroupCreation() ControlGroupCreator {
	return ControlGroupCreation{
		dataHandler:  NewCGDataHandling(),
		cgRandomizer: NewCGroupRandomization(),
	}
}

func (cgc ControlGroupCreation) CreateMultipleControlData(inputItems []StandardInputItem, multiplicity int) []StandardInputItem {
	allControlData := make([]StandardInputItem, 0)
	for i := 0; i < multiplicity; i++ {
		controlDataForOneSet := cgc.createControlData(inputItems, i)
		allControlData = append(allControlData, controlDataForOneSet...)
	}
	return allControlData
}

func (cgc ControlGroupCreation) createControlData(inputItems []StandardInputItem, sequence int) []StandardInputItem {
	cgc.controlGroupItems = cgc.controlGroupItems[:0] // Clear slice
	cgc.processInputData(inputItems)
	cgc.sortDaysAndShuffleOtherItems()
	cgc.processData(sequence)
	return cgc.controlGroupItems
}

func (cgc ControlGroupCreation) processInputData(inputItems []StandardInputItem) {

	cgc.years = cgc.years[:0]
	cgc.months = cgc.months[:0]
	cgc.days = cgc.days[:0]
	cgc.hours = cgc.hours[:0]
	cgc.minutes = cgc.minutes[:0]
	cgc.seconds = cgc.seconds[:0]
	cgc.dsts = cgc.dsts[:0]
	cgc.zoneOffsets = cgc.zoneOffsets[:0]
	cgc.latitudes = cgc.latitudes[:0]
	cgc.longitudes = cgc.longitudes[:0]

	for _, inputItem := range inputItems {
		cgc.years = append(cgc.years, inputItem.DateTime.Year)
		cgc.months = append(cgc.months, inputItem.DateTime.Month)
		cgc.days = append(cgc.days, inputItem.DateTime.Day)
		cgc.hours = append(cgc.hours, inputItem.DateTime.Hour)
		cgc.minutes = append(cgc.minutes, inputItem.DateTime.Min)
		cgc.seconds = append(cgc.seconds, inputItem.DateTime.Sec)
		cgc.dsts = append(cgc.dsts, inputItem.DateTime.Dst)
		cgc.zoneOffsets = append(cgc.zoneOffsets, inputItem.DateTime.TZone)
		cgc.latitudes = append(cgc.latitudes, inputItem.GeoLatitude)
		cgc.longitudes = append(cgc.longitudes, inputItem.GeoLongitude)
	}
}

func (cgc ControlGroupCreation) sortDaysAndShuffleOtherItems() {
	// Note: You'll need to implement a sort.Sort interface or use sort.Slice for days
	// and implement the shuffle functionality in IControlGroupRng

	// Sort days in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(cgc.days)))

	cgc.cgRandomizer.ShuffleIntList(cgc.years)
	cgc.cgRandomizer.ShuffleIntList(cgc.months)
	cgc.cgRandomizer.ShuffleIntList(cgc.days)
	cgc.cgRandomizer.ShuffleIntList(cgc.hours)
	cgc.cgRandomizer.ShuffleIntList(cgc.minutes)
	cgc.cgRandomizer.ShuffleIntList(cgc.seconds)
	cgc.cgRandomizer.ShuffleFloatList(cgc.dsts)
	cgc.cgRandomizer.ShuffleFloatList(cgc.zoneOffsets)
	cgc.cgRandomizer.ShuffleFloatList(cgc.latitudes)
	cgc.cgRandomizer.ShuffleFloatList(cgc.longitudes)
}

func (cgc ControlGroupCreation) processData(sequence int) {
	counter := 0
	for len(cgc.years) > 0 {
		year := getFromIntList(&cgc.years)
		day := getFromIntList(&cgc.days)
		month := cgc.findMonth(day, year)
		hour := getFromIntList(&cgc.hours)
		minute := getFromIntList(&cgc.minutes)
		second := getFromIntList(&cgc.seconds)
		dst := getFromFloat64List(&cgc.dsts)
		//	zoneOffset := getFromFloat64List(&cgc.zoneOffsets)
		latitude := getFromFloat64List(&cgc.latitudes)
		longitude := getFromFloat64List(&cgc.longitudes)

		dateTime := domain.DateTimeHms{Year: year, Month: month, Day: day, Hour: hour, Min: minute, Sec: second, Dst: dst}
		id := counter
		name := fmt.Sprintf("Controldata %d-%d", sequence, id)
		item := StandardInputItem{
			ID:           fmt.Sprintf("%d-%d", sequence, id),
			Name:         name,
			GeoLongitude: longitude,
			GeoLatitude:  latitude,
			DateTime:     dateTime,
		}
		cgc.controlGroupItems = append(cgc.controlGroupItems, item)
		counter++
	}
}

func (cgc ControlGroupCreation) findMonth(day, year int) int {
	for i, month := range cgc.months {
		if cgc.dataHandler.DayFitsInMonth(day, month, year) {
			// Remove the month from the slice
			cgc.months = append(cgc.months[:i], cgc.months[i+1:]...)
			return month
		}
	}
	return 0
}

func getFromIntList(list *[]int) int {
	if len(*list) == 0 {
		return 0
	}
	result := (*list)[0]
	*list = (*list)[1:]
	return result
}

func getFromFloat64List(list *[]float64) float64 {
	if len(*list) == 0 {
		return 0
	}
	result := (*list)[0]
	*list = (*list)[1:]
	return result
}
