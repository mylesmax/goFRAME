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
	Isi = goFRAME.Current{
		Current: IsiFunc,
		ID:      "si",
	}
	Ik = goFRAME.Current{
		Current: IkFunc,
		ID:      "K",
	}
	IK1 = goFRAME.Current{
		Current: Ik1Func,
		ID:      "K1",
	}
	IKp = goFRAME.Current{
		Current: IkpFunc,
		ID:      "Kp",
	}
	Ib = goFRAME.Current{
		Current: IbFunc,
		ID:      "b",
	}
)

func INaFunc(params ...interface{}) float64 {
	s, ok := params[0].(goFRAME.State)
	if !ok {
		fmt.Println("error in INa", ok)
		return 0
	}

	return s.GBar["Na"] * math.Pow(s.Gate["m"], 3) * s.Gate["h"] * s.Gate["j"] * (s.V - s.E["Na"])
}

func IsiFunc(params ...interface{}) float64 {
	s, ok := params[0].(goFRAME.State)
	if !ok {
		fmt.Println("error in Isi", ok)
		return 0
	}
	Esi := 7.7 - 13.0287*math.Log(s.ConcIn["Ca"]/s.ConcOut["Ca"]) //cur.out.Cai[cur.i]/c.Cao)

	return 0.09 * s.Gate["d"] * s.Gate["f"] * (s.V - Esi)
}

func IkFunc(params ...interface{}) float64 {
	s, ok := params[0].(goFRAME.State)
	if !ok {
		fmt.Println("error in Isi", ok)
		return 0
	}

	return s.GBar["K"] * s.Gate["X"] * s.Gate["Xi"] * (s.V - s.E["K"])
}

func Ik1Func(params ...interface{}) float64 {
	s, ok := params[0].(goFRAME.State)
	if !ok {
		fmt.Println("error in Isi", ok)
		return 0
	}

	return s.GBar["K1"] * s.Gate["K1"] * (s.V - s.E["K1"])
}

func IkpFunc(params ...interface{}) float64 {
	s, ok := params[0].(goFRAME.State)
	if !ok {
		fmt.Println("error in Isi", ok)
		return 0
	}

	return s.GBar["Kp"] * s.Gate["Kp"] * (s.V - s.E["Kp"])
}

func IbFunc(params ...interface{}) float64 {
	s, ok := params[0].(goFRAME.State)
	if !ok {
		fmt.Println("error in Isi", ok)
		return 0
	}

	return s.GBar["b"] * (s.V - s.E["b"])
}
