/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package apilocandzone

import (
	"enigma-ar/domain"
	"enigma-ar/internal/locandzone"
	"errors"
	"log/slog"
)

// TimeZoneServer provides services for time zones
type TimeZoneServer interface {
	ActualTimeZone(dateTime domain.DateTimeHms, tzIndication string) (locandzone.ZoneInfo, error)
}

type TimeZoneService struct {
	tzHandler locandzone.TimezoneHandler
}

func NewTimeZoneService() TimeZoneServer {
	return TimeZoneService{
		tzHandler: locandzone.NewTimezoneHandling(),
	}
}

// ActualTimeZone retrieves information about the actual time zone from the tz database
// PRE tzIndication contains at least 5 characters
// POST if no errors occurred --> returns offset with offset in seconds, indication of DST and the name of the zone,
// otherwise --> returns an empty zone info and an error
func (tzs TimeZoneService) ActualTimeZone(dateTime domain.DateTimeHms, tzIndication string) (locandzone.ZoneInfo, error) {
	emptyZone := locandzone.ZoneInfo{Offset: 0, ZoneName: "", DST: false}
	if len(tzIndication) < 5 {
		slog.Error("Received invalid tzIndication: " + tzIndication)
		return emptyZone, errors.New("invalid tzIndication")
	}
	offset, err := tzs.tzHandler.ActualTimezone(dateTime, tzIndication)
	if err != nil {
		// TODO handle error
		return emptyZone, err
	}
	return offset, nil
}
