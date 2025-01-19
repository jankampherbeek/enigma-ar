/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package persistency

import (
	"bufio"
	"enigma-ar/domain"
	"fmt"
	"os"
	"strings"
)

const (
	countriesFile = ".." + domain.PathSep + ".." + domain.PathSep + "data" + domain.PathSep + "countries.csv"
	citiesFile    = ".." + domain.PathSep + ".." + domain.PathSep + "data" + domain.PathSep + "cities.csv"
	regionsFile   = ".." + domain.PathSep + ".." + domain.PathSep + "data" + domain.PathSep + "regions.csv"
	itemSeparator = ";"
)

type LocationHandler interface {
	Countries() ([]domain.Country, error)
	Cities(countryCode string) ([]domain.City, error)
}

type LocationHandling struct{}

func NewLocationHandling() LocationHandler {
	return LocationHandling{}
}

// Countries returns all available countries
func (lh LocationHandling) Countries() ([]domain.Country, error) {
	file, err := os.Open(countriesFile)
	if err != nil {
		fmt.Printf("Error opening countries file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	var countries []domain.Country
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), itemSeparator)
		if len(fields) == 3 { // there is a field for continent that is currently not used
			country := domain.Country{
				Code: strings.TrimSpace(fields[0]),
				Name: strings.TrimSpace(fields[1]),
			}
			countries = append(countries, country)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading countries file: %v\n", err)
		return nil, err
	}

	return countries, nil
}

// Cities returns all the cities for a given country
func (lh LocationHandling) Cities(countryCode string) ([]domain.City, error) {
	file, err := os.Open(citiesFile)
	if err != nil {
		fmt.Printf("Error opening cities file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	var cities []domain.City
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), itemSeparator)
		if len(fields) == 7 && fields[0] == countryCode {
			regionCode := strings.TrimSpace(fields[0] + "." + fields[4])
			regionName, err := lh.findRegionName(regionCode)
			if err != nil {
				// TODO handle error
			}
			city := domain.City{
				Country:      countryCode,
				Name:         strings.TrimSpace(fields[1]),
				GeoLat:       strings.TrimSpace(fields[2]),
				GeoLong:      strings.TrimSpace(fields[3]),
				Region:       regionName,
				Elevation:    strings.TrimSpace(fields[5]),
				IndicationTz: strings.TrimSpace(fields[6]),
			}
			cities = append(cities, city)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading cities file: %v\n", err)
		return nil, err
	}

	return cities, nil

}

// findRegionName looks up the name for a region, using the region code.
func (lh LocationHandling) findRegionName(regionCode string) (string, error) {
	file, err := os.Open(regionsFile)
	if err != nil {
		fmt.Printf("Error opening regions file: %v\n", err)
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), itemSeparator)
		if len(fields) == 2 && fields[0] == regionCode {
			return strings.TrimSpace(fields[1]), nil
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading cities file: %v\n", err)
		return "", err
	}
	return "", nil
}
