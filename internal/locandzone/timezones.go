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
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const TzFile = ".." + domain.PathSep + ".." + domain.PathSep + "data" + domain.PathSep + "tz.txt"

type TimeZoneHandler interface {
	ActualTimeZone(dateTime domain.DateTimeHms, tzIndication string) (string, int, error)
}

type TimeZoneHandling struct{}

func NewTimeZoneHandling() TimeZoneHandler {
	return TimeZoneHandling{}
}

func (tzh TimeZoneHandling) ActualTimeZone(dt domain.DateTimeHms, tzIndication string) (string, int, error) {
	// Load the timezone data
	err := LoadTimeZoneData(TzFile)
	if err != nil {
		fmt.Printf("Error loading timezone data: %v\n", err)
		return "", 0, err
	}

	// Example usage
	targetTime := time.Date(1940, 5, 15, 12, 0, 0, 0, time.UTC)
	zoneName := "Europe/Amsterdam"

	rule, err := GetTimeZoneInfo(zoneName, targetTime)
	if err != nil {
		fmt.Printf("Error getting timezone info: %v\n", err)
		return "", 0, err
	}

	fmt.Printf("Timezone information for %s at %s:\n", zoneName, targetTime)
	fmt.Printf("Offset: %d seconds\n", rule.OffsetSec)
	fmt.Printf("Name: %s\n", rule.Name)
	fmt.Printf("DST: %v\n", rule.IsDST)

	return zoneName, rule.OffsetSec, nil

}

// TimeZoneRule represents a single timezone rule
type TimeZoneRule struct {
	From      time.Time
	To        time.Time
	OffsetSec int
	Name      string
	IsDST     bool
}

// TimeZoneData holds all rules for a specific zone
type TimeZoneData struct {
	ZoneName string
	Rules    []TimeZoneRule
}

var timeZones = make(map[string]*TimeZoneData)

// parseOffset enhanced to handle more formats
func parseOffset(offset string) (int, error) {
	// Handle combined offset format (e.g., "+0020/+0120")
	if strings.Contains(offset, "/") {
		// Use the first offset in the pair
		parts := strings.Split(offset, "/")
		offset = parts[0]
	}

	// Handle H:M:S format
	if strings.Contains(offset, ":") {
		parts := strings.Split(offset, ":")
		var hours, minutes, seconds int
		var err error

		sign := 1
		if strings.HasPrefix(parts[0], "-") {
			sign = -1
			parts[0] = strings.TrimPrefix(parts[0], "-")
		}

		if hours, err = strconv.Atoi(parts[0]); err != nil {
			return 0, err
		}
		if len(parts) > 1 {
			if minutes, err = strconv.Atoi(parts[1]); err != nil {
				return 0, err
			}
		}
		if len(parts) > 2 {
			if seconds, err = strconv.Atoi(parts[2]); err != nil {
				return 0, err
			}
		}
		return sign * (hours*3600 + minutes*60 + seconds), nil
	}

	// Rest of the existing parseOffset logic...
	// Remove any leading plus sign
	offset = strings.TrimPrefix(offset, "+")

	// Handle simple integer format
	if simpleOffset, err := strconv.Atoi(offset); err == nil {
		return simpleOffset * 3600, nil
	}

	// Handle compact format
	if len(offset) < 4 {
		return 0, fmt.Errorf("invalid offset format: %s", offset)
	}

	sign := 1
	startIdx := 0
	if offset[0] == '-' {
		sign = -1
		startIdx = 1
	}

	hours, err := strconv.Atoi(offset[startIdx : startIdx+2])
	if err != nil {
		return 0, err
	}

	minutes := 0
	if len(offset[startIdx:]) >= 4 {
		minutes, err = strconv.Atoi(offset[startIdx+2 : startIdx+4])
		if err != nil {
			return 0, err
		}
	}

	return sign * (hours*3600 + minutes*60), nil
}

// formatZoneName formats the zone name, replacing %z with the offset in hours
// enhanced to handle more placeholders
func formatZoneName(name string, offsetSec int) string {
	switch {
	case name == "%z":
		hours := offsetSec / 3600
		if hours >= 0 {
			return fmt.Sprintf("+%02d", hours)
		}
		return fmt.Sprintf("-%02d", -hours)
	case name == "LMT":
		return "LMT"
	case strings.Contains(name, "%s"):
		// Replace %s with "S" for standard time
		return strings.ReplaceAll(name, "%s", "S")
	default:
		return name
	}
}

// parseDateTime now handles UTC suffix 'u'
func parseDateTime(fields []string, startIdx int) (time.Time, int, error) {
	if startIdx >= len(fields) {
		return time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC), startIdx, nil
	}

	// Check for continuation line marker
	if fields[startIdx] == "-" {
		return time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC), startIdx + 1, nil
	}

	// Initialize defaults
	year := 1
	month := time.January
	day := 1
	hour := 0
	min := 0
	sec := 0

	fieldIdx := startIdx

	// Try to parse year
	if fieldIdx < len(fields) {
		if y, err := strconv.Atoi(fields[fieldIdx]); err == nil {
			year = y
			fieldIdx++
		}
	}

	// Try to parse month if available
	if fieldIdx < len(fields) {
		monthStr := strings.ToLower(fields[fieldIdx])
		switch monthStr {
		case "jan":
			month = time.January
		case "feb":
			month = time.February
		case "mar":
			month = time.March
		case "apr":
			month = time.April
		case "may":
			month = time.May
		case "jun":
			month = time.June
		case "jul":
			month = time.July
		case "aug":
			month = time.August
		case "sep", "sept":
			month = time.September
		case "oct":
			month = time.October
		case "nov":
			month = time.November
		case "dec":
			month = time.December
		default:
			if m, err := strconv.Atoi(monthStr); err == nil && m >= 1 && m <= 12 {
				month = time.Month(m)
			}
		}
		fieldIdx++
	}

	// Try to parse day if available
	if fieldIdx < len(fields) {
		if d, err := strconv.Atoi(fields[fieldIdx]); err == nil && d >= 1 && d <= 31 {
			day = d
			fieldIdx++
		}
	}

	// Try to parse time if available
	if fieldIdx < len(fields) {
		timeStr := fields[fieldIdx]
		// Remove 'u' suffix if present
		if strings.HasSuffix(timeStr, "u") {
			timeStr = strings.TrimSuffix(timeStr, "u")
		}

		if strings.Contains(timeStr, ":") {
			timeParts := strings.Split(timeStr, ":")
			if len(timeParts) >= 1 {
				if h, err := strconv.Atoi(timeParts[0]); err == nil {
					hour = h
				}
			}
			if len(timeParts) >= 2 {
				if m, err := strconv.Atoi(timeParts[1]); err == nil {
					min = m
				}
			}
			if len(timeParts) >= 3 {
				if s, err := strconv.Atoi(timeParts[2]); err == nil {
					sec = s
				}
			}
			fieldIdx++
		}
	}

	return time.Date(year, month, day, hour, min, sec, 0, time.UTC), fieldIdx, nil
}

func LoadTimeZoneData(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// ----
	scanner := bufio.NewScanner(file)
	var currentZone string
	var lastRule *TimeZoneRule

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) < 3 {
			continue
		}

		// Handle Zone lines and continuation lines
		if fields[0] == "Zone" || (lastRule != nil && fields[0] == "-") {
			// For continuation lines, skip the Zone name parsing
			if fields[0] == "Zone" {
				if len(fields) < 3 {
					continue // Skip invalid zone lines
				}
				currentZone = fields[1]
				fields = fields[2:] // Remove "Zone" and zone name
			} else {
				fields = fields[1:] // Remove continuation marker
			}

			if currentZone == "" {
				continue // Skip if no valid zone name
			}

			if _, exists := timeZones[currentZone]; !exists {
				timeZones[currentZone] = &TimeZoneData{
					ZoneName: currentZone,
					Rules:    make([]TimeZoneRule, 0),
				}
			}

			// Parse the offset
			if len(fields) == 0 {
				continue
			}

			offset, err := parseOffset(fields[0])
			if err != nil {
				return fmt.Errorf("error parsing offset for zone %s: %v", currentZone, err)
			}

			// Parse the until date if present
			from := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
			if lastRule != nil {
				from = lastRule.To
			}

			to := time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)
			if len(fields) > 2 {
				to, _, err = parseDateTime(fields, 2)
				if err != nil {
					return fmt.Errorf("error parsing date for zone %s: %v", currentZone, err)
				}
			}

			// Get the name, defaulting to UTC if not specified
			name := "UTC"
			if len(fields) > 1 {
				name = fields[len(fields)-1]
			}

			rule := TimeZoneRule{
				From:      from,
				To:        to,
				OffsetSec: offset,
				Name:      formatZoneName(name, offset),
				IsDST: strings.Contains(name, "DST") ||
					(strings.Contains(name, "%s") && strings.HasSuffix(name, "T")),
			}

			timeZones[currentZone].Rules = append(timeZones[currentZone].Rules, rule)
			lastRule = &rule
		} else {
			lastRule = nil
		}
	}

	return scanner.Err()
}

func GetTimeZoneInfo(zoneName string, t time.Time) (*TimeZoneRule, error) {
	zoneData, exists := timeZones[zoneName]
	if !exists {
		return nil, fmt.Errorf("timezone %s not found", zoneName)
	}

	// Find the most recent rule that applies
	var matchingRule *TimeZoneRule
	for i := range zoneData.Rules {
		rule := &zoneData.Rules[i]
		// A rule applies if:
		// 1. The time is after or equal to the rule's start time
		// 2. AND either it's before the rule's end time OR it's the last rule
		if !t.Before(rule.From) && (t.Before(rule.To) || i == len(zoneData.Rules)-1) {
			matchingRule = rule
			break
		}
	}

	if matchingRule == nil {
		return nil, fmt.Errorf("no matching rule found for %s at %s", zoneName, t)
	}

	return matchingRule, nil
}
