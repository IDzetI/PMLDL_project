package main

import (
	"fmt"
	"strings"
)

func loadCordFile(fileName string)[][3]float64 {
	var result [][3]float64
	data := strings.Split(ReadAll(fileName),"\n")
	for i:=0; i < len(data); i++ {
		point, err := stringToCoordinates(data[i])
		if err || !checkCoordinate(point){
			fmt.Println("ERROR with " + data[i])
			continue
		}
		result = append(result,point)
		fmt.Println(point)
	}
	fmt.Println("Loading file complete")
	return result
}
