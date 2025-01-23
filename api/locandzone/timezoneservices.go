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
	ActualTimeZone(dateTime domain.DateTimeHms, tzIndication string) (string, int, error)
}

type TimeZoneService struct {
	tzHandler locandzone.TimeZoneHandler
}

func NewTimeZoneService() TimeZoneServer {
	return TimeZoneService{
		tzHandler: locandzone.NewTimeZoneHandling(),
	}
}

// ActualTimeZone retrieves information about the actual time zone from the tz database
// PRE tzIndication contains at least 5 characters
// POST if no errors occurred --> returns abbreviated name of zone and the offset in seconds, otherwise --> returns an
// empty string, a zero and an error
func (tzs TimeZoneService) ActualTimeZone(dateTime domain.DateTimeHms, tzIndication string) (string, int, error) {
	if len(tzIndication) < 5 {
		slog.Error("Received invalid tzIndication: " + tzIndication)
		return "", 0, errors.New("invalid tzIndication")
	}
	name, offset, err := tzs.tzHandler.ActualTimeZone(dateTime, tzIndication)
	if err != nil {
		// TODO handle error
		return "", 0, err
	}
	return name, offset, nil
}
