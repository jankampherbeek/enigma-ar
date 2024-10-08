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

func NewChartsView(gm *GuiMgr) fyne.CanvasObject {
	header := widget.NewLabelWithStyle(gm.Rosetta.GetText("btnTutorial"), fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	button := widget.NewButton("Back to Home", func() {
		gm.Show("home")
	})
	content := container.NewVBox(
		header,
		button,
	)
	return content
}
