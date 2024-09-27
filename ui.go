package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type Circle struct {
	widget.BaseWidget
}

func NewCircle() *Circle {
	c := &Circle{}
	c.ExtendBaseWidget(c)
	return c
}

func (c *Circle) CreateRenderer() fyne.WidgetRenderer {
	circle := canvas.NewCircle(color.NRGBA{R: 0, G: 0, B: 255, A: 255})
	return &circleRenderer{circle: circle}
}

type circleRenderer struct {
	circle *canvas.Circle
}

func (r *circleRenderer) Layout(size fyne.Size) {
	r.circle.Resize(size)
}

func (r *circleRenderer) MinSize() fyne.Size {
	return fyne.NewSize(50, 50) // Minimum size for the circle
}

func (r *circleRenderer) Refresh() {
	canvas.Refresh(r.circle)
}

func (r *circleRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *circleRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.circle}
}

func (r *circleRenderer) Destroy() {}

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
	enigmaHead := canvas.NewText("Enigma 1.0", color.Black)
	enigmaHead.TextSize = 36
	labelGlobalBtns := widget.NewLabel("Modules")
	labelGlobalBtns.Importance = widget.MediumImportance
	labelLocalBtns := widget.NewLabel("Charts")
	labelLocalBtns.Importance = widget.MediumImportance
	labelDescription := widget.NewLabel("Description of chart: namew, date, time, coordinates")
	labelDescription.Importance = widget.MediumImportance
	boxGlobal := container.New(layout.NewVBoxLayout(), enigmaHead, labelGlobalBtns, button1, button2, button3, button4, button5)
	boxLocal := container.New(layout.NewVBoxLayout(), labelLocalBtns, btnNewChart, btnSearchChart)
	leftBox := container.New(layout.NewVBoxLayout(), boxGlobal, boxLocal)
	return leftBox
}

func handleUi() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Resizable Circle")
	labelDescription := widget.NewLabel("Description of chart: name, date, time, coordinates")
	labelDescription.Importance = widget.MediumImportance
	labelSettings := widget.NewLabel("Placeholder for settings: housesystem, parallax, ayanamsha bodies included etc.")
	circle := NewCircle()
	content := container.NewBorder(
		labelDescription,
		labelSettings,
		makeBox(),
		nil,
		circle)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(900, 700))
	myWindow.ShowAndRun()
}
