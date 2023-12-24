package goFRAME

import (
	"fmt"
	"github.com/xuri/excelize"
	"os"
	"strconv"
	"time"
)

func WriteExcel(O Out, filePath string) error {
	f := excelize.NewFile()

	sheetName := time.Now().Local().Format("2006-01-02 03-04-05 PM")

	index, err := f.NewSheet(sheetName)
	if err != nil {
		return err
	}

	f.SetActiveSheet(index)
	err = f.DeleteSheet("Sheet1")
	if err != nil {
		return err
	}
	currentColumnIndex := 1
	headerIndexMap := make(map[string]int)

	writeHeader := func(header string) {
		cell, _ := excelize.CoordinatesToCellName(currentColumnIndex, 1)
		err := f.SetCellValue(sheetName, cell, header)
		if err != nil {
			return
		}
		headerIndexMap[header] = currentColumnIndex
		currentColumnIndex++
	}

	writeHeader("Index")
	writeHeader("t")
	writeHeader("V")
	writeHeader("stim")
	writeHeader("Cm")
	writeHeader("RTF")
	for _, state := range O {
		for key := range state.E {
			writeHeader("E" + key)
		}
		for key := range state.GBar {
			writeHeader("GBar" + key)
		}
		for key := range state.I {
			writeHeader("I" + key)
		}
		for key := range state.Gate {
			writeHeader("Gate" + key)
		}
		for key := range state.ConcOut {
			writeHeader("ConcOut" + key)
		}
		for key := range state.ConcIn {
			writeHeader("ConcIn" + key)
		}
		for key := range state.Misc {
			writeHeader("Misc" + key)
		}
		break
	}

	for rowIndex, state := range O {
		rowNumber := rowIndex + 2
		err := f.SetCellValue(sheetName, "A"+strconv.Itoa(rowNumber), state.Index)
		if err != nil {
			return err
		}
		err = f.SetCellValue(sheetName, "B"+strconv.Itoa(rowNumber), state.T)
		if err != nil {
			return err
		}
		err = f.SetCellValue(sheetName, "C"+strconv.Itoa(rowNumber), state.V)
		if err != nil {
			return err
		}
		err = f.SetCellValue(sheetName, "D"+strconv.Itoa(rowNumber), state.Stim)
		if err != nil {
			return err
		}
		err = f.SetCellValue(sheetName, "E"+strconv.Itoa(rowNumber), state.Cm)
		if err != nil {
			return err
		}
		err = f.SetCellValue(sheetName, "F"+strconv.Itoa(rowNumber), state.RTF)
		if err != nil {
			return err
		}

		writeMapData := func(m map[string]float64, prefix string) {
			for key, value := range m {
				colIndex := headerIndexMap[prefix+key]
				cell, _ := excelize.CoordinatesToCellName(colIndex, rowNumber)
				err := f.SetCellValue(sheetName, cell, value)
				if err != nil {
					return
				}
			}
		}

		writeMapData(state.E, "E")
		writeMapData(state.GBar, "GBar")
		writeMapData(state.I, "I")
		writeMapData(state.Gate, "Gate")
		writeMapData(state.ConcOut, "ConcOut")
		writeMapData(state.ConcIn, "ConcIn")
		writeMapData(state.Misc, "Misc")
	}

	if _, err := os.Stat(filePath); err == nil {
		if removeErr := os.Remove(filePath); removeErr != nil {
			fmt.Printf("Unable to remove existing file: %s\n", removeErr)
			return removeErr
		}
	}

	if err := f.SaveAs(filePath); err != nil {
		fmt.Printf("Error saving file: %s\n", err)
		return err
	}

	return nil
}
