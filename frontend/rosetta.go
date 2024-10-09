/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"enigma-ar/api"
	"fmt"
	"log"
	"strings"
	"sync"
)

// TODO create test for Rosetta
// TODO check comment about singleton in Rosetta
// Rosetta is a singleton that keeps track of the current language and retrieves texts for this language.
type Rosetta struct {
	persApi     api.PersistencyServer
	currentLang string
	texts       map[string]string
}

var (
	instance *Rosetta
	once     sync.Once
)

// NewRosetta initializes the Rosetta singleton.
func NewRosetta() *Rosetta {
	lang := "en" // TODO read language from configuration
	once.Do(func() {
		instance = &Rosetta{
			persApi:     api.NewPersistencyService(),
			currentLang: lang,
			texts:       make(map[string]string),
		}
		instance.readTextsForLanguage()
	})
	return instance
}

// SetLanguage sets the preferred language.
func (r *Rosetta) SetLanguage(newLang string) {
	if newLang != "en" && newLang != "nl" && newLang != "fr" && newLang != "ge" {
		return
	}
	r.currentLang = newLang
	r.readTextsForLanguage()
}

// GetLanguage returns the currently selected language.
func (r *Rosetta) GetLanguage() string {
	return r.currentLang
}

// GetText retrieves text in the current active language.
func (r *Rosetta) GetText(rbKey string) string {
	if text, found := r.texts[rbKey]; found {
		return text
	}
	return "-- Text not found --"
}

func (r *Rosetta) readTextsForLanguage() {

	relativePath := fmt.Sprintf("./translations/%s.txt", r.currentLang)
	lines, err := r.persApi.ReadLines(relativePath)
	if err != nil {
		log.Printf("rosetta.readTextsForLanguage: failed to read lines from file: %s", err)
	}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if !strings.Contains(line, "#") {
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				r.texts[key] = value
			}
		}
	}

}
