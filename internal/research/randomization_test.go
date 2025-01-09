/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package research

import (
	"math"
	"testing"
)

func TestGetIntegersHappyFlow(t *testing.T) {
	start := 10
	end := 50
	count := 100
	r := NewCGroupRandomization()
	result, err := r.GetIntegers(start, end, count)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if len(result) != count {
		t.Errorf("GetIntegers() returned %d results, expected %d", len(result), count)
	}
	for i := range result {
		if result[i] < start || result[i] >= end {
			t.Errorf("GetIntegerHappyFlow returned value that is out of range: %d", result[i])
		}
	}
}

// TestGetIntegersDifferenceHappyFlow should return two different list, however there is a very small
// possibility that both lists are equal. If that occurs, repeat the test.
func TestGetIntegersDifferenceHappyFlow(t *testing.T) {
	start := 10
	end := 50
	count := 100
	r := NewCGroupRandomization()
	result1, err := r.GetIntegers(start, end, count)
	if err != nil {
		t.Errorf("GetIntegers() returned error for result1: %v", err)
	}
	result2, err := r.GetIntegers(start, end, count)
	if err != nil {
		t.Errorf("GetIntegers() returned error for result2: %v", err)
	}
	listsAreEqual := true
	for i := range result1 {
		if result1[i] != result2[i] {
			listsAreEqual = false
		}
	}
	if listsAreEqual {
		t.Errorf("GetIntegers() returned twice the same list. Try to repeat the test.")
	}
}

func TestGetIntegersWrongSequenceOfParameters(t *testing.T) {
	r := NewCGroupRandomization()
	result, err := r.GetIntegers(10, 2, 30)
	if result != nil {
		t.Errorf("GetIntegers() returned %d, expected nil", result)
	}
	if err == nil {
		t.Errorf("GetIntegers() returned no error, expected error")
	}
}

func TestGetIntegersInvalidCount(t *testing.T) {
	r := NewCGroupRandomization()
	result, err := r.GetIntegers(10, 2, -1)
	if result != nil {
		t.Errorf("GetIntegers() returned %d while count was invalid, expected nil", result)
	}
	if err == nil {
		t.Errorf("GetIntegers() returned no error while count was invalid")
	}
}

const DELTA = 0.0001 // Define the delta constant for float comparisons

func TestShuffleIntList(t *testing.T) {
	dataItems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	data := make([]int, len(dataItems))
	copy(data, dataItems)
	r := NewCGroupRandomization()
	r.ShuffleIntList(data)
	if len(data) != len(dataItems) {
		t.Errorf("Expected length %d, got %d", len(dataItems), len(data))
	}
	// Test if at least one of the first 6 elements is in a different position
	if data[0] == 1 && data[1] == 2 && data[2] == 3 &&
		data[3] == 4 && data[4] == 5 && data[5] == 6 {
		t.Error("List appears to not be shuffled")
	}
}

func TestShuffleDoubleList(t *testing.T) {
	dataItems := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0, 11.1, 12.2}
	data := make([]float64, len(dataItems))
	copy(data, dataItems)
	r := NewCGroupRandomization()
	r.ShuffleFloatList(data)
	if len(data) != len(dataItems) {
		t.Errorf("Expected length %d, got %d", len(dataItems), len(data))
	}

	// Test if at least one of the first 6 elements is in a different position
	if math.Abs(data[0]-1.1) <= DELTA &&
		math.Abs(data[1]-2.2) <= DELTA &&
		math.Abs(data[2]-3.3) <= DELTA &&
		math.Abs(data[3]-4.4) <= DELTA &&
		math.Abs(data[4]-5.5) <= DELTA &&
		math.Abs(data[5]-6.6) <= DELTA {
		t.Error("List appears to not be shuffled")
	}
}
