/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
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
