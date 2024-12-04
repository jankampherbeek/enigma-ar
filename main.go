/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package main

import (
	_ "embed"
	"enigma-ar/frontend"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

// Embedding font files
//
//go:embed fonts/EnigmaAstrology2.ttf
var fontData []byte

type myTheme struct{}

func (m *myTheme) Font(style fyne.TextStyle) fyne.Resource { // Textstyle is ignored
	return fyne.NewStaticResource("EnigmaAstrology2.ttf", fontData)
}

func (m *myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m *myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m *myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func main() {
	enigmaApp := app.NewWithID("com.radixpro.enigma")
	enigmaApp.Settings().SetTheme(&myTheme{})
	enigmaApp.SetIcon(resourceIconPng)
	frontend.MakeUI(enigmaApp)
}
