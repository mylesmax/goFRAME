package main

import "goFRAME"

func init() {
	initial = goFRAME.State{
		V:     0,
		Cm:    1,
		Stim:  0,
		T:     0,
		Index: 0,
		E: map[string]float64{
			"Na": 115,
			"K":  -12,
		},
		GBar: map[string]float64{
			"Na":   70.7,
			"K":    24,
			"Leak": 0.3,
		},
		I: map[string]float64{
			INa.ID:   0,
			Ik.ID:    0,
			ILeak.ID: 0}}

	initial.Gate = map[string]float64{
		m.ID: m.Ss(initial),
		n.ID: n.Ss(initial),
		h.ID: h.Ss(initial),
	}
}
