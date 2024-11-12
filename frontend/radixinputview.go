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

//type RadixInput interface {
//	RadixInputView(r Rosetta, w fyne.Window)
//}

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

// RadixInputView handles the input of data for a new horoscope calculation.
func RadixInputView() fyne.Container {
	//sm := GetSmInstance()
	r := GetRosetta()
	gm := GetGuiMgr()
	w := gm.window
	dvRadix := GetDataVaultRadix()
	dvRadix.completed = false
	ValidData := ValidRadixInputData{}
	//var popupInput *widget.PopUp

	// Texts
	// Title
	txtTitle := canvas.NewText(r.GetText("v_input_radix_title"), color.Gray16{})
	txtTitle.TextSize = 24
	txtTitle.TextStyle = fyne.TextStyle{Bold: true}
	txtTitle.Alignment = fyne.TextAlignCenter

	// Subtitle location
	txtSectionLocation := canvas.NewText("Location", color.Gray16{})
	txtSectionLocation.TextSize = 18
	txtSectionLocation.TextStyle = fyne.TextStyle{Bold: true}

	// Subtitle date and time
	txtSectionDateTime := canvas.NewText("Date and time", color.Gray16{})
	txtSectionDateTime.TextSize = 18
	txtSectionDateTime.TextStyle = fyne.TextStyle{Bold: true}

	// Labels
	lblName := widget.NewLabel(r.GetText("v_input_radix_name"))
	lblName.Importance = widget.HighImportance
	lblDescription := widget.NewLabel(r.GetText("v_input_radix_description"))
	lblSource := widget.NewLabel(r.GetText("v_input_radix_source"))
	lblCatChart := widget.NewLabel(r.GetText("v_input_radix_catchart"))
	lblRating := widget.NewLabel(r.GetText("v_input_radix_rating"))
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

	// Entries
	entryNameId := widget.NewEntry()
	entryDescription := widget.NewEntry()
	entrySource := widget.NewEntry()
	entryLocation := widget.NewEntry()
	// todo change entry for location into select, based on the country (radix input)
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

	// Selections
	// define Chart category
	var optionsChartCat []string
	for _, value := range domain.AllChartCats() {
		optionsChartCat = append(optionsChartCat, r.GetText(value.TextId))
	}
	selBoxChartCat := widget.NewSelect(optionsChartCat, func(selected string) {})
	selBoxChartCat.SetSelected(r.GetText("r_cc_unknown"))

	// define Rodden rating
	var optionRating []string
	for _, value := range domain.AllRatings() {
		optionRating = append(optionRating, r.GetText(value.TextId))
	}
	selBoxRating := widget.NewSelect(optionRating, func(selected string) {})
	selBoxRating.SetSelected(r.GetText("r_rr_unknown"))

	// define country
	optionsCountry := []string{"US", "NL"}
	// TODO use geonames database to populate list of countreies in radix input
	selBoxCountry := widget.NewSelect(optionsCountry, func(selected string) {})

	// define calendar
	var optionsCalendar []string
	for _, value := range domain.AllCalendars() {
		optionsCalendar = append(optionsCalendar, r.GetText(value.TextId))
	}
	selBoxCalendar := widget.NewSelect(optionsCalendar, func(selected string) {})
	selBoxCalendar.SetSelected(r.GetText("r_cal_gregorian"))

	// define timezone
	var optionsTimeZone []string
	// todo use tz database to suggest a timezone
	for _, value := range domain.AllTimeZones() {
		optionsTimeZone = append(optionsTimeZone, r.GetText(value.TextId))
	}
	selBoxTimeZone := widget.NewSelect(optionsTimeZone, func(selected string) {})
	selBoxTimeZone.SetSelected(r.GetText("r_tz_ut"))

	//	define the checkbox for DST
	checkDst := widget.NewCheck("", func(b bool) {
	})

	// Closure to validate and process data
	processInput := func() {
		// TODO implement activities for radix input
		// validate input
		// show any errors
		// fill RadixInputData
		if len(entryNameId.Text) > 0 {
			ValidData.NameId = entryNameId.Text
		} else {
			// handle error
		}
		if len(entryDescription.Text) > 0 {
			ValidData.Description = entryDescription.Text
		} else {
			ValidData.Description = r.GetText("") // todo key for 'No description'
		}
		if len(entrySource.Text) > 0 {
			ValidData.Source = entrySource.Text
		} else {
			ValidData.Source = r.GetText("") // todo key for 'No source'
		}

		ratingId := selBoxRating.SelectedIndex()
		ValidData.Rating = domain.Rating(ratingId)
		chartCatId := selBoxChartCat.SelectedIndex()
		ValidData.ChartCat = domain.ChartCat(chartCatId)

		lang := r.GetLanguage()
		gLongVal := NewGeoLongValidator()

		ok, geoLong := gLongVal.CheckGeoLong(entryGeoLong.Text, lang)
		if !ok {
			// handle error
		}
		ValidData.GeoLong = geoLong

		gLatVal := NewGeoLatValidator()
		ok, geoLat := gLatVal.CheckGeoLat(entryGeoLat.Text, lang)
		if !ok {
			// handle error
		}
		ValidData.GeoLat = geoLat

		dateVal := NewDateValidator()
		dateOk, y, m, d := dateVal.CheckDate(entryDate.Text, domain.Calendar(selBoxCalendar.SelectedIndex()))
		if dateOk {
			ValidData.Year = y
			ValidData.Month = m
			ValidData.Day = d
		} else {
			// handle error
		}
		timeVal := NewTimeValidator()
		timeOk, h, m, s := timeVal.CheckTime(entryTime.Text)
		if timeOk {
			ValidData.Hour = h
			ValidData.Minute = m
			ValidData.Second = s
		} else {
			// handle error
		}

		var geoLongLmt = 0.0
		if len(entryGeoLongLmt.Text) > 0 {
			ok, gLLmt := gLongVal.CheckGeoLong(entryGeoLongLmt.Text, lang)
			if !ok {
				// handle error
			} else {
				geoLongLmt = gLLmt
			}
		}
		ValidData.GeoLongLmt = geoLongLmt

		// todo define UT
		dt := domain.DateTime{
			Year:  ValidData.Year,
			Month: ValidData.Month,
			Day:   ValidData.Day,
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
			GeoLong:   ValidData.GeoLong,
			GeoLat:    ValidData.GeoLat,
		}
		fcServer := api.NewFullChartServer()
		fcResponse, err := fcServer.DefineFullChart(fcRequest)
		if err == nil {
			//dvRadix.Response = fcResponse
			dvRadix.AddCalculatedChart(fcResponse)
			dvRadix.completed = true
			fmt.Println("In closure: dvRadix.completed: ")
			fmt.Println(dvRadix.completed)
			fmt.Println(fcResponse)
		} else {
			dvRadix.completed = false
		}
	}

	// Form container
	formContainer := container.New(layout.NewFormLayout(),
		lblName,
		entryNameId,
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

	// define buttons
	txtCalc := r.GetText("v_input_radix_calc")
	txtClose := r.GetText("g_btn_close")
	txtHelp := r.GetText("g_btn_help")

	btnCalc := widget.NewButton(txtCalc, func() {
		processInput()
		changeState(chartCalcCompleted)
		//	sm.NewChartState(NewChartCompleted)
		//	popupInput.Hide()

	})
	btnCalc.Importance = widget.HighImportance

	btnClose := widget.NewButton(txtClose, func() {
		//popupInput.Hide()
	})

	btnHelp := widget.NewButton(txtHelp, func() {
		// TODO create help page for radix input
		ShowHelpWindow("input_radix", r.GetLanguage(), w)
	})
	buttonBar := container.NewHBox(layout.NewSpacer(), btnClose, btnHelp, btnCalc)

	viewContent := container.NewVBox(
		txtTitle,
		formContainer,
		buttonBar,
	)

	return *viewContent

	//popupInput = widget.NewModalPopUp(viewContent, w.Canvas())
	//popupInput.Resize(fyne.NewSize(500, 800))
	//popupInput.Show()
}
