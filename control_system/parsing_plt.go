package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadPltFile(inputFileName string, outputFileName string, currentPos [3]float64) {
	zUP := 100.0
	zZERO := 0.0
	zDOWN := -50.0

	fmt.Println("load file from" + inputFileName)
	pltData := strings.Split(ReadAll(inputFileName), "\n")
	f, _ := os.Create(outputFileName)
	defer f.Close()
	if currentPos[2] <= zZERO {
		currentPos = [3]float64{currentPos[0], currentPos[1], zUP}
		writeCord(f, currentPos)
	}

	for i := 0; i < len(pltData); i++ {
		switch pltData[i][:2] {
		case "PU":
			x, y := pltLineToXY(pltData[i][2:])
			if currentPos[2] != zUP {
				writeCord(f, [3]float64{currentPos[0], currentPos[1], zUP})
			}
			currentPos = [3]float64{x, y, zUP}
			writeCord(f, currentPos)
			break
		case "PD":
			x, y := pltLineToXY(pltData[i][2:])
			if currentPos[2] != zDOWN {
				writeCord(f, [3]float64{currentPos[0], currentPos[1], zDOWN})
			}
			currentPos = [3]float64{x, y, zDOWN}
			writeCord(f, currentPos)
			break
		}
	}
	fmt.Println("Parsing file complete")
	if currentPos[2] < zZERO {
		writeCord(f, [3]float64{currentPos[0], currentPos[1], zUP})
	}
	_, _ = f.WriteString("0 0 1000")
}

func pltLineToXY(line string) (float64, float64) {
	xy := strings.Split(line, " ")
	x, _ := strconv.ParseFloat(xy[0], 64)
	y, _ := strconv.ParseFloat(xy[1][:len(xy[1])-2], 64)
	return x, y
}

func writeCord(f *os.File, cord [3]float64) {
	prec := 0
	line := ""
	line += strconv.FormatFloat(cord[0], 'f', prec, 64) + " "
	line += strconv.FormatFloat(cord[1], 'f', prec, 64) + " "
	line += strconv.FormatFloat(cord[2], 'f', prec, 64) + "\n"

	_, _ = f.WriteString(line)
}
