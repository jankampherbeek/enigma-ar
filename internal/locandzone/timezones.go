/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package locandzone

import (
	"enigma-ar/domain"
	"time"
	_ "time/tzdata"
)

type TimeZoneHandler interface {
	ActualTimeZone(dateTime domain.DateTimeHms, tzIndication string) (string, int, error)
}

type TimeZoneHandling struct{}

func NewTimeZoneHandling() TimeZoneHandler {
	return TimeZoneHandling{}
}

// ActualTimeZone reads time zone information from the IANA tz database that is embedded in the library tzdata.
// Currently, there is data missing, e.g. for the Netherlands. I filed a bug report.
// TODO check results of bug report
func (tzh TimeZoneHandling) ActualTimeZone(dt domain.DateTimeHms, tzIndication string) (string, int, error) {

	loc, err := time.LoadLocation(tzIndication)
	if err != nil {
		panic(err)
	}
	historicalTime := time.Date(dt.Year, time.Month(dt.Month), dt.Day, dt.Hour, dt.Min, dt.Sec, 0, loc)
	name, offset := historicalTime.Zone()
	return name, offset, nil
}
