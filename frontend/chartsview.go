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
	diameter := fyne.Min(size.Width, size.Height)
	r.circle.Resize(fyne.NewSize(diameter, diameter))
	r.circle.Move(fyne.NewPos((size.Width-diameter)/2, (size.Height-diameter)/2))
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

func createButtons() *fyne.Container {
	btnHelp := widget.NewButton("Help", func() {})
	btnCancel := widget.NewButton("Cancel", func() {})
	btnExit := widget.NewButton("Exit", func() {})
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer(), btnHelp, btnCancel, btnExit)
}

func NewChartsView(gm *GuiMgr) fyne.CanvasObject {
	toolBar := gm.CreateToolBar()
	circle := NewCircle() // todo replace circle with real chart, probably in svg format
	buttons := createButtons()
	circleContainer := container.NewStack(circle)
	ariesGlyph := '\uE000'
	taurusGlyph := '\uE001'
	tempLabel := fmt.Sprintf("Temporary label for right part %c %c", ariesGlyph, taurusGlyph)
	mainPart := container.NewHSplit(
		circleContainer,
		widget.NewLabel(tempLabel),
	)
	content := container.NewBorder(toolBar, buttons, nil, nil, mainPart)

	return content
}

func UpdateChartsView(newContainer *fyne.Container) fyne.CanvasObject {
	gm := GetGuiMgr()
	toolBar := gm.CreateToolBar()
	circle := NewCircle() // todo replace circle with real chart, probably in svg format
	buttons := createButtons()
	circleContainer := container.NewStack(circle)
	mainPart := container.NewHSplit(
		circleContainer,
		newContainer,
	)
	content := container.NewBorder(toolBar, buttons, nil, nil, mainPart)
	return content

}
