/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package research

import (
	"enigma-ar/domain"
	"testing"
)

func TestCreateControlGroupTooFewItems(t *testing.T) {
	multiplFactor := 1
	inputItems := []domain.StandardInputItem{
		{
			ID:           "1",
			Name:         "Number 1",
			GeoLongitude: 7.0,
			GeoLatitude:  -8.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
	}
	cgs := NewCGroupServices()
	result, err := cgs.CreateControlGroup(inputItems, multiplFactor)
	if result != nil {
		t.Errorf("CreateControlGroup returned a non-nil result for too few items")
	}
	if err == nil {
		t.Errorf("CreateControlGroup returned a nil error for too few items")
	}
}

func TestCreateControlGroupMultiplFactorTooSmall(t *testing.T) {
	multiplFactor := -1
	inputItems := []domain.StandardInputItem{
		{
			ID:           "1",
			Name:         "Number 1",
			GeoLongitude: 7.0,
			GeoLatitude:  -8.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
		{
			ID:           "2",
			Name:         "Number 2",
			GeoLongitude: 7.0,
			GeoLatitude:  -8.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
	}
	cgs := NewCGroupServices()
	result, err := cgs.CreateControlGroup(inputItems, multiplFactor)
	if result != nil {
		t.Errorf("CreateControlGroup returned a non-nil result for a multiplication factor that is too small")
	}
	if err == nil {
		t.Errorf("CreateControlGroup returned a nil error a multiplication factor that is too small")
	}
}

func TestCreateControlGroupMultiplFactorTooLarge(t *testing.T) {
	multiplFactor := domain.MaxMultiplicationCGroups + 1
	inputItems := []domain.StandardInputItem{
		{
			ID:           "1",
			Name:         "Number 1",
			GeoLongitude: 7.0,
			GeoLatitude:  -8.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
		{
			ID:           "2",
			Name:         "Number 2",
			GeoLongitude: 7.0,
			GeoLatitude:  -8.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
	}
	cgs := NewCGroupServices()
	result, err := cgs.CreateControlGroup(inputItems, multiplFactor)
	if result != nil {
		t.Errorf("CreateControlGroup returned a non-nil result for a multiplication factor that is too large")
	}
	if err == nil {
		t.Errorf("CreateControlGroup returned a nil error a multiplication factor that is too large")
	}
}

func TestCreateControlNonUniqueId(t *testing.T) {
	multiplFactor := 10
	inputItems := []domain.StandardInputItem{
		{
			ID:           "1",
			Name:         "Number 1",
			GeoLongitude: 7.0,
			GeoLatitude:  -8.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
		{
			ID:           "2",
			Name:         "Number 2",
			GeoLongitude: 7.0,
			GeoLatitude:  -8.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
		{
			ID:           "1",
			Name:         "Another number 1",
			GeoLongitude: 7.0,
			GeoLatitude:  -8.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
	}
	cgs := NewCGroupServices()
	result, err := cgs.CreateControlGroup(inputItems, multiplFactor)
	if result != nil {
		t.Errorf("CreateControlGroup returned a non-nil result for a multiplication factor that is too large")
	}
	if err == nil {
		t.Errorf("CreateControlGroup returned a nil error for a multiplication factor that is too large")
	}
}

func TestCreateControlGroupLongitudeTooLarge(t *testing.T) {
	multiplFactor := 10
	inputItems := []domain.StandardInputItem{
		{
			ID:           "1",
			Name:         "Number 1",
			GeoLongitude: 7.0,
			GeoLatitude:  -8.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
		{
			ID:           "2",
			Name:         "Number 2",
			GeoLongitude: 187.0,
			GeoLatitude:  -8.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
	}
	cgs := NewCGroupServices()
	result, err := cgs.CreateControlGroup(inputItems, multiplFactor)

	if result != nil {
		t.Errorf("CreateControlGroup returned a nil error for a longitude that is too large")
	}
	if err == nil {
		t.Errorf("CreateControlGroup returned a nil error for a longitude that is too large")
	}
}

func TestCreateControlGroupLongitudeTooSmall(t *testing.T) {
	multiplFactor := 10
	inputItems := []domain.StandardInputItem{
		{
			ID:           "1",
			Name:         "Number 1",
			GeoLongitude: 7.0,
			GeoLatitude:  -8.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
		{
			ID:           "2",
			Name:         "Number 2",
			GeoLongitude: -187.0,
			GeoLatitude:  -8.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
	}
	cgs := NewCGroupServices()
	result, err := cgs.CreateControlGroup(inputItems, multiplFactor)
	if result != nil {
		t.Errorf("CreateControlGroup returned a non-nil result for a longitude that is too small")
	}
	if err == nil {
		t.Errorf("CreateControlGroup returned a nil error for a longitude that is too small")
	}
}

func TestCreateControlGroupLatitudeTooLarge(t *testing.T) {
	multiplFactor := 10
	inputItems := []domain.StandardInputItem{
		{
			ID:           "1",
			Name:         "Number 1",
			GeoLongitude: 7.0,
			GeoLatitude:  -8.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
		{
			ID:           "2",
			Name:         "Number 2",
			GeoLongitude: 7.0,
			GeoLatitude:  90.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
	}
	cgs := NewCGroupServices()
	result, err := cgs.CreateControlGroup(inputItems, multiplFactor)
	if result != nil {
		t.Errorf("CreateControlGroup returned a non-nil result for a latitude that is too large")
	}
	if err == nil {
		t.Errorf("CreateControlGroup returned a nil error for a latitude that is too large")
	}
}

func TestCreateControlGroupLatitudeTooSmall(t *testing.T) {
	multiplFactor := 10
	inputItems := []domain.StandardInputItem{
		{
			ID:           "1",
			Name:         "Number 1",
			GeoLongitude: 7.0,
			GeoLatitude:  -98.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
		{
			ID:           "2",
			Name:         "Number 2",
			GeoLongitude: 7.0,
			GeoLatitude:  9.0,
			DateTime: domain.DateTimeHms{
				Year:  2000,
				Month: 1,
				Day:   1,
				Hour:  1,
				Min:   1,
				Sec:   1,
				Greg:  true,
				Dst:   0,
				TZone: 0,
			},
		},
	}
	cgs := NewCGroupServices()
	result, err := cgs.CreateControlGroup(inputItems, multiplFactor)
	if result != nil {
		t.Errorf("CreateControlGroup returned a non-nil result for a latitude that is too small")
	}
	if err == nil {
		t.Errorf("CreateControlGroup returned a nil error for a latitude that is too small")
	}
}
