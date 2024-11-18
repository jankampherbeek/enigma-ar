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
	"golang.org/x/text/language"
	"log"
	"sync"
)

type GuiMgr struct {
	App     fyne.App
	Log     *log.Logger
	Rosetta *Rosetta
	DvRadix *DataVaultRadix
	window  fyne.Window
	views   map[string]fyne.CanvasObject
}

var (
	gmInstance *GuiMgr
	gmOnce     sync.Once
)

func NewGuiMgr(app fyne.App, window fyne.Window) *GuiMgr {

	gmOnce.Do(func() {
		gmInstance = &GuiMgr{
			App:     app,
			window:  window,
			Rosetta: NewRosetta(app),
			DvRadix: GetDataVaultRadix(),
			views:   make(map[string]fyne.CanvasObject),
		}
	})
	return gmInstance
}

func GetGuiMgr() *GuiMgr {
	if gmInstance == nil {
		panic("Gui manager not initialized")
	}
	return gmInstance
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

func (gm *GuiMgr) SaveLanguage(lang language.Tag) {
	gm.App.Preferences().SetString("language", lang.String())
}

// TODO change to support multiple charts
type DataVaultRadix struct {
	Request   domain.FullChartRequest
	Response  domain.FullChartResponse
	Meta      domain.FullChartMeta
	completed bool
}

var (
	dvrInstance *DataVaultRadix
	dvrOnce     sync.Once
)

func GetDataVaultRadix() *DataVaultRadix {

	dvrOnce.Do(func() {
		dvrInstance = &DataVaultRadix{
			Request: domain.FullChartRequest{
				Points:    nil,
				HouseSys:  0,
				Ayanamsha: 0,
				CoordSys:  0,
				ObsPos:    0,
				ProjType:  0,
				Jd:        0,
				GeoLong:   0,
				GeoLat:    0,
			},
			Response: domain.FullChartResponse{
				Points:    nil,
				Mc:        domain.HousePosResult{},
				Asc:       domain.HousePosResult{},
				Vertex:    domain.HousePosResult{},
				EastPoint: domain.HousePosResult{},
				Cusps:     nil,
			},
			completed: false,
		}
	})
	return dvrInstance
}

// TODO add request and meta
func (dvr *DataVaultRadix) AddCalculatedChart(request domain.FullChartRequest, response domain.FullChartResponse,
	meta domain.FullChartMeta) {
	dvr.Request = request
	dvr.Response = response
	dvr.Meta = meta
}

/*
type RadixInputData struct {
	NameId      string
	Description string
	Source      string
	ChartCat   domain.ChartCat
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
*/
