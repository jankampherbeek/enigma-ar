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
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func createBottomPart() *fyne.Container {
	bottomLineTxt := "Enigma Astrology Research - version 1.0"
	bottomLine := widget.NewLabelWithStyle(bottomLineTxt, fyne.TextAlignCenter, fyne.TextStyle{
		Bold:      true,
		Italic:    false,
		Monospace: false,
		Symbol:    false,
		TabWidth:  0,
	})
	return container.New(layout.NewHBoxLayout(), bottomLine)
}

func createMainPart() *fyne.Container {
	placeHolder := widget.NewLabel("Placeholder for main container")
	return container.New(layout.NewHBoxLayout(), placeHolder)
}

func NewHomeView(gm *GuiMgr) fyne.CanvasObject {
	toolBar := gm.createToolBar()
	ContentBottom := createBottomPart()
	ContentCenter := createMainPart()
	button := widget.NewButton("Charts", func() {
		gm.Show("charts")
	})
	content := container.NewBorder(
		toolBar,
		ContentBottom,
		makeBox(),
		button,
		ContentCenter)
	return content
}
