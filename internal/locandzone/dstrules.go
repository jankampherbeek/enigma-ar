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
	"enigma-ar/internal/calc"
	"enigma-ar/internal/calc/conversion"
	"enigma-ar/internal/se"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
)

const filePathRules = ".." + domain.PathSep + ".." + domain.PathSep + "data" + domain.PathSep + "rules.csv"

// dstTextLine represents textual data for dst, the names of the fields correspond with definitions in the tz database
type dstElementsLine struct {
	Name   string
	From   int
	To     int
	In     int
	On     string
	At     float64
	Save   float64
	Letter string
}

type dstLine struct {
	startJd float64
	offset  float64
	letter  string
}

type dstInfo struct {
	Offset float64
	Letter string
}

type DstHandler interface {
	CurrentDst(dateTime domain.DateTimeHms, dstRule string) (dstInfo, error)
}

type DstHandling struct {
	jdCalc  calc.JulDayCalculator
	dowCalc se.SwephDayOfWeekCalculator
}

func NewDstHandling() DstHandler {
	calc := calc.NewJulDayCalculation()
	dowCalc := se.NewSwephDayOfWeekCalculation()
	return DstHandling{
		jdCalc:  calc,
		dowCalc: dowCalc,
	}
}

func (dh DstHandling) CurrentDst(dateTime domain.DateTimeHms, dstRule string) (dstInfo, error) {
	emptyDstInfo := dstInfo{
		Offset: 0,
		Letter: "",
	}
	var actDstLine dstLine
	dstLines, err := dh.dstData(dateTime, dstRule)
	if err != nil {
		return emptyDstInfo, err
	}
	sort.Slice(dstLines, func(i, j int) bool {
		return dstLines[i].startJd < dstLines[j].startJd
	})
	clockTime := float64(dateTime.Hour) + float64(dateTime.Min)/60.0 + float64(dateTime.Sec)/3600.0
	jd := dh.jdCalc.CalcJd(dateTime.Year, dateTime.Month, dateTime.Day, clockTime, true) // always use Gregorian cal.
	if jd < dstLines[0].startJd {
		return emptyDstInfo, nil
	} else {
		prevDstLine := dstLines[0]
		for _, line := range dstLines {
			if line.startJd < jd {
				actDstLine = prevDstLine
			}
		}
	}
	newDstInfo := dstInfo{
		Offset: actDstLine.offset,
		Letter: actDstLine.letter,
	}
	return newDstInfo, nil
}

func (dh DstHandling) dstData(dateTime domain.DateTimeHms, dstRule string) ([]dstLine, error) {

	dstTxtLines, err := dh.readDstLines(dstRule)
	if err != nil {
		slog.Error("Reading lines from the dst file returns an error")
		return nil, err
	}
	dstElementsLines, err := dh.parseDstElementsLines(dstTxtLines)
	if err != nil {
		slog.Error("Parsing lines from the dst file returns an error")
		return nil, err
	}
	dstLines, err := dh.parseDstLines(dstElementsLines)
	if err != nil {
		slog.Error("Parsing dstLines from dstElementsLines returns an error")
		return nil, err
	}
	return dstLines, nil
}

func (dh DstHandling) readDstLines(ruleName string) ([]string, error) {
	var dstTxtLines []string
	dstFile, err := os.Open(filePathRules)
	if err != nil {
		fmt.Errorf("error opening dst file: %v", err)
	}
	defer dstFile.Close()
	scanner := bufio.NewScanner(dstFile)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ruleName) {
			dstTxtLines = append(dstTxtLines, line)
		}
	}
	return dstTxtLines, nil
}

func (dh DstHandling) parseDstElementsLines(lines []string) ([]dstElementsLine, error) {
	parsedLines := make([]dstElementsLine, 0)
	for _, line := range lines {
		dataLine := line
		items := strings.Split(dataLine, ";")
		if len(items) < 12 {
			return nil, fmt.Errorf("invalid dataLine: %s", dataLine)
		}
		from, err := strconv.Atoi(items[1])
		if err != nil {
			return nil, fmt.Errorf("invalid value for from in dataLine: %s", dataLine)
		}
		to, err := strconv.Atoi(items[2])
		if err != nil {
			return nil, fmt.Errorf("invalid value for to in dataLine: %s", dataLine)
		}
		in, err := strconv.Atoi(items[3])
		if err != nil {
			return nil, fmt.Errorf("invalid value for in in dataLine: %s", dataLine)
		}

		sdt, err := conversion.ParseDateTimeFromText(items[6:9])
		if err != nil {
			return nil, fmt.Errorf("invalid value for sdt in dataLine: %s", dataLine)
		}
		startTime := float64(sdt.Hour) + float64(sdt.Min)/60.0 + float64(sdt.Sec)/3600.0
		os, err := conversion.ParseDateTimeFromText(items[8:11])
		if err != nil {
			return nil, fmt.Errorf("invalid value for offset in dataLine: %s", dataLine)
		}
		offset := float64(os.Hour) + float64(os.Min)/60.0 + float64(os.Sec)/3600.0
		dstLine := dstElementsLine{
			Name:   items[0],
			From:   from,
			To:     to,
			In:     in,
			On:     items[4],
			At:     startTime,
			Save:   offset,
			Letter: items[11],
		}
		parsedLines = append(parsedLines, dstLine)
	}
	return parsedLines, nil
}

func (dh DstHandling) parseDstLines(lines []dstElementsLine) ([]dstLine, error) {
	parsedLines := make([]dstLine, 0)
	for _, line := range lines {
		startYear := line.From
		endYear := line.To
		for year := startYear; year <= endYear; year++ {
			newLine, err := dh.createSingleDstLine(line)
			if err != nil {
				return nil, err
			}
			parsedLines = append(parsedLines, newLine)
		}
	}
	return parsedLines, nil
}

func (dh DstHandling) createSingleDstLine(line dstElementsLine) (dstLine, error) {
	emptyDstLine := dstLine{
		startJd: 0,
		offset:  0,
		letter:  "",
	}
	day, err := dh.dayFromDefinition(line.From, line.In, line.On) // resp. year, month and day definition
	if err != nil {
		slog.Error("Error getting day definition")
		return emptyDstLine, err
	}

	jd := dh.jdCalc.CalcJd(line.From, line.In, day, line.At, true) // alwyas Gregorian
	newDstLine := dstLine{
		startJd: jd,
		offset:  line.Save,
		letter:  line.Letter,
	}
	return newDstLine, nil
}

func (dh DstHandling) dayFromDefinition(year, month int, def string) (int, error) {
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
	switchDay, err := strconv.Atoi(defDay)
	if err != nil {
		slog.Error("could not parse defDay: " + def)
		return -1, errors.New("could not parse DefDay: " + def)
	}
	jd := dh.jdCalc.CalcJd(year, month, 1, 12.0, true) // jd for first day of month
	firstDOW := dh.dowCalc.DayOfWeek(jd)               // index for first day of month, Mon=0...Sun=7
	var actualDay int
	switch defType {
	case "last":
		m31 := []int{1, 3, 5, 7, 8, 10, 12}
		if dh.contains(m31, month) {
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

// contains is a helper function for dayFromdefinition()
func (dh DstHandling) contains(numbers []int, num int) bool {
	for _, n := range numbers {
		if n == num {
			return true
		}
	}
	return false
}
