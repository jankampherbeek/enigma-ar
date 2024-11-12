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
	guiMgr := NewGuiMgr(app, mainWindow)
	settings := NewSettings()

	guiMgr.Rosetta.SetLanguage(settings.GetLanguage())
	mainWindow.Resize(fyne.NewSize(1200, 900))
	mainWindow.SetMaster()
	mainWindow.SetMainMenu(CreateChartsMenu(guiMgr))

	homeView := NewChartsView(guiMgr)
	guiMgr.Register("home", homeView)
	guiMgr.Register("charts", NewChartsView(guiMgr))
	guiMgr.Register("config", NewConfigView(guiMgr))
	guiMgr.Register("counts", NewCountsView(guiMgr))
	guiMgr.Register("cycles", NewCyclesView(guiMgr))
	guiMgr.Register("calculators", NewCalculatorsView(guiMgr))
	guiMgr.Register("manual", NewManualView(guiMgr))

	// begin debug
	/*	lblChicken := widget.NewLabel("Chicken")
		lblEgg := widget.NewLabel("Egg")
		entryChicken := widget.NewEntry()
		entryChicken.PlaceHolder = "What about the chicken?"
		entryEgg := widget.NewEntry()
		entryEgg.PlaceHolder = ".... and the egg?"
		formContainer := container.New(layout.NewFormLayout(),
			lblChicken,
			entryChicken,
			lblEgg,
			entryEgg,
		)
		mainWindow.SetContent(formContainer)*/

	// end debug

	mainWindow.SetContent(homeView)

	mainWindow.ShowAndRun()
}
