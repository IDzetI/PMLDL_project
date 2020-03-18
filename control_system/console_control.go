package main

import (
	"./fins"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)


//Function to control position of endofunctor with speed control
func consolePositionControlWithSmoothMoving(robot *fins.Client) {

	//set default values
	speed := 300.0
	acceleration := 100.0
	deceleration := 100.0
	minSpeed := 5.0
	position := [3]float64{0, 0, 0}
	var cordFile [][3]float64
	cordFilePointer := 0

	//create reader for reed from console
	reader := bufio.NewReader(os.Stdin)

	//main loop
	for true {
		//read line
		line, err := reader.ReadString('\n')
		check(err)

		line = strings.Replace(line, "\n", "", -1)
		data := strings.Split(line, " ")

		switch data[0] {
		case "dnow":
			fmt.Println( "degree: ", readCurrentDegrees(robot))
		case "set":
			switch data[1] {
			case "cur":
				position = getPosition(data[2:5], position, "current position = ")
				break
			case "speed":
				speed = getValue(data[2], speed, "speed", 1, 500)
				break
			case "acc":
				acceleration = getValue(data[2], speed, "acceleration", 1, 1000)
				break
			case "dec":
				deceleration = getValue(data[2], deceleration, "deceleration", 1, 1000)
				break
			case "minspeed":
				minSpeed = getValue(data[2], minSpeed, "minimum speed", 0, 10)
				break
			default:
				break
			}
			break

		case "reset":
			//for reset
			if err := robot.WriteDNoResponse(420, PacToUint([4]float64{0, 0, 0, 0})); err != nil {
				fmt.Println(err.Error())
			}
			if err := robot.WriteDNoResponse(440, PacToUint([4]float64{0, 0, 0, 0})); err != nil {
				fmt.Println(err.Error())
			}
			if err := robot.WriteDNoResponse(460, PacToUint([4]float64{0, 0, 0, 0})); err != nil {
				fmt.Println(err.Error())
			}
			if err := robot.WriteDNoResponse(480, PacToUint([4]float64{0, 0, 0, 0})); err != nil {
				fmt.Println(err.Error())
			}
			if err := robot.WriteDNoResponse(500, PacToUint([4]float64{0, 0, 0, 0})); err != nil {
				fmt.Println(err.Error())
			}

			bytes, _ := robot.ReadD(420, 16)
			fmt.Println([4]float64{Uint16Float64(bytes[:4]), Uint16Float64(bytes[4:8]), Uint16Float64(bytes[8:12]), Uint16Float64(bytes[12:16])})

			bytes, _ = robot.ReadD(440, 16)
			fmt.Println([4]float64{Uint16Float64(bytes[:4]), Uint16Float64(bytes[4:8]), Uint16Float64(bytes[8:12]), Uint16Float64(bytes[12:16])})

			bytes, _ = robot.ReadD(460, 16)
			fmt.Println([4]float64{Uint16Float64(bytes[:4]), Uint16Float64(bytes[4:8]), Uint16Float64(bytes[8:12]), Uint16Float64(bytes[12:16])})

			bytes, _ = robot.ReadD(480, 16)
			fmt.Println([4]float64{Uint16Float64(bytes[:4]), Uint16Float64(bytes[4:8]), Uint16Float64(bytes[8:12]), Uint16Float64(bytes[12:16])})

			bytes, _ = robot.ReadD(500, 16)
			fmt.Println([4]float64{Uint16Float64(bytes[:4]), Uint16Float64(bytes[4:8]), Uint16Float64(bytes[8:12]), Uint16Float64(bytes[12:16])})

			position = [3]float64{0, 0, 0}
		case "init":
			position = movementInJointSpace(robot, getPosition(data[1:4], position, "move in joint space to position = "), speed, acceleration, deceleration, minSpeed)
		case "move":
			position = movementInCartesianSpace(robot, position, getPosition(data[1:4], position, "move in cartesian space to position = "), speed, acceleration, deceleration, minSpeed)
			break
		case "stop":
			stopMotors(robot)
			break
		case "start":
			turnONMotors(robot)
			break
		case "control":
			switch data[1] {
			case "on":
				ONControlMode(robot)
				break
			case "off":
				OFFControlMode(robot)
				break
			case "reset":
				OFFControlMode(robot)
				ONControlMode(robot)
			}
			break
		case "f":
			switch data[1] {
			case "parse":
				loadPltFile("trajectories/plt/"+data[2]+".plt", "testing/target_data/"+data[2]+".txt", [3]float64{0, 0, 0})
				break
			case "load":
				cordFile = loadCordFile("testing/target_data/" + data[2] + ".txt")
				break
			case "init":
				fmt.Println("moving to ", cordFile[cordFilePointer])
				position = movementInJointSpace(robot, cordFile[cordFilePointer], speed, acceleration, deceleration, minSpeed)
				break
			case "next":
				cordFilePointer = (cordFilePointer + 1) % len(cordFile)
				fmt.Println("moving to ", cordFile[cordFilePointer])
				position = movementInCartesianSpace(robot, position, cordFile[cordFilePointer], speed, acceleration, deceleration, minSpeed)
				break
			case "cur":
				fmt.Println("moving to ", cordFile[cordFilePointer])
				position = movementInCartesianSpace(robot, position, cordFile[cordFilePointer], speed, acceleration, deceleration, minSpeed)
				break
			case "prev":
				cordFilePointer = (cordFilePointer + (len(cordFile) - 1)) % len(cordFile)
				fmt.Println("moving to ", cordFile[cordFilePointer])
				position = movementInCartesianSpace(robot, position, cordFile[cordFilePointer], speed, acceleration, deceleration, minSpeed)
				break
			case "set":
				cordFilePointer = int(getValue(data[2], float64(cordFilePointer), "current point in coordinate file", 0, math.MaxInt64))
				break
			case "go":
				for i := cordFilePointer; i < len(cordFile); i++ {
					cordFilePointer = (cordFilePointer + 1) % len(cordFile)
					fmt.Println("moving to ", cordFilePointer, cordFile[cordFilePointer])
					position = movementInCartesianSpace(robot, position, cordFile[cordFilePointer], speed, acceleration, deceleration, minSpeed)
					readCurrentDegrees(robot)
					lineGo, err := reader.ReadString('\n')
					check(err)
					lineGo = strings.Replace(lineGo, "\n", "", -1)
					if lineGo == "p" {
						cordFilePointer = (cordFilePointer + (len(cordFile) - 2)) % len(cordFile)
						i -= 1
					}
					fmt.Println(cordFilePointer, i)
				}
				break
			case "run":
				for i := cordFilePointer; i < len(cordFile)-1; i++ {
					cordFilePointer = (cordFilePointer + 1) % len(cordFile)
					fmt.Println("moving to ", cordFilePointer, cordFile[cordFilePointer])
					position = movementInCartesianSpace(robot, position, cordFile[cordFilePointer], speed, acceleration, deceleration, minSpeed)
					readCurrentDegrees(robot)
					if readError(robot) {
						waitEnter()
					}
				}
				break
			}
			break
		case "exit":
			return
		default:
			break
		}
	}
}
