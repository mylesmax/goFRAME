package main

import (
	"fmt"
	"goFRAME"
	"reflect"
	"runtime"
)

var (
	ITemplate = goFRAME.Current{
		Current: ITemplateFunc,
		ID:      "Template",
	}
)

func ITemplateFunc(params ...interface{}) float64 {
	//Uncomment if you need current state
	s, ok := params[0].(goFRAME.State)
	if !ok {
		fmt.Println("error in s of", runtime.FuncForPC(reflect.ValueOf(ITemplateFunc).Pointer()).Name())
		return 0
	}

	//Uncomment if you need all saved states
	//out, ok := params[1].(goFRAME.Out)
	//if !ok {
	//	fmt.Println("error in out of", runtime.FuncForPC(reflect.ValueOf(ITemplateFunc).Pointer()).Name())
	//	return 0
	//}

	//Uncomment if you need index
	//i, ok := params[2].(int)
	//if !ok {
	//	fmt.Println("error in i of", runtime.FuncForPC(reflect.ValueOf(ITemplateFunc).Pointer()).Name())
	//	return 0
	//}

	//Getting a gate: s.Gate["__"] where you enter the gate ID as a string
	//Getting a conductance; s.GBar["__"] where you enter the conductance ID as a string
	//Getting a nernst potential: s.E["__"] where you enter the Nernst potential ID as a string
	//Getting a Conc. In or Conc. Out: s.ConcIn["__"] or s.ConcOut["__"] where you enter the ion ID as a string
	//Getting voltage: s.V

	return s.V //return current here
}
