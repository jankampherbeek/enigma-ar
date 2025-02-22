/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package locandzone

import (
	"enigma-ar/internal/calc"
	"enigma-ar/internal/calc/conversion"
	"enigma-ar/internal/se"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

// DstParser handles parsing dst lines.
type DstParser interface {
	ProcessDstLines(lines []string) ([]dstLine, error)
}

type DstParsing struct {
	jdCalc    calc.JulDayCalculator
	dowCalc   se.SwephDayOfWeekCalculator
	dayNrCalc DayDefHandler
}

func NewDstParsing() DstParser {
	return DstParsing{
		jdCalc:    calc.NewJulDayCalculation(),
		dowCalc:   se.NewSwephDayOfWeekCalculation(),
		dayNrCalc: NewDayDefHandling(),
	}
}

func (dp DstParsing) ProcessDstLines(lines []string) ([]dstLine, error) {
	elementsLines, err := dp.parseDstElementsLines(lines)
	if err != nil {
		return nil, err
	}
	processedLines, err := dp.parseDstLines(elementsLines)
	if err != nil {
		return nil, err
	}
	return processedLines, nil
}

func (dp DstParsing) parseDstElementsLines(lines []string) ([]dstElementsLine, error) {
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
		oset, err := conversion.ParseDateTimeFromText(items[8:11])
		if err != nil {
			return nil, fmt.Errorf("invalid value for offset in dataLine: %s", dataLine)
		}
		offset := float64(oset.Hour) + float64(oset.Min)/60.0 + float64(oset.Sec)/3600.0
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

func (dp DstParsing) parseDstLines(lines []dstElementsLine) ([]dstLine, error) {
	parsedLines := make([]dstLine, 0)
	for _, line := range lines {
		startYear := line.From
		endYear := line.To
		for year := startYear; year <= endYear; year++ {
			newLine, err := dp.createSingleDstLine(line)
			if err != nil {
				return nil, err
			}
			parsedLines = append(parsedLines, newLine)
		}
	}
	return parsedLines, nil
}

func (dp DstParsing) createSingleDstLine(line dstElementsLine) (dstLine, error) {
	emptyDstLine := dstLine{
		startJd: 0,
		offset:  0,
		letter:  "",
	}
	day, err := dp.dayNrCalc.DayFromDefinition(line.From, line.In, line.On) // resp. year, month and day definition
	if err != nil {
		slog.Error("Error getting day definition")
		return emptyDstLine, err
	}

	jd := dp.jdCalc.CalcJd(line.From, line.In, day, line.At, true) // always Gregorian
	newDstLine := dstLine{
		startJd: jd,
		offset:  line.Save,
		letter:  line.Letter,
	}
	return newDstLine, nil
}
