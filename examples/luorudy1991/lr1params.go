package main

import (
	"goFRAME"
	"math"
)

func init() {
	initial = goFRAME.State{
		V:     -85,
		Cm:    1,
		Stim:  0,
		T:     0,
		Index: 0,
		RTF:   (8.3145) * (37 + 273.15) / (96.485),
		GBar: map[string]float64{
			"Na": 23,
			"Kp": 0.0183,
			"b":  0.03921,
		},
		I: map[string]float64{
			INa.ID: 0,
			Isi.ID: 0,
			Ik.ID:  0,
			IK1.ID: 0,
			IKp.ID: 0,
			Ib.ID:  0},
		ConcOut: map[string]float64{
			"K":  5.4,
			"Na": 140,
			"Ca": 1.8,
		},
		ConcIn: map[string]float64{
			"K":  145,
			"Na": 18,
			"Ca": 2 * math.Pow10(-4),
		},
		Misc: map[string]float64{
			"PRNaK": 0.01833,
		},
	}
	initial.GBar["K"] = 0.282 * math.Sqrt(initial.ConcOut["K"]/5.4)
	initial.GBar["K1"] = 0.6047 * math.Sqrt(initial.ConcOut["K"]/5.4)

	initial.E = map[string]float64{
		"Na": initial.RTF * math.Log(initial.ConcOut["Na"]/initial.ConcIn["Na"]),
		"K":  initial.RTF * math.Log((initial.ConcOut["K"]+initial.Misc["PRNaK"]*initial.ConcOut["Na"])/(initial.ConcIn["K"]+initial.Misc["PRNaK"]*initial.ConcIn["Na"])),
		"K1": initial.RTF * math.Log(initial.ConcOut["K"]/initial.ConcIn["K"]),
		"Kp": initial.RTF * math.Log(initial.ConcOut["K"]/initial.ConcIn["K"]),
		"b":  -59.87,
	}
	initial.Gate = map[string]float64{
		Xi.ID: Xi.Ss(initial),
		X.ID:  X.Ss(initial),
		K1.ID: K1.Ss(initial),
		Kp.ID: Kp.Ss(initial),
		m.ID:  m.Ss(initial),
		h.ID:  h.Ss(initial),
		j.ID:  j.Ss(initial),
		d.ID:  d.Ss(initial),
		f.ID:  f.Ss(initial),
	}

}
