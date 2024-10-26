/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"enigma-ar/domain"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func RadixInput(r Rosetta, w fyne.Window) {
	var popupInput *widget.PopUp

	// define title and subtitles
	txtTitle := canvas.NewText(r.GetText("v_input_radix_title"), color.Gray16{})
	txtTitle.TextSize = 24
	txtTitle.TextStyle = fyne.TextStyle{Bold: true}
	txtTitle.Alignment = fyne.TextAlignCenter
	txtSectionLocation := canvas.NewText("Location", color.Gray16{})
	txtSectionLocation.TextSize = 18
	txtSectionLocation.TextStyle = fyne.TextStyle{Bold: true}
	txtSectionDateTime := canvas.NewText("Date and time", color.Gray16{})
	txtSectionDateTime.TextSize = 18
	txtSectionDateTime.TextStyle = fyne.TextStyle{Bold: true}

	// define all labels
	lblName := widget.NewLabel(r.GetText("v_input_radix_name"))
	lblName.Importance = widget.HighImportance
	lblDescription := widget.NewLabel(r.GetText("v_input_radix_description"))
	lblSource := widget.NewLabel(r.GetText("v_input_radix_source"))
	lblRating := widget.NewLabel(r.GetText("v_input_radix_rating"))
	lblCatChart := widget.NewLabel(r.GetText("v_input_radix_catchart"))
	lblLocation := widget.NewLabel(r.GetText("v_input_radix_locname"))
	lblCountry := widget.NewLabel(r.GetText("v_input_radix_country"))
	lblGeoLong := widget.NewLabel(r.GetText("v_input_radix_geolong"))
	lblGeoLong.Importance = widget.HighImportance
	lblGeoLat := widget.NewLabel(r.GetText("v_input_radix_geolat"))
	lblGeoLat.Importance = widget.HighImportance
	lblDate := widget.NewLabel(r.GetText("v_input_radix_date"))
	lblDate.Importance = widget.HighImportance
	lblCalendar := widget.NewLabel(r.GetText("v_input_radix_calendar"))
	lblCalendar.Importance = widget.HighImportance
	lblTime := widget.NewLabel(r.GetText("v_input_radix_time"))
	lblTime.Importance = widget.HighImportance
	lblTimeZone := widget.NewLabel(r.GetText("v_input_radix_timezone"))
	lblTimeZone.Importance = widget.HighImportance
	lblDst := widget.NewLabel(r.GetText("v_input_radix_dst"))
	lblGeoLongLmt := widget.NewLabel(r.GetText("v_input_radix_geolonglmt"))

	// define input elements: entries
	entryName := widget.NewEntry()
	entryDescription := widget.NewEntry()
	entrySource := widget.NewEntry()
	entryLocation := widget.NewEntry()
	// todo change entry for location into select based on the country (radix input)
	entryGeoLong := widget.NewEntry()
	entryGeoLong.PlaceHolder = r.GetText("v_input_radix_geolong_placeholder")
	entryGeoLat := widget.NewEntry()
	entryGeoLat.PlaceHolder = r.GetText("v_input_radix_geolat_placeholder")
	entryTime := widget.NewEntry()
	entryTime.PlaceHolder = r.GetText("v_input_radix_time_placeholder")
	entryDate := widget.NewEntry()
	entryDate.PlaceHolder = r.GetText("v_input_radix_date_placeholder")
	entryGeoLongLmt := widget.NewEntry()
	entryGeoLongLmt.PlaceHolder = r.GetText("v_input_radix_geolong_placeholder")

	// define input elements: selects
	var optionsCalendar []string
	for _, value := range domain.AllCalendars() {
		optionsCalendar = append(optionsCalendar, r.GetText(value.TextId))
	}
	selBoxCalendar := widget.NewSelect(optionsCalendar, func(selected string) {})
	selBoxCalendar.SetSelected(r.GetText("r_cal_gregorian"))

	var optionRating []string
	for _, value := range domain.AllRatings() {
		optionRating = append(optionRating, r.GetText(value.TextId))
	}
	selBoxRating := widget.NewSelect(optionRating, func(selected string) {})
	selBoxRating.SetSelected(r.GetText("r_rr_unknown"))

	var optionsChartCat []string
	for _, value := range domain.AllChartCats() {
		optionsChartCat = append(optionsChartCat, r.GetText(value.TextId))
	}
	selBoxChartCat := widget.NewSelect(optionsChartCat, func(selected string) {})
	selBoxChartCat.SetSelected(r.GetText("r_cc_unknown"))

	var optionsTimeZone []string
	// todo use tz database to suggest a timezone
	for _, value := range domain.AllTimeZones() {
		optionsTimeZone = append(optionsTimeZone, r.GetText(value.TextId))
	}
	selBoxTimeZone := widget.NewSelect(optionsTimeZone, func(selected string) {})
	selBoxTimeZone.SetSelected(r.GetText("r_tz_ut"))

	optionsCountry := []string{"US", "NL"}
	// TODO use geonames database to populate list of countreies in radix input
	selBoxCountry := widget.NewSelect(optionsCountry, func(selected string) {})

	// define the checkbox for DST
	checkDst := widget.NewCheck("", func(b bool) {

	})

	// define buttons
	txtCalc := r.GetText("v_input_radix_calc")
	txtClose := r.GetText("g_btn_close")
	txtHelp := r.GetText("g_btn_help")

	btnCalc := widget.NewButton(txtCalc, func() {
		// TODO implement activities for radix input
		// validate input
		// show any errors
		// create RadixInputData
		// calculate chart by calling DataVault.DefineFullChart(inputData)
	})
	btnCalc.Importance = widget.HighImportance

	btnClose := widget.NewButton(txtClose, func() {
		popupInput.Hide()
	})

	btnHelp := widget.NewButton(txtHelp, func() {
		// TODO create help page for radix input
		ShowHelpWindow("input_radix", r.GetLanguage(), w)
	})
	buttonBar := container.NewHBox(layout.NewSpacer(), btnClose, btnHelp, btnCalc)

	// build formcontainer
	formContainer := container.New(layout.NewFormLayout(),
		lblName,
		entryName,
		lblDescription,
		entryDescription,
		lblSource,
		entrySource,
		lblCatChart,
		selBoxChartCat,
		lblRating,
		selBoxRating,
		txtSectionLocation,
		widget.NewLabel(""),
		lblCountry,
		selBoxCountry,
		lblLocation,
		entryLocation,
		lblGeoLong,
		entryGeoLong,
		lblGeoLat,
		entryGeoLat,
		txtSectionDateTime,
		widget.NewLabel(""),
		lblDate,
		entryDate,
		lblCalendar,
		selBoxCalendar,
		lblTime,
		entryTime,
		lblTimeZone,
		selBoxTimeZone,
		lblGeoLongLmt,
		entryGeoLongLmt,
		lblDst,
		checkDst,
	)

	// create content
	viewContent := container.NewVBox(
		txtTitle,
		formContainer,
		buttonBar,
	)

	// create popup
	popupInput = widget.NewModalPopUp(viewContent, w.Canvas())
	popupInput.Resize(fyne.NewSize(500, 800))
	popupInput.Show()

}
