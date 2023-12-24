package goFRAME

import (
	"fmt"
	"math"
	"strconv"
)

//GetStim : get the stimulus according to the time value
func GetStim(t float64, s Stims, tset TSet) (stimulus float64) {
	stimulus = 0

	for _, stim := range s {
		if math.Abs(t-stim.Start) < tset.Dt/2 {
			fmt.Println("Firing stimulus of", stim.Intensity, "mV at time", stim.Start, "ms with duration", strconv.FormatFloat(stim.End-stim.Start, 'f', 2, 64)+"ms.")
		}

		if t >= stim.Start && t <= stim.End || math.Abs(t-stim.Start) < tset.Dt/2 || math.Abs(t-stim.End) < tset.Dt/2 {
			stimulus += stim.Intensity
		}
	}

	return stimulus
}

func GetGateAB(s State, gAB GateAB, tset TSet) float64 {
	ss := gAB.Ss(s)
	tau := gAB.Tau(s)
	old := s.Gate[gAB.ID]

	return ss + (old-ss)*math.Exp(-tset.Dt/tau)
}

func GetGateDirect(s State, gDirect GateDirect) float64 {
	return gDirect.Ss(s)
}
