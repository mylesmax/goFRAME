package main

import (
	"goFRAME"
)

var (
	initial goFRAME.State
	out goFRAME.Out

	stims = goFRAME.Stims{
		goFRAME.Stim{
			Start:     1,
			End:       1.5,
			Intensity: 28,
		},
	}

	tset = goFRAME.TSet{
		Start: 0,
		Dt:    0.001,
		End:   20,
	}
)

func main() {
	out = []goFRAME.State{initial}

	solver := goFRAME.Solver{
		Out:         out,
		Tset:        tset,
		Stims:       stims,
		GatesAB:     []goFRAME.GateAB{m, n, h},
		GatesDirect: nil,
		Currents:    []goFRAME.Current{INa, Ik, ILeak},
		ToIntegrate: map[string]func(params ...interface{}) float64{
			"V": dvdt,
		},
		Method: "RK1",
	}

	out = solver.Solve()

	goFRAME.WriteExcel(out, "examples/hodgkinhuxley/out.xlsx")
}