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
	"log/slog"
	"os"
	"strings"
)

const filePathZones = ".." + domain.PathSep + ".." + domain.PathSep + "data" + domain.PathSep + "zones.cv"

type TzHandler interface {
	CurrentTime(dateTime domain.DateTimeHms, tzName string) (ZoneInfo, error)
}

type TzHandling struct {
	jdCalc     calc.JulDayCalculator
	dstHandler DstHandler
}

func NewTzHandling() TzHandler {
	jdCalc := calc.NewJulDayCalculation()
	dstHandler := NewDstHandling()
	return TzHandling{jdCalc: jdCalc,
		dstHandler: dstHandler}
}

// tzLine represents textual data for time zones, the names of the fields correspond with definitions in the tz database
type tzLine struct {
	Name   string
	StdOff float64
	Rules  string
	Format string
	Until  float64 // date/time converted to Julian day number
}

type ZoneInfo struct {
	Offset float64
	TzName string
	Dst    bool
}

func (tzh TzHandling) CurrentTime(dateTime domain.DateTimeHms, tzGroupName string) (ZoneInfo, error) {
	// find Offset, TzName and dstRuleName for given dateTime
	emptyZoneInfo := ZoneInfo{
		Offset: 0,
		TzName: "",
		Dst:    false,
	}
	dstOffset := 0.0
	dstUsed := false
	zoneOffset, tzName, dstRule, err := tzh.zoneData(dateTime, tzGroupName)
	if err != nil {
		slog.Error("Could not define time zone")
		return emptyZoneInfo, err
	}
	if len(dstRule) >= 2 { // ignoring hyphen and empty string
		dstUsed = true
		dst, err := tzh.dstHandler.CurrentDst(dateTime, dstRule)
		if err != nil {
			return emptyZoneInfo, err
		}
		dstOffset = dst.Offset
		strings.Replace(tzName, "%s", dst.Letter, 1)
	}
	if strings.Contains(tzName, "%z") {
		tzName = conversion.ParseSexTextFromFloat(zoneOffset)
	}
	zoneInfo := ZoneInfo{
		Offset: zoneOffset + dstOffset,
		TzName: tzName,
		Dst:    dstUsed,
	}

	return zoneInfo, nil
}

// zoneData calculates the values for a time zone, it returns the offset, the name and the dst rule name
func (tzh TzHandling) zoneData(dateTime domain.DateTimeHms, tzName string) (float64, string, string, error) {
	zoneTxtLines, err := tzh.readTzLines(tzName)
	if err != nil {
		slog.Error("Reading lines from the tz file returns an error")
		return 0.0, "", "", err
	}
	zoneLines, err := tzh.parseTzLines(zoneTxtLines, tzName)
	if err != nil {
		slog.Error("Parsing lines from the tz file returns an error")
		return 0.0, "", "", err
	}
	actualZone, err := tzh.findZone(dateTime, zoneLines)
	if err != nil {
		slog.Error("Finding zone from the tz file returns an error")
		return 0.0, "", "", err
	}
	offset := actualZone.StdOff
	name := actualZone.Name
	dstRule := actualZone.Rules
	return offset, name, dstRule, nil

}

// readTzLines reads all lines for a given time zone
func (tzh TzHandling) readTzLines(tzName string) ([]string, error) {
	searchTxt1 := "Zone\t" + tzName
	searchTxt2 := "Zone " + tzName
	var tzLines []string
	tzFile, err := os.Open(filePathZones)
	if err != nil {
		slog.Error("Could not open tz file")
		return tzLines, err
	}
	defer func(tzFile *os.File) {
		err := tzFile.Close()
		if err != nil {
			slog.Error("Error closing tz file")
		}
	}(tzFile)
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

func (tzh TzHandling) parseTzLines(lines []string, name string) ([]tzLine, error) {
	parsedLines := make([]tzLine, 0)
	for _, line := range lines {
		dataLine := line
		if strings.HasPrefix(line, "Zone;") {
			dataLine = strings.TrimPrefix(line, "Zone;")
			index := strings.Index(dataLine, ";")
			dataLine = dataLine[index+1:] // remove tz name
		}
		items := strings.Split(dataLine, ";")
		offset := conversion.ParseHmsFromText(items[0], items[1], items[2])
		sdt, err := conversion.ParseDateTimeFromText(items[3:])
		if err != nil {
			return nil, err
		}
		ut := float64(sdt.Hour) + float64(sdt.Min)/60.0 + float64(sdt.Sec)/3600.0
		until := tzh.jdCalc.CalcJd(sdt.Hour, sdt.Min, sdt.Month, ut, true) // always Gregorian

		tzLine := tzLine{
			Name:   name,
			StdOff: offset,
			Rules:  items[3],
			Format: items[4],
			Until:  until,
		}
		parsedLines = append(parsedLines, tzLine)
	}
	return parsedLines, nil
}

func (tzh TzHandling) findZone(dateTime domain.DateTimeHms, lines []tzLine) (tzLine, error) {
	time := float64(dateTime.Hour) + float64(dateTime.Min)/60.0 + float64(dateTime.Sec)/3600.0
	jd := tzh.jdCalc.CalcJd(dateTime.Year, dateTime.Month, dateTime.Day, time, true)
	counter := 0
	line := lines[0]
	for _, newLine := range lines[1:] {
		if newLine.Until < jd {
			line = lines[counter]
			continue
		}
		counter++
	}
	return line, nil
}
