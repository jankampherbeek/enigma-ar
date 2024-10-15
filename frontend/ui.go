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

func CreateMenu(gm *GuiMgr) *fyne.MainMenu {
	r := gm.Rosetta
	languageItem := fyne.NewMenuItem("Language", func() {
		fmt.Println("Language clicked.")
	})

	langEnItem := fyne.NewMenuItem("English", func() {
		r.SetLanguage("en")
		gm.Refresh("Home")
		// TODO persist language
	})
	langDuItem := fyne.NewMenuItem("Nederlands / Dutch", func() {
		r.SetLanguage("nl")
		gm.Refresh("Home")
		// TODO persist language
	})
	langGeItem := fyne.NewMenuItem("Deutsch / German", func() {
		r.SetLanguage("ge")
		gm.Refresh("Home")
		// TODO persist language
	})
	langFrItem := fyne.NewMenuItem("Français / French", func() {
		r.SetLanguage("fr")
		gm.Refresh("Home")
		// TODO persist language
	})
	languageItem.ChildMenu = fyne.NewMenu("", langEnItem, langDuItem, langGeItem, langFrItem)

	settingsItem := fyne.NewMenuItem("Settings", func() {
		fmt.Println("Settings clicked.")
	})
	configItem := fyne.NewMenuItem("Configuration", func() {
		fmt.Println("Configuration clicked.")
	})
	generalMenu := fyne.NewMenu("General", languageItem, settingsItem, configItem)

	newChartItem := fyne.NewMenuItem("New chart", func() {
		fmt.Println("New chart clicked.")
	})
	searchChartItem := fyne.NewMenuItem("Search chart", func() {
		fmt.Println("Search chart clicked.")
	})
	importChartsItem := fyne.NewMenuItem("Import charts", func() {
		fmt.Println("Import charts clicked.")
	})
	chartMenu := fyne.NewMenu("Charts", newChartItem, searchChartItem, importChartsItem)

	aspectsItem := fyne.NewMenuItem("Aspects", func() {
		fmt.Println("Aspects clicked.")
	})
	harmonicsItem := fyne.NewMenuItem("Harmonics", func() {
		fmt.Println("Harmonics clicked.")
	})
	midpointsItem := fyne.NewMenuItem("Midpoints", func() {
		fmt.Println("Midpoints clicked.")
	})
	declDiagramItem := fyne.NewMenuItem("Declination diagram", func() {
		fmt.Println("Declination diagram clicked.")
	})
	declStripItem := fyne.NewMenuItem("Declination strip", func() {
		fmt.Println("Declination strip clicked.")
	})
	declParallelsItem := fyne.NewMenuItem("Parallels", func() {
		fmt.Println("Parallels clicked.")
	})
	declLongEquivItem := fyne.NewMenuItem("Longitude equivalents", func() {
		fmt.Println("Longitude equivalents clicked.")
	})
	declinationItem := fyne.NewMenuItem("Declination", func() {
		fmt.Println("Declination clicked.")
	})
	declinationItem.ChildMenu = fyne.NewMenu("", declDiagramItem, declStripItem, declParallelsItem, declLongEquivItem)
	analysisMenu := fyne.NewMenu("Analysis", aspectsItem, harmonicsItem, midpointsItem, declinationItem)

	newProgEventItem := fyne.NewMenuItem("New event", func() {
		fmt.Println("New event clicked.")
	})
	searchProgEventItem := fyne.NewMenuItem("Search event", func() {
		fmt.Println("Search event clicked.")
	})
	primDirMenuItem := fyne.NewMenuItem("Primary directions", func() {
		fmt.Println("Primary directions clicked.")
	})
	secDirMenuItem := fyne.NewMenuItem("Secondary directions", func() {
		fmt.Println("Secondary directions clicked.")
	})
	symDirMenuItem := fyne.NewMenuItem("Symbolic directions", func() {
		fmt.Println("Symbolic directions clicked.")
	})
	transitMenuItem := fyne.NewMenuItem("Transits", func() {
		fmt.Println("Transits clicked.")
	})
	oobCalMenuItem := fyne.NewMenuItem("OOB Calendar", func() {
		fmt.Println("OOB Calendar clicked.")
	})
	progressiveMenu := fyne.NewMenu("Progressive", newProgEventItem, searchProgEventItem, primDirMenuItem, secDirMenuItem, symDirMenuItem, transitMenuItem, oobCalMenuItem)

	availableResearchDataItem := fyne.NewMenuItem("Available research data", func() {
		fmt.Println("Available research data clicked.")
	})
	addResearchData := fyne.NewMenuItem("Add research data", func() {
		fmt.Println("Add research data clicked.")
	})
	removeResearchData := fyne.NewMenuItem("Remove research data", func() {
		fmt.Println("Remove research data clicked.")
	})
	researchDataMenu := fyne.NewMenu("Research data", availableResearchDataItem, addResearchData, removeResearchData)

	newResearchMenuItem := fyne.NewMenuItem("New research project", func() {
		fmt.Println("New research project clicked.")
	})
	searchResearchMenuItem := fyne.NewMenuItem("Search research project", func() {
		fmt.Println("Search research project clicked.")
	})
	deleteResearchMenuItem := fyne.NewMenuItem("Delete research project", func() {
		fmt.Println("Delete research project clicked.")
	})
	researchProjectMenu := fyne.NewMenu("Research projects", newResearchMenuItem, searchResearchMenuItem, deleteResearchMenuItem)

	newCycleMenuItem := fyne.NewMenuItem("New cycle", func() {
		fmt.Println("New cycle clicked.")
	})
	searchCycleMenuItem := fyne.NewMenuItem("Search cycle", func() {
		fmt.Println("Search cycle clicked.")
	})
	deleteCycleMenuItem := fyne.NewMenuItem("Delete cycle", func() {
		fmt.Println("Delete cycle clicked.")
	})
	cyclesMenu := fyne.NewMenu("Cycles", newCycleMenuItem, searchCycleMenuItem, deleteCycleMenuItem)

	calcJdNrMenuItem := fyne.NewMenuItem("Julian day number from date/time", func() {
		fmt.Println("Julian day number from date/timeclicked.")
	})
	calcDateMenuItem := fyne.NewMenuItem("Date/time from Julian date number", func() {
		fmt.Println("Date/time from Julian day number clicked.")
	})
	celcObliquityMenuItem := fyne.NewMenuItem("Celculate obliquity", func() {
		fmt.Println("Calculate obliquity clicked.")
	})
	calcMenu := fyne.NewMenu("Calculations", calcJdNrMenuItem, calcDateMenuItem, celcObliquityMenuItem)

	aboutMenuItem := fyne.NewMenuItem("About Enigma", func() {
		fmt.Println("About Enigma clicked.")
	})
	manualMenuItem := fyne.NewMenuItem("User manual", func() {
		fmt.Println("User manual clicked.")
	})
	whatsNewItem := fyne.NewMenuItem("Whats New", func() {
		fmt.Println("Whats New clicked.")
	})
	helpMenu := fyne.NewMenu("Help", aboutMenuItem, manualMenuItem, whatsNewItem)

	mainMenu := fyne.NewMainMenu(generalMenu, chartMenu, analysisMenu, progressiveMenu, researchDataMenu, researchProjectMenu, cyclesMenu, calcMenu, helpMenu)
	return mainMenu
}

func MakeUI(app fyne.App) {

	mainWindow := app.NewWindow("Enigma 1.0")
	guiMgr := NewGuiMgr(mainWindow)
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
