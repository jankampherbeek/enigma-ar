/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package api

import (
	"enigma-ar/domain"
	"enigma-ar/internal/persistency"
	"log/slog"
)

// PersistencyServer provides services to access files or databases.
type PersistencyServer interface {
	ReadLines(path string) ([]string, error)
	WriteLines(path string, lines []string) error
}

type PersistencyService struct{}

func NewPersistencyService() *PersistencyService {
	return &PersistencyService{}
}

// ReadLines returns the lines as read from the specified file.
func (ps PersistencyService) ReadLines(path string) ([]string, error) {
	return persistency.ReadTextLines(path)
}

// WriteLines creates a new file and writes the lines to that file.
func (ps PersistencyService) WriteLines(path string, lines []string) error {
	slog.Info("Writing lines")
	return persistency.WriteTextLines(path, lines)
}

func WriteChart(pcData domain.PersistableChart, pdlData domain.PersistableDateLocation) (int, int, error) {
	slog.Info("Writing chart")
	return persistency.SaveChartData(pcData, pdlData)
}
