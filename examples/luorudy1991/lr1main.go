package main

import (
	"goFRAME"
)

var (
	initial goFRAME.State
	out     = goFRAME.Out{initial}

	stims = goFRAME.Stims{
		goFRAME.Stim{
			Start:     30,
			End:       30.5,
			Intensity: 60,
		},
		goFRAME.Stim{
			Start:     450,
			End:       450.5,
			Intensity: 60,
		},
	}

	tset = goFRAME.TSet{
		Start: 0,
		Dt:    0.24,
		End:   1000,
	}
)

func main() {
	out = []goFRAME.State{initial}

	solver := goFRAME.Solver{
		Out:         out,
		Tset:        tset,
		Stims:       stims,
		GatesAB:     []goFRAME.GateAB{X, K1, m, h, j, d, f},
		GatesDirect: []goFRAME.GateDirect{Xi, Kp},
		Currents:    []goFRAME.Current{INa, Isi, Ik, IK1, IKp, Ib},
		ToIntegrate: map[string]func(params ...interface{}) float64{
			"V":              dvdt,
			"ConcIn[\"Ca\"]": dCaidt,
		},
		Method: "RK1",
	}

	out = solver.Solve()

	goFRAME.WriteExcel(out, "examples/luorudy1991/out.xlsx")
}
