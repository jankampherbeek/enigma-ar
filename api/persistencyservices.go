/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package api

import (
	"enigma-ar/internal/persistency"
)

// PersistencyServer provides services to access files or databases.
type PersistencyServer interface {
	ReadLines(path string) ([]string, error)
}

type PersistencyService struct{}

func NewPersistencyService() *PersistencyService {
	return &PersistencyService{}
}

// ReadLines returns the lines as read from the specified file.
func (ps PersistencyService) ReadLines(path string) ([]string, error) {
	lines, err := persistency.ReadTextLines(path)
	return lines, err
}
