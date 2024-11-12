/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (gm *GuiMgr) CreateToolBar() *widget.Toolbar {
	toolBar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			gm.Show("home")
		}),
		widget.NewToolbarAction(theme.SearchIcon(), func() {
			searchTermEntry := widget.NewEntry()
			searchForm := dialog.NewForm(
				"Search chart",
				"Search",
				"Cancel",
				[]*widget.FormItem{
					{Text: "Part of name", Widget: searchTermEntry},
				},
				func(valid bool) {
					if valid {
						// start search
					}
				},
				gm.window)
			searchForm.Resize(fyne.NewSize(400, 160))
			searchForm.Show()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			gm.Show("config")
		}),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			gm.Show("manual")
		}),
		widget.NewToolbarAction(theme.InfoIcon(), func() {}),
		widget.NewToolbarAction(theme.LogoutIcon(), func() {
		}),
	)
	return toolBar
}
