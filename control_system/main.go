package main

import (
	"./fins"
)

const (
	plcAddr = "192.168.250.1:9600"
)

func main() {

	//define robot
	robot := fins.NewClient(plcAddr)
	defer robot.CloseConnection()

	consolePositionControlWithSmoothMoving(robot)
}
