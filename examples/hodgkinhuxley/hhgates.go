package main

import (
	"goFRAME"
	"math"
)

var (
	h = goFRAME.GateAB{Alpha: ah, Beta: bh, ID: "h"}
	m = goFRAME.GateAB{Alpha: am, Beta: bm, ID: "m"}
	n = goFRAME.GateAB{Alpha: an, Beta: bn, ID: "n"}
)

func ah(s goFRAME.State) float64 {
	Vm := s.V
	return 0.07 * math.Exp(-Vm/20)
}
func bh(s goFRAME.State) float64 {
	Vm := s.V
	return 1 / (math.Exp((30-Vm)/10) + 1)
}

func am(s goFRAME.State) float64 {
	Vm := s.V
	if Vm == 25 {
		return 1
	}
	return (25 - Vm) / (10 * (math.Exp((25-Vm)/10) - 1))
}
func bm(s goFRAME.State) float64 {
	Vm := s.V
	return 4 * math.Exp(-Vm/18)
}

func an(s goFRAME.State) float64 {
	Vm := s.V
	if Vm == 10 {
		return 0.1
	}
	return (10 - Vm) / (100 * (math.Exp((10-Vm)/10) - 1))
}
func bn(s goFRAME.State) float64 {
	Vm := s.V
	return 0.125 * math.Exp(-Vm/80)
}
