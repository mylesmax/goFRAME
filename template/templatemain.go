package main

import (
	"fmt"
	"goFRAME"
)

var (
	initial goFRAME.State
	out     = goFRAME.Out{initial}

	stims = goFRAME.Stims{
		goFRAME.Stim{
			Start:     30,   //Start time of stimulus
			End:       30.5, //End time of stimulus
			Intensity: 60,   //Intensity of stimulus (units depend on model)
		},
	}

	tset = goFRAME.TSet{
		Start: 0,    //Start time of tspan
		Dt:    0.24, //Time step (choose wisely depending on numerical method)
		End:   1000, //End time of tspan
	}
)

func main() {
	out = []goFRAME.State{initial} //Set initial conditions as the first State in Out

	solver := goFRAME.Solver{
		Out:         out,
		Tset:        tset,
		Stims:       stims,
		GatesAB:     []goFRAME.GateAB{TemplateAB},         //Add all AB gates here
		GatesDirect: []goFRAME.GateDirect{TemplateDirect}, //Add all Direct gates here
		Currents:    []goFRAME.Current{ITemplate},         //Add all currents here
		ToIntegrate: map[string]func(params ...interface{}) float64{
			"V": dvdt, //integrate V with the function dvdt

			//If you need to integrate a variable that's within the State, look at this example below.
			//Below, we have inward calcium being dynamically modified by dCaidt
			//To state the variable itself, define it with similar State notation. So if s is State, then
			//s.ConcIn["Ca"] should be written as "s.ConcIn[\"Ca\"]". In practice this is:

			//"ConcIn[\"Ca\"]": dCaidt,
		},
		Method: "RK1", //See README on available methods
	}

	out = solver.Solve()

	err := goFRAME.WriteExcel(out, "PATH/out.xlsx") //Replace with path
	if err != nil {
		fmt.Println("Error writing excel", err)
		return
	}
}
