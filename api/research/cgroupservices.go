/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package research

import (
	"enigma-ar/domain"
	"enigma-ar/internal/research"
	"fmt"
	"log/slog"
)

type CGroupServer interface {
	CreateControlGroup(inputItems []domain.StandardInputItem, multiplicity int) ([]domain.StandardInputItem, error)
}

type CGroupService struct {
	cgCreator research.ControlGroupCreator
}

func NewCGroupServices() *CGroupService {
	return &CGroupService{
		research.NewControlGroupCreation(),
	}
}

// CreateControlGroup create a control group from the input data
// PRE length inputItems >= 2
// PRE 1 <= multiplicity <= 1000
// PRE ID for all inputItems is unique
// PRE for all input items: -180.0 <= geoLong <= 180.0
// PRE for all input items: -90.0 < geoLat , 90.0
// Dates are assumed ot be correct
// POST no errors -> returns control group. Error -> returns nil and error
func (cgs CGroupService) CreateControlGroup(inputItems []domain.StandardInputItem, multiplicity int) ([]domain.StandardInputItem, error) {
	if len(inputItems) < domain.MinSizeCGroups {
		slog.Error("inputItems < minSizeCGroups")
		return nil, fmt.Errorf("not enough inputItems")
	}
	if multiplicity < domain.MinMultiplicationCGroups || multiplicity > domain.MaxMultiplicationCGroups {
		slog.Error("multiplicity out of range")
		return nil, fmt.Errorf("multiplicity out of range")
	}
	if !cgs.isUnique(inputItems) {
		slog.Error("inputItems not unique")
		return nil, fmt.Errorf("inputItems not unique")
	}
	for _, value := range inputItems {
		if value.GeoLatitude >= domain.MaxGeoLat || value.GeoLatitude <= domain.MinGeoLat {
			slog.Error("geoLatitude out of range")
			return nil, fmt.Errorf("geoLatitude out of range")
		}
	}
	for _, value := range inputItems {
		if value.GeoLongitude > domain.MaxGeoLong || value.GeoLongitude <= domain.MinGeoLong {
			slog.Error("geoLongitude out of range")
			return nil, fmt.Errorf("geoLongitude out of range")
		}
	}

	return cgs.CreateControlGroup(inputItems, multiplicity)
}

func (cgs CGroupService) isUnique(inputItems []domain.StandardInputItem) bool {
	seen := make(map[string]bool)
	for _, value := range inputItems {
		if seen[value.ID] {
			return false
		}
		seen[value.ID] = true
	}
	return true
}
