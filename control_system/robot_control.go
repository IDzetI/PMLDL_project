package main

import (
	"./fins"
	"fmt"
)

func ONControlMode(robot *fins.Client) {
	fmt.Println("Try to turn ON the PC control mode")
	fmt.Println(robot.WriteD(403, []uint16{1}))
	fmt.Println("PC control mode is ON")
}

func OFFControlMode(robot *fins.Client) {
	fmt.Println("Try to turn OFF the PC control mode")
	fmt.Println(robot.WriteD(403, []uint16{0}))
	fmt.Println("PC control mode is OFF")
}

func stopMotors(robot *fins.Client) {
	_ = robot.WriteDNoResponse(400, []uint16{1})
	fmt.Println("Stop motors")
}

func turnONMotors(robot *fins.Client) {
	_ = robot.WriteDNoResponse(400, []uint16{0})
	fmt.Println("Turn ON motors")
}

func readCurrentDegrees(robot *fins.Client)[4]float64  {
	bytes, _ := robot.ReadD(480,16)
	degrees := [4]float64{Uint16Float64(bytes[:4]),Uint16Float64(bytes[4:8]),Uint16Float64(bytes[8:12]),Uint16Float64(bytes[12:16])}
	return degrees
}

func readError(robot *fins.Client) bool  {
	isError, _ := robot.ReadD(408,1)
	fmt.Print("Has error: ")
	fmt.Println(isError[0] != 0)
	return isError[0] != 0
}
