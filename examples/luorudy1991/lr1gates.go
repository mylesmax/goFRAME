package main

import (
	"goFRAME"
	"math"
)

var (
	Xi = goFRAME.GateDirect{
		Ss: XiFunc,
		ID: "Xi",
	}
	X = goFRAME.GateAB{
		Alpha: aX,
		Beta:  bX,
		ID:    "X",
	}

	K1 = goFRAME.GateAB{
		Alpha: aK1,
		Beta:  bK1,
		ID:    "K1",
	}

	Kp = goFRAME.GateDirect{
		Ss: KpFunc,
		ID: "Kp",
	}

	m = goFRAME.GateAB{
		Alpha: am,
		Beta:  bm,
		ID:    "m",
	}

	h = goFRAME.GateAB{
		Alpha: ah,
		Beta:  bh,
		ID:    "h",
	}

	j = goFRAME.GateAB{
		Alpha: aj,
		Beta:  bj,
		ID:    "j",
	}

	d = goFRAME.GateAB{
		Alpha: ad,
		Beta:  bd,
		ID:    "d",
	}

	f = goFRAME.GateAB{
		Alpha: af,
		Beta:  bf,
		ID:    "f",
	}
)

func XiFunc(s goFRAME.State) float64 {
	Vm := s.V
	if Vm > -100 {
		return 2.837 * (math.Exp(0.04*(Vm+77)) - 1) / ((Vm + 77) * (math.Exp(0.04 * (Vm + 35))))
	} else {
		return 1
	}
}

func aX(s goFRAME.State) float64 {
	Vm := s.V
	return 0.0005 * math.Exp(0.083*(Vm+50)) / (1 + math.Exp(0.057*(Vm+50)))
}
func bX(s goFRAME.State) float64 {
	Vm := s.V
	return 0.0013 * math.Exp(-0.06*(Vm+20)) / (1 + math.Exp(-0.04*(Vm+20)))
}

func aK1(s goFRAME.State) float64 {
	Vm := s.V
	return 1.02 / (1 + math.Exp(0.2385*(Vm-s.E["K1"]-59.215)))
}
func bK1(s goFRAME.State) float64 {
	Vm := s.V
	return (0.49124*math.Exp(0.08032*(Vm-s.E["K1"]+5.476)) + math.Exp(0.06175*(Vm-s.E["K1"]-594.31))) / (1 + math.Exp(-0.5143*(Vm-s.E["K1"]+4.753)))
}

func KpFunc(s goFRAME.State) float64 {
	Vm := s.V
	return 1 / (1 + math.Exp((7.488-Vm)/5.98))
}

func am(s goFRAME.State) float64 {
	Vm := s.V
	return 0.32 * (Vm + 47.13) / (1 - math.Exp(-0.1*(Vm+47.13)))
}
func bm(s goFRAME.State) float64 {
	Vm := s.V
	return 0.08 * math.Exp(-Vm/11)
}

func ah(s goFRAME.State) float64 {
	Vm := s.V
	if Vm >= -40 {
		return 0
	} else {
		return 0.135 * math.Exp((80+Vm)/-6.8)
	}
}
func bh(s goFRAME.State) float64 {
	Vm := s.V
	if Vm >= -40 {
		return 1 / (0.13 * (1 + math.Exp((Vm+10.66)/-11.1)))
	} else {
		return 3.56*math.Exp(0.079*Vm) + 3.1*math.Pow10(5)*math.Exp(0.35*Vm)
	}
}

func aj(s goFRAME.State) float64 {
	Vm := s.V
	if Vm >= -40 {
		return 0
	} else {
		return (-1.2714*math.Pow10(5)*math.Exp(0.2444*Vm) - 3.474*math.Pow10(-5)*math.Exp(-0.04391*Vm)) * (Vm + 37.78) / (1 + math.Exp(0.311*(Vm+79.23)))
	}
}
func bj(s goFRAME.State) float64 {
	Vm := s.V
	if Vm >= -40 {
		return 0.3 * math.Exp(-2.535*math.Pow10(-7)*Vm) / (1 + math.Exp(-0.1*(Vm+32)))
	} else {
		return 0.1212 * math.Exp(-0.01052*Vm) / (1 + math.Exp(-0.1378*(Vm+40.14)))
	}
}

func ad(s goFRAME.State) float64 {
	Vm := s.V
	return 0.095 * math.Exp(-0.01*(Vm-5)) / (1 + math.Exp(-0.072*(Vm-5)))
}
func bd(s goFRAME.State) float64 {
	Vm := s.V
	return 0.07 * math.Exp(-0.017*(Vm+44)) / (1 + math.Exp(0.05*(Vm+44)))
}

func af(s goFRAME.State) float64 {
	Vm := s.V
	return 0.012 * math.Exp(-0.008*(Vm+28)) / (1 + math.Exp(0.15*(Vm+28)))
}
func bf(s goFRAME.State) float64 {
	Vm := s.V
	return 0.0065 * math.Exp(-0.02*(Vm+30)) / (1 + math.Exp(-0.2*(Vm+30)))
}
