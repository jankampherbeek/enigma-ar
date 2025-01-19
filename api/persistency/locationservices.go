/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package persistency

import (
	"enigma-ar/domain"
	"enigma-ar/internal/persistency"
	"errors"
	"log/slog"
)

// LocationServer proviodes services for locations
type LocationServer interface {
	Countries() ([]domain.Country, error)
	Cities(countryCode string) ([]domain.City, error)
}

type LocationService struct {
	locHandler persistency.LocationHandler
}

func NewLocationService() LocationServer {
	return LocationService{
		locHandler: persistency.NewLocationHandling(),
	}
}

// Countries returns all available countries
// No PRE conditions
// POST returns a slice with all countries
func (ls LocationService) Countries() ([]domain.Country, error) {
	return ls.locHandler.Countries()
}

// Cities returns all cities for a given country
// PRE countryCode is a string of 2 characters
// POST of the country for the given countryCode is found: returns a slice with all cities in that country
// otherwise returns an empty slice
func (ls LocationService) Cities(countryCode string) ([]domain.City, error) {
	slog.Info("received request for Cities")
	if len(countryCode) != 2 {
		slog.Error("country code did not contain 2 characters")
		return nil, errors.New("wrong country code, should be 2 characters")
	}
	return ls.locHandler.Cities(countryCode)
}
