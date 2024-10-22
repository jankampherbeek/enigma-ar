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

//const dateSeparator = "/"
//const timeSeparator = ":"

func RadixInput(r Rosetta, w fyne.Window) {
	var inputPopUp *widget.PopUp
	title := widget.NewLabelWithStyle(r.GetText("v_input_radix_title"), fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	title.Importance = widget.HighImportance
	lblName := widget.NewLabel(r.GetText("v_input_radix_name"))
	lblDescription := widget.NewLabel(r.GetText("v_input_radix_description"))
	lblSource := widget.NewLabel(r.GetText("v_input_radix_source"))
	lblRating := widget.NewLabel(r.GetText("v_input_radix_rating"))
	lblCatChart := widget.NewLabel(r.GetText("v_input_radix_catchart"))
	lblLocation := widget.NewLabel(r.GetText("v_input_radix_locname"))
	lblCountry := widget.NewLabel(r.GetText("v_input_radix_country"))
	//mapText := r.GetText("v_input_radix_map")
	//lblGeoLong := widget.NewLabel(r.GetText("v_input_radix_geolong"))
	//txtGeoLongDirE := r.GetText("v_input_radix_geolongdire")
	//txtGeoLongDirW := r.GetText("v_input_radix_geolongw")
	//lblGeoLat := widget.NewLabel(r.GetText("v_input_radix_geolat"))
	//txtGeoLatDirN := r.GetText("v_input_radix_geolongn")
	//txtGeoLatDirS := r.GetText("v_input_radix_geolongs")
	//lblDate := widget.NewLabel(r.GetText("v_input_radix_date"))
	//lblCalendar := widget.NewLabel(r.GetText("v_input_radix_calendar"))
	//lblYearCount := widget.NewLabel(r.GetText("v_input_radix_yearcount"))
	//lblTime := widget.NewLabel(r.GetText("v_input_radix_time"))
	//lblTimeZone := widget.NewLabel(r.GetText("v_input_radix_timezone"))
	//lblDst := widget.NewLabel(r.GetText("v_input_radix_dst"))
	//lblGeoLongLmt := widget.NewLabel(r.GetText("v_input_radix_geolonglmt"))

	calcText := r.GetText("v_input_radix_calc")
	closeText := r.GetText("g_btn_close")
	helpText := r.GetText("g_btn_help")

	nameEntry := widget.NewEntry()
	descriptionEntry := widget.NewEntry()
	sourceEntry := widget.NewEntry()
	locationEntry := widget.NewEntry()
	locationEntry.SetPlaceHolder("Vul een plaats in")
	/*
		geoLongEntry := widget.NewEntry()
		geoLatEntry := widget.NewEntry()
		timeEntry := widget.NewEntry()
		dateEntry := widget.NewEntry()
		geoLongLmtEntry := widget.NewEntry()
	*/
	//	jdService := api.NewJulDayService()

	chartCatOptions := []string{"male", "female"}
	selBoxChartCat := widget.NewSelect(chartCatOptions, func(selected string) {})
	ratingOptions := []string{"AA", "A"}
	selBoxRating := widget.NewSelect(ratingOptions, func(selected string) {})
	countryOptions := []string{"US", "NL"}
	selBoxCountry := widget.NewSelect(countryOptions, func(selected string) {})

	catRatingContainer := container.NewHBox(selBoxChartCat, lblRating, selBoxRating)
	locationCountryContainer := container.NewHBox(selBoxCountry, lblLocation, locationEntry)

	/*


		longDirOptions := []string{txtGeoLongDirE, txtGeoLongDirW}
		latDirOptions := []string{txtGeoLatDirN, txtGeoLatDirS}
		timeZoneOptions := []string{"UT", "CET"}
		yearCountOptions := []string{"CE", "BCE", "Astronomical"}



		selBoxLongDir := widget.NewSelect(longDirOptions, func(selected string) {})
		selBoxLatDir := widget.NewSelect(latDirOptions, func(selected string) {})
		calEntry := widget.NewEntry()
		selBoxTimeZone := widget.NewSelect(timeZoneOptions, func(selected string) {})
		selBoxYearCount := widget.NewSelect(yearCountOptions, func(selected string) {})
		selBoxLongDirLmt := widget.NewSelect(longDirOptions, func(selected string) {})

		btnMap := widget.NewButton(mapText, func() {})*/
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

	formContainer := container.New(layout.NewFormLayout(),
		lblName,
		nameEntry,
		lblDescription,
		descriptionEntry,
		lblSource,
		sourceEntry,
		lblCatChart,
		catRatingContainer,
		lblCountry, locationCountryContainer,
	)

	viewContent := container.NewVBox(
		title,
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
