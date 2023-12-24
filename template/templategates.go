package main

import (
	"goFRAME"
)

var (
	TemplateDirect = goFRAME.GateDirect{
		Ss: DirectFunc,
		ID: "TemplateDirect",
	}
	TemplateAB = goFRAME.GateAB{
		Alpha: aTemp,
		Beta:  bTemp,
		ID:    "TemplateAB",
	}
)

func DirectFunc(s goFRAME.State) float64 {
	Vm := s.V

	return Vm //Gate function
}

func aTemp(s goFRAME.State) float64 {
	Vm := s.V

	return Vm //Alpha function
}
func bTemp(s goFRAME.State) float64 {
	Vm := s.V

	return Vm //Beta function
}
