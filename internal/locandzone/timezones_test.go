/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package locandzone

//func TestActualTimezone(t *testing.T) {
//
//	tzh := NewTimezoneHandling()
//	dt := domain.DateTimeHms{
//		Year:  1953,
//		Month: 1,
//		Day:   29,
//		Hour:  8,
//		Min:   37,
//		Sec:   30,
//		Greg:  true,
//		Dst:   0,
//		TZone: 0,
//	}
//	tzName := "Europe/Amsterdam"
//	result, err := tzh.ActualTimezone(dt, tzName)
//	if err != nil {
//		t.Errorf("ActualTimezone failed: %v", err)
//	}
//	if result.DST == true {
//		t.Errorf("ActualTimezone returned DST true")
//	}
//
//}
//
//func TestNumberFromDayTextHappyFlow(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	dayTxt := "Wed"
//	expected := 2
//	result, err := tzh.nrFromDayText(dayTxt)
//	if err != nil {
//		t.Errorf("Error parsing day text for day '%s': %s", dayTxt, err)
//	}
//	if result != expected {
//		t.Errorf("Expected %d, got %d", expected, result)
//	}
//}
//
//func TestNumberFromDayTextWrongInput(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	dayTxt := "Abc"
//	expected := -1
//	result, err := tzh.nrFromDayText(dayTxt)
//	if err == nil {
//		t.Errorf("Expected error for wrong input")
//	}
//	if result != expected {
//		t.Errorf("Expected %d for wrong input, got %d", expected, result)
//	}
//}
//
//func TestDayFromDefinitionLastSun(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	// jan 2025, last sunday 26
//	year := 2025
//	month := 1
//	def := "lastSun"
//	expected := 26
//	result, err := tzh.dayFromDefinition(year, month, def)
//	if err != nil {
//		t.Errorf("Error parsing day text for day '%s': %s", def, err)
//	}
//	if result != expected {
//		t.Errorf("Expected %d, got %d", expected, result)
//	}
//}
//
//func TestDayFromDefinitionFirstOccurrence(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	// jan 2025, first sunday 5
//	year := 2025
//	month := 1
//	def := "Sun>=1"
//	expected := 5
//	result, err := tzh.dayFromDefinition(year, month, def)
//	if err != nil {
//		t.Errorf("Error parsing day text for day '%s': %s", def, err)
//	}
//	if result != expected {
//		t.Errorf("Expected %d, got %d", expected, result)
//	}
//}
//
//func TestDayFromDefinitionSecondOccurrence(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	// jan 2025, 2nd sunday 12
//	year := 2025
//	month := 1
//	def := "Sun>=2"
//	expected := 12
//	result, err := tzh.dayFromDefinition(year, month, def)
//	if err != nil {
//		t.Errorf("Error parsing day text for day '%s': %s", def, err)
//	}
//	if result != expected {
//		t.Errorf("Expected %d, got %d", expected, result)
//	}
//}
//
//func TestDayFromDefinitionLastWed(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	// jan 2025, last wednesday 29
//	year := 2025
//	month := 1
//	def := "lastWed"
//	expected := 29
//	result, err := tzh.dayFromDefinition(year, month, def)
//	if err != nil {
//		t.Errorf("Error parsing day text for day '%s': %s", def, err)
//	}
//	if result != expected {
//		t.Errorf("Expected %d, got %d", expected, result)
//	}
//}
//
//func TestDayFromDefinitionFirstWed(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	// jan 2025, last wednesday 29
//	year := 2025
//	month := 1
//	def := "Wed>=1"
//	expected := 1
//	result, err := tzh.dayFromDefinition(year, month, def)
//	if err != nil {
//		t.Errorf("Error parsing day text for day '%s': %s", def, err)
//	}
//	if result != expected {
//		t.Errorf("Expected %d, got %d", expected, result)
//	}
//}
//
//func TestDayFromDefinitionSecondWed(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	// jan 2025, last wednesday 29
//	year := 2025
//	month := 1
//	def := "Wed>=2"
//	expected := 8
//	result, err := tzh.dayFromDefinition(year, month, def)
//	if err != nil {
//		t.Errorf("Error parsing day text for day '%s': %s", def, err)
//	}
//	if result != expected {
//		t.Errorf("Expected %d, got %d", expected, result)
//	}
//}
//
//func TestDayFromDefinitionWrongInput(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	year := 2025
//	month := 1
//	def := "dummy"
//	expected := -1
//	result, err := tzh.dayFromDefinition(year, month, def)
//	if err == nil {
//		t.Errorf("Expected error for wrong input")
//	}
//	if result != expected {
//		t.Errorf("Expected %d for wrong input, got %d", expected, result)
//	}
//}
//
//// TestReadDstRules might require a change if the rues for Mauritius will change
//func TestReadDstRules(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	ruleName := "Mauritius"
//	expected := 4
//	result, err := tzh.readDstRules(ruleName)
//	if err != nil {
//		t.Errorf("Error reading rules: %s", err)
//	}
//	if len(result) != expected {
//		t.Errorf("Expected %d, got %d", expected, len(result))
//	}
//}
//
//func TestReadDstRulesUnknownRule(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	ruleName := "abc"
//	expected := 0
//	result, err := tzh.readDstRules(ruleName)
//	if err != nil {
//		t.Errorf("Error reading rules: %s", err)
//	}
//	if len(result) != expected {
//		t.Errorf("Expected %d, got %d", expected, len(result))
//	}
//}
//
//func TestParseTimeZeroTime(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	time := "0:00"
//	expected := 0.0
//	result, err := tzh.parseTime(time)
//	if err != nil {
//		t.Errorf("Error parsing time: %s", err)
//	}
//	if result != expected {
//		t.Errorf("Expected %f, got %f", expected, result)
//	}
//}
//
//func TestParseTimeOnlyHours(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	time := "13"
//	expected := 13.0
//	result, err := tzh.parseTime(time)
//	if err != nil {
//		t.Errorf("Error parsing time: %s", err)
//	}
//	if math.Abs(result-expected) > 1e-8 {
//		t.Errorf("Expected %f, got %f", expected, result)
//	}
//}
//
//func TestParseTimeHoursMinutes(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	time := "10:30"
//	expected := 10.5
//	result, err := tzh.parseTime(time)
//	if err != nil {
//		t.Errorf("Error parsing time: %s", err)
//	}
//	if math.Abs(result-expected) > 1e-8 {
//		t.Errorf("Expected %f, got %f", expected, result)
//	}
//}
//
//func TestParseTimeSPostFix(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	time := "23:30s"
//	expected := 23.5
//	result, err := tzh.parseTime(time)
//	if err != nil {
//		t.Errorf("Error parsing time: %s", err)
//	}
//	if math.Abs(result-expected) > 1e-8 {
//		t.Errorf("Expected %f, got %f", expected, result)
//	}
//}
//
//func TestParseTimeDPostFix(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	time := "4:45d"
//	expected := 4.75
//	result, err := tzh.parseTime(time)
//	if err != nil {
//		t.Errorf("Error parsing time: %s", err)
//	}
//	if math.Abs(result-expected) > 1e-8 {
//		t.Errorf("Expected %f, got %f", expected, result)
//	}
//}
//
//func TestParseTimeError(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	time := "abc"
//	expected := 0.0
//	result, err := tzh.parseTime(time)
//	if err == nil {
//		t.Errorf("Expected error for wrong input: %s", err)
//	}
//	if math.Abs(result-expected) > 1e-8 {
//		t.Errorf("Expected %f, got %f", expected, result)
//	}
//}
//
//func TestParseDstLine(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	line := "Rule	Morocco	1940	only	-	Feb	25	 0:00	1:00	-" // tab separated
//
//	expected := dstDefLine{
//		StartYear:       1940,
//		To:              "only",
//		StartMonth:      2,
//		RuleForStartDay: "25",
//		TimeOfChange:    0.0,
//		OffsetLT:        1.0,
//		StartDst:        true,
//	}
//	result, err := tzh.parseDstLine(line)
//	if err != nil {
//		t.Errorf("Error parsing line: %s", err)
//	}
//	if result != expected {
//		t.Errorf("Expected %v, got %v", expected, result)
//	}
//}
//
//func TestReadTzLines(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	tzName := "America/Ensenada"
//	expected := 7
//	result, err := tzh.readTzLines(tzName)
//	if err != nil {
//		t.Errorf("Error reading timezones: %s", err)
//	}
//	if len(result) != expected {
//		t.Errorf("Expected %d, got %d", expected, len(result))
//	}
//}
//
//func TestReadTzLinesWithLink(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	tzName := "America/St_Thomas"
//	expected := 3
//	result, err := tzh.readTzLines(tzName)
//	if err != nil {
//		t.Errorf("Error reading timezones: %s", err)
//	}
//	if len(result) != expected {
//		t.Errorf("Expected %d, got %d", expected, len(result))
//	}
//}
//
//func TestReadTzLinesWrongName(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	tzName := "abc"
//	expected := 0
//	result, err := tzh.readTzLines(tzName)
//	if err != nil {
//		t.Errorf("Error reading timezones: %s", err)
//	}
//	if len(result) != expected {
//		t.Errorf("Expected %d, got %d", expected, len(result))
//	}
//}
//
//func TestParseTzTopLine(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	line := "Zone Pacific/Port_Moresby 9:48:40 -	LMT	1880"
//	expected := tzTopLine{
//		Name:      "Pacific/Port_Moresby",
//		OffsetUT:  9 + 48/60.0 + 40/3600.0,
//		Abbr:      "LMT",
//		StartYear: 1880,
//	}
//	result, err := tzh.parseTzTopLine(line)
//	if err != nil {
//		t.Errorf("Error parsing timezones: %s", err)
//	}
//	if result != expected {
//		t.Errorf("Expected %v, got %v", expected, result)
//	}
//}
//
//// TestParseTzTopLineWithRemark will become obsolete as a future solution of reading the lines will remove end remarks
//func TestParseTzTopLineWithRemark(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	line := "Zone Pacific/Guadalcanal 10:39:48 -\tLMT\t1912 Oct  1 # Honiara"
//	expected := tzTopLine{
//		Name:      "Pacific/Guadalcanal",
//		OffsetUT:  10 + 39/60.0 + 48/3600.0,
//		Abbr:      "LMT",
//		StartYear: 1912,
//	}
//	result, err := tzh.parseTzTopLine(line)
//	if err != nil {
//		t.Errorf("Error parsing timezones: %s", err)
//	}
//	if result != expected {
//		t.Errorf("Expected %v, got %v", expected, result)
//	}
//}
//
//func TestParsTzDefLineShortLine(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	line := "-8:00\tVanc\tP%sT\t1987"
//	expected := tzDefLine{
//		OffsetUT:        -8.0,
//		DstRuleName:     "Vanc",
//		Abbr:            "P%sT",
//		EndYear:         1987,
//		EndMonth:        0,
//		EndDay:          0,
//		EndTime:         0,
//		EndDateTimeText: "19870000",
//	}
//	result, err := tzh.parseTzDefLine(line)
//	if err != nil {
//		t.Errorf("Error parsing timezones: %s", err)
//	}
//	if result != expected {
//		t.Errorf("Expected %v, got %v", expected, result)
//	}
//
//}
//
//func TestParsTzDefLineLongLine(t *testing.T) {
//	tzh := NewTimezoneHandling()
//	line := "-5:00\tToronto\tE%sT\t1942 Feb  9  2:00s"
//	expected := tzDefLine{
//		OffsetUT:        -5.0,
//		DstRuleName:     "Toronto",
//		Abbr:            "E%sT",
//		EndYear:         1942,
//		EndMonth:        2,
//		EndDay:          9,
//		EndTime:         2.0,
//		EndDateTimeText: "19420209",
//	}
//	result, err := tzh.parseTzDefLine(line)
//	if err != nil {
//		t.Errorf("Error parsing timezones: %s", err)
//	}
//	if result != expected {
//		t.Errorf("Expected %v, got %v", expected, result)
//	}
//
//}
