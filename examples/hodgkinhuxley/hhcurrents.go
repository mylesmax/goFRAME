package main

import (
	"fmt"
	"goFRAME"
	"math"
)

var (
	INa = goFRAME.Current{
		Current: INaFunc,
		ID:      "Na",
	}
	Ik = goFRAME.Current{
		Current: IkFunc,
		ID:      "K",
	}
	ILeak = goFRAME.Current{
		Current: ILeakFunc,
		ID:      "Leak",
	}
)

func INaFunc(params ...interface{}) float64 {
	s, ok := params[0].(goFRAME.State)
	if !ok {
		fmt.Println("error in INa", ok)
		return 0
	}
	return s.GBar["Na"] * math.Pow(s.Gate["m"], 3) * s.Gate["h"] * (s.V - s.E["Na"])
}

func IkFunc(params ...interface{}) float64 {
	s, ok := params[0].(goFRAME.State)
	if !ok {
		fmt.Println("error in Ik", ok)
		return 0
	}
	return s.GBar["K"] * math.Pow(s.Gate["n"], 4) * (s.V - s.E["K"])
}

func ILeakFunc(params ...interface{}) float64 {
	out, ok := params[1].(goFRAME.Out)
	if !ok {
		fmt.Println("error in Out of ILeak", ok)
		return 0
	}

	i, ok := params[2].(int)
	if !ok {
		fmt.Println("error in i of ILeak", ok)
		return 0
	}
	return out[i].GBar["Leak"] * (out[i].V - out[0].V)
}
