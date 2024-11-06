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
	"fyne.io/fyne/v2/dialog"
)

// CreateMenu defines the global menu for the main window.
func CreateMenu(gm *GuiMgr) *fyne.MainMenu {
	r := gm.Rosetta
	s := NewSettings()

	menuGeneral := createMenuGeneral(r, s, *gm)
	menuCharts := createMenuCharts(r, gm.window)
	menuAnalysis := createMenuAnalysis(r)
	menuProgressive := createMenuProgressive(r)
	menuResearchData := createMenuResearchData(r)
	menuResearchProject := createMenuResearchProject(r)
	menuCycles := createMenuCycles(r)
	menuCalc := createMenuCalc(r, gm.window)
	menuHelp := createMenuHelp(r)
	mainMenu := fyne.NewMainMenu(menuGeneral, menuCharts, menuAnalysis, menuProgressive, menuResearchData, menuResearchProject, menuCycles, menuCalc, menuHelp)
	return mainMenu
}

func handleLangChange(r *Rosetta, s Settings, gm GuiMgr, lang string) {
	w := gm.window
	r.SetLanguage(lang)
	s.DefineLanguage(lang)
	dialog.NewInformation(r.GetText("v_main_language_changed_title"), r.GetText("v_main_language_changed"), w).Show()
	gm.window.Content().Refresh() // TODO try to refresh menu after changing language. Also change rbitems.

}

func createMenuGeneral(r *Rosetta, s Settings, gm GuiMgr) *fyne.Menu {

	languageItem := fyne.NewMenuItem(r.GetText("m_language"), func() {
		fmt.Println("Language clicked.")
	})

	langEnItem := fyne.NewMenuItem(r.GetText("m_lang_eng"), func() {
		handleLangChange(r, s, gm, "en")
	})
	langDuItem := fyne.NewMenuItem(r.GetText("m_lang_dutch"), func() {
		handleLangChange(r, s, gm, "nl")
	})
	langGeItem := fyne.NewMenuItem(r.GetText("m_lang_german"), func() {
		handleLangChange(r, s, gm, "de")
	})
	langFrItem := fyne.NewMenuItem(r.GetText("m_lang_french"), func() {
		handleLangChange(r, s, gm, "fr")
	})
	languageItem.ChildMenu = fyne.NewMenu("", langEnItem, langDuItem, langGeItem, langFrItem)

	settingsItem := fyne.NewMenuItem(r.GetText("m_general_settings"), func() {
		fmt.Println("DefinedSettings clicked.")
	})
	configItem := fyne.NewMenuItem(r.GetText("m_general_config"), func() {
		fmt.Println("Configuration clicked.")
	})
	return fyne.NewMenu(r.GetText("m_language"), languageItem, settingsItem, configItem)
}

func createMenuCharts(r *Rosetta, w fyne.Window) *fyne.Menu {
	//radixInputView := NewRadixInputView()
	newChartItem := fyne.NewMenuItem(r.GetText("m_charts_new"), func() {
		RadixInputView(*r, w)
	})

	searchChartItem := fyne.NewMenuItem(r.GetText("m_charts_search"), func() {
		fmt.Println("Search chart clicked.")
	})
	importChartsItem := fyne.NewMenuItem(r.GetText("m_charts_import"), func() {
		fmt.Println("Import charts clicked.")
	})
	return fyne.NewMenu(r.GetText("m_charts"), newChartItem, searchChartItem, importChartsItem)
}

func createMenuAnalysis(r *Rosetta) *fyne.Menu {
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
	return fyne.NewMenu(r.GetText("m_analysis"), aspectsItem, harmonicsItem, midpointsItem, declinationItem)
}

func createMenuProgressive(r *Rosetta) *fyne.Menu {

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
	return fyne.NewMenu(r.GetText("m_progressive"), newProgEventItem, searchProgEventItem, primDirMenuItem, secDirMenuItem, symDirMenuItem, transitMenuItem, oobCalMenuItem)
}

func createMenuResearchData(r *Rosetta) *fyne.Menu {
	availableResearchDataItem := fyne.NewMenuItem(r.GetText("m_res_data_available"), func() {
		fmt.Println("Available research data clicked.")
	})
	addResearchData := fyne.NewMenuItem(r.GetText("m_res_data_add"), func() {
		fmt.Println("Add research data clicked.")
	})
	removeResearchData := fyne.NewMenuItem(r.GetText("m_res_data_delete"), func() {
		fmt.Println("Remove research data clicked.")
	})
	return fyne.NewMenu(r.GetText("m_research_data"), availableResearchDataItem, addResearchData, removeResearchData)
}

func createMenuResearchProject(r *Rosetta) *fyne.Menu {
	newResearchMenuItem := fyne.NewMenuItem(r.GetText("m_res_proj_new"), func() {
		fmt.Println("New research project clicked.")
	})
	searchResearchMenuItem := fyne.NewMenuItem(r.GetText("m_res_proj_search"), func() {
		fmt.Println("Search research project clicked.")
	})
	deleteResearchMenuItem := fyne.NewMenuItem(r.GetText("m_res_proj_delete"), func() {
		fmt.Println("Delete research project clicked.")
	})
	return fyne.NewMenu(r.GetText("m_research_projects"), newResearchMenuItem, searchResearchMenuItem, deleteResearchMenuItem)
}

func createMenuCycles(r *Rosetta) *fyne.Menu {
	newCycleMenuItem := fyne.NewMenuItem(r.GetText("m_cycle_new"), func() {
		fmt.Println("New cycle clicked.")
	})
	searchCycleMenuItem := fyne.NewMenuItem(r.GetText("m_cycle_search"), func() {
		fmt.Println("Search cycle clicked.")
	})
	deleteCycleMenuItem := fyne.NewMenuItem(r.GetText("m_cycle_delete"), func() {
		fmt.Println("Delete cycle clicked.")
	})
	return fyne.NewMenu(r.GetText("m_cycles"), newCycleMenuItem, searchCycleMenuItem, deleteCycleMenuItem)
}

func createMenuCalc(r *Rosetta, w fyne.Window) *fyne.Menu {
	calcJdNrMenuItem := fyne.NewMenuItem(r.GetText("m_calc_jd"), func() {
		CalcJdView(*r, w)
	})
	calcDateMenuItem := fyne.NewMenuItem(r.GetText("m_calc_datetime"), func() {
		fmt.Println("Date/time from Julian day number clicked.")
	})
	celcObliquityMenuItem := fyne.NewMenuItem(r.GetText("m_calc_obliquity"), func() {
		fmt.Println("Calculate obliquity clicked.")
	})
	return fyne.NewMenu(r.GetText("m_calculations"), calcJdNrMenuItem, calcDateMenuItem, celcObliquityMenuItem)
}

func createMenuHelp(r *Rosetta) *fyne.Menu {

	aboutMenuItem := fyne.NewMenuItem(r.GetText("m_help_about"), func() {
		fmt.Println("About Enigma clicked.")
	})
	manualMenuItem := fyne.NewMenuItem(r.GetText("m_help_user_manual"), func() {
		fmt.Println("User manual clicked.")
	})
	whatsNewItem := fyne.NewMenuItem(r.GetText("m_help_whats_new"), func() {
		fmt.Println("Whats New clicked.")
	})
	return fyne.NewMenu(r.GetText("m_help"), aboutMenuItem, manualMenuItem, whatsNewItem)
}
