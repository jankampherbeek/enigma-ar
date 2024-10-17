/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"os"
)

func makeHeader() *widget.Label {
	aTitle := "Enigma Astrology Research - version 1.0"
	return widget.NewLabelWithStyle(aTitle, fyne.TextAlignCenter, fyne.TextStyle{
		Bold:      true,
		Italic:    false,
		Monospace: false,
		Symbol:    false,
		TabWidth:  0,
	})

}

func makeButton() *widget.Button {
	b := widget.NewButton("Tutorial",
		func() {
			showTutorial("Please study....")
		})
	b.Importance = widget.HighImportance
	return b

}

func showTutorial(text string) {
	fmt.Println(text)
}

func makeBox() *fyne.Container {
	//header := makeHeader()
	button1 := makeButton()
	button2 := widget.NewButton("Charts", func() { fmt.Println("Charts clicked.") })
	button2.Importance = widget.SuccessImportance
	button3 := widget.NewButton("Cycles", func() { fmt.Println("Cycles clicked.") })
	button3.Importance = widget.MediumImportance
	button4 := widget.NewButton("Calculators", func() { fmt.Println("Calculators clicked.") })
	button4.Importance = widget.MediumImportance
	button5 := widget.NewButton("Counting", func() { fmt.Println("Counting clicked.") })
	button5.Importance = widget.MediumImportance

	btnNewChart := widget.NewButton("New", func() { fmt.Println("New chart....") })
	btnNewChart.Importance = widget.HighImportance
	btnSearchChart := widget.NewButton("Search", func() { fmt.Println("Search chart....") })
	btnSearchChart.Importance = widget.MediumImportance

	text2 := canvas.NewText("Text 2", color.Black)
	text2.TextSize = 36
	labelGlobalBtns := widget.NewLabel("Modules")
	labelGlobalBtns.Importance = widget.MediumImportance
	labelLocalBtns := widget.NewLabel("Charts")
	labelLocalBtns.Importance = widget.MediumImportance
	labelDescription := widget.NewLabel("Description of chart: namew, date, time, coordinates")
	labelDescription.Importance = widget.MediumImportance
	boxGlobal := container.New(layout.NewVBoxLayout(), labelGlobalBtns, button1, button2, button3, button4, button5)
	boxLocal := container.New(layout.NewVBoxLayout(), labelLocalBtns, btnNewChart, btnSearchChart)
	leftBox := container.New(layout.NewVBoxLayout(), boxGlobal, boxLocal)
	return leftBox
}

func loadTranslation(path string) fyne.StaticResource {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to load translation file: %v", err)
	}
	return *fyne.NewStaticResource(path, data)
}

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
	mainWindow.Resize(fyne.NewSize(1024, 768))
	mainWindow.SetMaster()
	mainWindow.SetMainMenu(CreateMenu(guiMgr))

	homeView := NewHomeView(guiMgr)
	guiMgr.Register("home", homeView)
	guiMgr.Register("charts", NewChartsView(guiMgr))
	guiMgr.Register("config", NewConfigView(guiMgr))
	guiMgr.Register("calc", NewCalcView(guiMgr))
	guiMgr.Register("counts", NewCountsView(guiMgr))
	guiMgr.Register("cycles", NewCyclesView(guiMgr))
	guiMgr.Register("manual", NewManualView(guiMgr))

	mainWindow.SetContent(homeView)

	mainWindow.ShowAndRun()
}
