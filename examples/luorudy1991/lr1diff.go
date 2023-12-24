package main

import (
	"fmt"
	"goFRAME"
	"math"
)

func dvdt(params ...interface{}) float64 {
	solver, ok := params[0].(goFRAME.Solver)
	if !ok {
		fmt.Println("error in dvdt solver")
		return 0
	}
	i, ok := params[1].(int)
	if !ok {
		fmt.Println("error in dvdt i")
		return 0
	}
	curSum, ok := params[2].(float64)
	if !ok {
		fmt.Println("error in dvdt i")
		return 0
	}

	return -(1 / solver.Out[i].Cm) * (curSum)
}

func dCaidt(params ...interface{}) float64 {
	solver, ok := params[0].(goFRAME.Solver)
	if !ok {
		fmt.Println("error in dvdt solver")
		return 0
	}
	i, ok := params[1].(int)
	if !ok {
		fmt.Println("error in dvdt i")
		return 0
	}

	return -math.Pow10(-4)*solver.Out[i].I["si"] + 0.07*(math.Pow10(-4)-solver.Out[i].ConcIn["Ca"])
}
