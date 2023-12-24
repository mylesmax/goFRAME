package main

import (
	"fmt"
	"goFRAME"
	"reflect"
	"runtime"
)

//classic dvdt equation
func dvdt(params ...interface{}) float64 {
	solver, ok := params[0].(goFRAME.Solver)
	if !ok {
		fmt.Println("error in solver of", runtime.FuncForPC(reflect.ValueOf(ITemplateFunc).Pointer()).Name())
		return 0
	}
	i, ok := params[1].(int)
	if !ok {
		fmt.Println("error in i of", runtime.FuncForPC(reflect.ValueOf(ITemplateFunc).Pointer()).Name())
		return 0
	}
	curSum, ok := params[2].(float64)
	if !ok {
		fmt.Println("error in curSum of", runtime.FuncForPC(reflect.ValueOf(ITemplateFunc).Pointer()).Name())
		return 0
	}

	return -(1 / solver.Out[i].Cm) * (curSum)
}
