/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"enigma-ar/api"
	"enigma-ar/domain"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

const dateSeparator = "/"
const timeSeparator = ":"
const gregorian = "G"

func CalcJdView(r Rosetta, w fyne.Window) {

	var jdPopUp *widget.PopUp
	title := widget.NewLabel(r.GetText("v_calc_jd_title"))
	title.Importance = widget.HighImportance
	lblDate := widget.NewLabel(r.GetText("v_calc_jd_date"))
	lblTime := widget.NewLabel(r.GetText("v_calc_jd_time"))
	lblCalendar := widget.NewLabel(r.GetText("v_calc_jd_calendar"))
	lblResult := widget.NewLabel(r.GetText("v_calc_jd_result"))
	lblResult.Importance = widget.HighImportance
	calcText := r.GetText("v_calc_jd_btncalc")
	errorText := r.GetText("v_calc_jd_error")
	closeText := r.GetText("g_btn_close")
	helpText := r.GetText("g_btn_help")

	jdService := api.NewJulDayService()
	dateEntry := widget.NewEntry()
	timeEntry := widget.NewEntry()
	calEntry := widget.NewEntry()
	var result = widget.NewLabel("")

	btnCalc := widget.NewButton(calcText, func() {
		var dt domain.DateTime
		valid := true
		date := dateEntry.Text
		time := timeEntry.Text
		calendar := calEntry.Text
		dateItems := strings.Split(date, dateSeparator)
		timeItems := strings.Split(time, timeSeparator)
		year, err := strconv.Atoi(dateItems[0])
		if err != nil {
			valid = false
		}
		if len(dateItems) == 3 && len(timeItems) == 3 {
			month, err := strconv.Atoi(dateItems[1])
			if err != nil {
				valid = false
			}
			day, err := strconv.Atoi(timeItems[2])
			if err != nil {
				valid = false
			}
			hour, err := strconv.ParseFloat(timeItems[0], 64)
			if err != nil {
				valid = false
			}
			minute, err := strconv.ParseFloat(timeItems[1], 64)
			if err != nil {
				valid = false
			}
			second, err := strconv.ParseFloat(timeItems[2], 64)
			if err != nil {
				valid = false
			}
			greg := calendar == gregorian
			dt = domain.DateTime{
				Year:  year,
				Month: month,
				Day:   day,
				Ut:    hour + minute/60.0 + second/3600.0,
				Greg:  greg,
			}
		} else {
			valid = false
		}

		if valid {
			result.SetText(fmt.Sprintf("%f", jdService.JulDay(&dt)))
		} else {
			result.SetText(errorText)
		}

	})
	btnCalc.Importance = widget.HighImportance

	btnClose := widget.NewButton(closeText, func() {
		jdPopUp.Hide()
	})

	btnHelp := widget.NewButton(helpText, func() {
		ShowHelpWindow("calc_jd", r.GetLanguage(), w)
	})
	buttonBar := container.NewHBox(layout.NewSpacer(), btnClose, btnHelp)

	viewContent := container.NewVBox(
		title,
		lblDate,
		dateEntry,
		lblTime,
		timeEntry,
		lblCalendar,
		calEntry,
		btnCalc,
		lblResult,
		result,
		buttonBar,
	)

	jdPopUp = widget.NewModalPopUp(viewContent, w.Canvas())
	jdPopUp.Show()
}
