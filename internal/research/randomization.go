/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package research

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// CGroupRandomizer defines methods for generating random numbers and shuffling lists for control group construction.
// The standard True Random Number Generator from Go is used.
type CGroupRandomizer interface {
	GetIntegers(minInclusive, maxExclusive, count int) ([]int, error)
	GetIntegersWithMax(maxExclusive, count int) ([]int, error)
	ShuffleIntList(data []int) []int
	ShuffleFloatList(data []float64) []float64
}

type CGroupRandomization struct{}

func NewCGroupRandomization() CGroupRandomizer {
	return CGroupRandomization{}
}

func (cgr CGroupRandomization) GetIntegers(minInclusive, maxExclusive, count int) ([]int, error) {
	if minInclusive >= maxExclusive || count <= 0 {
		return nil, errors.New("wrong sequence of parameters")
	}
	result := make([]int, count)
	range_ := maxExclusive - minInclusive
	for i := 0; i < count; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(range_)))
		if err != nil {
			return nil, err
		}
		result[i] = int(n.Int64()) + minInclusive
	}
	return result, nil
}

func (cgr CGroupRandomization) GetIntegersWithMax(maxExclusive, count int) ([]int, error) {
	return cgr.GetIntegers(0, maxExclusive, count)
}

func (cgr CGroupRandomization) ShuffleIntList(data []int) []int {
	shuffle(data)
	return data
}

func (cgr CGroupRandomization) ShuffleFloatList(data []float64) []float64 {
	shuffle(data)
	return data
}

// shuffle is a generic function to shuffle any slice
func shuffle[T any](slice []T) {
	n := len(slice)
	if n <= 1 {
		return
	}
	rng := NewCGroupRandomization()
	randomNumbers, err := rng.GetIntegers(0, n-1, n-1)
	if err != nil {
		return
	}
	for i := n - 1; i > 0; i-- {
		k := randomNumbers[n-1-i]
		slice[i], slice[k] = slice[k], slice[i]
	}
}
