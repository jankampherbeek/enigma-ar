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

//const dateSeparator = "/"
//const timeSeparator = ":"

func RadixInput(r Rosetta, w fyne.Window) {
	var inputPopUp *widget.PopUp
	txtTitle := canvas.NewText(r.GetText("v_input_radix_title"), color.Gray16{})
	txtTitle.TextSize = 24
	txtTitle.TextStyle = fyne.TextStyle{Bold: true}
	txtTitle.Alignment = fyne.TextAlignCenter

	lblName := widget.NewLabel(r.GetText("v_input_radix_name"))
	lblName.Importance = widget.HighImportance
	lblDescription := widget.NewLabel(r.GetText("v_input_radix_description"))
	lblSource := widget.NewLabel(r.GetText("v_input_radix_source"))
	lblRating := widget.NewLabel(r.GetText("v_input_radix_rating"))
	lblCatChart := widget.NewLabel(r.GetText("v_input_radix_catchart"))
	lblLocation := widget.NewLabel(r.GetText("v_input_radix_locname"))
	lblCountry := widget.NewLabel(r.GetText("v_input_radix_country"))
	//	mapText := r.GetText("v_input_radix_map")
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

	calcText := r.GetText("v_input_radix_calc")
	closeText := r.GetText("g_btn_close")
	helpText := r.GetText("g_btn_help")

	nameEntry := widget.NewEntry()
	descriptionEntry := widget.NewEntry()
	sourceEntry := widget.NewEntry()
	locationEntry := widget.NewEntry()

	geoLongEntry := widget.NewEntry()
	geoLongEntry.PlaceHolder = r.GetText("v_input_radix_geolong_placeholder")
	geoLatEntry := widget.NewEntry()
	geoLatEntry.PlaceHolder = r.GetText("v_input_radix_geolat_placeholder")
	timeEntry := widget.NewEntry()
	timeEntry.PlaceHolder = r.GetText("v_input_radix_time_placeholder")
	dateEntry := widget.NewEntry()
	dateEntry.PlaceHolder = r.GetText("v_input_radix_date_placeholder")
	geoLongLmtEntry := widget.NewEntry()
	geoLongLmtEntry.PlaceHolder = r.GetText("v_input_radix_geolong_placeholder")
	//	jdService := api.NewJulDayService()

	calendarOptions := []string{}
	for _, value := range domain.AllCalendars() {
		calendarOptions = append(calendarOptions, r.GetText(value.TextId))
	}
	selBoxCalendar := widget.NewSelect(calendarOptions, func(selected string) {})
	selBoxCalendar.SetSelected(r.GetText("r_cal_gregorian"))

	ratingOptions := []string{}
	for _, value := range domain.AllRatings() {
		ratingOptions = append(ratingOptions, r.GetText(value.TextId))
	}
	selBoxRating := widget.NewSelect(ratingOptions, func(selected string) {})
	selBoxRating.SetSelected(r.GetText("r_rr_unknown"))

	chartCatOptions := []string{}
	for _, value := range domain.AllChartCats() {
		chartCatOptions = append(chartCatOptions, r.GetText(value.TextId))
	}
	selBoxChartCat := widget.NewSelect(chartCatOptions, func(selected string) {})
	selBoxChartCat.SetSelected(r.GetText("r_cc_unknown"))

	timeZoneOptions := []string{}
	for _, value := range domain.AllTimeZones() {
		timeZoneOptions = append(timeZoneOptions, r.GetText(value.TextId))
	}
	selBoxTimeZone := widget.NewSelect(timeZoneOptions, func(selected string) {})
	selBoxTimeZone.SetSelected(r.GetText("r_tz_ut"))

	countryOptions := []string{"US", "NL"}
	selBoxCountry := widget.NewSelect(countryOptions, func(selected string) {})

	dstCheck := widget.NewCheck("", func(b bool) {

	})

	btnCalc := widget.NewButton(calcText, func() {
	})
	btnCalc.Importance = widget.HighImportance

	btnClose := widget.NewButton(closeText, func() {
		inputPopUp.Hide()
	})

	btnHelp := widget.NewButton(helpText, func() {
		ShowHelpWindow("input_radix", r.GetLanguage(), w)
	})
	buttonBar := container.NewHBox(layout.NewSpacer(), btnClose, btnHelp, btnCalc)

	//lblSectionLocation := widget.NewLabelWithStyle("Location", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	txtSectionLocation := canvas.NewText("Location", color.Gray16{})
	txtSectionLocation.TextSize = 18
	txtSectionLocation.TextStyle = fyne.TextStyle{Bold: true}
	//lblSectionDateTime := widget.NewLabelWithStyle("Date and Time", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	//lblSectionLocation.Importance = widget.HighImportance
	//lblSectionDateTime.Importance = widget.HighImportance
	txtSectionDateTime := canvas.NewText("Date and time", color.Gray16{})
	txtSectionDateTime.TextSize = 18
	txtSectionDateTime.TextStyle = fyne.TextStyle{Bold: true}
	//lblSectionDateTime := widget.NewLabelWithStyle("Date and Time", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	formContainer := container.New(layout.NewFormLayout(),
		lblName,
		nameEntry,
		lblDescription,
		descriptionEntry,
		lblSource,
		sourceEntry,
		lblCatChart,
		selBoxChartCat,
		lblRating,
		selBoxRating,
		txtSectionLocation,
		widget.NewLabel(""),
		lblCountry,
		selBoxCountry,
		lblLocation,
		locationEntry,
		lblGeoLong,
		geoLongEntry,
		lblGeoLat,
		geoLatEntry,
		txtSectionDateTime,
		widget.NewLabel(""),
		lblDate,
		dateEntry,
		lblCalendar,
		selBoxCalendar,
		lblTime,
		timeEntry,
		lblTimeZone,
		selBoxTimeZone,
		lblGeoLongLmt,
		geoLongLmtEntry,
		lblDst,
		dstCheck,
	)

	viewContent := container.NewVBox(
		txtTitle,
		formContainer,
		/*		lblName,
				nameEntry,
				lblDescription,
				descriptionEntry,
				chartMetaContainer,
				locationContainer,
				coordinateContainer,
				dateCalContainer,
				fullTimeContainer,
				fullLmtLongContainer,*/
		buttonBar,
	)

	inputPopUp = widget.NewModalPopUp(viewContent, w.Canvas())
	inputPopUp.Resize(fyne.NewSize(500, 700))
	inputPopUp.Show()

}
