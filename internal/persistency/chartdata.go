/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package persistency

import (
	"encoding/csv"
	"enigma-ar/domain"
	"github.com/gocarina/gocsv"
	"io"
	"os"
)

func ReadAllChartData() []*domain.PersistableChart {

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '|'
		return r // Allows use pipe as delimiter
	})
	chartsPath := "./charts.csv"
	chartsFile, chartsFileError := os.OpenFile(chartsPath, os.O_RDWR, os.ModePerm)
	if chartsFileError != nil {
		panic(chartsFileError)
	}
	// Ensure the file is closed once the function returns
	defer chartsFile.Close()
	var charts []*domain.PersistableChart
	if unmarshalError := gocsv.UnmarshalFile(chartsFile, &charts); unmarshalError != nil {
		panic(unmarshalError)
	}
	return charts
}

func ReadAllDateLocationsChartData() []*domain.PersistableDateLocation {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '|'
		return r // Allows use pipe as delimiter
	})
	dateLocsPath := "./datelocs.csv"
	dateLocsFile, dateLocsFileError := os.OpenFile(dateLocsPath, os.O_RDWR, os.ModePerm)
	var dateLocs []*domain.PersistableDateLocation
	if dateLocsFileError != nil { // file does not exist
		return dateLocs // return empty slice
	}
	defer dateLocsFile.Close()
	if unmarshalError := gocsv.UnmarshalFile(dateLocsFile, &dateLocs); unmarshalError != nil {
		panic(unmarshalError)
	}
	return dateLocs
}

func SaveChartData(chart domain.PersistableChart, dateLoc domain.PersistableDateLocation) (int, int, error) {

	allCharts := ReadAllChartData()
	var pkChart = 0
	for _, chart := range allCharts {
		if chart.Id >= pkChart {
			pkChart = chart.Id + 1
		}
	}

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = '|'
		return gocsv.NewSafeCSVWriter(writer)
	})
	chartsPath := "./charts.csv"
	chart.Id = pkChart
	allCharts = append(allCharts, &chart)

	chartsFile, chartsFileError := os.OpenFile(chartsPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if chartsFileError != nil {
		panic(chartsFileError)
	}
	defer chartsFile.Close()

	if marshalFileError := gocsv.MarshalFile(&allCharts, chartsFile); marshalFileError != nil {
		panic(marshalFileError)
	}

	allDateLocs := ReadAllDateLocationsChartData()
	var pkDateLocs = 0
	for _, dateLoc := range allDateLocs {
		if dateLoc.Id >= pkDateLocs {
			pkDateLocs = dateLoc.Id + 1
		}
	}
	dateLoc.Id = pkDateLocs
	dateLoc.ChartId = pkChart
	allDateLocs = append(allDateLocs, &dateLoc)
	dateLocsPath := "./datelocs.csv"

	dateLocsFile, dateLocsFileError := os.OpenFile(dateLocsPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if dateLocsFileError != nil {
		panic(dateLocsFileError)
	}

	if marshalFileError := gocsv.MarshalFile(&allDateLocs, dateLocsFile); marshalFileError != nil {
		panic(marshalFileError)
	}
	return pkChart, pkDateLocs, nil
}
