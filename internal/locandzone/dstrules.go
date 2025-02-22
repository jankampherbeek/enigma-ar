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
	"enigma-ar/internal/se"
	"log/slog"
	"os"
	"sort"
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

type DstInfo struct {
	Offset float64
	Letter string
}

type DstHandler interface {
	CurrentDst(dateTime domain.DateTimeHms, dstRule string) (DstInfo, error)
}

type DstHandling struct {
	jdCalc         calc.JulDayCalculator
	dowCalc        se.SwephDayOfWeekCalculator
	dayNrCalc      DayDefHandler
	dstLinesParser DstParser
}

func NewDstHandling() DstHandler {
	return DstHandling{
		jdCalc:         calc.NewJulDayCalculation(),
		dowCalc:        se.NewSwephDayOfWeekCalculation(),
		dayNrCalc:      NewDayDefHandling(),
		dstLinesParser: NewDstParsing(),
	}
}

func (dh DstHandling) CurrentDst(dateTime domain.DateTimeHms, dstRule string) (DstInfo, error) {
	emptyDstInfo := DstInfo{
		Offset: 0,
		Letter: "",
	}
	var actDstLine dstLine
	dstLines, err := dh.dstData(dstRule)
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
	newDstInfo := DstInfo{
		Offset: actDstLine.offset,
		Letter: actDstLine.letter,
	}
	return newDstInfo, nil
}

func (dh DstHandling) dstData(dstRule string) ([]dstLine, error) {

	dstTxtLines, err := dh.readDstLines(dstRule)
	if err != nil {
		slog.Error("Reading lines from the dst file returns an error")
		return nil, err
	}
	processedLines, err2 := dh.dstLinesParser.ProcessDstLines(dstTxtLines)
	if err2 != nil {
		slog.Error("Processing lines from the dst file returns an error")
		return nil, err2
	}
	return processedLines, nil
}

func (dh DstHandling) readDstLines(ruleName string) ([]string, error) {
	var dstTxtLines []string
	dstFile, err := os.Open(filePathRules)
	if err != nil {
		return dstTxtLines, err
	}
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			slog.Error("Error closing dst file")
		}
	}(dstFile)
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
