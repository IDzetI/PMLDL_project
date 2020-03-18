package main

import (
	"./fins"
	"fmt"
	"math"
	"os"
)

/*
current position [x y z] in mm
new position [x y z] in mm
speed in mm/s
acceleration in mm/s^2
deceleration in mm/s^2
*/
func movementInCartesianSpace(robot *fins.Client, currentPoint [3]float64, newPoint [3]float64,
	speed float64, acceleration float64, deceleration float64, minSpeed float64) [3]float64 {

	//check input value
	if !checkCoordinate(newPoint) {
		return currentPoint
	}

	//init values
	var points [][3]float64

	//controller period
	period := 0.004

	//save next
	file, err := os.Create("next.txt")
	check(err)
	_, _ = fmt.Fprintln(file, formatstring3(newPoint))
	file.Close()

	//remaining distance
	S := distanceBeatenPoints(currentPoint, newPoint)

	// set speeds, acceleration and deceleration
	acc := [3]float64{}
	dec := [3]float64{}
	speedProj := [3]float64{}
	minSpeedProj := [3]float64{}
	for i := 0; i < 3; i++ {
		k := (newPoint[i] - currentPoint[i]) / S
		acc[i] = acceleration * k
		dec[i] = -deceleration * k
		speedProj[i] = speed * k
		minSpeedProj[i] = minSpeed * k
	}

	//init helped variables
	cSpeed := [3]float64{}
	stage := 0

	for S > 1e-3 {
		if stage == 0 {
			cSpeed = sumOfVectors(cSpeed, mulVectorByScalar(acc, period))
		}

		moduleOfCurSpeed := lengthOfVector(cSpeed)

		if moduleOfCurSpeed > speed {
			cSpeed = speedProj
			stage = 1
		}

		if S <= moduleOfCurSpeed*moduleOfCurSpeed/2/deceleration {
			stage = 2
		}
		if stage == 2 {
			cSpeed = sumOfVectors(cSpeed, mulVectorByScalar(dec, period))
		}

		if lengthOfVector(cSpeed) < minSpeed && stage == 2 {
			cSpeed = minSpeedProj
			stage = 3
		}

		currentPoint = sumOfVectors(currentPoint, mulVectorByScalar(cSpeed, period))
		points = append(points, currentPoint)

		if distanceBeatenPoints(currentPoint, newPoint) > S {
			break
		}
		S = distanceBeatenPoints(currentPoint, newPoint)
	}

	// last point
	points = append(points, newPoint)

	//generate sets of angles
	pac := make([][]uint16, len(points))
	for i := 0; i < len(points); i++ {
		pac[i] = PacToUint(xyzToDegrees(points[i]))
	}

	//send angles package
	sendTrajectory(robot, pac)
	return newPoint
}

/*
current position [x y z] in mm
new position [x y z] in mm
speed in mm/s
acceleration in mm/s^2
deceleration in mm/s^2
*/
func movementInJointSpace(robot *fins.Client, newPoint [3]float64,
	speed float64, acceleration float64, deceleration float64, minSpeed float64) [3]float64 {

	//controller period
	period := 0.004

	currentDegrees := readCurrentDegrees(robot)
	newDegrees := xyzToDegrees(newPoint)

	cableDegreeDistances := difOfVectors(newDegrees, currentDegrees)

	maxCableDist := math.Max(
		math.Max(math.Abs(cableDegreeDistances[0]), math.Abs(cableDegreeDistances[1])),
		math.Max(math.Abs(cableDegreeDistances[2]), math.Abs(cableDegreeDistances[3])))

	cableMinSpeed := [4]float64{}
	cableSpeed := [4]float64{}
	cableAcc := [4]float64{}
	cableDec := [4]float64{}

	for i := range cableDegreeDistances {
		k := cableDegreeDistances[i] / maxCableDist
		cableMinSpeed[i] = k * minSpeed
		cableSpeed[i] = k * speed
		cableAcc[i] = k * acceleration
		cableDec[i] = k * deceleration
	}

	S := distanceBeatenPoints4(newDegrees, currentDegrees)

	/*
		fmt.Println("degree now:",formatstring4(currentDegrees))
		fmt.Println("degree new:",formatstring4(newDegrees))
		fmt.Println("degree dist:",formatstring4(cableDegreeDistances))
		fmt.Println("   max dist:\t",maxCableDist)
		fmt.Println("min speed: ",formatstring4(cableMinSpeed))
		fmt.Println("    speed: ",formatstring4(cableSpeed))
		fmt.Println("      acc: ",formatstring4(cableMinSpeed))
		fmt.Println("      dec: ",formatstring4(cableMinSpeed))
	*/

	//init values
	var degrees [][4]float64
	cSpeed := [4]float64{}
	stage := 0

	for S > 1 {
		if stage == 0 {
			cSpeed = sumOfVectors4(cSpeed, mulVectorByScalar4(cableAcc, period))
		}

		moduleOfCurSpeed := lengthOfVector4(cSpeed)

		if moduleOfCurSpeed > speed {
			//cSpeed = cableSpeed
			stage = 1
		}

		if S+222 < moduleOfCurSpeed*moduleOfCurSpeed/2/deceleration {
			stage = 2
		}
		if stage == 2 {
			cSpeed = difOfVectors(cSpeed, mulVectorByScalar4(cableDec, period))
		}

		if lengthOfVector4(cSpeed) < minSpeed && stage == 2 {
			cSpeed = cableMinSpeed
			stage = 3
		}

		currentDegrees = sumOfVectors4(currentDegrees, mulVectorByScalar4(cSpeed, period))
		degrees = append(degrees, currentDegrees)
		//fmt.Println(stage,S,deceleration,lengthOfVector4(cSpeed))//strings.Replace(formatstring4(cSpeed),".",",",-1))
		//_, _ = fmt.Fprintln(file, formatstring4(currentDegrees)+"\t"+formatstring4(cSpeed))

		//if distanceBeatenPoints4(currentDegrees, newDegrees) > S && stage == 3 {
		//	fmt.Println("ALARM")
		//	break
		//}
		S = distanceBeatenPoints4(currentDegrees, newDegrees)
	}

	//degrees = append(degrees, newDegrees)

	//generate sets of angles
	pac := make([][]uint16, len(degrees))
	for i := 0; i < len(degrees); i++ {
		pac[i] = PacToUint(degrees[i])
	}

	//send angles package
	sendTrajectory(robot, pac)
	return newPoint
}
