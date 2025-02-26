/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package api

import (
	"enigma-ar/domain"
	"enigma-ar/internal/meta"
	"log/slog"
	"path/filepath"
	"strings"
)

// ConfigServer provides services for configurations.
type ConfigServer interface {
	DefaultConfig() domain.Config
	NamedConfig(name string) (domain.Config, error)
	SaveConfig(name string, deltas []string) error
	ExistingConfigs() []string
}

type ConfigService struct {
	persService PersistencyServer
}

func NewConfig() *ConfigService {
	return &ConfigService{
		persService: NewPersistencyService(),
	}
}

// DefaultConfig returns the default configuration.
func (cs ConfigService) DefaultConfig() domain.Config {
	slog.Info("Retrieving default configuration")
	cfg := meta.DefaultConfig()
	return cfg
}

func (cs ConfigService) NamedConfig(name string) (domain.Config, error) {
	slog.Info("Retrieving named configuration for name: " + name)
	path := name + ".cfg"
	deltas, err := cs.persService.ReadLines(path)
	if err != nil {
		return domain.Config{}, err
	}
	nConfig, err := meta.ActualConfig(deltas)
	if err != nil {
		return domain.Config{}, err
	}
	return nConfig, nil
}

func (cs ConfigService) SaveConfig(name string, deltas []string) error {
	slog.Info("Saving configuration for name: " + name)
	path := name + ".cfg"
	err := cs.persService.WriteLines(path, deltas)
	if err != nil {
		return err
	}
	return nil
}

func (cs ConfigService) ExistingConfigs() []string {
	slog.Info("Retrieving existing configurations")
	matches, _ := filepath.Glob("*.cfg")
	names := make([]string, len(matches))
	for i, match := range matches {
		names[i] = strings.TrimSuffix(match, ".cfg")
	}
	return names
}
