/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"os"
)

func ShowHelpWindow(viewName string, lang string, w fyne.Window) {

	var helpPopUp *widget.PopUp
	fileName := "./help/" + viewName + "_" + lang + ".md"
	helpBytes, err := os.ReadFile(fileName)
	if err != nil {
		// handle error
	}
	helpContent := string(helpBytes)
	richText := widget.NewRichTextFromMarkdown(helpContent)
	richText.Wrapping = fyne.TextWrapWord
	closeButton := widget.NewButton("Close", func() {
		helpPopUp.Hide()
	})
	helpContainer := container.NewVBox(
		richText,
		closeButton)
	scroll := container.NewScroll(helpContainer)
	scroll.SetMinSize(fyne.NewSize(500, 400))

	bg := canvas.NewRectangle(theme.BackgroundColor())
	bg.SetMinSize(fyne.NewSize(500, 400))
	content := container.NewStack(bg, scroll)

	helpPopUp = widget.NewModalPopUp(content, w.Canvas())
	helpPopUp.Show()

}
