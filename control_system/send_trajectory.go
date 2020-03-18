package main

import (
	"./fins"
	"fmt"
	"time"
)


//Function send one robot package per period
func sendTrajectory(robot *fins.Client, packages [][]uint16) {
	rows := len(packages)

	sendPackage(robot,packages[0])
	ONControlMode(robot)

	//create timer
	period := 4*time.Millisecond
	done := make(chan bool, 1)
	ticker := time.NewTicker(period)

	// current line
	counter := 0
	fmt.Println("robot move")
	go func () {
		for counter < rows{
			select {
			case <- ticker.C:
				sendPackage(robot,packages[counter])
				counter ++
			}
		}
		done <- true
	}()
	<- done
	fmt.Println("robot stop")
}
