package goFRAME

import (
	"fmt"
	"github.com/mohae/deepcopy"
	"reflect"
	"strings"
)

type Solver struct {
	Out         Out
	Tset        TSet
	Stims       Stims
	GatesAB     []GateAB
	GatesDirect []GateDirect
	Currents    []Current
	ToIntegrate map[string]func(params ...interface{}) float64
	Method      string
}

func (solver Solver) Solve() Out {
	var s2 State
	tspan := solver.Tset.GenerateTSpan()

	for i, t := range tspan {
		s2 = deepcopy.Copy(solver.Out[i]).(State)
		if i == 0 {
			solver.Out = append(solver.Out, s2)
			continue
		}

		//update index and time
		s2.Index = i
		s2.T = t

		//Update all AB gates
		if len(solver.GatesAB) != 0 {
			for _, gate := range solver.GatesAB {
				s2.Gate[gate.ID] = GetGateAB(s2, gate, solver.Tset)
			}
		}

		//update direct gates
		if len(solver.GatesDirect) != 0 {
			for _, gate := range solver.GatesDirect {
				s2.Gate[gate.ID] = GetGateDirect(s2, gate)
			}
		}

		//update all currents
		for _, cur := range solver.Currents {
			s2.I[cur.ID] = cur.Current(s2, solver.Out, i)
		}

		s2.Stim = GetStim(t, solver.Stims, solver.Tset)

		var curSum float64
		for _, cur := range s2.I {
			curSum += cur
		}
		curSum -= s2.Stim

		switch solver.Method {
		case "RK1":
			//Runge kutta 1 (euler)
			for x, dv := range solver.ToIntegrate {
				//here we check if it is something other than V
				if strings.Contains(x, "[") {
					parts := strings.Split(x, "[")
					section := parts[0]
					mapKey := strings.Trim(parts[1], "\"[]")

					mapField := reflect.ValueOf(&s2).Elem().FieldByName(section)
					if mapField.IsValid() && mapField.Kind() == reflect.Map {
						currentVal := mapField.MapIndex(reflect.ValueOf(mapKey)).Float()
						newVal := currentVal + dv(solver, i, curSum)*solver.Tset.Dt
						mapField.SetMapIndex(reflect.ValueOf(mapKey), reflect.ValueOf(newVal))
					}
				} else {
					field := reflect.ValueOf(&s2).Elem().FieldByName(x)
					if field.IsValid() && field.CanSet() {
						currentVal := field.Float()
						newVal := currentVal + dv(solver, i, curSum)*solver.Tset.Dt
						field.SetFloat(newVal)
					}
				}
			}
		//TODO case "RK2":
		default:
			fmt.Println("ERROR! Incorrect integration method. Please see documentation for usable methods.")
			return solver.Out
		}

		//update out
		solver.Out = append(solver.Out, s2)
		solver.Out[i] = s2
	}

	return solver.Out
}
