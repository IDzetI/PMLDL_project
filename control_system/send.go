package main

import (
	"./fins"
	"log"
)


/*
Functions send package to robot
 */
func sendPackage(robot *fins.Client, pac []uint16)  {
	err := robot.WriteDNoResponse(420,pac)
	if err != nil {
		log.Fatal(err)
	}
}

/*
Function convert angels package to robot package
 */
func PacToUint(pac [4]float64) []uint16 {
	res := make([]uint16, 4*len(pac))

	for i := 0; i < len(pac); i++ {
		current := Float64Uint16(pac[i])
		for j := 0; j < len(current); j++ {
			res[4*i + j] = current[j]
		}
	}
	return res
}
