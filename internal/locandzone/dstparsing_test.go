/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package locandzone

import (
	"math"
	"testing"
)

func TestProcessedDstLines(t *testing.T) {

	testLines := []string{
		"Neth;1916;1916;5;1;0;00;0;1;00;0;NST",
		"Neth;1916;1916;10;1;0;00;0;0;0;0;AMT",
		"Neth;1917;1917;4;16;2;00;0;1;00;0;NST",
		"Neth;1917;1917;9;17;2;00;0;0;0;0;AMT",
		"Neth;1918;1921;4;0>=1;2;00;0;1;00;0;NST",
		"Neth;1918;1921;9;last0;2;00;0;0;0;0;AMT",
	}

	dp := NewDstParsing()
	result, err := dp.ProcessDstLines(testLines)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(result) != 12 {
		t.Fatalf("Unexpected result length: %v", len(result))
	}
	if result[0].letter != "NST" {
		t.Fatalf("Unexpected letter: %v", result[0].letter)
	}
	if math.Abs(result[0].offset-0.0) > 1e-8 {
		t.Fatalf("Unexpected offset: %v", result[0].offset)
	}

}
