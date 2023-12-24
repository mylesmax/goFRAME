package goFRAME

type TSet struct {
	Start, Dt, End float64
}

type State struct {
	V, Cm, RTF, Stim, T float64
	Index                int
	//Nernst Potentials
	E map[string]float64

	//conductances
	GBar map[string]float64

	//currents
	I map[string]float64

	//gates
	Gate map[string]float64

	//ion concentrations
	ConcOut map[string]float64
	ConcIn map[string]float64

	//misc
	Misc map[string]float64
}

type Out []State

type Stim struct{
	Start, End, Intensity float64
}

type Stims []Stim

//GenerateTSpan : generate an array of the stiff time span
func (tsp TSet) GenerateTSpan() (tspan []float64) {
	for t := tsp.Start; t <= tsp.End; t += tsp.Dt {
		tspan = append(tspan, t)
	}

	return tspan
}