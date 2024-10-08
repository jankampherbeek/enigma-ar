/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// TODO create test for Rosetta
// TODO check comment about singleton in Rosetta
// Rosetta is a singleton that keeps track of the current language and retrieves texts for this language.
type Rosetta struct {
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
	file, err := os.Open(relativePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	r.texts = make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			r.texts[key] = value
		} else {
			log.Fatalf("Rosetta: No = sign found in line: %s", line)
		}
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
}
