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
	"fyne.io/fyne/v2/dialog"
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

func handleLangChange(r *Rosetta, s Settings, w fyne.Window) {
	r.SetLanguage("en")
	s.DefineLanguage("en")
	dialog.NewInformation(r.GetText("v_main_language_changed_title"), r.GetText("v_main_language_changed"), w).Show()
}

func CreateMenu(gm *GuiMgr) *fyne.MainMenu {
	r := gm.Rosetta
	s := NewSettings()
	languageItem := fyne.NewMenuItem(r.GetText("m_language"), func() {
		fmt.Println("Language clicked.")
	})

	langEnItem := fyne.NewMenuItem(r.GetText("m_lang_eng"), func() {
		handleLangChange(r, s, gm.window)
	})
	langDuItem := fyne.NewMenuItem(r.GetText("m_lang_dutch"), func() {
		handleLangChange(r, s, gm.window)
	})
	langGeItem := fyne.NewMenuItem(r.GetText("m_lang_german"), func() {
		handleLangChange(r, s, gm.window)
	})
	langFrItem := fyne.NewMenuItem(r.GetText("m_lang_french"), func() {
		handleLangChange(r, s, gm.window)
	})
	languageItem.ChildMenu = fyne.NewMenu("", langEnItem, langDuItem, langGeItem, langFrItem)

	settingsItem := fyne.NewMenuItem(r.GetText("m_general_settings"), func() {
		fmt.Println("DefinedSettings clicked.")
	})
	configItem := fyne.NewMenuItem(r.GetText("m_general_config"), func() {
		fmt.Println("Configuration clicked.")
	})
	generalMenu := fyne.NewMenu(r.GetText("m_language"), languageItem, settingsItem, configItem)

	newChartItem := fyne.NewMenuItem(r.GetText("m_charts_new"), func() {
		fmt.Println("New chart clicked.")
	})
	searchChartItem := fyne.NewMenuItem(r.GetText("m_charts_search"), func() {
		fmt.Println("Search chart clicked.")
	})
	importChartsItem := fyne.NewMenuItem(r.GetText("m_charts_import"), func() {
		fmt.Println("Import charts clicked.")
	})
	chartMenu := fyne.NewMenu(r.GetText("m_charts"), newChartItem, searchChartItem, importChartsItem)

	aspectsItem := fyne.NewMenuItem(r.GetText("m_analysis_aspects"), func() {
		fmt.Println("Aspects clicked.")
	})
	harmonicsItem := fyne.NewMenuItem(r.GetText("m_analysis_harmonics"), func() {
		fmt.Println("Harmonics clicked.")
	})
	midpointsItem := fyne.NewMenuItem(r.GetText("m_analysis_midpoints"), func() {
		fmt.Println("Midpoints clicked.")
	})
	declDiagramItem := fyne.NewMenuItem(r.GetText("m_analysis_decl_diagram"), func() {
		fmt.Println("Declination diagram clicked.")
	})
	declStripItem := fyne.NewMenuItem(r.GetText("m_analysis_decl_strip"), func() {
		fmt.Println("Declination strip clicked.")
	})
	declParallelsItem := fyne.NewMenuItem(r.GetText("m_analysis_decl_parallels"), func() {
		fmt.Println("Parallels clicked.")
	})
	declLongEquivItem := fyne.NewMenuItem(r.GetText("m_analysis_decl_long_equiv"), func() {
		fmt.Println("Longitude equivalents clicked.")
	})
	declinationItem := fyne.NewMenuItem(r.GetText("m_analysis_declinations"), func() {
		fmt.Println("Declination clicked.")
	})
	declinationItem.ChildMenu = fyne.NewMenu("", declDiagramItem, declStripItem, declParallelsItem, declLongEquivItem)
	analysisMenu := fyne.NewMenu(r.GetText("m_analysis"), aspectsItem, harmonicsItem, midpointsItem, declinationItem)

	newProgEventItem := fyne.NewMenuItem(r.GetText("m_prog_new_event"), func() {
		fmt.Println("New event clicked.")
	})
	searchProgEventItem := fyne.NewMenuItem(r.GetText("m_prog_search_event"), func() {
		fmt.Println("Search event clicked.")
	})
	primDirMenuItem := fyne.NewMenuItem(r.GetText("m_prog_prim_dir"), func() {
		fmt.Println("Primary directions clicked.")
	})
	secDirMenuItem := fyne.NewMenuItem(r.GetText("m_prog_sec_dir"), func() {
		fmt.Println("Secondary directions clicked.")
	})
	symDirMenuItem := fyne.NewMenuItem(r.GetText("m_prog_sym_dir"), func() {
		fmt.Println("Symbolic directions clicked.")
	})
	transitMenuItem := fyne.NewMenuItem(r.GetText("m_prog_transits"), func() {
		fmt.Println("Transits clicked.")
	})
	oobCalMenuItem := fyne.NewMenuItem(r.GetText("m_prog_oob_calendar"), func() {
		fmt.Println("OOB Calendar clicked.")
	})
	progressiveMenu := fyne.NewMenu(r.GetText("m_progressive"), newProgEventItem, searchProgEventItem, primDirMenuItem, secDirMenuItem, symDirMenuItem, transitMenuItem, oobCalMenuItem)

	availableResearchDataItem := fyne.NewMenuItem(r.GetText("m_res_data_available"), func() {
		fmt.Println("Available research data clicked.")
	})
	addResearchData := fyne.NewMenuItem(r.GetText("m_res_data_add"), func() {
		fmt.Println("Add research data clicked.")
	})
	removeResearchData := fyne.NewMenuItem(r.GetText("m_res_data_delete"), func() {
		fmt.Println("Remove research data clicked.")
	})
	researchDataMenu := fyne.NewMenu(r.GetText("m_research_data"), availableResearchDataItem, addResearchData, removeResearchData)

	newResearchMenuItem := fyne.NewMenuItem(r.GetText("m_res_proj_new"), func() {
		fmt.Println("New research project clicked.")
	})
	searchResearchMenuItem := fyne.NewMenuItem(r.GetText("m_res_proj_search"), func() {
		fmt.Println("Search research project clicked.")
	})
	deleteResearchMenuItem := fyne.NewMenuItem(r.GetText("m_res_proj_delete"), func() {
		fmt.Println("Delete research project clicked.")
	})
	researchProjectMenu := fyne.NewMenu(r.GetText("m_research_projects"), newResearchMenuItem, searchResearchMenuItem, deleteResearchMenuItem)

	newCycleMenuItem := fyne.NewMenuItem(r.GetText("m_cycle_new"), func() {
		fmt.Println("New cycle clicked.")
	})
	searchCycleMenuItem := fyne.NewMenuItem(r.GetText("m_cycle_search"), func() {
		fmt.Println("Search cycle clicked.")
	})
	deleteCycleMenuItem := fyne.NewMenuItem(r.GetText("m_cycle_delete"), func() {
		fmt.Println("Delete cycle clicked.")
	})
	cyclesMenu := fyne.NewMenu(r.GetText("m_cycles"), newCycleMenuItem, searchCycleMenuItem, deleteCycleMenuItem)

	calcJdNrMenuItem := fyne.NewMenuItem(r.GetText("m_calc_jd"), func() {
		fmt.Println("Julian day number from date/timeclicked.")
	})
	calcDateMenuItem := fyne.NewMenuItem(r.GetText("m_calc_datetime"), func() {
		fmt.Println("Date/time from Julian day number clicked.")
	})
	celcObliquityMenuItem := fyne.NewMenuItem(r.GetText("m_calc_obliquity"), func() {
		fmt.Println("Calculate obliquity clicked.")
	})
	calcMenu := fyne.NewMenu(r.GetText("m_calculations"), calcJdNrMenuItem, calcDateMenuItem, celcObliquityMenuItem)

	aboutMenuItem := fyne.NewMenuItem(r.GetText("m_help_about"), func() {
		fmt.Println("About Enigma clicked.")
	})
	manualMenuItem := fyne.NewMenuItem(r.GetText("m_help_user_manual"), func() {
		fmt.Println("User manual clicked.")
	})
	whatsNewItem := fyne.NewMenuItem(r.GetText("m_help_whats_new"), func() {
		fmt.Println("Whats New clicked.")
	})
	helpMenu := fyne.NewMenu(r.GetText("m_help"), aboutMenuItem, manualMenuItem, whatsNewItem)

	mainMenu := fyne.NewMainMenu(generalMenu, chartMenu, analysisMenu, progressiveMenu, researchDataMenu, researchProjectMenu, cyclesMenu, calcMenu, helpMenu)
	return mainMenu
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
