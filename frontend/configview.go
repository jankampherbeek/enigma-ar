/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"enigma-ar/domain"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func NewConfigView(gm *GuiMgr) fyne.CanvasObject {
	r := GetRosetta()
	header := widget.NewLabelWithStyle("Configuration", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// define buttons
	txtSave := r.GetText("g_btn_save")
	txtClose := r.GetText("g_btn_close")
	txtHelp := r.GetText("g_btn_help")

	btnSave := widget.NewButton(txtSave, func() {
	})
	btnSave.Importance = widget.HighImportance

	btnClose := widget.NewButton(txtClose, func() {
		gm.Show("home")
	})

	btnHelp := widget.NewButton(txtHelp, func() {
		// TODO create help page for configuration

	})
	buttonBar := container.NewHBox(layout.NewSpacer(), btnClose, btnHelp, btnSave)

	configTabs := container.NewAppTabs(
		container.NewTabItem("General", createTabGeneralContent()),
		container.NewTabItem("Points", createTabPointsContent()),
		container.NewTabItem("Aspects", createTabAspectsContent()),
		container.NewTabItem("Glyphs", createTabGlyphsContent()),
		container.NewTabItem("Colors", createTabColorsContent()),
		container.NewTabItem("Progressive", createTabProgressiveContent()),
	)
	configTabs.SetTabLocation(container.TabLocationLeading)

	configView := container.NewVBox(
		header,
		configTabs,
		buttonBar)
	return configView
}

func createTabGeneralContent() fyne.CanvasObject {
	r := GetRosetta()
	generalTitle := canvas.NewText("Configuration - general", color.Gray16{})
	generalTitle.TextSize = 24
	generalTitle.TextStyle = fyne.TextStyle{Bold: true}
	generalTitle.Alignment = fyne.TextAlignCenter
	txtEplanation := widget.NewRichTextFromMarkdown(createExplanation())
	txtEplanation.Wrapping = fyne.TextWrapWord

	lblHouses := widget.NewLabel("Houses")
	lblAyanamsha := widget.NewLabel("Ayanamsha")
	lblObserverPos := widget.NewLabel("Observer position")
	lblProjectionType := widget.NewLabel("Projection type")
	var optionHouses []string
	for _, value := range domain.AllHouseSystems() {
		optionHouses = append(optionHouses, r.GetText(value.TextId))
	}
	selBoxHouses := widget.NewSelect(optionHouses, func(selected string) {})
	selBoxHouses.SetSelected(r.GetText("r_hs_placidus"))

	var optionAyanamsha []string
	for _, value := range domain.AllAyanamshas() {
		optionAyanamsha = append(optionAyanamsha, r.GetText(value.TextId))
	}
	selBoxAyanamsha := widget.NewSelect(optionAyanamsha, func(selected string) {})
	selBoxAyanamsha.SetSelected(r.GetText("r_ay_none"))

	var optionObserverPos []string
	for _, value := range domain.AllObserverPositions() {
		optionObserverPos = append(optionObserverPos, r.GetText(value.TextId))
	}
	selBoxObserverPos := widget.NewSelect(optionObserverPos, func(selected string) {})
	selBoxObserverPos.SetSelected(r.GetText("r_op_geocentric"))

	var optionProjectionType []string
	for _, value := range domain.AllProjectionTypes() {
		optionProjectionType = append(optionProjectionType, r.GetText(value.TextId))
	}
	selBoxProjectionType := widget.NewSelect(optionProjectionType, func(selected string) {})
	selBoxProjectionType.SetSelected(r.GetText("r_pt_2d"))

	leftView := container.New(layout.NewFormLayout(),
		lblHouses,
		selBoxHouses,
		lblAyanamsha,
		selBoxAyanamsha,
		lblObserverPos,
		selBoxObserverPos,
		lblProjectionType,
		selBoxProjectionType)
	rightView := container.New(layout.NewVBoxLayout(),
		generalTitle,
		txtEplanation,
	)

	leftRightView := container.New(layout.NewHBoxLayout(),
		leftView, rightView)

	generalView := container.New(layout.NewVBoxLayout(),
		leftRightView)

	return generalView
}

type TableData struct {
	data    [][]string
	headers []string
}

func (t *TableData) Length() int {
	return len(t.data)
}

func (t *TableData) CreateCell() fyne.CanvasObject {
	return widget.NewLabel("Template")
}

func (t *TableData) UpdateCell(cell fyne.CanvasObject, id widget.TableCellID) {
	label := cell.(*widget.Label)
	if id.Row == 0 {
		label.SetText(t.headers[id.Col])
	} else {
		label.SetText(t.data[id.Row-1][id.Col])
	}
}

func createTabPointsContent() fyne.CanvasObject {
	generalTitle := canvas.NewText("Configuration - celestial points", color.Gray16{})
	generalTitle.TextSize = 24
	generalTitle.TextStyle = fyne.TextStyle{Bold: true}
	generalTitle.Alignment = fyne.TextAlignCenter

	type TableRow struct {
		name   string
		age    int
		active bool
	}

	data := []TableRow{
		{"John Doe", 30, true},
		{"Jane Smith", 25, false},
		{"Bob Johnson", 35, true},
	}

	table := widget.NewTable(
		// Size function - returns total rows and columns
		func() (int, int) {
			return len(data) + 1, 3 // +1 for header row
		},
		// Create template cell
		func() fyne.CanvasObject {
			return container.NewMax(
				widget.NewLabel(""),      // for text columns
				widget.NewCheck("", nil), // for checkbox column
			)
		},
		// Update cell content
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			// Get the container
			cont := cell.(*fyne.Container)
			label := cont.Objects[0].(*widget.Label)
			check := cont.Objects[1].(*widget.Check)

			// Hide both by default
			label.Hide()
			check.Hide()

			if id.Row == 0 {
				// Header row
				label.Show()
				label.TextStyle = fyne.TextStyle{Bold: true}
				switch id.Col {
				case 0:
					label.SetText("Name")
				case 1:
					label.SetText("Age")
				case 2:
					label.SetText("Active")
				}
			} else {
				// Data rows
				row := data[id.Row-1]
				switch id.Col {
				case 0:
					// Name column
					label.Show()
					label.SetText(row.name)
				case 1:
					// Age column
					label.Show()
					label.SetText(fmt.Sprintf("%d", row.age))
				case 2:
					// Active column
					check.Show()
					check.SetChecked(row.active)
					check.OnChanged = func(b bool) {
						data[id.Row-1].active = b
					}
				}
			}
		},
	)

	table.SetColumnWidth(0, 150)
	table.SetColumnWidth(1, 80)
	table.SetColumnWidth(2, 120)

	pointsContent := container.NewVBox(
		container.NewPadded(table))
	return pointsContent
}

func createTabAspectsContent() fyne.CanvasObject {
	aspectsContent := widget.NewLabel("Aspects content")
	return aspectsContent
}

func createTabGlyphsContent() fyne.CanvasObject {
	glyphsContent := widget.NewLabel("Glyphs content")
	return glyphsContent
}

func createTabColorsContent() fyne.CanvasObject {
	colorsContent := widget.NewLabel("Colors content")
	return colorsContent
}

func createTabProgressiveContent() fyne.CanvasObject {
	progressiveContent := widget.NewLabel("Primdir content")
	return progressiveContent
}

func createExplanation() string {
	return "## Configuration - general\n\nTo edit the current configuration, you can change the values in the left part of the screen. \n\nClicking one of the tabs *Points*, *Aspects*, *Glyphs*, or *Progressive* will give you access to other parts of the configuration.\n\nPlease note that you are changing the current configuration.\n\nYou can also create a number of new configurations and select which one you want to use. "
}
