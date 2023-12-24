package goFRAME

// GateAB : if a gate is defined by alpha and beta values
type GateAB struct {
	Alpha func(State) float64
	Beta  func(State) float64
	ID string
}

func (g GateAB) Ss(s State) float64 {
	return g.Alpha(s) / (g.Alpha(s) + g.Beta(s))
}

func (g GateAB) Tau(s State) float64 {
	return 1 / (g.Alpha(s) + g.Beta(s))
}

// GateDirect : if a gate is defined by an analytical solution
type GateDirect struct {
	Ss func(State) float64
	ID string
}