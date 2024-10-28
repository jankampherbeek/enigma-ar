/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"encoding/json"
	"enigma-ar/api"
	"fyne.io/fyne/v2"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
	"os"
	"path/filepath"
	"sync"
)

// TODO create test for Rosetta
// TODO check comment about singleton in Rosetta
// Rosetta is a singleton that keeps track of the current language and retrieves texts for this language.
type Rosetta struct {
	persApi     api.PersistencyServer
	currentLang language.Tag
	langCat     *catalog.Builder
	printer     *message.Printer
}

var (
	instance *Rosetta
	once     sync.Once
)

// NewRosetta initializes the Rosetta singleton.
func NewRosetta(a fyne.App) *Rosetta {

	// On startup, load the user's preferred language
	langPref := a.Preferences().StringWithFallback("language", language.English.String())
	rLang := language.Make(langPref)
	rLangCat := loadTranslations()
	rPrinter := message.NewPrinter(rLang, message.Catalog(rLangCat))

	once.Do(func() {
		instance = &Rosetta{
			persApi:     api.NewPersistencyService(),
			currentLang: rLang,
			langCat:     rLangCat,
			printer:     rPrinter,
		}
	})
	return instance
}

// SetLanguage sets the preferred language.
func (r *Rosetta) SetLanguage(newLang string) {
	if newLang != "en" && newLang != "nl" && newLang != "fr" && newLang != "de" {
		return
	}
	r.currentLang = language.English
	if newLang == "nl" {
		r.currentLang = language.Dutch
	}
	if newLang == "fr" {
		r.currentLang = language.French
	}
	if newLang == "de" {
		r.currentLang = language.German
	}
	r.printer = message.NewPrinter(r.currentLang, message.Catalog(r.langCat))
	//r.readTextsForLanguage()
}

// GetLanguage returns the currently selected language. If the lanuage is not found, "en" is returned.
func (r *Rosetta) GetLanguage() string {
	if r.currentLang == language.French {
		return "fr"
	}
	if r.currentLang == language.German {
		return "ge"
	}
	if r.currentLang == language.Dutch {
		return "nl"
	}
	return "en"
}

// GetText retrieves text in the current active language.
func (r *Rosetta) GetText(rbKey string) string {
	return r.printer.Sprintf(rbKey)
}

func loadTranslations() *catalog.Builder {
	builder := catalog.NewBuilder()
	// Load translation files from the locales directory
	localesDir := "locales"
	files, err := os.ReadDir(localesDir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		// Determine the language from the filename
		langTag := language.MustParse(file.Name()[:len(file.Name())-len(".json")])

		// Open and decode the JSON file
		fpath := filepath.Join(localesDir, file.Name())
		fileContent, err := os.ReadFile(fpath)
		if err != nil {
			panic(err)
		}

		messages := make(map[string]string)
		if err := json.Unmarshal(fileContent, &messages); err != nil {
			panic(err)
		}

		// Add messages to the catalog
		for key, msg := range messages {
			err := builder.SetString(langTag, key, msg)
			if err != nil {
				// TODO log error for unknown entry in Rosetta
			}
		}
	}

	return builder
}
