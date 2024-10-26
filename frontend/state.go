/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"enigma-ar/domain"
	"fyne.io/fyne/v2"
	"log"
)

type GuiMgr struct {
	App     fyne.App
	Log     *log.Logger
	Rosetta *Rosetta
	window  fyne.Window
	views   map[string]fyne.CanvasObject
}

func NewGuiMgr(window fyne.Window) *GuiMgr {
	return &GuiMgr{
		window:  window,
		Rosetta: NewRosetta(),
		views:   make(map[string]fyne.CanvasObject),
	}
}

func (gm *GuiMgr) Register(name string, view fyne.CanvasObject) {
	gm.views[name] = view
}

func (gm *GuiMgr) Show(name string) {
	if view, ok := gm.views[name]; ok {
		gm.window.SetContent(view)
		gm.window.Show()
	}
}

func (gm *GuiMgr) Refresh(name string) {
	if view, ok := gm.views[name]; ok {
		view.Refresh()
	}
}

type RadixInputData struct {
	NameId      string
	Description string
	Source      string
	Categorie   domain.ChartCat
	Rating      domain.Rating
	Country     string
	Location    string
	GeoLong     float64
	GeoLat      float64
	Year        int
	Month       int
	Day         int
	Calendar    domain.Calendar
	Hour        int
	Minute      int
	Second      int
	TimeZone    domain.TimeZone
	GeoLongLmt  float64 // zero if not applicable
	Dst         bool
}

type DataVault struct {
	InputData RadixInputData
	Request   domain.PointPositionsRequest
	Response  domain.FullChartResponse
}

func (dv DataVault) DefineChartForInput(inputData RadixInputData) {
	dv.InputData = inputData
	// define request
	// fire request
	// handle response
}
