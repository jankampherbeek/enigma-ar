/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"fmt"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type state int

const (
	chartCalcCompleted state = iota
	calcNewChart
	editConfig
)

func changeState(s state) {
	switch s {
	case chartCalcCompleted:
		handleChartCompleted()
	case calcNewChart:
		handleNewChart()
	case editConfig:
		handleNewconfig()
	}
}

// handleNewChart takes care of the input and calculation of a new chart.
func handleNewChart() {
	dvRadix := GetDataVaultRadix()
	gm := GetGuiMgr()
	newContainer := RadixInputView()
	gm.Register("charts", UpdateChartsView(&newContainer))
	gm.Show("charts")

	fmt.Println("In handleNewChart(Menu) dvRadix.completed: ")
	fmt.Println(dvRadix.completed)
	fmt.Println(dvRadix.Response)
	// TODO save chart in database
	// TODO show chart and data
	//
	//tempContainer := container.NewVBox(
	//	widget.NewLabel("Tijdelijk label: de container is veranderd"),
	//)
	//gm.Register("charts", UpdateChartsView(tempContainer))
	//gm.Show("charts")
}

func handleChartCompleted() {
	persistCurrentChart()
	gm := GetGuiMgr()
	tempContainer := container.NewVBox(
		widget.NewLabel("Tijdelijk label: de container is veranderd"),
	)
	gm.Register("charts", UpdateChartsView(tempContainer))
	gm.Show("charts")

}

func handleNewconfig() {
	gm := GetGuiMgr()
	gm.Show("config")

}

func handleLangChange(r *Rosetta, s Settings, gm GuiMgr, lang string) {
	w := gm.window
	r.SetLanguage(lang)
	s.DefineLanguage(lang)
	dialog.NewInformation(r.GetText("v_main_language_changed_title"), r.GetText("v_main_language_changed"), w).Show()
	gm.window.Content().Refresh() // TODO try to refresh menu after changing language. Also change rbitems.

}
