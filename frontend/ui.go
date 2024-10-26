/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (gm *GuiMgr) createChartsMain() *fyne.Container {
	content := container.NewCenter(
		widget.NewLabel(gm.Rosetta.GetText("btnTutorial")),
	)
	return content
}

func MakeUI(app fyne.App) {

	mainWindow := app.NewWindow("Enigma 1.0")
	guiMgr := NewGuiMgr(mainWindow)
	settings := NewSettings()
	guiMgr.Rosetta.SetLanguage(settings.GetLanguage())
	mainWindow.Resize(fyne.NewSize(1200, 900))
	mainWindow.SetMaster()
	mainWindow.SetMainMenu(CreateMenu(guiMgr))

	homeView := NewHomeView(guiMgr)
	guiMgr.Register("home", homeView)
	guiMgr.Register("charts", NewChartsView(guiMgr))
	guiMgr.Register("config", NewConfigView(guiMgr))
	guiMgr.Register("counts", NewCountsView(guiMgr))
	guiMgr.Register("cycles", NewCyclesView(guiMgr))
	guiMgr.Register("manual", NewManualView(guiMgr))

	mainWindow.SetContent(homeView)

	mainWindow.ShowAndRun()
}
