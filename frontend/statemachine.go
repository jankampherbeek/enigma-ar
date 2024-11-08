/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import "sync"

type MainState int

const (
	Started MainState = iota
	Terminatated
	ChartViewOpened
)

type ChartState int

const (
	ChartStarted ChartState = iota
	NewChartRequested
	NewChartCancelled
	NewChartCompleted
)

type StateMachine struct {
	mainState          MainState
	previousMainState  MainState
	chartState         ChartState
	previousChartState ChartState
}

var smInstance *StateMachine

var smOnce sync.Once

// GetSmInstance returns an rInstance of StateMachine which is always a singleton.
func GetSmInstance() *StateMachine {
	smOnce.Do(func() {
		smInstance = &StateMachine{
			// initialize
		}
	})
	return smInstance
}

// NewMainState updates the state for the main screens.
func (sm StateMachine) NewMainState(newState MainState) {
	sm.previousMainState = sm.mainState
	sm.mainState = newState
	sm.handleChangeMain()
}

// NewChartState updates the state for the chart screens.
func (sm StateMachine) NewChartState(newState ChartState) {
	sm.previousChartState = sm.chartState
	sm.chartState = newState
	sm.handleChangeCharts()
}

// handleChangeMain manages the required actions after a change of one of the main states.
func (sm StateMachine) handleChangeMain() {
	if sm.mainState != sm.previousMainState {
		// handle main screens
	}
}

// handleChangeCharts manages the required actions after a change of one of the chart states.
func (sm StateMachine) handleChangeCharts() {
	if sm.chartState != sm.previousChartState {
		if sm.chartState == NewChartRequested {
			//RadixInputView(r, w)
		}
		if sm.chartState == NewChartCancelled {
			//close
		}
		if sm.chartState == NewChartCompleted {
			// show results
		}
	}
}
