/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package locandzone

import (
	"bufio"
	"enigma-ar/domain"
	"enigma-ar/internal/se"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

const (
	filePathZones = ".." + domain.PathSep + ".." + domain.PathSep + "data" + domain.PathSep + "zones.txt"
	filePathRules = ".." + domain.PathSep + ".." + domain.PathSep + "data" + domain.PathSep + "rules.txt"
)

type dstDefLine struct {
	StartYear       int
	To              string
	StartMonth      int
	RuleForStartDay string
	TimeOfChange    float64
	OffsetLT        float64
	StartDst        bool
}

type tzDefLine struct {
	OffsetUT        float64
	DstRuleName     string
	Abbr            string
	EndYear         int
	EndMonth        int
	EndDay          int
	EndTime         float64
	EndDateTimeText string
}

type tzTopLine struct {
	Name      string
	OffsetUT  float64
	Abbr      string
	StartYear int
}

type dateTimeForDst struct {
	Year   int
	Month  int
	Day    int
	Time   float64
	Offset float64
}

type TimezoneHandler interface {
	ActualTimezone(dateTime domain.DateTimeHms, tzName string) (ZoneInfo, error)
	constructDst(dt domain.DateTimeHms, dstLines []string) (float64, error)
	nrFromDayText(dayText string) (int, error)
	defineDstStartEndDate(year int, month, onValue string) (string, error)
	dayFromDefinition(year, month int, def string) (int, error)
	readDstRules(ruleName string) ([]string, error)
	readTzLines(tzName string) ([]string, error)
	parseDstLine(line string) (dstDefLine, error)
	parseTzTopLine(line string) (tzTopLine, error)
	parseTzDefLine(line string) (tzDefLine, error)
	parseTime(time string) (float64, error)
	createTextFromDateTime(date dateTimeForDst) string
	firstDateIsGreater(date1 dateTimeForDst, date2 dateTimeForDst) bool
}

type TimezoneHandling struct {
	dowCalc se.SwephDayOfWeekCalculator
	jdCalc  se.SwephJulDayCalculator
}

func NewTimezoneHandling() TimezoneHandler {
	return TimezoneHandling{
		dowCalc: se.NewSwephDayOfWeekCalculation(),
		jdCalc:  se.NewSwephJulDayCalculation(),
	}
}

type ZoneInfo struct {
	Offset   float64
	ZoneName string
	DST      bool
	CalcLmt  bool
}

// ActualTimezone reads time zone information from IANA tz database files.
func (tzh TimezoneHandling) ActualTimezone(dt domain.DateTimeHms, tzName string) (ZoneInfo, error) {

	tzLines, err := tzh.readTzLines(tzName)
	if err != nil {
		return ZoneInfo{}, err
	}
	tzTopLine, err := tzh.parseTzTopLine(tzLines[0])
	if err != nil {
		return ZoneInfo{}, err
	}
	var tzDefLines []tzDefLine
	for _, tzLine := range tzLines[1:] {
		tzDefLine, err := tzh.parseTzDefLine(tzLine)
		if err != nil {
			return ZoneInfo{}, err
		}
		tzDefLines = append(tzDefLines, tzDefLine)
	}
	if dt.Year < tzTopLine.StartYear { // Before start of tz definitions
		return ZoneInfo{
			Offset:   0.0,
			ZoneName: "LMT",
			DST:      false,
			CalcLmt:  true,
		}, nil
	}
	monthPrefix := ""
	if dt.Month < 10 {
		monthPrefix = "0"
	}
	dayPrefix := ""
	if dt.Day < 10 {
		dayPrefix = "0"
	}
	dtText := strconv.Itoa(dt.Year) + monthPrefix + strconv.Itoa(dt.Month) + dayPrefix + strconv.Itoa(dt.Day)

	prevDef := ZoneInfo{
		Offset:   tzTopLine.OffsetUT,
		ZoneName: tzTopLine.Abbr,
		DST:      false,
		CalcLmt:  false,
	}
	for _, tzDefLine := range tzDefLines {

		if dtText < tzDefLine.EndDateTimeText {
			dstRule := tzDefLine.DstRuleName
			dstLines, err := tzh.readDstRules(dstRule)
			if err != nil {
				return ZoneInfo{}, err
			}
			dstOffset := 0.0
			if len(dstLines) > 0 {
				dstOffset, err = tzh.constructDst(dt, dstLines)
				if err != nil {
					return ZoneInfo{}, err
				}
			}
			return ZoneInfo{
				Offset:   prevDef.Offset + dstOffset,
				ZoneName: prevDef.ZoneName,
				DST:      dstOffset > 0.0,
				CalcLmt:  false,
			}, nil
		}

	}

	// find effective rule in timezone
	// check for DST
	// if dst: read rules for DST
	//         apply DST rules
	// construct ZonerInfo and return it

	return tzh.DefineZoneForDate(dt, tzName, filePathZones, filePathRules)
}

func (tzh TimezoneHandling) constructDst(dt domain.DateTimeHms, dstLines []string) (float64, error) {

	var dstDefs []dstDefLine
	var dstOffset float64
	for _, line := range dstLines {
		dstDef, err := tzh.parseDstLine(line)
		if err != nil {
			return 0.0, err
		}
		dstDefs = append(dstDefs, dstDef)
	}
	var startFound, endFound bool
	var startDT, endDT, actDT dateTimeForDst
	actDT = dateTimeForDst{
		Year:   dt.Year,
		Month:  dt.Sec,
		Day:    dt.Day,
		Time:   float64(dt.Hour) + float64(dt.Min)*60.0 + float64(dt.Sec)*3600.0,
		Offset: 0.0,
	}
	for _, dstDef := range dstDefs {
		year1 := dstDef.StartYear
		year2 := 0
		if dstDef.To == "only" {
			year2 = year1
		} else if dstDef.To == "max" {
			year2 = 100_000 // there is no end year, this number should suffice
		} else {
			y, err := strconv.Atoi(dstDef.To)
			if err != nil {
				return 0.0, err
			}
			year2 = y
		}
		day, err := tzh.dayFromDefinition(year1, dstDef.StartMonth, dstDef.RuleForStartDay)
		if err != nil {
			return 0.0, err
		}
		if dt.Year > year1 && dt.Month > dstDef.StartMonth && dt.Day > day {
			if !startFound && dstDef.OffsetLT > 0.000001 {
				startFound = true
				startDT = dateTimeForDst{
					Year:   year1,
					Month:  dstDef.StartMonth,
					Day:    day,
					Time:   dstDef.TimeOfChange,
					Offset: dstDef.OffsetLT,
				}
			}
			if !endFound && dstDef.OffsetLT < 0.000001 {
				endFound = true
				endDT = dateTimeForDst{
					Year:   year2,
					Month:  dstDef.StartMonth,
					Day:    day,
					Time:   dstDef.TimeOfChange,
					Offset: dstDef.OffsetLT,
				}
			}
		}
	}
	if tzh.firstDateIsGreater(actDT, startDT) && tzh.firstDateIsGreater(endDT, actDT) {
		dstOffset = startDT.Offset
	} else {
		dstOffset = endDT.Offset
	}
	return dstOffset, nil
}

func (tzh TimezoneHandling) firstDateIsGreater(date1 dateTimeForDst, date2 dateTimeForDst) bool {
	dTxt1 := tzh.createTextFromDateTime(date1)
	dTxt2 := tzh.createTextFromDateTime(date2)
	if dTxt1 > dTxt2 {
		return true
	}
	return false
}

func (tzh TimezoneHandling) createTextFromDateTime(date dateTimeForDst) string {
	var dSpacer, mSpacer string
	if date.Month < 10 {
		mSpacer = " "
	}
	if date.Day < 10 {
		dSpacer = " "
	}
	yTxt := strconv.Itoa(date.Year)
	mTxt := mSpacer + strconv.Itoa(date.Month)
	dTxt := dSpacer + strconv.Itoa(date.Day)
	return yTxt + mTxt + dTxt
}

func (tzh TimezoneHandling) DefineZoneForDate(dt domain.DateTimeHms, tzName, fpZones, fpRules string) (ZoneInfo, error) {
	var zoneLines []string
	searchTerm := "Zone\t" + tzName
	emptyZone := ZoneInfo{Offset: 0, ZoneName: "", DST: false}
	inputFile, err := os.Open(fpZones)
	if err != nil {
		return emptyZone, err
	}
	defer inputFile.Close()
	zoneFound := false
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		// Remove trailing whitespace
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, searchTerm) && !zoneFound {
			zoneLines = append(zoneLines, line)
			zoneFound = true
		}
		if zoneFound {
			if strings.HasPrefix(line, "Zone") {
				break
			}
		}
	}
	// find value for actual date

	// check for DST rules
	// create result and return it

	return emptyZone, err
}

func (tzh TimezoneHandling) defineDstStartEndDate(year int, month, onValue string) (string, error) {
	// Try to check for explicit date
	num, err := strconv.Atoi(onValue)
	var m, dm string
	if err != nil {
		// startdate is explicitly defined
		m, err = tzh.monthIdFromText(month)
		if err != nil {
			slog.Error("impossible month/onValue for DST rule in timezones: %s/%s", month, onValue)
			return "", err
		}
		prefix := ""
		if num < 10 {
			prefix = "0"
		}
		dm = prefix + onValue + "." + m
	} else {
		// handle definition
		monthNr, err := tzh.monthIdFromText(month)
		if err != nil {
			return "", err
		}
		monthId, err := strconv.Atoi(monthNr)
		if err != nil {
			return "", err
		}
		dayNr, err := tzh.dayFromDefinition(year, monthId, onValue)
		spacer := ""
		if dayNr < 10 {
			spacer = "0"
		}
		dayTxt := spacer + strconv.Itoa(dayNr)
		dm = monthNr + "." + dayTxt
	}
	return dm, nil
}

func (tzh TimezoneHandling) monthIdFromText(month string) (string, error) {
	monthId := ""
	switch month {
	case "Jan":
		monthId = "01"
	case "Feb":
		monthId = "02"
	case "Mar":
		monthId = "03"
	case "Apr":
		monthId = "04"
	case "May":
		monthId = "05"
	case "Jun":
		monthId = "06"
	case "Jul":
		monthId = "07"
	case "Aug":
		monthId = "08"
	case "Sep":
		monthId = "09"
	case "Oct":
		monthId = "10"
	case "Nov":
		monthId = "11"
	case "Dec":
		monthId = "12"
	default:
		return "", errors.New("invalid month " + month)
	}
	return monthId, nil
}

func (tzh TimezoneHandling) dayFromDefinition(year, month int, def string) (int, error) {
	var defDay, defType string
	if strings.HasPrefix(def, "last") {
		defDay = def[4:]
		defType = "last"
	} else if def[3:] == ">=1" {
		defDay = def[:3]
		defType = ">=1"
	} else if def[3:] == ">=2" {
		defDay = def[:3]
		defType = ">=2"
	} else {
		// unknown deftype
		slog.Error("encountered unknown def: " + def)
		return -1, errors.New("unknown def: " + def)
	}
	switchDay, err := tzh.nrFromDayText(defDay)
	if err != nil {
		return -1, err
	}
	jd := tzh.jdCalc.CalcJd(year, month, 1, 12.0, 1) // jd for first day of month
	firstDOW := tzh.dowCalc.DayOfWeek(jd)            // index for first day of month, Mon=0...Sun=7
	var actualDay int
	switch defType {
	case "last":
		m31 := []int{1, 3, 5, 7, 8, 10, 12}
		if tzh.contains(m31, month) {
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
	case ">=1":
		diff := switchDay - firstDOW
		actualDay = 1 + diff
	case ">=2":
		diff := switchDay - firstDOW
		actualDay = 8 + diff
	}
	return actualDay, nil
}

// NrFromDayText is a helper function for dayFromDefinition
func (tzh TimezoneHandling) nrFromDayText(dayText string) (int, error) {
	switch dayText {
	case "Mon":
		return 0, nil
	case "Tue":
		return 1, nil
	case "Wed":
		return 2, nil
	case "Thu":
		return 3, nil
	case "Fri":
		return 4, nil
	case "Sat":
		return 5, nil
	case "Sun":
		return 6, nil
	default:
		return -1, errors.New("day text is invalid")
	}
}

// contains is a helper function for dayFromdefinition()
func (tzh TimezoneHandling) contains(numbers []int, num int) bool {
	for _, n := range numbers {
		if n == num {
			return true
		}
	}
	return false
}

// readDstRules reads the lines for the given ruleName from the DST rules file
func (tzh TimezoneHandling) readDstRules(ruleName string) ([]string, error) {
	searchTxt := "Rule " + ruleName
	var ruleLines []string
	rulesFile, err := os.Open(filePathRules)
	if err != nil {
		fmt.Errorf("error opening rules file: %v", err)
	}

	defer rulesFile.Close()

	scanner := bufio.NewScanner(rulesFile)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, searchTxt) {
			ruleLines = append(ruleLines, line)
		}
	}
	return ruleLines, nil
}

// readTzLines reads the lines for the given timezone from the timezones file
func (tzh TimezoneHandling) readTzLines(tzName string) ([]string, error) {
	searchTxt1 := "Zone\t" + tzName
	searchTxt2 := "Zone " + tzName
	var tzLines []string
	tzFile, err := os.Open(filePathZones)
	if err != nil {
		fmt.Errorf("error opening tz file: %v", err)
	}
	defer tzFile.Close()
	startLineFound := false
	scanner := bufio.NewScanner(tzFile)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, searchTxt1) || strings.HasPrefix(line, searchTxt2) {
			tzLines = append(tzLines, line)
			startLineFound = true
		} else {
			if startLineFound {
				if !strings.HasPrefix(line, "Zone") {
					tzLines = append(tzLines, line)
				} else {
					startLineFound = false
				}
			}
		}
	}
	return tzLines, nil
}

func (tzh TimezoneHandling) parseDstLine(line string) (dstDefLine, error) {
	emptyDef := dstDefLine{
		StartYear:       0,
		To:              "",
		StartMonth:      0,
		RuleForStartDay: "",
		TimeOfChange:    0,
		OffsetLT:        0,
		StartDst:        false,
	}
	// Example of rule line with indexes:
	// 0    1       2       3       4   5   6   7       8       9
	// Rule	Neth	1917	only	-	Apr	16	2:00s	1:00	NST

	items := strings.Split(line, "\t")
	y, err := strconv.Atoi(items[2])
	if err != nil {
		return emptyDef, err
	}
	mId, err := tzh.monthIdFromText(items[5])
	if err != nil {
		return emptyDef, err
	}
	m, err := strconv.Atoi(strings.TrimSpace(mId))
	if err != nil {
		return emptyDef, err
	}
	time, err := tzh.parseTime(items[7])
	if err != nil {
		return emptyDef, err
	}
	offset, err := tzh.parseTime(items[8])
	if err != nil {
		return emptyDef, err
	}
	started := offset > 0.0000001 // minimal above zero to prevent rounding problems
	def := dstDefLine{
		StartYear:       y,
		To:              items[3],
		StartMonth:      m,
		RuleForStartDay: items[6],
		TimeOfChange:    time,
		OffsetLT:        offset,
		StartDst:        started,
	}
	return def, nil
}

func (tzh TimezoneHandling) parseTzTopLine(line string) (tzTopLine, error) {

	// TODO handle lines with year mont and day as startingpoint
	// e.g. Zone Pacific/Guadalcanal 10:39:48 -	LMT	1912 Oct  1 # Honiara

	emptyTzTopLine := tzTopLine{
		Name:      "",
		OffsetUT:  0,
		Abbr:      "",
		StartYear: 0,
	}
	strippedLine := strings.TrimSpace(line)
	strippedLine = strings.ReplaceAll(strippedLine, "\t", " ")
	items := strings.Split(strippedLine, " ")

	// example of tz top line with indexes
	// 0    1                   2       3   4   5
	// Zone Europe/Amsterdam	0:19:32 -	LMT	1835
	offsetUT, err := tzh.parseTime(items[2])
	if err != nil {
		return emptyTzTopLine, err
	}
	startYear, err := strconv.Atoi(items[5])
	if err != nil {
		return emptyTzTopLine, err
	}
	tzTop := tzTopLine{
		Name:      items[1],
		OffsetUT:  offsetUT,
		Abbr:      items[4],
		StartYear: startYear,
	}
	return tzTop, nil
}

func (tzh TimezoneHandling) parseTzDefLine(line string) (tzDefLine, error) {
	emptyTzDefLine := tzDefLine{
		OffsetUT:        0,
		DstRuleName:     "",
		Abbr:            "",
		EndYear:         0,
		EndMonth:        0,
		EndDay:          0,
		EndTime:         0,
		EndDateTimeText: "",
	}
	strippedLine := strings.TrimSpace(line)
	// line can have a mix of tabs, spaces and multiple spaces
	strippedLine = strings.ReplaceAll(strippedLine, "\t", " ")
	strippedLine = strings.ReplaceAll(strippedLine, "     ", " ") // assuming max 5 spaces
	strippedLine = strings.ReplaceAll(strippedLine, "    ", " ")
	strippedLine = strings.ReplaceAll(strippedLine, "   ", " ")
	strippedLine = strings.ReplaceAll(strippedLine, "  ", " ")

	items := strings.Split(strippedLine, " ")
	// example tzDefLine with indexes
	// 0    1       2       3    4    5  6       Items 4, 5 and 6 are optional
	// 1:00	C-Eur	CE%sT	1945 Apr  2  2:00
	offsetUT, err := tzh.parseTime(items[0])
	if err != nil {
		return emptyTzDefLine, err
	}
	endYear := 0
	if len(items) > 3 {
		endYear, err = strconv.Atoi(items[3])
		if err != nil {
			return emptyTzDefLine, err
		}
	}
	endMonth := 0
	endMonthTxt := "00"
	if len(items) > 4 {
		endMonthTxt, err = tzh.monthIdFromText(items[4])
		if err != nil {
			return emptyTzDefLine, err
		}
		endMonth, err = strconv.Atoi(endMonthTxt)
		if err != nil {
			return emptyTzDefLine, err
		}
	}
	endDay := 0
	endDayTxt := "00"
	spacer := ""
	if len(items) > 5 {
		endDayTxt = items[5]
		endDay, err = strconv.Atoi(endDayTxt)
		if err != nil {
			return emptyTzDefLine, err
		}
		if len(items[5]) == 1 {
			spacer = "0"
		}
	}
	endTime := 0.0

	if len(items) > 6 {
		endTime, err = tzh.parseTime(items[6])
		if err != nil {
			return emptyTzDefLine, err
		}
	}
	endDateTimeTxt := items[3] + endMonthTxt + spacer + endDayTxt
	tzDefLine := tzDefLine{
		OffsetUT:        offsetUT,
		DstRuleName:     items[1],
		Abbr:            items[2],
		EndYear:         endYear,
		EndMonth:        endMonth,
		EndDay:          endDay,
		EndTime:         endTime,
		EndDateTimeText: endDateTimeTxt,
	}
	return tzDefLine, nil
}

func (tzh TimezoneHandling) parseTime(time string) (float64, error) {
	items := strings.Split(time, ":")
	h, err := strconv.Atoi(strings.TrimSpace(items[0]))
	if err != nil {
		return 0.0, err
	}
	m := 0
	s := 0
	if len(items) > 1 {
		if strings.Contains(items[1], "s") {
			items[1] = strings.Replace(items[1], "s", "", -1)
		}
		if strings.Contains(items[1], "d") {
			items[1] = strings.Replace(items[1], "d", "", -1)
		}
		m, err = strconv.Atoi(strings.TrimSpace(items[1]))
		if err != nil {
			return 0.0, err
		}
	}
	if len(items) > 2 {
		s, err = strconv.Atoi(strings.TrimSpace(items[2]))
		if err != nil {
			return 0.0, err
		}
	}
	return float64(h) + float64(m)/60.0 + float64(s)/3600, nil
}
