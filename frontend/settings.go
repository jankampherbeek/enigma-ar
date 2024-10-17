/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"enigma-ar/api"
	"log"
	"os"
	"strconv"
	"strings"
)

const settingsPath = "./settings.txt"

// Settings defines the global settings for the application.
type Settings interface {
	DefineLanguage(lang string)
	DefineWorkFolder(lang string)
	DefineMode(darkMode bool)
	GetLanguage() string
	GetWorkFolder() string
	GetDarkMode() bool
}

// DefinedSettings defines the actual global settings for the application.
type DefinedSettings struct {
	persApi    api.PersistencyServer
	lang       string
	workFolder string
	darkMode   bool
}

func NewSettings() Settings {
	pApi := api.NewPersistencyService()
	ds := DefinedSettings{
		persApi:    pApi,
		lang:       "en",
		workFolder: "/enigma-ar",
		darkMode:   false,
	}
	ds.readSettingsFromFile()
	return ds
}

func (s DefinedSettings) DefineLanguage(lang string) {
	if lang == "en" || lang == "ge" || lang == "nl" || lang == "fr" {
		s.lang = lang
		s.writeSettingsToFile()
	} else {
		log.Printf("settings.DefineLanguage. Did not recognize language \"%s\".", lang)
	}
}

func (s DefinedSettings) DefineWorkFolder(folder string) {
	if len(folder) != 0 {
		s.workFolder = folder
		s.writeSettingsToFile()
	} else {
		log.Printf("settings.DefineWorkFolder. Refused to set workfolder with length zero.")
	}
}

func (s DefinedSettings) DefineMode(darkMode bool) {
	s.darkMode = darkMode
}

func (s DefinedSettings) GetLanguage() string { return s.lang }

func (s DefinedSettings) GetWorkFolder() string { return s.workFolder }

func (s DefinedSettings) GetDarkMode() bool { return s.darkMode }

// readSettingsFromFile reads the lines from the settings file
func (s *DefinedSettings) readSettingsFromFile() {
	_, err := os.Stat(settingsPath)
	if os.IsNotExist(err) {
		log.Printf("settings.ReadSettingsFromFile. Settings file \"%s\" does not exist. Creating default settings file.", settingsPath)
		s.defineDefaultSettings()
		s.writeSettingsToFile()
	}
	lines, err := s.persApi.ReadLines(settingsPath)
	if err != nil {
		log.Printf("No settings, creating default settings.")
		s.defineDefaultSettings()
		s.writeSettingsToFile()
	} else {
		for i := 0; i < len(lines); i++ {
			line := lines[i]
			if !strings.Contains(line, "#") {
				parts := strings.Split(line, "=")
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])
					if key == "lang" {
						s.lang = value
					}
					if key == "workFolder" {
						s.workFolder = value
					}
					if key == "darkMode" {
						s.darkMode = value == "true"
					}
				}
			}
		}
	}
}

// writeSettingsToFile writes lines to the settigns file
func (s *DefinedSettings) writeSettingsToFile() {
	var lines []string
	lines = append(lines, "lang="+s.lang)
	lines = append(lines, "workFolder="+s.workFolder)
	lines = append(lines, "darkMode="+strconv.FormatBool(s.darkMode))
	err := s.persApi.WriteLines(settingsPath, lines)
	if err != nil {
		log.Printf("Error writing settings to file: %v", err)
	}
}

// defineDefaultSettings if no settings file is found, default settings will be applied
func (s *DefinedSettings) defineDefaultSettings() {
	s.lang = "en"
	s.workFolder = "/enigma-ra"
	s.darkMode = false
}
