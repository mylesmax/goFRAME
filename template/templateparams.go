package main

import (
	"goFRAME"
	"math"
)

func init() {
	initial = goFRAME.State{
		V:     0, //This is Vrest
		Cm:    0, //Membrane capacitance
		Stim:  0, //Initial stimulus
		T:     0, //Initial time value
		Index: 0, //Initial index (set to 0)
		RTF:   (8.3145) * (37 + 273.15) / (96.485),
		GBar: map[string]float64{
			"Template": 0, //GBar["Temp"] = 0, "Temp" Conductance
		},
		I: map[string]float64{
			ITemplate.ID: 0, //Set initial currents by mapping ID: value (Normally currents start at 0)
		},
		ConcOut: map[string]float64{
			"Template": 0, //Set initial extracellular ionic concentration
		},
		ConcIn: map[string]float64{
			"Template": 0, //Set initial intracellular ionic concentration
		},
		Misc: map[string]float64{
			"Template": 0, //Other parameters, such as permeability constants
		},
	}

	//For all initial variables that have a dependence on other functions (i.e., not a simple number), call them
	//outside the initial range. Be careful to not overwrite the already-defined variables above.
	initial.GBar["Template2"] = initial.ConcOut["Template"] / initial.ConcIn["Template"]

	initial.E = map[string]float64{
		"Template": initial.RTF * math.Log(initial.ConcOut["Template"]/initial.ConcIn["Template"]),
	}

	//Now we will calculate the initial value of the gates
	//Call the SS (steady state) modifier to compute
	initial.Gate = map[string]float64{
		TemplateDirect.ID: TemplateDirect.Ss(initial),
		TemplateAB.ID:     TemplateAB.Ss(initial),
	}

}
