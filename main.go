/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
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

	btnNewChart := widget.NewButton(("New"), func() { fmt.Println("New chart....") })
	btnNewChart.Importance = widget.HighImportance
	btnSearchChart := widget.NewButton(("Search"), func() { fmt.Println("Search chart....") })
	btnSearchChart.Importance = widget.MediumImportance

	text2 := canvas.NewText("Text 2", color.Black)
	text2.TextSize = 36
	//	text3 := canvas.NewText("Text 3", color.Black)
	//	data := widget.NewEntry()
	enigmaHead := canvas.NewText("Enigma 1.0", color.Black)
	enigmaHead.TextSize = 36
	labelGlobalBtns := widget.NewLabel("Modules")
	labelGlobalBtns.Importance = widget.MediumImportance
	labelLocalBtns := widget.NewLabel("Charts")
	labelLocalBtns.Importance = widget.MediumImportance
	circle := canvas.NewCircle(color.RGBA{R: 126, G: 126, B: 255, A: 255})
	circle.Resize(fyne.NewSize(600, 600))
	content := container.NewWithoutLayout(circle)
	boxGlobal := container.New(layout.NewVBoxLayout(), enigmaHead, labelGlobalBtns, button1, button2, button3, button4, button5)
	boxLocal := container.New(layout.NewVBoxLayout(), labelLocalBtns, btnNewChart, btnSearchChart)

	leftBox := container.New(layout.NewVBoxLayout(), boxGlobal, boxLocal)
	//	grid := container.New(layout.NewGridLayout(2), header, label1, text2, text3, data)
	totalBox := container.New(layout.NewHBoxLayout(), leftBox, content)
	return totalBox

}

func main() {

	a := app.New()
	w := a.NewWindow("Enigma")
	w.Resize(fyne.NewSize(1024, 768))
	w.SetContent(makeBox())
	w.ShowAndRun()

}
