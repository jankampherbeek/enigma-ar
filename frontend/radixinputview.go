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
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type RadixInput interface {
	RadixInputView(r Rosetta, w fyne.Window)
}

type ValidRadixInputData struct {
	NameId      string
	Description string
	Source      string
	ChartCat    domain.ChartCat
	Rating      domain.Rating
	Country     string
	Location    string
	GeoLong     float64
	GeoLat      float64
	Year        int
	Month       int
	Day         int
	Ut          float64
	Calendar    domain.Calendar
	Hour        int
	Minute      int
	Second      int
	TimeZone    domain.TimeZone
	GeoLongLmt  float64 // zero if not applicable
	Dst         bool
}

type RadixInputData struct {
	EntryNameId      *widget.Entry
	EntryDescription *widget.Entry
	EntrySource      *widget.Entry
	EntryLocation    *widget.Entry
	EntryGeoLong     *widget.Entry
	EntryGeoLat      *widget.Entry
	EntryDate        *widget.Entry
	EntryTime        *widget.Entry
	EntryGeoLongLmt  *widget.Entry
	SelBoxCalendar   *widget.Select
	SelBoxChartCat   *widget.Select
	SelBoxRating     *widget.Select
	SelBoxTimeZone   *widget.Select
	ValidData        ValidRadixInputData
}

func NewRadixInputView() RadixInputData {
	return RadixInputData{
		EntryNameId:      nil,
		EntryDescription: nil,
		EntrySource:      nil,
		EntryGeoLong:     nil,
		EntryGeoLat:      nil,
		SelBoxCalendar:   nil,
		SelBoxChartCat:   nil,
		SelBoxRating:     nil,
		SelBoxTimeZone:   nil,
		ValidData:        ValidRadixInputData{},
	}
}

func (rid RadixInputData) RadixInputView(r Rosetta, w fyne.Window) {
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
	rid.EntryNameId = widget.NewEntry()
	rid.EntryDescription = widget.NewEntry()
	rid.EntrySource = widget.NewEntry()
	rid.EntryLocation = widget.NewEntry()
	// todo change entry for location into select, based on the country (radix input)
	rid.EntryGeoLong = widget.NewEntry()
	rid.EntryGeoLong.PlaceHolder = r.GetText("v_input_radix_geolong_placeholder")
	rid.EntryGeoLat = widget.NewEntry()
	rid.EntryGeoLat.PlaceHolder = r.GetText("v_input_radix_geolat_placeholder")
	rid.EntryTime = widget.NewEntry()
	rid.EntryTime.PlaceHolder = r.GetText("v_input_radix_time_placeholder")
	rid.EntryDate = widget.NewEntry()
	rid.EntryDate.PlaceHolder = r.GetText("v_input_radix_date_placeholder")
	rid.EntryGeoLongLmt = widget.NewEntry()
	rid.EntryGeoLongLmt.PlaceHolder = r.GetText("v_input_radix_geolong_placeholder")

	// define input elements: selects
	var optionsCalendar []string
	for _, value := range domain.AllCalendars() {
		optionsCalendar = append(optionsCalendar, r.GetText(value.TextId))
	}
	rid.SelBoxCalendar = widget.NewSelect(optionsCalendar, func(selected string) {})
	rid.SelBoxCalendar.SetSelected(r.GetText("r_cal_gregorian"))

	var optionRating []string
	for _, value := range domain.AllRatings() {
		optionRating = append(optionRating, r.GetText(value.TextId))
	}
	rid.SelBoxRating = widget.NewSelect(optionRating, func(selected string) {})
	rid.SelBoxRating.SetSelected(r.GetText("r_rr_unknown"))

	var optionsChartCat []string
	for _, value := range domain.AllChartCats() {
		optionsChartCat = append(optionsChartCat, r.GetText(value.TextId))
	}
	rid.SelBoxChartCat = widget.NewSelect(optionsChartCat, func(selected string) {})
	rid.SelBoxChartCat.SetSelected(r.GetText("r_cc_unknown"))

	var optionsTimeZone []string
	// todo use tz database to suggest a timezone
	for _, value := range domain.AllTimeZones() {
		optionsTimeZone = append(optionsTimeZone, r.GetText(value.TextId))
	}
	rid.SelBoxTimeZone = widget.NewSelect(optionsTimeZone, func(selected string) {})
	rid.SelBoxTimeZone.SetSelected(r.GetText("r_tz_ut"))

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
		rid.processInput(r)
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
		rid.EntryNameId,
		lblDescription,
		rid.EntryDescription,
		lblSource,
		rid.EntrySource,
		lblCatChart,
		rid.SelBoxChartCat,
		lblRating,
		rid.SelBoxRating,
		txtSectionLocation,
		widget.NewLabel(""),
		lblCountry,
		selBoxCountry,
		lblLocation,
		rid.EntryLocation,
		lblGeoLong,
		rid.EntryGeoLong,
		lblGeoLat,
		rid.EntryGeoLat,
		txtSectionDateTime,
		widget.NewLabel(""),
		lblDate,
		rid.EntryDate,
		lblCalendar,
		rid.SelBoxCalendar,
		lblTime,
		rid.EntryTime,
		lblTimeZone,
		rid.SelBoxTimeZone,
		lblGeoLongLmt,
		rid.EntryGeoLongLmt,
		lblGeoLongLmt,
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

func (rid RadixInputData) processInput(r Rosetta) {
	// TODO implement activities for radix input
	// validate input
	// show any errors
	// fill RadixInputData
	if len(rid.EntryNameId.Text) > 0 {
		rid.ValidData.NameId = rid.EntryNameId.Text
	} else {
		// handle error
	}
	if len(rid.EntryDescription.Text) > 0 {
		rid.ValidData.Description = rid.EntryDescription.Text
	} else {
		rid.ValidData.Description = r.GetText("") // todo key for 'No description'
	}
	if len(rid.EntrySource.Text) > 0 {
		rid.ValidData.Source = rid.EntrySource.Text
	} else {
		rid.ValidData.Source = r.GetText("") // todo key for 'No source'
	}

	ratingId := rid.SelBoxRating.SelectedIndex()
	rid.ValidData.Rating = domain.Rating(ratingId)
	chartCatId := rid.SelBoxChartCat.SelectedIndex()
	rid.ValidData.ChartCat = domain.ChartCat(chartCatId)

	lang := r.GetLanguage()
	gLongVal := NewGeoLongValidator()

	ok, geoLong := gLongVal.CheckGeoLong(rid.EntryGeoLong.Text, lang)
	if !ok {
		// handle error
	}
	rid.ValidData.GeoLong = geoLong

	gLatVal := NewGeoLatValidator()
	ok, geoLat := gLatVal.CheckGeoLat(rid.EntryGeoLat.Text, lang)
	if !ok {
		// handle error
	}
	rid.ValidData.GeoLat = geoLat

	dateVal := NewDateValidator()
	dateOk, y, m, d := dateVal.CheckDate(rid.EntryDate.Text, domain.Calendar(rid.SelBoxCalendar.SelectedIndex()))
	if dateOk {
		rid.ValidData.Year = y
		rid.ValidData.Month = m
		rid.ValidData.Day = d
	} else {
		// handle error
	}
	timeVal := NewTimeValidator()
	timeOk, h, m, s := timeVal.CheckTime(rid.EntryTime.Text)
	if timeOk {
		rid.ValidData.Hour = h
		rid.ValidData.Minute = m
		rid.ValidData.Second = s
	} else {
		// handle error
	}

	var geoLongLmt = 0.0
	if len(rid.EntryGeoLongLmt.Text) > 0 {
		ok, gLLmt := gLongVal.CheckGeoLong(rid.EntryGeoLongLmt.Text, lang)
		if !ok {
			// handle error
		} else {
			geoLongLmt = gLLmt
		}
	}
	rid.ValidData.GeoLongLmt = geoLongLmt

	// todo define UT
	dt := domain.DateTime{
		Year:  rid.ValidData.Year,
		Month: rid.ValidData.Month,
		Day:   rid.ValidData.Day,
		Ut:    0.0,
		Greg:  true,
	}
	jdServer := api.NewJulDayService()
	jd := jdServer.JulDay(&dt)

	var points []domain.ChartPoint
	points = make([]domain.ChartPoint, 3)
	points[0] = domain.Sun
	points[1] = domain.Moon
	points[2] = domain.Mercury

	fcRequest := domain.FullChartRequest{
		Points:    points,
		HouseSys:  domain.HousesPlacidus,
		Ayanamsha: domain.AyanNone,
		CoordSys:  domain.CoordEcliptical,
		ObsPos:    domain.ObsPosGeocentric,
		ProjType:  domain.ProjType2D,
		Jd:        jd,
		GeoLong:   rid.ValidData.GeoLong,
		GeoLat:    rid.ValidData.GeoLat,
	}
	fcServer := api.NewFullChartServer()
	fcResponse, err := fcServer.DefineFullChart(fcRequest)
	if err == nil {
		fmt.Println(fcResponse)
	}

}
