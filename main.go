/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package main

import (
	"enigma-ar/frontend"
	"fyne.io/fyne/v2/app"
)

func main() {
	enigmaApp := app.NewWithID("com.radixpro.enigma")
	enigmaApp.SetIcon(resourceIconPng)
	frontend.MakeUI(enigmaApp)
}
